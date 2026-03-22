package compiler

import (
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/amezianechayer/aurex-vm/core"
	"github.com/amezianechayer/aurex-vm/script/parser"
	"github.com/amezianechayer/aurex-vm/vm/program"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type parseVisitor struct {
	elistener       *ErrorListener
	instructions    []byte
	constants       []core.Value
	variables       map[string]program.VarInfo
	needed_balances map[core.Address]map[core.Address]struct{}
}

func (p *parseVisitor) AllocateConstant(v core.Value) (core.Address, error) {
	for i := 0; i < len(p.constants); i++ {
		if core.ValueEquals(p.constants[i], v) {
			return core.Address(i), nil
		}
	}
	if len(p.constants) >= 32768 {
		return 0, errors.New("number of unique constants exceeded 32768")
	}
	p.constants = append(p.constants, v)
	return core.Address(len(p.constants) - 1), nil
}

func (p *parseVisitor) PushValue(val core.Value) (*core.Address, error) {
	switch val := val.(type) {
	case core.Account, core.Asset, core.Monetary, core.Allotment, core.Portion:
		p.instructions = append(p.instructions, program.OP_APUSH)
		addr, err := p.AllocateConstant(val)
		if err != nil {
			return nil, err
		}
		bytes := addr.ToBytes()
		p.instructions = append(p.instructions, bytes...)
		return &addr, nil
	case core.Number:
		p.instructions = append(p.instructions, program.OP_IPUSH)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(val))
		p.instructions = append(p.instructions, bytes...)
		return nil, nil
	default:
		panic("internal compiler error")
	}
}

func (p *parseVisitor) isMonetaryAll(addr core.Address) bool {
	if addr.IsConstant() {
		idx := addr.ToIdx()
		if idx < len(p.constants) {
			if mon, ok := p.constants[idx].(core.Monetary); ok {
				return mon.Amount.All
			}
		}
	}
	return false
}

func (p *parseVisitor) isWorld(addr core.Address) bool {
	if addr.IsConstant() {
		idx := addr.ToIdx()
		if idx < len(p.constants) {
			if acc, ok := p.constants[idx].(core.Account); ok {
				if string(acc) == "@world" {
					return true
				}
			}
		}
	}
	return false
}

func (p *parseVisitor) VisitVariable(c parser.IExpressionContext) (core.Type, *core.Address, error) {
	if ctx, ok := c.(*parser.ExprVariableContext); ok {
		name := ctx.GetVariable().GetText()[1:]
		if info, ok := p.variables[name]; ok {
			p.instructions = append(p.instructions, program.OP_APUSH)
			bytes := info.Addr.ToBytes()
			p.instructions = append(p.instructions, bytes...)
			return info.Ty, &info.Addr, nil
		}
		return 0, nil, fmt.Errorf("variable not declared : %v", name)
	}
	return 0, nil, errors.New("expected variable")
}

func (p *parseVisitor) VisitScript(c parser.IScriptContext) error {
	switch c := c.(type) {
	case *parser.ScriptContext:
		for _, v := range c.AllVarDecl() {
			err := p.VisitVarDecl(v)
			if err != nil {
				return err
			}
		}
		for _, stmt := range c.GetStmts() {
			switch c := stmt.(type) {
			case *parser.PrintContext:
				err := p.VisitPrint(c)
				if err != nil {
					p.elistener.LogicError(c.GetParser().GetCurrentToken(), err)
					return err
				}
			case *parser.FailContext:
				p.instructions = append(p.instructions, program.OP_FAIL)
			case *parser.TransferSimpleContext:
				err := p.VisitTransferSimple(c)
				if err != nil {
					p.elistener.LogicError(c.GetParser().GetCurrentToken(), err)
					return err
				}
			case *parser.TransferWithDestContext:
				err := p.VisitTransferWithDest(c)
				if err != nil {
					p.elistener.LogicError(c.GetParser().GetCurrentToken(), err)
					return err
				}
			default:
				panic("Invalid context")
			}
		}
	default:
		panic("Invalid context")
	}
	return nil
}

func (p *parseVisitor) VisitVarDecl(v parser.IVarDeclContext) error {
	switch v := v.(type) {
	case *parser.VarTypedContext:
		name := v.GetName().GetText()[1:]
		if _, ok := p.variables[name]; ok {
			return fmt.Errorf("duplicate variable: %v", name)
		}
		var ty core.Type
		switch v.GetTy().GetText() {
		case "account":
			ty = core.TYPE_ACCOUNT
		case "asset":
			ty = core.TYPE_ASSET
		case "number":
			ty = core.TYPE_NUMBER
		case "monetary":
			ty = core.TYPE_MONETARY
		case "portion":
			ty = core.TYPE_PORTION
		default:
			return errors.New("internal compiler error: unsupported type")
		}
		addr := core.NewVarAddress(uint16(len(p.variables)))
		p.variables[name] = program.VarInfo{
			Ty:   ty,
			Addr: addr,
		}
	default:
		panic("internal compiler error: unsupported var decl")
	}
	return nil
}

