// Code generated from FaRl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FaRl

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 49, 267,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 3, 2, 5, 2, 51, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 63, 10, 3, 3, 4, 3, 4, 3, 4, 3, 4, 5,
	4, 69, 10, 4, 3, 5, 3, 5, 3, 5, 5, 5, 74, 10, 5, 3, 5, 3, 5, 3, 5, 7, 5,
	79, 10, 5, 12, 5, 14, 5, 82, 11, 5, 3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 88, 10,
	6, 3, 6, 3, 6, 3, 6, 5, 6, 93, 10, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3,
	7, 3, 7, 6, 7, 102, 10, 7, 13, 7, 14, 7, 103, 3, 7, 3, 7, 3, 8, 3, 8, 5,
	8, 110, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 6, 9, 119, 10,
	9, 13, 9, 14, 9, 120, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 6, 11, 135, 10, 11, 13, 11, 14, 11, 136,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5, 12, 144, 10, 12, 3, 13, 3, 13, 5,
	13, 148, 10, 13, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15,
	3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 164, 10, 16, 3, 17, 3,
	17, 5, 17, 168, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 5, 19, 180, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	196, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 6, 19, 215,
	10, 19, 13, 19, 14, 19, 216, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19,
	3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19, 231, 10, 19, 3, 20, 7,
	20, 234, 10, 20, 12, 20, 14, 20, 237, 11, 20, 3, 20, 3, 20, 6, 20, 241,
	10, 20, 13, 20, 14, 20, 242, 7, 20, 245, 10, 20, 12, 20, 14, 20, 248, 11,
	20, 3, 20, 3, 20, 6, 20, 252, 10, 20, 13, 20, 14, 20, 253, 3, 20, 7, 20,
	257, 10, 20, 12, 20, 14, 20, 260, 11, 20, 3, 20, 5, 20, 263, 10, 20, 3,
	20, 3, 20, 3, 20, 2, 3, 8, 21, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22,
	24, 26, 28, 30, 32, 34, 36, 38, 2, 4, 3, 2, 30, 31, 3, 2, 24, 29, 2, 281,
	2, 50, 3, 2, 2, 2, 4, 62, 3, 2, 2, 2, 6, 68, 3, 2, 2, 2, 8, 73, 3, 2, 2,
	2, 10, 92, 3, 2, 2, 2, 12, 94, 3, 2, 2, 2, 14, 109, 3, 2, 2, 2, 16, 111,
	3, 2, 2, 2, 18, 124, 3, 2, 2, 2, 20, 129, 3, 2, 2, 2, 22, 143, 3, 2, 2,
	2, 24, 147, 3, 2, 2, 2, 26, 149, 3, 2, 2, 2, 28, 151, 3, 2, 2, 2, 30, 157,
	3, 2, 2, 2, 32, 167, 3, 2, 2, 2, 34, 169, 3, 2, 2, 2, 36, 230, 3, 2, 2,
	2, 38, 235, 3, 2, 2, 2, 40, 41, 7, 34, 2, 2, 41, 42, 7, 48, 2, 2, 42, 43,
	7, 40, 2, 2, 43, 44, 7, 49, 2, 2, 44, 45, 7, 49, 2, 2, 45, 51, 7, 35, 2,
	2, 46, 47, 7, 34, 2, 2, 47, 48, 7, 48, 2, 2, 48, 49, 7, 49, 2, 2, 49, 51,
	7, 35, 2, 2, 50, 40, 3, 2, 2, 2, 50, 46, 3, 2, 2, 2, 51, 3, 3, 2, 2, 2,
	52, 53, 7, 34, 2, 2, 53, 54, 7, 48, 2, 2, 54, 55, 7, 40, 2, 2, 55, 56,
	7, 49, 2, 2, 56, 57, 7, 42, 2, 2, 57, 63, 7, 35, 2, 2, 58, 59, 7, 34, 2,
	2, 59, 60, 7, 48, 2, 2, 60, 61, 7, 42, 2, 2, 61, 63, 7, 35, 2, 2, 62, 52,
	3, 2, 2, 2, 62, 58, 3, 2, 2, 2, 63, 5, 3, 2, 2, 2, 64, 69, 7, 47, 2, 2,
	65, 69, 7, 48, 2, 2, 66, 69, 7, 49, 2, 2, 67, 69, 5, 2, 2, 2, 68, 64, 3,
	2, 2, 2, 68, 65, 3, 2, 2, 2, 68, 66, 3, 2, 2, 2, 68, 67, 3, 2, 2, 2, 69,
	7, 3, 2, 2, 2, 70, 71, 8, 5, 1, 2, 71, 74, 5, 6, 4, 2, 72, 74, 7, 46, 2,
	2, 73, 70, 3, 2, 2, 2, 73, 72, 3, 2, 2, 2, 74, 80, 3, 2, 2, 2, 75, 76,
	12, 5, 2, 2, 76, 77, 9, 2, 2, 2, 77, 79, 5, 8, 5, 6, 78, 75, 3, 2, 2, 2,
	79, 82, 3, 2, 2, 2, 80, 78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 9, 3, 2,
	2, 2, 82, 80, 3, 2, 2, 2, 83, 93, 7, 45, 2, 2, 84, 87, 7, 49, 2, 2, 85,
	86, 7, 40, 2, 2, 86, 88, 7, 49, 2, 2, 87, 85, 3, 2, 2, 2, 87, 88, 3, 2,
	2, 2, 88, 89, 3, 2, 2, 2, 89, 93, 7, 32, 2, 2, 90, 93, 7, 46, 2, 2, 91,
	93, 7, 33, 2, 2, 92, 83, 3, 2, 2, 2, 92, 84, 3, 2, 2, 2, 92, 90, 3, 2,
	2, 2, 92, 91, 3, 2, 2, 2, 93, 11, 3, 2, 2, 2, 94, 95, 7, 36, 2, 2, 95,
	101, 7, 5, 2, 2, 96, 97, 5, 10, 6, 2, 97, 98, 7, 13, 2, 2, 98, 99, 5, 8,
	5, 2, 99, 100, 7, 5, 2, 2, 100, 102, 3, 2, 2, 2, 101, 96, 3, 2, 2, 2, 102,
	103, 3, 2, 2, 2, 103, 101, 3, 2, 2, 2, 103, 104, 3, 2, 2, 2, 104, 105,
	3, 2, 2, 2, 105, 106, 7, 37, 2, 2, 106, 13, 3, 2, 2, 2, 107, 110, 5, 8,
	5, 2, 108, 110, 5, 12, 7, 2, 109, 107, 3, 2, 2, 2, 109, 108, 3, 2, 2, 2,
	110, 15, 3, 2, 2, 2, 111, 112, 7, 36, 2, 2, 112, 118, 7, 5, 2, 2, 113,
	114, 5, 10, 6, 2, 114, 115, 7, 12, 2, 2, 115, 116, 5, 22, 12, 2, 116, 117,
	7, 5, 2, 2, 117, 119, 3, 2, 2, 2, 118, 113, 3, 2, 2, 2, 119, 120, 3, 2,
	2, 2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 122, 3, 2, 2, 2,
	122, 123, 7, 37, 2, 2, 123, 17, 3, 2, 2, 2, 124, 125, 7, 14, 2, 2, 125,
	126, 5, 8, 5, 2, 126, 127, 7, 12, 2, 2, 127, 128, 5, 22, 12, 2, 128, 19,
	3, 2, 2, 2, 129, 130, 7, 36, 2, 2, 130, 134, 7, 5, 2, 2, 131, 132, 5, 22,
	12, 2, 132, 133, 7, 5, 2, 2, 133, 135, 3, 2, 2, 2, 134, 131, 3, 2, 2, 2,
	135, 136, 3, 2, 2, 2, 136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137,
	138, 3, 2, 2, 2, 138, 139, 7, 37, 2, 2, 139, 21, 3, 2, 2, 2, 140, 144,
	5, 8, 5, 2, 141, 144, 5, 18, 10, 2, 142, 144, 5, 20, 11, 2, 143, 140, 3,
	2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 142, 3, 2, 2, 2, 144, 23, 3, 2, 2,
	2, 145, 148, 5, 22, 12, 2, 146, 148, 5, 16, 9, 2, 147, 145, 3, 2, 2, 2,
	147, 146, 3, 2, 2, 2, 148, 25, 3, 2, 2, 2, 149, 150, 9, 3, 2, 2, 150, 27,
	3, 2, 2, 2, 151, 152, 7, 3, 2, 2, 152, 153, 5, 8, 5, 2, 153, 154, 7, 4,
	2, 2, 154, 155, 7, 44, 2, 2, 155, 156, 7, 39, 2, 2, 156, 29, 3, 2, 2, 2,
	157, 158, 7, 17, 2, 2, 158, 159, 7, 46, 2, 2, 159, 160, 7, 41, 2, 2, 160,
	163, 5, 26, 14, 2, 161, 162, 7, 43, 2, 2, 162, 164, 5, 28, 15, 2, 163,
	161, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 31, 3, 2, 2, 2, 165, 168, 5,
	8, 5, 2, 166, 168, 7, 45, 2, 2, 167, 165, 3, 2, 2, 2, 167, 166, 3, 2, 2,
	2, 168, 33, 3, 2, 2, 2, 169, 170, 7, 44, 2, 2, 170, 171, 7, 43, 2, 2, 171,
	172, 5, 32, 17, 2, 172, 35, 3, 2, 2, 2, 173, 174, 7, 9, 2, 2, 174, 231,
	5, 8, 5, 2, 175, 231, 7, 10, 2, 2, 176, 179, 7, 11, 2, 2, 177, 180, 5,
	8, 5, 2, 178, 180, 5, 4, 3, 2, 179, 177, 3, 2, 2, 2, 179, 178, 3, 2, 2,
	2, 180, 181, 3, 2, 2, 2, 181, 182, 7, 38, 2, 2, 182, 195, 7, 5, 2, 2, 183,
	184, 7, 12, 2, 2, 184, 185, 5, 24, 13, 2, 185, 186, 7, 5, 2, 2, 186, 187,
	7, 13, 2, 2, 187, 188, 5, 14, 8, 2, 188, 196, 3, 2, 2, 2, 189, 190, 7,
	13, 2, 2, 190, 191, 5, 14, 8, 2, 191, 192, 7, 5, 2, 2, 192, 193, 7, 12,
	2, 2, 193, 194, 5, 24, 13, 2, 194, 196, 3, 2, 2, 2, 195, 183, 3, 2, 2,
	2, 195, 189, 3, 2, 2, 2, 196, 197, 3, 2, 2, 2, 197, 198, 7, 5, 2, 2, 198,
	199, 7, 39, 2, 2, 199, 231, 3, 2, 2, 2, 200, 201, 7, 18, 2, 2, 201, 202,
	7, 19, 2, 2, 202, 203, 7, 20, 2, 2, 203, 204, 7, 44, 2, 2, 204, 205, 7,
	43, 2, 2, 205, 231, 5, 32, 17, 2, 206, 207, 7, 18, 2, 2, 207, 208, 7, 19,
	2, 2, 208, 209, 7, 20, 2, 2, 209, 210, 7, 36, 2, 2, 210, 214, 7, 5, 2,
	2, 211, 212, 5, 34, 18, 2, 212, 213, 7, 5, 2, 2, 213, 215, 3, 2, 2, 2,
	214, 211, 3, 2, 2, 2, 215, 216, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216,
	217, 3, 2, 2, 2, 217, 218, 3, 2, 2, 2, 218, 219, 7, 37, 2, 2, 219, 231,
	3, 2, 2, 2, 220, 221, 7, 18, 2, 2, 221, 222, 7, 24, 2, 2, 222, 223, 7,
	20, 2, 2, 223, 224, 7, 21, 2, 2, 224, 225, 5, 8, 5, 2, 225, 226, 7, 22,
	2, 2, 226, 227, 7, 44, 2, 2, 227, 228, 7, 43, 2, 2, 228, 229, 5, 32, 17,
	2, 229, 231, 3, 2, 2, 2, 230, 173, 3, 2, 2, 2, 230, 175, 3, 2, 2, 2, 230,
	176, 3, 2, 2, 2, 230, 200, 3, 2, 2, 2, 230, 206, 3, 2, 2, 2, 230, 220,
	3, 2, 2, 2, 231, 37, 3, 2, 2, 2, 232, 234, 7, 5, 2, 2, 233, 232, 3, 2,
	2, 2, 234, 237, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235, 236, 3, 2, 2, 2,
	236, 246, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 238, 240, 5, 30, 16, 2, 239,
	241, 7, 5, 2, 2, 240, 239, 3, 2, 2, 2, 241, 242, 3, 2, 2, 2, 242, 240,
	3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 245, 3, 2, 2, 2, 244, 238, 3, 2,
	2, 2, 245, 248, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2, 246, 247, 3, 2, 2, 2,
	247, 249, 3, 2, 2, 2, 248, 246, 3, 2, 2, 2, 249, 258, 5, 36, 19, 2, 250,
	252, 7, 5, 2, 2, 251, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 251,
	3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255, 257, 5, 36,
	19, 2, 256, 251, 3, 2, 2, 2, 257, 260, 3, 2, 2, 2, 258, 256, 3, 2, 2, 2,
	258, 259, 3, 2, 2, 2, 259, 262, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 261,
	263, 7, 5, 2, 2, 262, 261, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 264,
	3, 2, 2, 2, 264, 265, 7, 2, 2, 3, 265, 39, 3, 2, 2, 2, 27, 50, 62, 68,
	73, 80, 87, 92, 103, 109, 120, 136, 143, 147, 163, 167, 179, 195, 216,
	230, 235, 242, 246, 253, 258, 262,
}
var literalNames = []string{
	"", "'meta('", "','", "", "", "", "", "'print'", "'fail'", "'transfer'",
	"'from'", "'to'", "'max'", "'keep'", "'all'", "'var'", "'set'", "'transaction'",
	"'metadata'", "'of'", "'key'", "'meta'", "'account'", "'asset'", "'number'",
	"'monetary'", "'portion'", "'string'", "'+'", "'-'", "'%'", "'remaining'",
	"'['", "']'", "'{'", "'}'", "'('", "')'", "'.'", "':'", "'*'", "'='",
}
var symbolicNames = []string{
	"", "", "", "NEWLINE", "WHITESPACE", "MULTILINE_COMMENT", "LINE_COMMENT",
	"PRINT", "FAIL", "TRANSFER", "FROM", "TO", "MAX", "KEEP", "ALL", "VAR",
	"SET", "TRANSACTION", "METADATA", "OF", "KEY", "META", "TY_ACCOUNT", "TY_ASSET",
	"TY_NUMBER", "TY_MONETARY", "TY_PORTION", "TY_STRING", "OP_ADD", "OP_SUB",
	"PERCENT", "PORTION_REMAINING", "LBRACK", "RBRACK", "LBRACE", "RBRACE",
	"LPAREN", "RPAREN", "DOT", "COLON", "STAR", "EQ", "STRING", "RATIO", "VARIABLE_NAME",
	"ACCOUNT", "ASSET", "NUMBER",
}

