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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 18, 107,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 3,
	2, 6, 2, 37, 10, 2, 13, 2, 14, 2, 38, 3, 3, 6, 3, 42, 10, 3, 13, 3, 14,
	3, 43, 3, 3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4,
	3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7,
	3, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 9, 3, 9, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 7, 14, 89,
	10, 14, 12, 14, 14, 14, 92, 11, 14, 3, 15, 3, 15, 7, 15, 96, 10, 15, 12,
	15, 14, 15, 99, 11, 15, 3, 16, 3, 16, 3, 17, 6, 17, 104, 10, 17, 13, 17,
	14, 17, 105, 2, 2, 18, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17,
	10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16, 31, 17, 33, 18, 3,
	2, 9, 4, 2, 12, 12, 15, 15, 4, 2, 11, 11, 34, 34, 4, 2, 97, 97, 99, 124,
	5, 2, 50, 60, 97, 97, 99, 124, 3, 2, 67, 92, 4, 2, 50, 59, 67, 92, 3, 2,
	50, 59, 2, 111, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2,
	9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2,
	2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2,
	2, 2, 25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2,
	2, 2, 2, 33, 3, 2, 2, 2, 3, 36, 3, 2, 2, 2, 5, 41, 3, 2, 2, 2, 7, 47, 3,
	2, 2, 2, 9, 56, 3, 2, 2, 2, 11, 62, 3, 2, 2, 2, 13, 67, 3, 2, 2, 2, 15,
	72, 3, 2, 2, 2, 17, 75, 3, 2, 2, 2, 19, 77, 3, 2, 2, 2, 21, 79, 3, 2, 2,
	2, 23, 81, 3, 2, 2, 2, 25, 83, 3, 2, 2, 2, 27, 85, 3, 2, 2, 2, 29, 93,
	3, 2, 2, 2, 31, 100, 3, 2, 2, 2, 33, 103, 3, 2, 2, 2, 35, 37, 9, 2, 2,
	2, 36, 35, 3, 2, 2, 2, 37, 38, 3, 2, 2, 2, 38, 36, 3, 2, 2, 2, 38, 39,
	3, 2, 2, 2, 39, 4, 3, 2, 2, 2, 40, 42, 9, 3, 2, 2, 41, 40, 3, 2, 2, 2,
	42, 43, 3, 2, 2, 2, 43, 41, 3, 2, 2, 2, 43, 44, 3, 2, 2, 2, 44, 45, 3,
	2, 2, 2, 45, 46, 8, 3, 2, 2, 46, 6, 3, 2, 2, 2, 47, 48, 7, 118, 2, 2, 48,
	49, 7, 116, 2, 2, 49, 50, 7, 99, 2, 2, 50, 51, 7, 112, 2, 2, 51, 52, 7,
	117, 2, 2, 52, 53, 7, 104, 2, 2, 53, 54, 7, 103, 2, 2, 54, 55, 7, 116,
	2, 2, 55, 8, 3, 2, 2, 2, 56, 57, 7, 114, 2, 2, 57, 58, 7, 116, 2, 2, 58,
	59, 7, 107, 2, 2, 59, 60, 7, 112, 2, 2, 60, 61, 7, 118, 2, 2, 61, 10, 3,
	2, 2, 2, 62, 63, 7, 104, 2, 2, 63, 64, 7, 99, 2, 2, 64, 65, 7, 107, 2,
	2, 65, 66, 7, 110, 2, 2, 66, 12, 3, 2, 2, 2, 67, 68, 7, 104, 2, 2, 68,
	69, 7, 116, 2, 2, 69, 70, 7, 113, 2, 2, 70, 71, 7, 111, 2, 2, 71, 14, 3,
	2, 2, 2, 72, 73, 7, 118, 2, 2, 73, 74, 7, 113, 2, 2, 74, 16, 3, 2, 2, 2,
	75, 76, 7, 45, 2, 2, 76, 18, 3, 2, 2, 2, 77, 78, 7, 47, 2, 2, 78, 20, 3,
	2, 2, 2, 79, 80, 7, 93, 2, 2, 80, 22, 3, 2, 2, 2, 81, 82, 7, 95, 2, 2,
	82, 24, 3, 2, 2, 2, 83, 84, 7, 48, 2, 2, 84, 26, 3, 2, 2, 2, 85, 86, 7,
	66, 2, 2, 86, 90, 9, 4, 2, 2, 87, 89, 9, 5, 2, 2, 88, 87, 3, 2, 2, 2, 89,
	92, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 90, 91, 3, 2, 2, 2, 91, 28, 3, 2, 2,
	2, 92, 90, 3, 2, 2, 2, 93, 97, 9, 6, 2, 2, 94, 96, 9, 7, 2, 2, 95, 94,
	3, 2, 2, 2, 96, 99, 3, 2, 2, 2, 97, 95, 3, 2, 2, 2, 97, 98, 3, 2, 2, 2,
	98, 30, 3, 2, 2, 2, 99, 97, 3, 2, 2, 2, 100, 101, 9, 8, 2, 2, 101, 32,
	3, 2, 2, 2, 102, 104, 9, 8, 2, 2, 103, 102, 3, 2, 2, 2, 104, 105, 3, 2,
	2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 34, 3, 2, 2, 2,
	8, 2, 38, 43, 90, 97, 105, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "", "", "'transfer'", "'print'", "'fail'", "'from'", "'to'", "'+'",
	"'-'", "'['", "']'", "'.'",
}

var lexerSymbolicNames = []string{
	"", "NEWLINE", "WHITESPACE", "TRANSFER", "PRINT", "FAIL", "FROM", "TO",
	"OP_ADD", "OP_SUB", "LBRACK", "RBRACK", "DOT", "IDENTIFIER", "ASSET", "PRECISION",
	"NUMBER",
}

var lexerRuleNames = []string{
	"NEWLINE", "WHITESPACE", "TRANSFER", "PRINT", "FAIL", "FROM", "TO", "OP_ADD",
	"OP_SUB", "LBRACK", "RBRACK", "DOT", "IDENTIFIER", "ASSET", "PRECISION",
	"NUMBER",
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
	FaRlLexerNEWLINE    = 1
	FaRlLexerWHITESPACE = 2
	FaRlLexerTRANSFER   = 3
	FaRlLexerPRINT      = 4
	FaRlLexerFAIL       = 5
	FaRlLexerFROM       = 6
	FaRlLexerTO         = 7
	FaRlLexerOP_ADD     = 8
	FaRlLexerOP_SUB     = 9
	FaRlLexerLBRACK     = 10
	FaRlLexerRBRACK     = 11
	FaRlLexerDOT        = 12
	FaRlLexerIDENTIFIER = 13
	FaRlLexerASSET      = 14
	FaRlLexerPRECISION  = 15
	FaRlLexerNUMBER     = 16
)