func (p *parseVisitor) VisitPrint(ctx *parser.PrintContext) error {
	_, _, err := p.VisitExpr(ctx.GetExpr())
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_PRINT)
	return nil
}

func (p *parseVisitor) VisitExpr(ctx parser.IExpressionContext) (core.Type, *core.Address, error) {
	switch ctx := ctx.(type) {
	case *parser.ExprAddSubContext:
		ty, _, err := p.VisitExpr(ctx.GetLhs())
		if err != nil {
			return 0, nil, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, nil, errors.New("tried to do arithmetic with wrong type")
		}
		ty, _, err = p.VisitExpr(ctx.GetRhs())
		if err != nil {
			return 0, nil, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, nil, errors.New("tried to do arithmetic with wrong type")
		}
		switch ctx.GetOp().GetTokenType() {
		case parser.FaRlLexerOP_ADD:
			p.instructions = append(p.instructions, program.OP_IADD)
		case parser.FaRlLexerOP_SUB:
			p.instructions = append(p.instructions, program.OP_ISUB)
		}
		return core.TYPE_NUMBER, nil, nil
	case *parser.ExprLiteralContext:
		ty, addr, err := p.VisitLit(ctx.GetLit())
		if err != nil {
			return 0, nil, err
		}
		return ty, addr, nil
	case *parser.ExprVariableContext:
		name := ctx.GetVariable().GetText()[1:]
		if info, ok := p.variables[name]; ok {
			p.instructions = append(p.instructions, program.OP_APUSH)
			bytes := info.Addr.ToBytes()
			p.instructions = append(p.instructions, bytes...)
			return info.Ty, &info.Addr, nil
		} else {
			return 0, nil, fmt.Errorf("variable not declared : %v", name)
		}
	default:
		panic("internal compiler error")
	}
}

func (p *parseVisitor) VisitLit(c parser.ILiteralContext) (core.Type, *core.Address, error) {
	switch c := c.(type) {
	case *parser.LitAccountContext:
		account := core.Account(c.ACCOUNT().GetText())
		addr, err := p.PushValue(account)
		if err != nil {
			return 0, nil, err
		}
		return core.TYPE_ACCOUNT, addr, nil
	case *parser.LitAssetContext:
		asset := core.Asset(c.GetText())
		addr, err := p.PushValue(asset)
		if err != nil {
			return 0, nil, err
		}
		return core.TYPE_ASSET, addr, nil
	case *parser.LitNumberContext:
		n, err := strconv.ParseUint(c.GetText(), 10, 64)
		if err != nil {
			return 0, nil, err
		}
		p.PushValue(core.Number(n))
		return core.TYPE_NUMBER, nil, nil
	case *parser.LitMonetaryContext:
		switch m := c.Monetary().(type) {
		case *parser.MonetaryLitContext:
			asset := m.GetAsset().GetText()
			precision := m.GetPrecision().GetText()
			a, err := strconv.ParseUint(m.GetAmount().GetText(), 10, 64)
			if err != nil {
				return 0, nil, err
			}
			monetary := core.Monetary{
				Asset:  core.Asset(asset + "." + precision),
				Amount: core.NewAmountSpecific(a),
			}
			addr, err := p.PushValue(monetary)
			if err != nil {
				return 0, nil, err
			}
			return core.TYPE_MONETARY, addr, nil
		case *parser.MonetaryAllContext:
			asset := m.GetAsset().GetText()
			precision := m.GetPrecision().GetText()
			monetary := core.Monetary{
				Asset:  core.Asset(asset + "." + precision),
				Amount: core.NewAmountAll(),
			}
			addr, err := p.PushValue(monetary)
			if err != nil {
				return 0, nil, err
			}
			return core.TYPE_MONETARY, addr, nil
		case *parser.MonetaryNoPrecisionContext:
			asset := m.GetAsset().GetText()
			a, err := strconv.ParseUint(m.GetAmount().GetText(), 10, 64)
			if err != nil {
				return 0, nil, err
			}
			monetary := core.Monetary{
				Asset:  core.Asset(asset),
				Amount: core.NewAmountSpecific(a),
			}
			addr, err := p.PushValue(monetary)
			if err != nil {
				return 0, nil, err
			}
			return core.TYPE_MONETARY, addr, nil
		case *parser.MonetaryNoPrecisionAllContext:
			asset := m.GetAsset().GetText()
			monetary := core.Monetary{
				Asset:  core.Asset(asset),
				Amount: core.NewAmountAll(),
			}
			addr, err := p.PushValue(monetary)
			if err != nil {
				return 0, nil, err
			}
			return core.TYPE_MONETARY, addr, nil
		default:
			panic("internal compiler error")
		}
	default:
		panic("internal compiler error")
	}
}

