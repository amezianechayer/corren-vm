// Code generated from FaRl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FaRl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FaRlListener is a complete listener for a parse tree produced by FaRlParser.
type FaRlListener interface {
	antlr.ParseTreeListener

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)
}