var ruleNames = []string{
	"monetary", "monetaryAll", "literal", "expression", "allotmentPortion",
	"destinationAllotment", "destination", "sourceAllotment", "sourceMaxed",
	"sourceInOrder", "source", "valueAwareSource", "type_", "origin", "varDecl",
	"metadataValue", "metadataEntry", "statement", "script",
}

type FaRlParser struct {
	*antlr.BaseParser
}

// NewFaRlParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *FaRlParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewFaRlParser(input antlr.TokenStream) *FaRlParser {
	this := new(FaRlParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "FaRl.g4"

	return this
}

// FaRlParser tokens.
const (
	FaRlParserEOF               = antlr.TokenEOF
	FaRlParserT__0              = 1
	FaRlParserT__1              = 2
	FaRlParserNEWLINE           = 3
	FaRlParserWHITESPACE        = 4
	FaRlParserMULTILINE_COMMENT = 5
	FaRlParserLINE_COMMENT      = 6
	FaRlParserPRINT             = 7
	FaRlParserFAIL              = 8
	FaRlParserTRANSFER          = 9
	FaRlParserFROM              = 10
	FaRlParserTO                = 11
	FaRlParserMAX               = 12
	FaRlParserKEEP              = 13
	FaRlParserALL               = 14
	FaRlParserVAR               = 15
	FaRlParserSET               = 16
	FaRlParserTRANSACTION       = 17
	FaRlParserMETADATA          = 18
	FaRlParserOF                = 19
	FaRlParserKEY               = 20
	FaRlParserMETA              = 21
	FaRlParserTY_ACCOUNT        = 22
	FaRlParserTY_ASSET          = 23
	FaRlParserTY_NUMBER         = 24
	FaRlParserTY_MONETARY       = 25
	FaRlParserTY_PORTION        = 26
	FaRlParserTY_STRING         = 27
	FaRlParserOP_ADD            = 28
	FaRlParserOP_SUB            = 29
	FaRlParserPERCENT           = 30
	FaRlParserPORTION_REMAINING = 31
	FaRlParserLBRACK            = 32
	FaRlParserRBRACK            = 33
	FaRlParserLBRACE            = 34
	FaRlParserRBRACE            = 35
	FaRlParserLPAREN            = 36
	FaRlParserRPAREN            = 37
	FaRlParserDOT               = 38
	FaRlParserCOLON             = 39
	FaRlParserSTAR              = 40
	FaRlParserEQ                = 41
	FaRlParserSTRING            = 42
	FaRlParserRATIO             = 43
	FaRlParserVARIABLE_NAME     = 44
	FaRlParserACCOUNT           = 45
	FaRlParserASSET             = 46
	FaRlParserNUMBER            = 47
)

// FaRlParser rules.
const (
	FaRlParserRULE_monetary             = 0
	FaRlParserRULE_monetaryAll          = 1
	FaRlParserRULE_literal              = 2
	FaRlParserRULE_expression           = 3
	FaRlParserRULE_allotmentPortion     = 4
	FaRlParserRULE_destinationAllotment = 5
	FaRlParserRULE_destination          = 6
	FaRlParserRULE_sourceAllotment      = 7
	FaRlParserRULE_sourceMaxed          = 8
	FaRlParserRULE_sourceInOrder        = 9
	FaRlParserRULE_source               = 10
	FaRlParserRULE_valueAwareSource     = 11
	FaRlParserRULE_type_                = 12
	FaRlParserRULE_origin               = 13
	FaRlParserRULE_varDecl              = 14
	FaRlParserRULE_metadataValue        = 15
	FaRlParserRULE_metadataEntry        = 16
	FaRlParserRULE_statement            = 17
	FaRlParserRULE_script               = 18
)

// IMonetaryContext is an interface to support dynamic dispatch.
type IMonetaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonetaryContext differentiates from other interfaces.
	IsMonetaryContext()
}

type MonetaryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonetaryContext() *MonetaryContext {
	var p = new(MonetaryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_monetary
	return p
}

func (*MonetaryContext) IsMonetaryContext() {}

func NewMonetaryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonetaryContext {
	var p = new(MonetaryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_monetary

	return p
}

func (s *MonetaryContext) GetParser() antlr.Parser { return s.parser }

func (s *MonetaryContext) CopyFrom(ctx *MonetaryContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MonetaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MonetaryNoPrecisionContext struct {
	*MonetaryContext
	asset  antlr.Token
	amount antlr.Token
}

func NewMonetaryNoPrecisionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MonetaryNoPrecisionContext {
	var p = new(MonetaryNoPrecisionContext)

	p.MonetaryContext = NewEmptyMonetaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MonetaryContext))

	return p
}

func (s *MonetaryNoPrecisionContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryNoPrecisionContext) GetAmount() antlr.Token { return s.amount }

func (s *MonetaryNoPrecisionContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryNoPrecisionContext) SetAmount(v antlr.Token) { s.amount = v }

func (s *MonetaryNoPrecisionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryNoPrecisionContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACK, 0)
}

func (s *MonetaryNoPrecisionContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACK, 0)
}

func (s *MonetaryNoPrecisionContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *MonetaryNoPrecisionContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *MonetaryNoPrecisionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMonetaryNoPrecision(s)
	}
}

func (s *MonetaryNoPrecisionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMonetaryNoPrecision(s)
	}
}

type MonetaryLitContext struct {
	*MonetaryContext
	asset     antlr.Token
	precision antlr.Token
	amount    antlr.Token
}

func NewMonetaryLitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MonetaryLitContext {
	var p = new(MonetaryLitContext)

	p.MonetaryContext = NewEmptyMonetaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MonetaryContext))

	return p
}

func (s *MonetaryLitContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryLitContext) GetPrecision() antlr.Token { return s.precision }

func (s *MonetaryLitContext) GetAmount() antlr.Token { return s.amount }

func (s *MonetaryLitContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryLitContext) SetPrecision(v antlr.Token) { s.precision = v }

func (s *MonetaryLitContext) SetAmount(v antlr.Token) { s.amount = v }

func (s *MonetaryLitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryLitContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACK, 0)
}

func (s *MonetaryLitContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaRlParserDOT, 0)
}

func (s *MonetaryLitContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACK, 0)
}

func (s *MonetaryLitContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *MonetaryLitContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNUMBER)
}

func (s *MonetaryLitContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, i)
}

func (s *MonetaryLitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMonetaryLit(s)
	}
}

func (s *MonetaryLitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMonetaryLit(s)
	}
}

func (p *FaRlParser) Monetary() (localctx IMonetaryContext) {
	localctx = NewMonetaryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, FaRlParserRULE_monetary)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(48)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMonetaryLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(38)
			p.Match(FaRlParserLBRACK)
		}
		{
			p.SetState(39)

			var _m = p.Match(FaRlParserASSET)

			localctx.(*MonetaryLitContext).asset = _m
		}
		{
			p.SetState(40)
			p.Match(FaRlParserDOT)
		}
		{
			p.SetState(41)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*MonetaryLitContext).precision = _m
		}
		{
			p.SetState(42)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*MonetaryLitContext).amount = _m
		}
		{
			p.SetState(43)
			p.Match(FaRlParserRBRACK)
		}

	case 2:
		localctx = NewMonetaryNoPrecisionContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(44)
			p.Match(FaRlParserLBRACK)
		}
		{
			p.SetState(45)

			var _m = p.Match(FaRlParserASSET)

			localctx.(*MonetaryNoPrecisionContext).asset = _m
		}
		{
			p.SetState(46)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*MonetaryNoPrecisionContext).amount = _m
		}
		{
			p.SetState(47)
			p.Match(FaRlParserRBRACK)
		}

	}

	return localctx
}

// IMonetaryAllContext is an interface to support dynamic dispatch.
type IMonetaryAllContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonetaryAllContext differentiates from other interfaces.
	IsMonetaryAllContext()
}

type MonetaryAllContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonetaryAllContext() *MonetaryAllContext {
	var p = new(MonetaryAllContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_monetaryAll
	return p
}

func (*MonetaryAllContext) IsMonetaryAllContext() {}

func NewMonetaryAllContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonetaryAllContext {
	var p = new(MonetaryAllContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_monetaryAll

	return p
}

func (s *MonetaryAllContext) GetParser() antlr.Parser { return s.parser }

func (s *MonetaryAllContext) CopyFrom(ctx *MonetaryAllContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MonetaryAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryAllContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MonetaryAllPrecContext struct {
	*MonetaryAllContext
	asset     antlr.Token
	precision antlr.Token
}

func NewMonetaryAllPrecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MonetaryAllPrecContext {
	var p = new(MonetaryAllPrecContext)

	p.MonetaryAllContext = NewEmptyMonetaryAllContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MonetaryAllContext))

	return p
}

func (s *MonetaryAllPrecContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryAllPrecContext) GetPrecision() antlr.Token { return s.precision }

func (s *MonetaryAllPrecContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryAllPrecContext) SetPrecision(v antlr.Token) { s.precision = v }

func (s *MonetaryAllPrecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryAllPrecContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACK, 0)
}

func (s *MonetaryAllPrecContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaRlParserDOT, 0)
}

func (s *MonetaryAllPrecContext) STAR() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTAR, 0)
}

func (s *MonetaryAllPrecContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACK, 0)
}

func (s *MonetaryAllPrecContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *MonetaryAllPrecContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *MonetaryAllPrecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMonetaryAllPrec(s)
	}
}

func (s *MonetaryAllPrecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMonetaryAllPrec(s)
	}
}

type MonetaryAllNoPrecContext struct {
	*MonetaryAllContext
	asset antlr.Token
}

func NewMonetaryAllNoPrecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MonetaryAllNoPrecContext {
	var p = new(MonetaryAllNoPrecContext)

	p.MonetaryAllContext = NewEmptyMonetaryAllContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MonetaryAllContext))

	return p
}

func (s *MonetaryAllNoPrecContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryAllNoPrecContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryAllNoPrecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryAllNoPrecContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACK, 0)
}

func (s *MonetaryAllNoPrecContext) STAR() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTAR, 0)
}

func (s *MonetaryAllNoPrecContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACK, 0)
}

func (s *MonetaryAllNoPrecContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *MonetaryAllNoPrecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMonetaryAllNoPrec(s)
	}
}

func (s *MonetaryAllNoPrecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMonetaryAllNoPrec(s)
	}
}

func (p *FaRlParser) MonetaryAll() (localctx IMonetaryAllContext) {
	localctx = NewMonetaryAllContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, FaRlParserRULE_monetaryAll)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMonetaryAllPrecContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.Match(FaRlParserLBRACK)
		}
		{
			p.SetState(51)

			var _m = p.Match(FaRlParserASSET)

			localctx.(*MonetaryAllPrecContext).asset = _m
		}
		{
			p.SetState(52)
			p.Match(FaRlParserDOT)
		}
		{
			p.SetState(53)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*MonetaryAllPrecContext).precision = _m
		}
		{
			p.SetState(54)
			p.Match(FaRlParserSTAR)
		}
		{
			p.SetState(55)
			p.Match(FaRlParserRBRACK)
		}

	case 2:
		localctx = NewMonetaryAllNoPrecContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(56)
			p.Match(FaRlParserLBRACK)
		}
		{
			p.SetState(57)

			var _m = p.Match(FaRlParserASSET)

			localctx.(*MonetaryAllNoPrecContext).asset = _m
		}
		{
			p.SetState(58)
			p.Match(FaRlParserSTAR)
		}
		{
			p.SetState(59)
			p.Match(FaRlParserRBRACK)
		}

	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) CopyFrom(ctx *LiteralContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type LitAccountContext struct {
	*LiteralContext
}

func NewLitAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitAccountContext {
	var p = new(LitAccountContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitAccountContext) ACCOUNT() antlr.TerminalNode {
	return s.GetToken(FaRlParserACCOUNT, 0)
}

func (s *LitAccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterLitAccount(s)
	}
}

func (s *LitAccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitLitAccount(s)
	}
}

type LitAssetContext struct {
	*LiteralContext
}

func NewLitAssetContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitAssetContext {
	var p = new(LitAssetContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitAssetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitAssetContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *LitAssetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterLitAsset(s)
	}
}

func (s *LitAssetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitLitAsset(s)
	}
}

type LitMonetaryContext struct {
	*LiteralContext
}

func NewLitMonetaryContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitMonetaryContext {
	var p = new(LitMonetaryContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitMonetaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitMonetaryContext) Monetary() IMonetaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryContext)
}

func (s *LitMonetaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterLitMonetary(s)
	}
}

func (s *LitMonetaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitLitMonetary(s)
	}
}

type LitNumberContext struct {
	*LiteralContext
}

func NewLitNumberContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LitNumberContext {
	var p = new(LitNumberContext)

	p.LiteralContext = NewEmptyLiteralContext()
	p.parser = parser
	p.CopyFrom(ctx.(*LiteralContext))

	return p
}

func (s *LitNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LitNumberContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *LitNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterLitNumber(s)
	}
}

func (s *LitNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitLitNumber(s)
	}
}

func (p *FaRlParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, FaRlParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(66)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(62)
			p.Match(FaRlParserACCOUNT)
		}

	case FaRlParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(FaRlParserASSET)
		}

	case FaRlParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Match(FaRlParserNUMBER)
		}

	case FaRlParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(65)
			p.Monetary()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExprAddSubContext struct {
	*ExpressionContext
	lhs IExpressionContext
	op  antlr.Token
	rhs IExpressionContext
}

func NewExprAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprAddSubContext {
	var p = new(ExprAddSubContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprAddSubContext) GetOp() antlr.Token { return s.op }

func (s *ExprAddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *ExprAddSubContext) GetLhs() IExpressionContext { return s.lhs }

func (s *ExprAddSubContext) GetRhs() IExpressionContext { return s.rhs }

func (s *ExprAddSubContext) SetLhs(v IExpressionContext) { s.lhs = v }

func (s *ExprAddSubContext) SetRhs(v IExpressionContext) { s.rhs = v }

func (s *ExprAddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprAddSubContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExprAddSubContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExprAddSubContext) OP_ADD() antlr.TerminalNode {
	return s.GetToken(FaRlParserOP_ADD, 0)
}

func (s *ExprAddSubContext) OP_SUB() antlr.TerminalNode {
	return s.GetToken(FaRlParserOP_SUB, 0)
}

func (s *ExprAddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterExprAddSub(s)
	}
}

func (s *ExprAddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitExprAddSub(s)
	}
}

type ExprLiteralContext struct {
	*ExpressionContext
	lit ILiteralContext
}

func NewExprLiteralContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprLiteralContext {
	var p = new(ExprLiteralContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprLiteralContext) GetLit() ILiteralContext { return s.lit }

func (s *ExprLiteralContext) SetLit(v ILiteralContext) { s.lit = v }

func (s *ExprLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprLiteralContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *ExprLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterExprLiteral(s)
	}
}