func (p *parseVisitor) VisitSource(ctx parser.ISourceContext) ([]core.Address, bool, error) {
	needed_accounts := []core.Address{}
	bottomless := false
	switch ctx := ctx.(type) {
	case *parser.SrcSimpleContext:
		ty, addr, err := p.VisitExpr(ctx.Expression())
		if err != nil {
			return nil, false, err
		}
		if ty != core.TYPE_ACCOUNT {
			return nil, false, errors.New("expected account as source")
		}
		bottomless = bottomless || p.isWorld(*addr)
		needed_accounts = append(needed_accounts, *addr)
		p.PushValue(core.Number(1))
	case *parser.SrcCascadeContext:
		sources := collectCascade(ctx)
		n := len(sources)
		for i := n - 1; i >= 0; i-- {
			ty, addr, err := p.VisitExpr(sources[i])
			if err != nil {
				return nil, false, err
			}
			if ty != core.TYPE_ACCOUNT {
				return nil, false, errors.New("expected only accounts in sources")
			}
			bottomless = bottomless || p.isWorld(*addr)
			needed_accounts = append(needed_accounts, *addr)
		}
		p.PushValue(core.Number(n))
		p.instructions = append(p.instructions, program.OP_SOURCE)
	default:
		panic("internal compiler error: unsupported source type")
	}
	return needed_accounts, bottomless, nil
}

func collectCascade(ctx parser.ISourceContext) []parser.IExpressionContext {
	switch ctx := ctx.(type) {
	case *parser.SrcCascadeContext:
		left := collectCascade(ctx.Source())
		right := ctx.Expression()
		return append(left, right)
	case *parser.SrcSimpleContext:
		return []parser.IExpressionContext{ctx.Expression()}
	default:
		return []parser.IExpressionContext{}
	}
}

func (p *parseVisitor) VisitTransferSimple(ctx *parser.TransferSimpleContext) error {
	ty, mon_addr, err := p.VisitExpr(ctx.GetAmount())
	if err != nil {
		return err
	}
	if ty != core.TYPE_MONETARY {
		return errors.New("wrong type for monetary value")
	}

	needed_accounts, bottomless, err := p.VisitSource(ctx.GetSrc())
	if err != nil {
		return err
	}
	if p.isMonetaryAll(*mon_addr) && bottomless {
		return errors.New("cannot take all balance of @world")
	}

	for _, acc := range needed_accounts {
		if b, ok := p.needed_balances[acc]; ok {
			b[*mon_addr] = struct{}{}
		} else {
			p.needed_balances[acc] = map[core.Address]struct{}{
				*mon_addr: {},
			}
		}
	}

	dstTy, _, err := p.VisitExpr(ctx.GetDest())
	if err != nil {
		return err
	}
	if dstTy != core.TYPE_ACCOUNT {
		return errors.New("wrong type for destination")
	}

	p.instructions = append(p.instructions, program.OP_SEND)
	return nil
}

func (p *parseVisitor) VisitTransferWithDest(ctx *parser.TransferWithDestContext) error {
	ty, mon_addr, err := p.VisitExpr(ctx.GetAmount())
	if err != nil {
		return err
	}
	if ty != core.TYPE_MONETARY {
		return errors.New("wrong type for monetary value")
	}

	needed_accounts, bottomless, err := p.VisitSource(ctx.GetSrc())
	if err != nil {
		return err
	}
	if p.isMonetaryAll(*mon_addr) && bottomless {
		return errors.New("cannot take all balance of @world")
	}

	for _, acc := range needed_accounts {
		if b, ok := p.needed_balances[acc]; ok {
			b[*mon_addr] = struct{}{}
		} else {
			p.needed_balances[acc] = map[core.Address]struct{}{
				*mon_addr: {},
			}
		}
	}

	err = p.VisitAllocation(ctx.GetSends())
	if err != nil {
		return err
	}

	return nil
}

func Compile(input string) (*program.Program, error) {
	elistener := &ErrorListener{}

	is := antlr.NewInputStream(input)
	lexer := parser.NewFaRlLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(elistener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewFaRlParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(elistener)

	p.BuildParseTrees = true

	tree := p.Script()

	if len(elistener.Errors) != 0 {
		return nil, (*CompileErrorList)(&elistener.Errors)
	}

	visitor := parseVisitor{
		elistener:       elistener,
		instructions:    make([]byte, 0),
		constants:       make([]core.Value, 0),
		variables:       make(map[string]program.VarInfo),
		needed_balances: make(map[core.Address]map[core.Address]struct{}),
	}

	_ = visitor.VisitScript(tree)

	if len(elistener.Errors) != 0 {
		return nil, (*CompileErrorList)(&elistener.Errors)
	}

	return &program.Program{
		Instructions:   visitor.instructions,
		Constants:      visitor.constants,
		Variables:      visitor.variables,
		NeededBalances: visitor.needed_balances,
	}, nil
}
