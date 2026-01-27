// Code generated from FaRl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FaRl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFaRlListener is a complete listener for a parse tree produced by FaRlParser.
type BaseFaRlListener struct{}

var _ FaRlListener = &BaseFaRlListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFaRlListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFaRlListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFaRlListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFaRlListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseFaRlListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseFaRlListener) ExitNumber(ctx *NumberContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseFaRlListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseFaRlListener) ExitAddSub(ctx *AddSubContext) {}

// EnterScript is called when production script is entered.
func (s *BaseFaRlListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseFaRlListener) ExitScript(ctx *ScriptContext) {}