func (s *ExprLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitExprLiteral(s)
	}
}

type ExprVariableContext struct {
	*ExpressionContext
	variable antlr.Token
}

func NewExprVariableContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprVariableContext {
	var p = new(ExprVariableContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ExprVariableContext) GetVariable() antlr.Token { return s.variable }

func (s *ExprVariableContext) SetVariable(v antlr.Token) { s.variable = v }

func (s *ExprVariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprVariableContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(FaRlParserVARIABLE_NAME, 0)
}

func (s *ExprVariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterExprVariable(s)
	}
}

func (s *ExprVariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitExprVariable(s)
	}
}

func (p *FaRlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *FaRlParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, FaRlParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserLBRACK, FaRlParserACCOUNT, FaRlParserASSET, FaRlParserNUMBER:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(69)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case FaRlParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(70)

			var _m = p.Match(FaRlParserVARIABLE_NAME)

			localctx.(*ExprVariableContext).variable = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*ExprAddSubContext).lhs = _prevctx

			p.PushNewRecursionContext(localctx, _startState, FaRlParserRULE_expression)
			p.SetState(73)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(74)

				var _lt = p.GetTokenStream().LT(1)

				localctx.(*ExprAddSubContext).op = _lt

				_la = p.GetTokenStream().LA(1)

				if !(_la == FaRlParserOP_ADD || _la == FaRlParserOP_SUB) {
					var _ri = p.GetErrorHandler().RecoverInline(p)

					localctx.(*ExprAddSubContext).op = _ri
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(75)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext())
	}

	return localctx
}

// IAllotmentPortionContext is an interface to support dynamic dispatch.
type IAllotmentPortionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAllotmentPortionContext differentiates from other interfaces.
	IsAllotmentPortionContext()
}

type AllotmentPortionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAllotmentPortionContext() *AllotmentPortionContext {
	var p = new(AllotmentPortionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_allotmentPortion
	return p
}

func (*AllotmentPortionContext) IsAllotmentPortionContext() {}

func NewAllotmentPortionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AllotmentPortionContext {
	var p = new(AllotmentPortionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_allotmentPortion

	return p
}

func (s *AllotmentPortionContext) GetParser() antlr.Parser { return s.parser }

func (s *AllotmentPortionContext) CopyFrom(ctx *AllotmentPortionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AllotmentPortionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type AllotmentPortionRemainingContext struct {
	*AllotmentPortionContext
}

func NewAllotmentPortionRemainingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllotmentPortionRemainingContext {
	var p = new(AllotmentPortionRemainingContext)

	p.AllotmentPortionContext = NewEmptyAllotmentPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllotmentPortionContext))

	return p
}

func (s *AllotmentPortionRemainingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionRemainingContext) PORTION_REMAINING() antlr.TerminalNode {
	return s.GetToken(FaRlParserPORTION_REMAINING, 0)
}

func (s *AllotmentPortionRemainingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterAllotmentPortionRemaining(s)
	}
}

func (s *AllotmentPortionRemainingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitAllotmentPortionRemaining(s)
	}
}

type AllotmentPortionConstContext struct {
	*AllotmentPortionContext
}

func NewAllotmentPortionConstContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllotmentPortionConstContext {
	var p = new(AllotmentPortionConstContext)

	p.AllotmentPortionContext = NewEmptyAllotmentPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllotmentPortionContext))

	return p
}

func (s *AllotmentPortionConstContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionConstContext) RATIO() antlr.TerminalNode {
	return s.GetToken(FaRlParserRATIO, 0)
}

func (s *AllotmentPortionConstContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterAllotmentPortionConst(s)
	}
}

func (s *AllotmentPortionConstContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitAllotmentPortionConst(s)
	}
}

type AllotmentPortionConstPercentContext struct {
	*AllotmentPortionContext
}

func NewAllotmentPortionConstPercentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllotmentPortionConstPercentContext {
	var p = new(AllotmentPortionConstPercentContext)

	p.AllotmentPortionContext = NewEmptyAllotmentPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllotmentPortionContext))

	return p
}

func (s *AllotmentPortionConstPercentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionConstPercentContext) AllNUMBER() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNUMBER)
}

func (s *AllotmentPortionConstPercentContext) NUMBER(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, i)
}

func (s *AllotmentPortionConstPercentContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(FaRlParserPERCENT, 0)
}

func (s *AllotmentPortionConstPercentContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaRlParserDOT, 0)
}

func (s *AllotmentPortionConstPercentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterAllotmentPortionConstPercent(s)
	}
}

func (s *AllotmentPortionConstPercentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitAllotmentPortionConstPercent(s)
	}
}

type AllotmentPortionVarContext struct {
	*AllotmentPortionContext
	por antlr.Token
}

func NewAllotmentPortionVarContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AllotmentPortionVarContext {
	var p = new(AllotmentPortionVarContext)

	p.AllotmentPortionContext = NewEmptyAllotmentPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AllotmentPortionContext))

	return p
}

func (s *AllotmentPortionVarContext) GetPor() antlr.Token { return s.por }

func (s *AllotmentPortionVarContext) SetPor(v antlr.Token) { s.por = v }

func (s *AllotmentPortionVarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AllotmentPortionVarContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(FaRlParserVARIABLE_NAME, 0)
}

func (s *AllotmentPortionVarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterAllotmentPortionVar(s)
	}
}

func (s *AllotmentPortionVarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitAllotmentPortionVar(s)
	}
}

func (p *FaRlParser) AllotmentPortion() (localctx IAllotmentPortionContext) {
	localctx = NewAllotmentPortionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FaRlParserRULE_allotmentPortion)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(90)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserRATIO:
		localctx = NewAllotmentPortionConstContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(81)
			p.Match(FaRlParserRATIO)
		}

	case FaRlParserNUMBER:
		localctx = NewAllotmentPortionConstPercentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(82)
			p.Match(FaRlParserNUMBER)
		}
		p.SetState(85)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == FaRlParserDOT {
			{
				p.SetState(83)
				p.Match(FaRlParserDOT)
			}
			{
				p.SetState(84)
				p.Match(FaRlParserNUMBER)
			}

		}
		{
			p.SetState(87)
			p.Match(FaRlParserPERCENT)
		}

	case FaRlParserVARIABLE_NAME:
		localctx = NewAllotmentPortionVarContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)

			var _m = p.Match(FaRlParserVARIABLE_NAME)

			localctx.(*AllotmentPortionVarContext).por = _m
		}

	case FaRlParserPORTION_REMAINING:
		localctx = NewAllotmentPortionRemainingContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(FaRlParserPORTION_REMAINING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDestinationAllotmentContext is an interface to support dynamic dispatch.
type IDestinationAllotmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_allotmentPortion returns the _allotmentPortion rule contexts.
	Get_allotmentPortion() IAllotmentPortionContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_allotmentPortion sets the _allotmentPortion rule contexts.
	Set_allotmentPortion(IAllotmentPortionContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IAllotmentPortionContext

	// GetDests returns the dests rule context list.
	GetDests() []IExpressionContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllotmentPortionContext)

	// SetDests sets the dests rule context list.
	SetDests([]IExpressionContext)

	// IsDestinationAllotmentContext differentiates from other interfaces.
	IsDestinationAllotmentContext()
}

type DestinationAllotmentContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_allotmentPortion IAllotmentPortionContext
	portions          []IAllotmentPortionContext
	_expression       IExpressionContext
	dests             []IExpressionContext
}

func NewEmptyDestinationAllotmentContext() *DestinationAllotmentContext {
	var p = new(DestinationAllotmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_destinationAllotment
	return p
}

func (*DestinationAllotmentContext) IsDestinationAllotmentContext() {}

func NewDestinationAllotmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationAllotmentContext {
	var p = new(DestinationAllotmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_destinationAllotment

	return p
}

func (s *DestinationAllotmentContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationAllotmentContext) Get_allotmentPortion() IAllotmentPortionContext {
	return s._allotmentPortion
}

func (s *DestinationAllotmentContext) Get_expression() IExpressionContext { return s._expression }

func (s *DestinationAllotmentContext) Set_allotmentPortion(v IAllotmentPortionContext) {
	s._allotmentPortion = v
}

func (s *DestinationAllotmentContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *DestinationAllotmentContext) GetPortions() []IAllotmentPortionContext { return s.portions }

func (s *DestinationAllotmentContext) GetDests() []IExpressionContext { return s.dests }

func (s *DestinationAllotmentContext) SetPortions(v []IAllotmentPortionContext) { s.portions = v }

func (s *DestinationAllotmentContext) SetDests(v []IExpressionContext) { s.dests = v }

func (s *DestinationAllotmentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACE, 0)
}

func (s *DestinationAllotmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *DestinationAllotmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *DestinationAllotmentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACE, 0)
}

func (s *DestinationAllotmentContext) AllTO() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserTO)
}

func (s *DestinationAllotmentContext) TO(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserTO, i)
}

func (s *DestinationAllotmentContext) AllAllotmentPortion() []IAllotmentPortionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem())
	var tst = make([]IAllotmentPortionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllotmentPortionContext)
		}
	}

	return tst
}

func (s *DestinationAllotmentContext) AllotmentPortion(i int) IAllotmentPortionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllotmentPortionContext)
}

func (s *DestinationAllotmentContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *DestinationAllotmentContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DestinationAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationAllotmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DestinationAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterDestinationAllotment(s)
	}
}

func (s *DestinationAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitDestinationAllotment(s)
	}
}

