package compiler

import (
	"fmt"
	"strconv"

	"github.com/amezianechayer/aurex-vm/script/parser"
	"github.com/amezianechayer/aurex-vm/vm/program"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type parseListener struct {
	*parser.BaseFaRlListener
	program program.Program
}

func (p *parseListener) ExitFail(c *parser.FailContext) {
	p.program = append(p.program, program.OP_FAIL)
}

func (p *parseListener) ExitCalc(c *parser.CalcContext) {
	p.program = append(p.program, program.OP_PRINT)
}

func (p *parseListener) ExitAddSub(c *parser.AddSubContext) {
	switch c.GetOp().GetTokenType() {
	case parser.FaRlParserOP_ADD:
		p.program = append(p.program, program.OP_IADD)
	case parser.FaRlLexerOP_SUB:
		p.program = append(p.program, program.OP_ISUB)
	}
}

func (p *parseListener) ExitNumber(c *parser.NumberContext) {
	a, _ := strconv.Atoi(c.GetText())
	p.program = append(p.program, program.OP_IPUSH, byte(a))
}

type SyntaxError struct {
	line, column int
	msg          string
}

type CompileError []SyntaxError

func (c *CompileError) Error() string {
	s := ""
	for _, e := range *c {
		s = fmt.Sprintf("%v\nerror:%v:%v   %v", s, e.line, e.column, e.msg)
	}
	return s
}

type ErrorListener struct {
	*antlr.DefaultErrorListener
	Errors []SyntaxError
}

func (l *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.Errors = append(l.Errors, SyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}

func Compile(input string) (program.Program, error) {
	listener := parseListener{
		program: program.Program{},
	}

	elistener := &ErrorListener{}

	is := antlr.NewInputStream(input)
	lexer := parser.NewFaRlLexer(is)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(elistener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewFaRlParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(elistener)

	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Script())

	if len(elistener.Errors) != 0 {
		return nil, (*CompileError)(&elistener.Errors)
	}

	return listener.program, nil
}
