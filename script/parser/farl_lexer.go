// Code generated from FaRl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 27, 173,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 3, 2, 6, 2, 55, 10, 2,
	13, 2, 14, 2, 56, 3, 3, 6, 3, 60, 10, 3, 13, 3, 14, 3, 61, 3, 3, 3, 3,
	3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 16,
	3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 7, 22, 147, 10, 22, 12, 22, 14, 22, 150, 11, 22,
	3, 23, 3, 23, 3, 23, 7, 23, 155, 10, 23, 12, 23, 14, 23, 158, 11, 23, 3,
	24, 3, 24, 7, 24, 162, 10, 24, 12, 24, 14, 24, 165, 11, 24, 3, 25, 3, 25,
	3, 26, 6, 26, 170, 10, 26, 13, 26, 14, 26, 171, 2, 2, 27, 3, 3, 5, 4, 7,
	5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27,
	15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45,
	24, 47, 25, 49, 26, 51, 27, 3, 2, 10, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11,
	34, 34, 4, 2, 97, 97, 99, 124, 5, 2, 50, 59, 97, 97, 99, 124, 5, 2, 50,
	60, 97, 97, 99, 124, 3, 2, 67, 92, 4, 2, 50, 59, 67, 92, 3, 2, 50, 59,
	2, 178, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 3, 54, 3, 2, 2, 2, 5, 59, 3,
	2, 2, 2, 7, 65, 3, 2, 2, 2, 9, 71, 3, 2, 2, 2, 11, 76, 3, 2, 2, 2, 13,
	85, 3, 2, 2, 2, 15, 90, 3, 2, 2, 2, 17, 93, 3, 2, 2, 2, 19, 97, 3, 2, 2,
	2, 21, 105, 3, 2, 2, 2, 23, 111, 3, 2, 2, 2, 25, 118, 3, 2, 2, 2, 27, 127,
	3, 2, 2, 2, 29, 129, 3, 2, 2, 2, 31, 131, 3, 2, 2, 2, 33, 133, 3, 2, 2,
	2, 35, 135, 3, 2, 2, 2, 37, 137, 3, 2, 2, 2, 39, 139, 3, 2, 2, 2, 41, 141,
	3, 2, 2, 2, 43, 143, 3, 2, 2, 2, 45, 151, 3, 2, 2, 2, 47, 159, 3, 2, 2,
	2, 49, 166, 3, 2, 2, 2, 51, 169, 3, 2, 2, 2, 53, 55, 9, 2, 2, 2, 54, 53,
	3, 2, 2, 2, 55, 56, 3, 2, 2, 2, 56, 54, 3, 2, 2, 2, 56, 57, 3, 2, 2, 2,
	57, 4, 3, 2, 2, 2, 58, 60, 9, 3, 2, 2, 59, 58, 3, 2, 2, 2, 60, 61, 3, 2,
	2, 2, 61, 59, 3, 2, 2, 2, 61, 62, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 64,
	8, 3, 2, 2, 64, 6, 3, 2, 2, 2, 65, 66, 7, 114, 2, 2, 66, 67, 7, 116, 2,
	2, 67, 68, 7, 107, 2, 2, 68, 69, 7, 112, 2, 2, 69, 70, 7, 118, 2, 2, 70,
	8, 3, 2, 2, 2, 71, 72, 7, 104, 2, 2, 72, 73, 7, 99, 2, 2, 73, 74, 7, 107,
	2, 2, 74, 75, 7, 110, 2, 2, 75, 10, 3, 2, 2, 2, 76, 77, 7, 118, 2, 2, 77,
	78, 7, 116, 2, 2, 78, 79, 7, 99, 2, 2, 79, 80, 7, 112, 2, 2, 80, 81, 7,
	117, 2, 2, 81, 82, 7, 104, 2, 2, 82, 83, 7, 103, 2, 2, 83, 84, 7, 116,
	2, 2, 84, 12, 3, 2, 2, 2, 85, 86, 7, 104, 2, 2, 86, 87, 7, 116, 2, 2, 87,
	88, 7, 113, 2, 2, 88, 89, 7, 111, 2, 2, 89, 14, 3, 2, 2, 2, 90, 91, 7,
	118, 2, 2, 91, 92, 7, 113, 2, 2, 92, 16, 3, 2, 2, 2, 93, 94, 7, 120, 2,
	2, 94, 95, 7, 99, 2, 2, 95, 96, 7, 116, 2, 2, 96, 18, 3, 2, 2, 2, 97, 98,
	7, 99, 2, 2, 98, 99, 7, 101, 2, 2, 99, 100, 7, 101, 2, 2, 100, 101, 7,
	113, 2, 2, 101, 102, 7, 119, 2, 2, 102, 103, 7, 112, 2, 2, 103, 104, 7,
	118, 2, 2, 104, 20, 3, 2, 2, 2, 105, 106, 7, 99, 2, 2, 106, 107, 7, 117,
	2, 2, 107, 108, 7, 117, 2, 2, 108, 109, 7, 103, 2, 2, 109, 110, 7, 118,
	2, 2, 110, 22, 3, 2, 2, 2, 111, 112, 7, 112, 2, 2, 112, 113, 7, 119, 2,
	2, 113, 114, 7, 111, 2, 2, 114, 115, 7, 100, 2, 2, 115, 116, 7, 103, 2,
	2, 116, 117, 7, 116, 2, 2, 117, 24, 3, 2, 2, 2, 118, 119, 7, 111, 2, 2,
	119, 120, 7, 113, 2, 2, 120, 121, 7, 112, 2, 2, 121, 122, 7, 103, 2, 2,
	122, 123, 7, 118, 2, 2, 123, 124, 7, 99, 2, 2, 124, 125, 7, 116, 2, 2,
	125, 126, 7, 123, 2, 2, 126, 26, 3, 2, 2, 2, 127, 128, 7, 45, 2, 2, 128,
	28, 3, 2, 2, 2, 129, 130, 7, 47, 2, 2, 130, 30, 3, 2, 2, 2, 131, 132, 7,
	93, 2, 2, 132, 32, 3, 2, 2, 2, 133, 134, 7, 95, 2, 2, 134, 34, 3, 2, 2,
	2, 135, 136, 7, 125, 2, 2, 136, 36, 3, 2, 2, 2, 137, 138, 7, 127, 2, 2,
	138, 38, 3, 2, 2, 2, 139, 140, 7, 48, 2, 2, 140, 40, 3, 2, 2, 2, 141, 142,
	7, 60, 2, 2, 142, 42, 3, 2, 2, 2, 143, 144, 7, 38, 2, 2, 144, 148, 9, 4,
	2, 2, 145, 147, 9, 5, 2, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2, 2,
	148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 44, 3, 2, 2, 2, 150, 148,
	3, 2, 2, 2, 151, 152, 7, 66, 2, 2, 152, 156, 9, 4, 2, 2, 153, 155, 9, 6,
	2, 2, 154, 153, 3, 2, 2, 2, 155, 158, 3, 2, 2, 2, 156, 154, 3, 2, 2, 2,
	156, 157, 3, 2, 2, 2, 157, 46, 3, 2, 2, 2, 158, 156, 3, 2, 2, 2, 159, 163,
	9, 7, 2, 2, 160, 162, 9, 8, 2, 2, 161, 160, 3, 2, 2, 2, 162, 165, 3, 2,
	2, 2, 163, 161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 48, 3, 2, 2, 2,
	165, 163, 3, 2, 2, 2, 166, 167, 9, 9, 2, 2, 167, 50, 3, 2, 2, 2, 168, 170,
	9, 9, 2, 2, 169, 168, 3, 2, 2, 2, 170, 171, 3, 2, 2, 2, 171, 169, 3, 2,
	2, 2, 171, 172, 3, 2, 2, 2, 172, 52, 3, 2, 2, 2, 9, 2, 56, 61, 148, 156,
	163, 171, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "'print'", "'fail'", "'transfer'", "'from'", "'to'", "'var'",
	"'account'", "'asset'", "'number'", "'monetary'", "'+'", "'-'", "'['",
	"']'", "'{'", "'}'", "'.'", "':'",
}