func (p *FaRlParser) DestinationAllotment() (localctx IDestinationAllotmentContext) {
	localctx = NewDestinationAllotmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FaRlParserRULE_destinationAllotment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(92)
		p.Match(FaRlParserLBRACE)
	}
	{
		p.SetState(93)
		p.Match(FaRlParserNEWLINE)
	}
	p.SetState(99)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(FaRlParserPORTION_REMAINING-31))|(1<<(FaRlParserRATIO-31))|(1<<(FaRlParserVARIABLE_NAME-31))|(1<<(FaRlParserNUMBER-31)))) != 0) {
		{
			p.SetState(94)

			var _x = p.AllotmentPortion()

			localctx.(*DestinationAllotmentContext)._allotmentPortion = _x
		}
		localctx.(*DestinationAllotmentContext).portions = append(localctx.(*DestinationAllotmentContext).portions, localctx.(*DestinationAllotmentContext)._allotmentPortion)
		{
			p.SetState(95)
			p.Match(FaRlParserTO)
		}
		{
			p.SetState(96)

			var _x = p.expression(0)

			localctx.(*DestinationAllotmentContext)._expression = _x
		}
		localctx.(*DestinationAllotmentContext).dests = append(localctx.(*DestinationAllotmentContext).dests, localctx.(*DestinationAllotmentContext)._expression)
		{
			p.SetState(97)
			p.Match(FaRlParserNEWLINE)
		}

		p.SetState(101)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(103)
		p.Match(FaRlParserRBRACE)
	}

	return localctx
}

// IDestinationContext is an interface to support dynamic dispatch.
type IDestinationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDestinationContext differentiates from other interfaces.
	IsDestinationContext()
}

type DestinationContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDestinationContext() *DestinationContext {
	var p = new(DestinationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_destination
	return p
}

func (*DestinationContext) IsDestinationContext() {}

func NewDestinationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DestinationContext {
	var p = new(DestinationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_destination

	return p
}

func (s *DestinationContext) GetParser() antlr.Parser { return s.parser }

func (s *DestinationContext) CopyFrom(ctx *DestinationContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *DestinationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestinationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type DestAccountContext struct {
	*DestinationContext
}

func NewDestAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestAccountContext {
	var p = new(DestAccountContext)

	p.DestinationContext = NewEmptyDestinationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DestinationContext))

	return p
}

func (s *DestAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestAccountContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DestAccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterDestAccount(s)
	}
}

func (s *DestAccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitDestAccount(s)
	}
}

type DestAllotmentContext struct {
	*DestinationContext
}

func NewDestAllotmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DestAllotmentContext {
	var p = new(DestAllotmentContext)

	p.DestinationContext = NewEmptyDestinationContext()
	p.parser = parser
	p.CopyFrom(ctx.(*DestinationContext))

	return p
}

func (s *DestAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DestAllotmentContext) DestinationAllotment() IDestinationAllotmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationAllotmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationAllotmentContext)
}

func (s *DestAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterDestAllotment(s)
	}
}

func (s *DestAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitDestAllotment(s)
	}
}

func (p *FaRlParser) Destination() (localctx IDestinationContext) {
	localctx = NewDestinationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FaRlParserRULE_destination)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(107)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserLBRACK, FaRlParserVARIABLE_NAME, FaRlParserACCOUNT, FaRlParserASSET, FaRlParserNUMBER:
		localctx = NewDestAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(105)
			p.expression(0)
		}

	case FaRlParserLBRACE:
		localctx = NewDestAllotmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(106)
			p.DestinationAllotment()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ISourceAllotmentContext is an interface to support dynamic dispatch.
type ISourceAllotmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_allotmentPortion returns the _allotmentPortion rule contexts.
	Get_allotmentPortion() IAllotmentPortionContext

	// Get_source returns the _source rule contexts.
	Get_source() ISourceContext

	// Set_allotmentPortion sets the _allotmentPortion rule contexts.
	Set_allotmentPortion(IAllotmentPortionContext)

	// Set_source sets the _source rule contexts.
	Set_source(ISourceContext)

	// GetPortions returns the portions rule context list.
	GetPortions() []IAllotmentPortionContext

	// GetSources returns the sources rule context list.
	GetSources() []ISourceContext

	// SetPortions sets the portions rule context list.
	SetPortions([]IAllotmentPortionContext)

	// SetSources sets the sources rule context list.
	SetSources([]ISourceContext)

	// IsSourceAllotmentContext differentiates from other interfaces.
	IsSourceAllotmentContext()
}

type SourceAllotmentContext struct {
	*antlr.BaseParserRuleContext
	parser            antlr.Parser
	_allotmentPortion IAllotmentPortionContext
	portions          []IAllotmentPortionContext
	_source           ISourceContext
	sources           []ISourceContext
}

func NewEmptySourceAllotmentContext() *SourceAllotmentContext {
	var p = new(SourceAllotmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_sourceAllotment
	return p
}

func (*SourceAllotmentContext) IsSourceAllotmentContext() {}

func NewSourceAllotmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceAllotmentContext {
	var p = new(SourceAllotmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_sourceAllotment

	return p
}

func (s *SourceAllotmentContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceAllotmentContext) Get_allotmentPortion() IAllotmentPortionContext {
	return s._allotmentPortion
}

func (s *SourceAllotmentContext) Get_source() ISourceContext { return s._source }

func (s *SourceAllotmentContext) Set_allotmentPortion(v IAllotmentPortionContext) {
	s._allotmentPortion = v
}

func (s *SourceAllotmentContext) Set_source(v ISourceContext) { s._source = v }

func (s *SourceAllotmentContext) GetPortions() []IAllotmentPortionContext { return s.portions }

func (s *SourceAllotmentContext) GetSources() []ISourceContext { return s.sources }

func (s *SourceAllotmentContext) SetPortions(v []IAllotmentPortionContext) { s.portions = v }

func (s *SourceAllotmentContext) SetSources(v []ISourceContext) { s.sources = v }

func (s *SourceAllotmentContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACE, 0)
}

func (s *SourceAllotmentContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *SourceAllotmentContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *SourceAllotmentContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACE, 0)
}

func (s *SourceAllotmentContext) AllFROM() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserFROM)
}

func (s *SourceAllotmentContext) FROM(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, i)
}

func (s *SourceAllotmentContext) AllAllotmentPortion() []IAllotmentPortionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem())
	var tst = make([]IAllotmentPortionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAllotmentPortionContext)
		}
	}

	return tst
}

func (s *SourceAllotmentContext) AllotmentPortion(i int) IAllotmentPortionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAllotmentPortionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAllotmentPortionContext)
}

func (s *SourceAllotmentContext) AllSource() []ISourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceContext)(nil)).Elem())
	var tst = make([]ISourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceContext)
		}
	}

	return tst
}

func (s *SourceAllotmentContext) Source(i int) ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceAllotmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSourceAllotment(s)
	}
}

func (s *SourceAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSourceAllotment(s)
	}
}

func (p *FaRlParser) SourceAllotment() (localctx ISourceAllotmentContext) {
	localctx = NewSourceAllotmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, FaRlParserRULE_sourceAllotment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(109)
		p.Match(FaRlParserLBRACE)
	}
	{
		p.SetState(110)
		p.Match(FaRlParserNEWLINE)
	}
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-31)&-(0x1f+1)) == 0 && ((1<<uint((_la-31)))&((1<<(FaRlParserPORTION_REMAINING-31))|(1<<(FaRlParserRATIO-31))|(1<<(FaRlParserVARIABLE_NAME-31))|(1<<(FaRlParserNUMBER-31)))) != 0) {
		{
			p.SetState(111)

			var _x = p.AllotmentPortion()

			localctx.(*SourceAllotmentContext)._allotmentPortion = _x
		}
		localctx.(*SourceAllotmentContext).portions = append(localctx.(*SourceAllotmentContext).portions, localctx.(*SourceAllotmentContext)._allotmentPortion)
		{
			p.SetState(112)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(113)

			var _x = p.Source()

			localctx.(*SourceAllotmentContext)._source = _x
		}
		localctx.(*SourceAllotmentContext).sources = append(localctx.(*SourceAllotmentContext).sources, localctx.(*SourceAllotmentContext)._source)
		{
			p.SetState(114)
			p.Match(FaRlParserNEWLINE)
		}

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(120)
		p.Match(FaRlParserRBRACE)
	}

	return localctx
}

// ISourceMaxedContext is an interface to support dynamic dispatch.
type ISourceMaxedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMax returns the max rule contexts.
	GetMax() IExpressionContext

	// GetSrc returns the src rule contexts.
	GetSrc() ISourceContext

	// SetMax sets the max rule contexts.
	SetMax(IExpressionContext)

	// SetSrc sets the src rule contexts.
	SetSrc(ISourceContext)

	// IsSourceMaxedContext differentiates from other interfaces.
	IsSourceMaxedContext()
}

type SourceMaxedContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	max    IExpressionContext
	src    ISourceContext
}

func NewEmptySourceMaxedContext() *SourceMaxedContext {
	var p = new(SourceMaxedContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_sourceMaxed
	return p
}

func (*SourceMaxedContext) IsSourceMaxedContext() {}

func NewSourceMaxedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceMaxedContext {
	var p = new(SourceMaxedContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_sourceMaxed

	return p
}

func (s *SourceMaxedContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceMaxedContext) GetMax() IExpressionContext { return s.max }

func (s *SourceMaxedContext) GetSrc() ISourceContext { return s.src }

func (s *SourceMaxedContext) SetMax(v IExpressionContext) { s.max = v }

func (s *SourceMaxedContext) SetSrc(v ISourceContext) { s.src = v }

func (s *SourceMaxedContext) MAX() antlr.TerminalNode {
	return s.GetToken(FaRlParserMAX, 0)
}

func (s *SourceMaxedContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SourceMaxedContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SourceMaxedContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceMaxedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceMaxedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceMaxedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSourceMaxed(s)
	}
}

func (s *SourceMaxedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSourceMaxed(s)
	}
}

func (p *FaRlParser) SourceMaxed() (localctx ISourceMaxedContext) {
	localctx = NewSourceMaxedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, FaRlParserRULE_sourceMaxed)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(122)
		p.Match(FaRlParserMAX)
	}
	{
		p.SetState(123)

		var _x = p.expression(0)

		localctx.(*SourceMaxedContext).max = _x
	}
	{
		p.SetState(124)
		p.Match(FaRlParserFROM)
	}
	{
		p.SetState(125)

		var _x = p.Source()

		localctx.(*SourceMaxedContext).src = _x
	}

	return localctx
}

