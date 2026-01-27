package compiler

import (
	"strconv"

	"aurex-vm/script/parser"
	"aurex-vm/vm/program"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type parseListener struct {
	*parser.BaseFaRlListener // Gardez BaseFaRlListener si c'est généré par ANTLR
	program                  program.Program
}

func (p *parseListener) ExitAddSub(c *parser.AddSubContext) {
	switch c.GetOp().GetTokenType() {
	case parser.FaRlParserOP_ADD: // Gardez FaRlParser si c'est le nom généré
		p.program = append(p.program, program.OP_IADD)
	case parser.FaRlLexerOP_SUB: // Changez NumScriptLexer en FaRlLexer
		p.program = append(p.program, program.OP_ISUB)
	}
}

func (p *parseListener) ExitNumber(c *parser.NumberContext) {
	a, _ := strconv.Atoi(c.GetText())
	p.program = append(p.program, program.OP_IPUSH, byte(a))
}

func (p *parseListener) ExitScript(c *parser.ScriptContext) {
	p.program = append(p.program, program.OP_PRINT)
}

func Compile(input string) program.Program {
	listener := parseListener{
		program: program.Program{},
	}

	is := antlr.NewInputStream(input)
	lexer := parser.NewFaRlLexer(is) // Changez NumScriptLexer en FaRlLexer
	stream := antlr.NewCommonTokenStream(lexer, antlr.LexerDefaultTokenChannel)
	p := parser.NewFaRlParser(stream)
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Script())

	return listener.program
}
