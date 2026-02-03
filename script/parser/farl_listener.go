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

	// EnterCalc is called when entering the Calc production.
	EnterCalc(c *CalcContext)

	// EnterFail is called when entering the Fail production.
	EnterFail(c *FailContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitCalc is called when exiting the Calc production.
	ExitCalc(c *CalcContext)

	// ExitFail is called when exiting the Fail production.
	ExitFail(c *FailContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)
}