// ISourceInOrderContext is an interface to support dynamic dispatch.
type ISourceInOrderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_source returns the _source rule contexts.
	Get_source() ISourceContext

	// Set_source sets the _source rule contexts.
	Set_source(ISourceContext)

	// GetSources returns the sources rule context list.
	GetSources() []ISourceContext

	// SetSources sets the sources rule context list.
	SetSources([]ISourceContext)

	// IsSourceInOrderContext differentiates from other interfaces.
	IsSourceInOrderContext()
}

type SourceInOrderContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	_source ISourceContext
	sources []ISourceContext
}

func NewEmptySourceInOrderContext() *SourceInOrderContext {
	var p = new(SourceInOrderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_sourceInOrder
	return p
}

func (*SourceInOrderContext) IsSourceInOrderContext() {}

func NewSourceInOrderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceInOrderContext {
	var p = new(SourceInOrderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_sourceInOrder

	return p
}

func (s *SourceInOrderContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceInOrderContext) Get_source() ISourceContext { return s._source }

func (s *SourceInOrderContext) Set_source(v ISourceContext) { s._source = v }

func (s *SourceInOrderContext) GetSources() []ISourceContext { return s.sources }

func (s *SourceInOrderContext) SetSources(v []ISourceContext) { s.sources = v }

func (s *SourceInOrderContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACE, 0)
}

func (s *SourceInOrderContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *SourceInOrderContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *SourceInOrderContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACE, 0)
}

func (s *SourceInOrderContext) AllSource() []ISourceContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISourceContext)(nil)).Elem())
	var tst = make([]ISourceContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISourceContext)
		}
	}

	return tst
}

func (s *SourceInOrderContext) Source(i int) ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SourceInOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceInOrderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceInOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSourceInOrder(s)
	}
}

func (s *SourceInOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSourceInOrder(s)
	}
}

func (p *FaRlParser) SourceInOrder() (localctx ISourceInOrderContext) {
	localctx = NewSourceInOrderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FaRlParserRULE_sourceInOrder)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(127)
		p.Match(FaRlParserLBRACE)
	}
	{
		p.SetState(128)
		p.Match(FaRlParserNEWLINE)
	}
	p.SetState(132)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == FaRlParserMAX || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(FaRlParserLBRACK-32))|(1<<(FaRlParserLBRACE-32))|(1<<(FaRlParserVARIABLE_NAME-32))|(1<<(FaRlParserACCOUNT-32))|(1<<(FaRlParserASSET-32))|(1<<(FaRlParserNUMBER-32)))) != 0) {
		{
			p.SetState(129)

			var _x = p.Source()

			localctx.(*SourceInOrderContext)._source = _x
		}
		localctx.(*SourceInOrderContext).sources = append(localctx.(*SourceInOrderContext).sources, localctx.(*SourceInOrderContext)._source)
		{
			p.SetState(130)
			p.Match(FaRlParserNEWLINE)
		}

		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(136)
		p.Match(FaRlParserRBRACE)
	}

	return localctx
}

// ISourceContext is an interface to support dynamic dispatch.
type ISourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSourceContext differentiates from other interfaces.
	IsSourceContext()
}

type SourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceContext() *SourceContext {
	var p = new(SourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_source
	return p
}

func (*SourceContext) IsSourceContext() {}

func NewSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceContext {
	var p = new(SourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_source

	return p
}

func (s *SourceContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceContext) CopyFrom(ctx *SourceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SrcAccountContext struct {
	*SourceContext
}

func NewSrcAccountContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcAccountContext {
	var p = new(SrcAccountContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcAccountContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcAccountContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcAccountContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcAccount(s)
	}
}

func (s *SrcAccountContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcAccount(s)
	}
}

type SrcMaxedContext struct {
	*SourceContext
}

func NewSrcMaxedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcMaxedContext {
	var p = new(SrcMaxedContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcMaxedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcMaxedContext) SourceMaxed() ISourceMaxedContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceMaxedContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceMaxedContext)
}

func (s *SrcMaxedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcMaxed(s)
	}
}

func (s *SrcMaxedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcMaxed(s)
	}
}

type SrcInOrderContext struct {
	*SourceContext
}

func NewSrcInOrderContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcInOrderContext {
	var p = new(SrcInOrderContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcInOrderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcInOrderContext) SourceInOrder() ISourceInOrderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceInOrderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceInOrderContext)
}

func (s *SrcInOrderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcInOrder(s)
	}
}

func (s *SrcInOrderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcInOrder(s)
	}
}

func (p *FaRlParser) Source() (localctx ISourceContext) {
	localctx = NewSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, FaRlParserRULE_source)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(141)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserLBRACK, FaRlParserVARIABLE_NAME, FaRlParserACCOUNT, FaRlParserASSET, FaRlParserNUMBER:
		localctx = NewSrcAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(138)
			p.expression(0)
		}

	case FaRlParserMAX:
		localctx = NewSrcMaxedContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(139)
			p.SourceMaxed()
		}

	case FaRlParserLBRACE:
		localctx = NewSrcInOrderContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(140)
			p.SourceInOrder()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IValueAwareSourceContext is an interface to support dynamic dispatch.
type IValueAwareSourceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueAwareSourceContext differentiates from other interfaces.
	IsValueAwareSourceContext()
}

type ValueAwareSourceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueAwareSourceContext() *ValueAwareSourceContext {
	var p = new(ValueAwareSourceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_valueAwareSource
	return p
}

func (*ValueAwareSourceContext) IsValueAwareSourceContext() {}

func NewValueAwareSourceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueAwareSourceContext {
	var p = new(ValueAwareSourceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_valueAwareSource

	return p
}

func (s *ValueAwareSourceContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueAwareSourceContext) CopyFrom(ctx *ValueAwareSourceContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ValueAwareSourceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueAwareSourceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SrcContext struct {
	*ValueAwareSourceContext
}

func NewSrcContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcContext {
	var p = new(SrcContext)

	p.ValueAwareSourceContext = NewEmptyValueAwareSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueAwareSourceContext))

	return p
}

func (s *SrcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SrcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrc(s)
	}
}

func (s *SrcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrc(s)
	}
}

type SrcAllotmentContext struct {
	*ValueAwareSourceContext
}

func NewSrcAllotmentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcAllotmentContext {
	var p = new(SrcAllotmentContext)

	p.ValueAwareSourceContext = NewEmptyValueAwareSourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ValueAwareSourceContext))

	return p
}

func (s *SrcAllotmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcAllotmentContext) SourceAllotment() ISourceAllotmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceAllotmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceAllotmentContext)
}

func (s *SrcAllotmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcAllotment(s)
	}
}

func (s *SrcAllotmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcAllotment(s)
	}
}

func (p *FaRlParser) ValueAwareSource() (localctx IValueAwareSourceContext) {
	localctx = NewValueAwareSourceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, FaRlParserRULE_valueAwareSource)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(145)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSrcContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(143)
			p.Source()
		}

	case 2:
		localctx = NewSrcAllotmentContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(144)
			p.SourceAllotment()
		}

	}

	return localctx
}

// IType_Context is an interface to support dynamic dispatch.
type IType_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsType_Context differentiates from other interfaces.
	IsType_Context()
}

type Type_Context struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyType_Context() *Type_Context {
	var p = new(Type_Context)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_type_
	return p
}

func (*Type_Context) IsType_Context() {}

func NewType_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Type_Context {
	var p = new(Type_Context)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_type_

	return p
}

func (s *Type_Context) GetParser() antlr.Parser { return s.parser }

func (s *Type_Context) TY_ACCOUNT() antlr.TerminalNode {
	return s.GetToken(FaRlParserTY_ACCOUNT, 0)
}

func (s *Type_Context) TY_ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserTY_ASSET, 0)
}

func (s *Type_Context) TY_NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserTY_NUMBER, 0)
}

func (s *Type_Context) TY_MONETARY() antlr.TerminalNode {
	return s.GetToken(FaRlParserTY_MONETARY, 0)
}

func (s *Type_Context) TY_PORTION() antlr.TerminalNode {
	return s.GetToken(FaRlParserTY_PORTION, 0)
}

func (s *Type_Context) TY_STRING() antlr.TerminalNode {
	return s.GetToken(FaRlParserTY_STRING, 0)
}

func (s *Type_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Type_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Type_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterType_(s)
	}
}

func (s *Type_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitType_(s)
	}
}

func (p *FaRlParser) Type_() (localctx IType_Context) {
	localctx = NewType_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, FaRlParserRULE_type_)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(147)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaRlParserTY_ACCOUNT)|(1<<FaRlParserTY_ASSET)|(1<<FaRlParserTY_NUMBER)|(1<<FaRlParserTY_MONETARY)|(1<<FaRlParserTY_PORTION)|(1<<FaRlParserTY_STRING))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IOriginContext is an interface to support dynamic dispatch.
type IOriginContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetKey returns the key token.
	GetKey() antlr.Token

	// SetKey sets the key token.
	SetKey(antlr.Token)

	// GetAcc returns the acc rule contexts.
	GetAcc() IExpressionContext

	// SetAcc sets the acc rule contexts.
	SetAcc(IExpressionContext)

	// IsOriginContext differentiates from other interfaces.
	IsOriginContext()
}

type OriginContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	acc    IExpressionContext
	key    antlr.Token
}

func NewEmptyOriginContext() *OriginContext {
	var p = new(OriginContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_origin
	return p
}

func (*OriginContext) IsOriginContext() {}

func NewOriginContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OriginContext {
	var p = new(OriginContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_origin

	return p
}

func (s *OriginContext) GetParser() antlr.Parser { return s.parser }

func (s *OriginContext) GetKey() antlr.Token { return s.key }

func (s *OriginContext) SetKey(v antlr.Token) { s.key = v }

func (s *OriginContext) GetAcc() IExpressionContext { return s.acc }

func (s *OriginContext) SetAcc(v IExpressionContext) { s.acc = v }

func (s *OriginContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FaRlParserRPAREN, 0)
}

func (s *OriginContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OriginContext) STRING() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTRING, 0)
}

