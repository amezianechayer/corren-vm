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
	elistener    *ErrorListener
	instructions []byte
	constants    []string /* must not exceed 32768 elements */
}

func (p *parseVisitor) AllocateString(str string) (uint16, error) {
	for i := 0; i < len(p.constants); i++ {
		if p.constants[i] == str {
			return uint16(i), nil
		}
	}
	if len(p.constants) >= 32768 {
		return 0, errors.New("number of unique constants exceeded 32768")
	}
	p.constants = append(p.constants, str)
	return uint16(len(p.constants) - 1), nil
}

func (p *parseVisitor) PushValue(val core.Value) error {
	switch val := val.(type) {
	case *core.Address:
		p.instructions = append(p.instructions, program.OP_PUSH2)
		addr, err := p.AllocateString(string(*val))
		if err != nil {
			return err
		}
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, addr)
		p.instructions = append(p.instructions, bytes...)
	case *core.Asset:
		p.instructions = append(p.instructions, program.OP_PUSH2)
		addr, err := p.AllocateString(string(*val))
		if err != nil {
			return err
		}
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, addr)
		p.instructions = append(p.instructions, bytes...)
	case *core.Number:
		p.instructions = append(p.instructions, program.OP_PUSH8)
		bytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, uint64(*val))
		p.instructions = append(p.instructions, bytes...)
	case *core.Monetary:
		p.instructions = append(p.instructions, program.OP_PUSH2)
		addr, err := p.AllocateString(val.Asset)
		if err != nil {
			return err
		}
		bytes := make([]byte, 2)
		binary.LittleEndian.PutUint16(bytes, addr)
		p.instructions = append(p.instructions, bytes...)

		p.instructions = append(p.instructions, program.OP_PUSH8)
		bytes = make([]byte, 8)
		binary.LittleEndian.PutUint64(bytes, val.Amount)
		p.instructions = append(p.instructions, bytes...)
	default:
		panic("unreachable")
	}
	return nil
}

func (p *parseVisitor) VisitScript(c parser.IScriptContext) error {
	switch c := c.(type) {
	case *parser.ScriptContext:
		for _, stmt := range c.GetStmts() {
			switch c := stmt.(type) {
			case *parser.PrintContext:
				err := p.VisitPrint(c)
				if err != nil {
					p.elistener.LogicError(c.GetParser().GetCurrentToken(), err)
					return err
				}
			case *parser.FailContext:
				p.VisitFail(c)
			case *parser.TransferContext:
				err := p.VisitTransfer(c)
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

func (p *parseVisitor) VisitPrint(ctx *parser.PrintContext) error {
	ty, err := p.VisitExpr(ctx.GetExpr())
	if err != nil {
		return err
	}
	if ty != core.TYPE_NUMBER {
		return errors.New("print can only print numbers")
	}
	p.instructions = append(p.instructions, program.OP_PRINT)
	return nil
}

func (p *parseVisitor) VisitExpr(ctx parser.IExpressionContext) (core.ValueType, error) {
	switch ctx := ctx.(type) {
	case *parser.ExprAddSubContext:
		ty, err := p.VisitExpr(ctx.GetLhs())
		if err != nil {
			return 0, err
		}
		if ty != core.TYPE_NUMBER {
			return 0, errors.New("tried to do arithmetic with wrong type")
		}
		ty, err = p.VisitExpr(ctx.GetRhs())
		if err != nil {
			return 0, nil
		}
		if ty != core.TYPE_NUMBER {
			return 0, errors.New("tried to do arithmetic with wrong type")
		}
		switch ctx.GetOp().GetTokenType() {
		case parser.FaRlLexerOP_ADD:
			p.instructions = append(p.instructions, program.OP_IADD)
		case parser.FaRlLexerOP_SUB:
			p.instructions = append(p.instructions, program.OP_ISUB)
		}
		return core.TYPE_NUMBER, nil
	case *parser.ExprLiteralContext:
		val, err := p.VisitLit(ctx.GetLit())
		if err != nil {
			return 0, err
		}
		err = p.PushValue(val)
		if err != nil {
			return 0, err
		}
		return val.GetType(), nil
	default:
		panic("unreachable")
	}
}

func (p *parseVisitor) VisitLit(c parser.ILiteralContext) (core.Value, error) {
	switch c := c.(type) {
	case *parser.LitAddressContext:
		addr := core.Address(c.GetText())
		return &addr, nil
	case *parser.LitAssetContext:
		asset := core.Asset(c.GetText())
		return &asset, nil
	case *parser.LitNumberContext:
		n, err := strconv.ParseUint(c.GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		number := core.Number(n)
		return &number, nil
	case *parser.LitMonetaryContext:
		asset := c.Monetary().GetAsset().GetText()
		amount, err := strconv.ParseUint(c.Monetary().GetAmount().GetText(), 10, 64)
		if err != nil {
			return nil, err
		}
		monetary := core.Monetary{
			Asset:  asset,
			Amount: amount,
		}
		return &monetary, nil
	default:
		panic("unreachable")
	}
}

func (p *parseVisitor) VisitFail(ctx *parser.FailContext) {
	p.instructions = append(p.instructions, program.OP_FAIL)
}

func (p *parseVisitor) VisitTransfer(ctx *parser.TransferContext) error {
	monetary_ctx := ctx.GetAmount()
	asset := monetary_ctx.GetAsset().GetText()
	amount, err := strconv.ParseUint(monetary_ctx.GetAmount().GetText(), 10, 64)
	if err != nil {
		return err
	}
	mon := &core.Monetary{
		Asset:  asset,
		Amount: amount,
	}
	src, err := p.VisitLit(ctx.GetSource())
	if err != nil {
		return err
	}
	dst, err := p.VisitLit(ctx.GetDest())
	if err != nil {
		return err
	}
	err = p.PushValue(mon)
	if err != nil {
		return err
	}
	err = p.PushValue(src)
	if err != nil {
		return err
	}
	err = p.PushValue(dst)
	if err != nil {
		return err
	}
	p.instructions = append(p.instructions, program.OP_SEND)
	return nil
}

type CompileError struct {
	line, column int
	msg          string
}

type CompileErrorList []CompileError

func (c *CompileErrorList) Error() string {
	s := ""
	for _, e := range *c {
		s = fmt.Sprintf("%v\nerror:%v:%v   %v", s, e.line, e.column, e.msg)
	}
	return s
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors CompileErrorList
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, CompileError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func (l *ErrorListener) LogicError(token antlr.Token, err error) {
	l.Errors = append(l.Errors, CompileError{
		line:   token.GetLine(),
		column: token.GetColumn(),
		msg:    err.Error(),
	})
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
		elistener:    elistener,
		instructions: make([]byte, 0),
		constants:    make([]string, 0),
	}

	_ = visitor.VisitScript(tree)

	if len(elistener.Errors) != 0 {
		return nil, (*CompileErrorList)(&elistener.Errors)
	}

	return &program.Program{
		Instructions: visitor.instructions,
		Constants:    visitor.constants,
		Variables:    make([]string, 0),
	}, nil
}
