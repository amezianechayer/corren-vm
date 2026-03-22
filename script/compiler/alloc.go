package compiler

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/amezianechayer/aurex-vm/core"
	"github.com/amezianechayer/aurex-vm/script/parser"
	"github.com/amezianechayer/aurex-vm/vm/program"
)

func (p *parseVisitor) VisitAllocation(parts []parser.ISendClauseContext) *CompileError {
	hasDynamic := false
	for _, part := range parts {
		if send, ok := part.(*parser.SendToContext); ok {
			switch send.Portion().(type) {
			case *parser.PortionRemainingContext, *parser.PortionVarContext:
				hasDynamic = true
			}
		}
		if _, ok := part.(*parser.SendKeepContext); ok {
			hasDynamic = true
		}
	}

	if hasDynamic {
		return p.VisitAllocationDyn(parts)
	}
	return p.VisitAllocationConst(parts)
}

func (p *parseVisitor) VisitAllocationConst(parts []parser.ISendClauseContext) *CompileError {
	portions := []core.Portion{}
	for _, part := range parts {
		switch part := part.(type) {
		case *parser.SendToContext:
			switch por := part.Portion().(type) {
			case *parser.PortionPercentContext:
				pint := por.GetP().GetText()
				pfrac := "0"
				if por.GetPfrac() != nil {
					pfrac = por.GetPfrac().GetText()
				}
				portion, err := core.ParsePortionSpecific(pint + "." + pfrac + "%")
				if err != nil {
					return LogicError(part, err)
				}
				portions = append(portions, *portion)
			case *parser.PortionRatioContext:
				portion, err := core.ParsePortionSpecific(por.GetR().GetText())
				if err != nil {
					return LogicError(part, err)
				}
				portions = append(portions, *portion)
			case *parser.PortionRemainingContext:
				portions = append(portions, core.NewPortionRemaining())
			}
		case *parser.SendKeepContext:
			portions = append(portions, core.NewPortionRemaining())
		}
	}

	allotment, err := core.NewAllotment(portions)
	if err != nil {
		return LogicError(parts[0], err)
	}
	p.PushValue(*allotment)
	p.instructions = append(p.instructions, program.OP_ALLOC)
	return p.VisitAllocDestination(parts)
}

func (p *parseVisitor) VisitAllocationDyn(parts []parser.ISendClauseContext) *CompileError {
	total := big.NewRat(0, 1)
	has_remaining := false

	for i := len(parts) - 1; i >= 0; i-- {
		part := parts[i]
		switch part := part.(type) {
		case *parser.SendToContext:
			switch por := part.Portion().(type) {
			case *parser.PortionPercentContext:
				pint := por.GetP().GetText()
				pfrac := "0"
				if por.GetPfrac() != nil {
					pfrac = por.GetPfrac().GetText()
				}
				portion, err := core.ParsePortionSpecific(pint + "." + pfrac + "%")
				if err != nil {
					return LogicError(part, err)
				}
				rat := big.Rat(*portion.Specific)
				total.Add(&rat, total)
				p.PushValue(core.Portion(*portion))
			case *parser.PortionRatioContext:
				portion, err := core.ParsePortionSpecific(por.GetR().GetText())
				if err != nil {
					return LogicError(part, err)
				}
				rat := big.Rat(*portion.Specific)
				total.Add(&rat, total)
				p.PushValue(core.Portion(*portion))
			case *parser.PortionVarContext:
				name := por.GetV().GetText()[1:]
				if info, ok := p.variables[name]; ok {
					if info.Ty != core.TYPE_PORTION {
						return LogicError(part, fmt.Errorf("wrong type: expected type portion for variable: %v", info.Ty))
					}
					p.instructions = append(p.instructions, program.OP_APUSH)
					bytes := info.Addr.ToBytes()
					p.instructions = append(p.instructions, bytes...)
				} else {
					return LogicError(part, fmt.Errorf("variable not declared: %v", name))
				}
			case *parser.PortionRemainingContext:
				if has_remaining {
					return LogicError(part, errors.New("two uses of `remaining` in the same allocation"))
				}
				p.PushValue(core.NewPortionRemaining())
				has_remaining = true
			}
		case *parser.SendKeepContext:
			if has_remaining {
				return LogicError(part, errors.New("two uses of `remaining` in the same allocation"))
			}
			p.PushValue(core.NewPortionRemaining())
			has_remaining = true
		}
	}

	if !has_remaining {
		return LogicError(parts[0], errors.New("allocation has variable portions but no 'remaining'"))
	}
	if total.Cmp(big.NewRat(1, 1)) != -1 {
		return LogicError(parts[0], errors.New("sum of known portions is already equal or is greater than 100%"))
	}

	p.PushValue(core.Number(len(parts)))
	p.instructions = append(p.instructions, program.OP_MAKE_ALLOTMENT)
	p.instructions = append(p.instructions, program.OP_ALLOC)
	return p.VisitAllocDestination(parts)
}

func (p *parseVisitor) VisitAllocDestination(parts []parser.ISendClauseContext) *CompileError {
	for _, part := range parts {
		switch part := part.(type) {
		case *parser.SendToContext:
			ty, _, err := p.VisitExpr(part.Expression())
			if err != nil {
				return err
			}
			if ty != core.TYPE_ACCOUNT {
				return LogicError(part, errors.New("expected account as destination of allocation line"))
			}
			p.instructions = append(p.instructions, program.OP_SEND)
		case *parser.SendKeepContext:
			continue
		}
	}
	return nil
}