func (s *OriginContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OriginContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OriginContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterOrigin(s)
	}
}

func (s *OriginContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitOrigin(s)
	}
}

func (p *FaRlParser) Origin() (localctx IOriginContext) {
	localctx = NewOriginContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, FaRlParserRULE_origin)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(149)
		p.Match(FaRlParserT__0)
	}
	{
		p.SetState(150)

		var _x = p.expression(0)

		localctx.(*OriginContext).acc = _x
	}
	{
		p.SetState(151)
		p.Match(FaRlParserT__1)
	}
	{
		p.SetState(152)

		var _m = p.Match(FaRlParserSTRING)

		localctx.(*OriginContext).key = _m
	}
	{
		p.SetState(153)
		p.Match(FaRlParserRPAREN)
	}

	return localctx
}

// IVarDeclContext is an interface to support dynamic dispatch.
type IVarDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarDeclContext() *VarDeclContext {
	var p = new(VarDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_varDecl
	return p
}

func (*VarDeclContext) IsVarDeclContext() {}

func NewVarDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarDeclContext {
	var p = new(VarDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_varDecl

	return p
}

func (s *VarDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarDeclContext) CopyFrom(ctx *VarDeclContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type VarTypedContext struct {
	*VarDeclContext
	name antlr.Token
	ty   IType_Context
	orig IOriginContext
}

func NewVarTypedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarTypedContext {
	var p = new(VarTypedContext)

	p.VarDeclContext = NewEmptyVarDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VarDeclContext))

	return p
}

func (s *VarTypedContext) GetName() antlr.Token { return s.name }

func (s *VarTypedContext) SetName(v antlr.Token) { s.name = v }

func (s *VarTypedContext) GetTy() IType_Context { return s.ty }

func (s *VarTypedContext) GetOrig() IOriginContext { return s.orig }

func (s *VarTypedContext) SetTy(v IType_Context) { s.ty = v }

func (s *VarTypedContext) SetOrig(v IOriginContext) { s.orig = v }

func (s *VarTypedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarTypedContext) VAR() antlr.TerminalNode {
	return s.GetToken(FaRlParserVAR, 0)
}

func (s *VarTypedContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaRlParserCOLON, 0)
}

func (s *VarTypedContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(FaRlParserVARIABLE_NAME, 0)
}

func (s *VarTypedContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarTypedContext) EQ() antlr.TerminalNode {
	return s.GetToken(FaRlParserEQ, 0)
}

func (s *VarTypedContext) Origin() IOriginContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOriginContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOriginContext)
}

func (s *VarTypedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterVarTyped(s)
	}
}

func (s *VarTypedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitVarTyped(s)
	}
}

func (p *FaRlParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, FaRlParserRULE_varDecl)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewVarTypedContext(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(155)
		p.Match(FaRlParserVAR)
	}
	{
		p.SetState(156)

		var _m = p.Match(FaRlParserVARIABLE_NAME)

		localctx.(*VarTypedContext).name = _m
	}
	{
		p.SetState(157)
		p.Match(FaRlParserCOLON)
	}
	{
		p.SetState(158)

		var _x = p.Type_()

		localctx.(*VarTypedContext).ty = _x
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaRlParserEQ {
		{
			p.SetState(159)
			p.Match(FaRlParserEQ)
		}
		{
			p.SetState(160)

			var _x = p.Origin()

			localctx.(*VarTypedContext).orig = _x
		}

	}

	return localctx
}

// IMetadataValueContext is an interface to support dynamic dispatch.
type IMetadataValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetadataValueContext differentiates from other interfaces.
	IsMetadataValueContext()
}

type MetadataValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetadataValueContext() *MetadataValueContext {
	var p = new(MetadataValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_metadataValue
	return p
}

func (*MetadataValueContext) IsMetadataValueContext() {}

func NewMetadataValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetadataValueContext {
	var p = new(MetadataValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_metadataValue

	return p
}

func (s *MetadataValueContext) GetParser() antlr.Parser { return s.parser }

func (s *MetadataValueContext) CopyFrom(ctx *MetadataValueContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *MetadataValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetadataValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MetaValueRatioContext struct {
	*MetadataValueContext
}

func NewMetaValueRatioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MetaValueRatioContext {
	var p = new(MetaValueRatioContext)

	p.MetadataValueContext = NewEmptyMetadataValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MetadataValueContext))

	return p
}

func (s *MetaValueRatioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaValueRatioContext) RATIO() antlr.TerminalNode {
	return s.GetToken(FaRlParserRATIO, 0)
}

func (s *MetaValueRatioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMetaValueRatio(s)
	}
}

func (s *MetaValueRatioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMetaValueRatio(s)
	}
}

type MetaValueExprContext struct {
	*MetadataValueContext
}

func NewMetaValueExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MetaValueExprContext {
	var p = new(MetaValueExprContext)

	p.MetadataValueContext = NewEmptyMetadataValueContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MetadataValueContext))

	return p
}

func (s *MetaValueExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetaValueExprContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MetaValueExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMetaValueExpr(s)
	}
}

func (s *MetaValueExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMetaValueExpr(s)
	}
}

func (p *FaRlParser) MetadataValue() (localctx IMetadataValueContext) {
	localctx = NewMetadataValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, FaRlParserRULE_metadataValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserLBRACK, FaRlParserVARIABLE_NAME, FaRlParserACCOUNT, FaRlParserASSET, FaRlParserNUMBER:
		localctx = NewMetaValueExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(163)
			p.expression(0)
		}

	case FaRlParserRATIO:
		localctx = NewMetaValueRatioContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(164)
			p.Match(FaRlParserRATIO)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMetadataEntryContext is an interface to support dynamic dispatch.
type IMetadataEntryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMetadataEntryContext differentiates from other interfaces.
	IsMetadataEntryContext()
}

type MetadataEntryContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMetadataEntryContext() *MetadataEntryContext {
	var p = new(MetadataEntryContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_metadataEntry
	return p
}

func (*MetadataEntryContext) IsMetadataEntryContext() {}

func NewMetadataEntryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MetadataEntryContext {
	var p = new(MetadataEntryContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_metadataEntry

	return p
}

func (s *MetadataEntryContext) GetParser() antlr.Parser { return s.parser }

func (s *MetadataEntryContext) STRING() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTRING, 0)
}

func (s *MetadataEntryContext) EQ() antlr.TerminalNode {
	return s.GetToken(FaRlParserEQ, 0)
}

func (s *MetadataEntryContext) MetadataValue() IMetadataValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataValueContext)
}

func (s *MetadataEntryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MetadataEntryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MetadataEntryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMetadataEntry(s)
	}
}

func (s *MetadataEntryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMetadataEntry(s)
	}
}

func (p *FaRlParser) MetadataEntry() (localctx IMetadataEntryContext) {
	localctx = NewMetadataEntryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, FaRlParserRULE_metadataEntry)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(167)
		p.Match(FaRlParserSTRING)
	}
	{
		p.SetState(168)
		p.Match(FaRlParserEQ)
	}
	{
		p.SetState(169)
		p.MetadataValue()
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PrintContext struct {
	*StatementContext
	expr IExpressionContext
}

func NewPrintContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrintContext {
	var p = new(PrintContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *PrintContext) GetExpr() IExpressionContext { return s.expr }

func (s *PrintContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *PrintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrintContext) PRINT() antlr.TerminalNode {
	return s.GetToken(FaRlParserPRINT, 0)
}

func (s *PrintContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterPrint(s)
	}
}

func (s *PrintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitPrint(s)
	}
}

type TransferContext struct {
	*StatementContext
	mon    IExpressionContext
	monAll IMonetaryAllContext
	src    IValueAwareSourceContext
	dest   IDestinationContext
}

func NewTransferContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransferContext {
	var p = new(TransferContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TransferContext) GetMon() IExpressionContext { return s.mon }

func (s *TransferContext) GetMonAll() IMonetaryAllContext { return s.monAll }

func (s *TransferContext) GetSrc() IValueAwareSourceContext { return s.src }

func (s *TransferContext) GetDest() IDestinationContext { return s.dest }

func (s *TransferContext) SetMon(v IExpressionContext) { s.mon = v }

func (s *TransferContext) SetMonAll(v IMonetaryAllContext) { s.monAll = v }

func (s *TransferContext) SetSrc(v IValueAwareSourceContext) { s.src = v }

func (s *TransferContext) SetDest(v IDestinationContext) { s.dest = v }

func (s *TransferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferContext) TRANSFER() antlr.TerminalNode {
	return s.GetToken(FaRlParserTRANSFER, 0)
}

func (s *TransferContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(FaRlParserLPAREN, 0)
}

func (s *TransferContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *TransferContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *TransferContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(FaRlParserRPAREN, 0)
}

func (s *TransferContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *TransferContext) TO() antlr.TerminalNode {
	return s.GetToken(FaRlParserTO, 0)
}

func (s *TransferContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TransferContext) MonetaryAll() IMonetaryAllContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryAllContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryAllContext)
}

func (s *TransferContext) ValueAwareSource() IValueAwareSourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueAwareSourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueAwareSourceContext)
}

func (s *TransferContext) Destination() IDestinationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDestinationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDestinationContext)
}

func (s *TransferContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterTransfer(s)
	}
}

func (s *TransferContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitTransfer(s)
	}
}

type SetTxMetaBlockContext struct {
	*StatementContext
}

func NewSetTxMetaBlockContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetTxMetaBlockContext {
	var p = new(SetTxMetaBlockContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SetTxMetaBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetTxMetaBlockContext) SET() antlr.TerminalNode {
	return s.GetToken(FaRlParserSET, 0)
}

func (s *SetTxMetaBlockContext) TRANSACTION() antlr.TerminalNode {
	return s.GetToken(FaRlParserTRANSACTION, 0)
}

func (s *SetTxMetaBlockContext) METADATA() antlr.TerminalNode {
	return s.GetToken(FaRlParserMETADATA, 0)
}

func (s *SetTxMetaBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACE, 0)
}