var lexerSymbolicNames = []string{
	"", "NEWLINE", "WHITESPACE", "PRINT", "FAIL", "TRANSFER", "FROM", "TO",
	"VAR", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "OP_ADD",
	"OP_SUB", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "DOT", "COLON", "VARIABLE_NAME",
	"ACCOUNT", "ASSET", "PRECISION", "NUMBER",
}

var lexerRuleNames = []string{
	"NEWLINE", "WHITESPACE", "PRINT", "FAIL", "TRANSFER", "FROM", "TO", "VAR",
	"TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "OP_ADD", "OP_SUB",
	"LBRACK", "RBRACK", "LBRACE", "RBRACE", "DOT", "COLON", "VARIABLE_NAME",
	"ACCOUNT", "ASSET", "PRECISION", "NUMBER",
}

type FaRlLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewFaRlLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *FaRlLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFaRlLexer(input antlr.CharStream) *FaRlLexer {
	l := new(FaRlLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "FaRl.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// FaRlLexer tokens.
const (
	FaRlLexerNEWLINE       = 1
	FaRlLexerWHITESPACE    = 2
	FaRlLexerPRINT         = 3
	FaRlLexerFAIL          = 4
	FaRlLexerTRANSFER      = 5
	FaRlLexerFROM          = 6
	FaRlLexerTO            = 7
	FaRlLexerVAR           = 8
	FaRlLexerTY_ACCOUNT    = 9
	FaRlLexerTY_ASSET      = 10
	FaRlLexerTY_NUMBER     = 11
	FaRlLexerTY_MONETARY   = 12
	FaRlLexerOP_ADD        = 13
	FaRlLexerOP_SUB        = 14
	FaRlLexerLBRACK        = 15
	FaRlLexerRBRACK        = 16
	FaRlLexerLBRACE        = 17
	FaRlLexerRBRACE        = 18
	FaRlLexerDOT           = 19
	FaRlLexerCOLON         = 20
	FaRlLexerVARIABLE_NAME = 21
	FaRlLexerACCOUNT       = 22
	FaRlLexerASSET         = 23
	FaRlLexerPRECISION     = 24
	FaRlLexerNUMBER        = 25
)