func (s *SetTxMetaBlockContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *SetTxMetaBlockContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *SetTxMetaBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACE, 0)
}

func (s *SetTxMetaBlockContext) AllMetadataEntry() []IMetadataEntryContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMetadataEntryContext)(nil)).Elem())
	var tst = make([]IMetadataEntryContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMetadataEntryContext)
		}
	}

	return tst
}

func (s *SetTxMetaBlockContext) MetadataEntry(i int) IMetadataEntryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataEntryContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMetadataEntryContext)
}

func (s *SetTxMetaBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSetTxMetaBlock(s)
	}
}

func (s *SetTxMetaBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSetTxMetaBlock(s)
	}
}

type SetTxMetaContext struct {
	*StatementContext
}

func NewSetTxMetaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetTxMetaContext {
	var p = new(SetTxMetaContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SetTxMetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetTxMetaContext) SET() antlr.TerminalNode {
	return s.GetToken(FaRlParserSET, 0)
}

func (s *SetTxMetaContext) TRANSACTION() antlr.TerminalNode {
	return s.GetToken(FaRlParserTRANSACTION, 0)
}

func (s *SetTxMetaContext) METADATA() antlr.TerminalNode {
	return s.GetToken(FaRlParserMETADATA, 0)
}

func (s *SetTxMetaContext) STRING() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTRING, 0)
}

func (s *SetTxMetaContext) EQ() antlr.TerminalNode {
	return s.GetToken(FaRlParserEQ, 0)
}

func (s *SetTxMetaContext) MetadataValue() IMetadataValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataValueContext)
}

func (s *SetTxMetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSetTxMeta(s)
	}
}

func (s *SetTxMetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSetTxMeta(s)
	}
}

type FailContext struct {
	*StatementContext
}

func NewFailContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FailContext {
	var p = new(FailContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *FailContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FailContext) FAIL() antlr.TerminalNode {
	return s.GetToken(FaRlParserFAIL, 0)
}

func (s *FailContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterFail(s)
	}
}

func (s *FailContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitFail(s)
	}
}

type SetAccountMetaContext struct {
	*StatementContext
}

func NewSetAccountMetaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SetAccountMetaContext {
	var p = new(SetAccountMetaContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SetAccountMetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SetAccountMetaContext) SET() antlr.TerminalNode {
	return s.GetToken(FaRlParserSET, 0)
}

func (s *SetAccountMetaContext) TY_ACCOUNT() antlr.TerminalNode {
	return s.GetToken(FaRlParserTY_ACCOUNT, 0)
}

func (s *SetAccountMetaContext) METADATA() antlr.TerminalNode {
	return s.GetToken(FaRlParserMETADATA, 0)
}

func (s *SetAccountMetaContext) OF() antlr.TerminalNode {
	return s.GetToken(FaRlParserOF, 0)
}

func (s *SetAccountMetaContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SetAccountMetaContext) KEY() antlr.TerminalNode {
	return s.GetToken(FaRlParserKEY, 0)
}

func (s *SetAccountMetaContext) STRING() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTRING, 0)
}

func (s *SetAccountMetaContext) EQ() antlr.TerminalNode {
	return s.GetToken(FaRlParserEQ, 0)
}

func (s *SetAccountMetaContext) MetadataValue() IMetadataValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMetadataValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMetadataValueContext)
}

func (s *SetAccountMetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSetAccountMeta(s)
	}
}

func (s *SetAccountMetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSetAccountMeta(s)
	}
}

func (p *FaRlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, FaRlParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(171)
			p.Match(FaRlParserPRINT)
		}
		{
			p.SetState(172)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case 2:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(173)
			p.Match(FaRlParserFAIL)
		}

	case 3:
		localctx = NewTransferContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(174)
			p.Match(FaRlParserTRANSFER)
		}
		p.SetState(177)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(175)

				var _x = p.expression(0)

				localctx.(*TransferContext).mon = _x
			}

		case 2:
			{
				p.SetState(176)

				var _x = p.MonetaryAll()

				localctx.(*TransferContext).monAll = _x
			}

		}
		{
			p.SetState(179)
			p.Match(FaRlParserLPAREN)
		}
		{
			p.SetState(180)
			p.Match(FaRlParserNEWLINE)
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case FaRlParserFROM:
			{
				p.SetState(181)
				p.Match(FaRlParserFROM)
			}
			{
				p.SetState(182)

				var _x = p.ValueAwareSource()

				localctx.(*TransferContext).src = _x
			}
			{
				p.SetState(183)
				p.Match(FaRlParserNEWLINE)
			}
			{
				p.SetState(184)
				p.Match(FaRlParserTO)
			}
			{
				p.SetState(185)

				var _x = p.Destination()

				localctx.(*TransferContext).dest = _x
			}

		case FaRlParserTO:
			{
				p.SetState(187)
				p.Match(FaRlParserTO)
			}
			{
				p.SetState(188)

				var _x = p.Destination()

				localctx.(*TransferContext).dest = _x
			}
			{
				p.SetState(189)
				p.Match(FaRlParserNEWLINE)
			}
			{
				p.SetState(190)
				p.Match(FaRlParserFROM)
			}
			{
				p.SetState(191)

				var _x = p.ValueAwareSource()

				localctx.(*TransferContext).src = _x
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(195)
			p.Match(FaRlParserNEWLINE)
		}
		{
			p.SetState(196)
			p.Match(FaRlParserRPAREN)
		}

	case 4:
		localctx = NewSetTxMetaContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(198)
			p.Match(FaRlParserSET)
		}
		{
			p.SetState(199)
			p.Match(FaRlParserTRANSACTION)
		}
		{
			p.SetState(200)
			p.Match(FaRlParserMETADATA)
		}
		{
			p.SetState(201)
			p.Match(FaRlParserSTRING)
		}
		{
			p.SetState(202)
			p.Match(FaRlParserEQ)
		}
		{
			p.SetState(203)
			p.MetadataValue()
		}

	case 5:
		localctx = NewSetTxMetaBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(204)
			p.Match(FaRlParserSET)
		}
		{
			p.SetState(205)
			p.Match(FaRlParserTRANSACTION)
		}
		{
			p.SetState(206)
			p.Match(FaRlParserMETADATA)
		}
		{
			p.SetState(207)
			p.Match(FaRlParserLBRACE)
		}
		{
			p.SetState(208)
			p.Match(FaRlParserNEWLINE)
		}
		p.SetState(212)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FaRlParserSTRING {
			{
				p.SetState(209)
				p.MetadataEntry()
			}
			{
				p.SetState(210)
				p.Match(FaRlParserNEWLINE)
			}

			p.SetState(214)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(216)
			p.Match(FaRlParserRBRACE)
		}

	case 6:
		localctx = NewSetAccountMetaContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(218)
			p.Match(FaRlParserSET)
		}
		{
			p.SetState(219)
			p.Match(FaRlParserTY_ACCOUNT)
		}
		{
			p.SetState(220)
			p.Match(FaRlParserMETADATA)
		}
		{
			p.SetState(221)
			p.Match(FaRlParserOF)
		}
		{
			p.SetState(222)
			p.expression(0)
		}
		{
			p.SetState(223)
			p.Match(FaRlParserKEY)
		}
		{
			p.SetState(224)
			p.Match(FaRlParserSTRING)
		}
		{
			p.SetState(225)
			p.Match(FaRlParserEQ)
		}
		{
			p.SetState(226)
			p.MetadataValue()
		}

	}

	return localctx
}

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStmts returns the stmts rule context list.
	GetStmts() []IStatementContext

	// SetStmts sets the stmts rule context list.
	SetStmts([]IStatementContext)

	// IsScriptContext differentiates from other interfaces.
	IsScriptContext()
}

type ScriptContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_statement IStatementContext
	stmts      []IStatementContext
}

func NewEmptyScriptContext() *ScriptContext {
	var p = new(ScriptContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_script
	return p
}

func (*ScriptContext) IsScriptContext() {}

func NewScriptContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScriptContext {
	var p = new(ScriptContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_script

	return p
}

func (s *ScriptContext) GetParser() antlr.Parser { return s.parser }

func (s *ScriptContext) Get_statement() IStatementContext { return s._statement }

func (s *ScriptContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *ScriptContext) GetStmts() []IStatementContext { return s.stmts }

func (s *ScriptContext) SetStmts(v []IStatementContext) { s.stmts = v }

func (s *ScriptContext) EOF() antlr.TerminalNode {
	return s.GetToken(FaRlParserEOF, 0)
}

func (s *ScriptContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *ScriptContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ScriptContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *ScriptContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *ScriptContext) AllVarDecl() []IVarDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDeclContext)(nil)).Elem())
	var tst = make([]IVarDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDeclContext)
		}
	}

	return tst
}

func (s *ScriptContext) VarDecl(i int) IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *ScriptContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScriptContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScriptContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterScript(s)
	}
}

func (s *ScriptContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitScript(s)
	}
}

func (p *FaRlParser) Script() (localctx IScriptContext) {
	localctx = NewScriptContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, FaRlParserRULE_script)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(233)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaRlParserNEWLINE {
		{
			p.SetState(230)
			p.Match(FaRlParserNEWLINE)
		}

		p.SetState(235)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(244)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaRlParserVAR {
		{
			p.SetState(236)
			p.VarDecl()
		}
		p.SetState(238)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FaRlParserNEWLINE {
			{
				p.SetState(237)
				p.Match(FaRlParserNEWLINE)
			}

			p.SetState(240)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(247)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(256)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(249)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == FaRlParserNEWLINE {
				{
					p.SetState(248)
					p.Match(FaRlParserNEWLINE)
				}

				p.SetState(251)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(253)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(258)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext())
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaRlParserNEWLINE {
		{
			p.SetState(259)
			p.Match(FaRlParserNEWLINE)
		}

	}
	{
		p.SetState(262)
		p.Match(FaRlParserEOF)
	}

	return localctx
}

func (p *FaRlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *FaRlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
