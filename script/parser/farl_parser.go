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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 55, 274,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2,
	3, 2, 3, 2, 3, 2, 3, 2, 5, 2, 42, 10, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3,
	48, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 53, 10, 4, 3, 4, 3, 4, 3, 4, 7, 4, 58,
	10, 4, 12, 4, 14, 4, 61, 11, 4, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 67, 10, 5,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	3, 6, 3, 6, 3, 6, 3, 6, 5, 6, 109, 10, 6, 3, 6, 3, 6, 3, 6, 7, 6, 114,
	10, 6, 12, 6, 14, 6, 117, 11, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 145, 10, 8, 3, 9, 3,
	9, 5, 9, 149, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 6, 11, 170, 10, 11, 13, 11, 14, 11, 171, 5, 11, 174, 10, 11, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 192, 10, 12, 13, 12, 14, 12,
	193, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 6, 12, 222, 10, 12, 13, 12, 14,
	12, 223, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12,
	3, 12, 3, 12, 3, 12, 5, 12, 238, 10, 12, 3, 13, 7, 13, 241, 10, 13, 12,
	13, 14, 13, 244, 11, 13, 3, 13, 3, 13, 6, 13, 248, 10, 13, 13, 13, 14,
	13, 249, 7, 13, 252, 10, 13, 12, 13, 14, 13, 255, 11, 13, 3, 13, 3, 13,
	6, 13, 259, 10, 13, 13, 13, 14, 13, 260, 3, 13, 7, 13, 264, 10, 13, 12,
	13, 14, 13, 267, 11, 13, 3, 13, 5, 13, 270, 10, 13, 3, 13, 3, 13, 3, 13,
	2, 4, 6, 10, 14, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 2, 4, 3, 2,
	38, 39, 3, 2, 32, 37, 2, 299, 2, 41, 3, 2, 2, 2, 4, 47, 3, 2, 2, 2, 6,
	52, 3, 2, 2, 2, 8, 66, 3, 2, 2, 2, 10, 108, 3, 2, 2, 2, 12, 118, 3, 2,
	2, 2, 14, 144, 3, 2, 2, 2, 16, 148, 3, 2, 2, 2, 18, 150, 3, 2, 2, 2, 20,
	173, 3, 2, 2, 2, 22, 237, 3, 2, 2, 2, 24, 242, 3, 2, 2, 2, 26, 27, 7, 41,
	2, 2, 27, 28, 7, 53, 2, 2, 28, 29, 7, 45, 2, 2, 29, 30, 7, 54, 2, 2, 30,
	31, 7, 54, 2, 2, 31, 42, 7, 42, 2, 2, 32, 33, 7, 41, 2, 2, 33, 34, 7, 53,
	2, 2, 34, 35, 7, 45, 2, 2, 35, 36, 7, 54, 2, 2, 36, 37, 7, 48, 2, 2, 37,
	42, 7, 42, 2, 2, 38, 39, 7, 41, 2, 2, 39, 40, 7, 53, 2, 2, 40, 42, 7, 42,
	2, 2, 41, 26, 3, 2, 2, 2, 41, 32, 3, 2, 2, 2, 41, 38, 3, 2, 2, 2, 42, 3,
	3, 2, 2, 2, 43, 48, 7, 52, 2, 2, 44, 48, 7, 53, 2, 2, 45, 48, 7, 54, 2,
	2, 46, 48, 5, 2, 2, 2, 47, 43, 3, 2, 2, 2, 47, 44, 3, 2, 2, 2, 47, 45,
	3, 2, 2, 2, 47, 46, 3, 2, 2, 2, 48, 5, 3, 2, 2, 2, 49, 50, 8, 4, 1, 2,
	50, 53, 5, 4, 3, 2, 51, 53, 7, 51, 2, 2, 52, 49, 3, 2, 2, 2, 52, 51, 3,
	2, 2, 2, 53, 59, 3, 2, 2, 2, 54, 55, 12, 5, 2, 2, 55, 56, 9, 2, 2, 2, 56,
	58, 5, 6, 4, 6, 57, 54, 3, 2, 2, 2, 58, 61, 3, 2, 2, 2, 59, 57, 3, 2, 2,
	2, 59, 60, 3, 2, 2, 2, 60, 7, 3, 2, 2, 2, 61, 59, 3, 2, 2, 2, 62, 63, 7,
	54, 2, 2, 63, 67, 7, 40, 2, 2, 64, 67, 7, 50, 2, 2, 65, 67, 7, 13, 2, 2,
	66, 62, 3, 2, 2, 2, 66, 64, 3, 2, 2, 2, 66, 65, 3, 2, 2, 2, 67, 9, 3, 2,
	2, 2, 68, 69, 8, 6, 1, 2, 69, 70, 7, 8, 2, 2, 70, 109, 5, 6, 4, 2, 71,
	72, 7, 8, 2, 2, 72, 73, 5, 6, 4, 2, 73, 74, 7, 17, 2, 2, 74, 75, 7, 18,
	2, 2, 75, 109, 3, 2, 2, 2, 76, 77, 7, 8, 2, 2, 77, 78, 5, 6, 4, 2, 78,
	79, 7, 17, 2, 2, 79, 80, 7, 18, 2, 2, 80, 81, 7, 19, 2, 2, 81, 82, 7, 9,
	2, 2, 82, 83, 5, 2, 2, 2, 83, 109, 3, 2, 2, 2, 84, 85, 7, 8, 2, 2, 85,
	86, 5, 6, 4, 2, 86, 87, 7, 16, 2, 2, 87, 88, 5, 2, 2, 2, 88, 89, 7, 10,
	2, 2, 89, 90, 5, 10, 6, 7, 90, 109, 3, 2, 2, 2, 91, 92, 7, 15, 2, 2, 92,
	93, 7, 54, 2, 2, 93, 94, 7, 40, 2, 2, 94, 95, 7, 8, 2, 2, 95, 109, 5, 6,
	4, 2, 96, 97, 7, 15, 2, 2, 97, 98, 7, 54, 2, 2, 98, 99, 7, 40, 2, 2, 99,
	100, 7, 8, 2, 2, 100, 101, 5, 6, 4, 2, 101, 102, 7, 16, 2, 2, 102, 103,
	5, 2, 2, 2, 103, 109, 3, 2, 2, 2, 104, 105, 7, 15, 2, 2, 105, 106, 7, 13,
	2, 2, 106, 107, 7, 8, 2, 2, 107, 109, 5, 6, 4, 2, 108, 68, 3, 2, 2, 2,
	108, 71, 3, 2, 2, 2, 108, 76, 3, 2, 2, 2, 108, 84, 3, 2, 2, 2, 108, 91,
	3, 2, 2, 2, 108, 96, 3, 2, 2, 2, 108, 104, 3, 2, 2, 2, 109, 115, 3, 2,
	2, 2, 110, 111, 12, 6, 2, 2, 111, 112, 7, 10, 2, 2, 112, 114, 5, 6, 4,
	2, 113, 110, 3, 2, 2, 2, 114, 117, 3, 2, 2, 2, 115, 113, 3, 2, 2, 2, 115,
	116, 3, 2, 2, 2, 116, 11, 3, 2, 2, 2, 117, 115, 3, 2, 2, 2, 118, 119, 9,
	3, 2, 2, 119, 13, 3, 2, 2, 2, 120, 121, 7, 24, 2, 2, 121, 122, 7, 51, 2,
	2, 122, 123, 7, 46, 2, 2, 123, 145, 5, 12, 7, 2, 124, 125, 7, 24, 2, 2,
	125, 126, 7, 51, 2, 2, 126, 127, 7, 49, 2, 2, 127, 128, 7, 25, 2, 2, 128,
	129, 7, 26, 2, 2, 129, 130, 5, 6, 4, 2, 130, 131, 7, 21, 2, 2, 131, 132,
	7, 53, 2, 2, 132, 133, 7, 45, 2, 2, 133, 134, 7, 54, 2, 2, 134, 145, 3,
	2, 2, 2, 135, 136, 7, 24, 2, 2, 136, 137, 7, 51, 2, 2, 137, 138, 7, 49,
	2, 2, 138, 139, 7, 27, 2, 2, 139, 140, 7, 26, 2, 2, 140, 141, 5, 6, 4,
	2, 141, 142, 7, 28, 2, 2, 142, 143, 7, 55, 2, 2, 143, 145, 3, 2, 2, 2,
	144, 120, 3, 2, 2, 2, 144, 124, 3, 2, 2, 2, 144, 135, 3, 2, 2, 2, 145,
	15, 3, 2, 2, 2, 146, 149, 5, 6, 4, 2, 147, 149, 7, 50, 2, 2, 148, 146,
	3, 2, 2, 2, 148, 147, 3, 2, 2, 2, 149, 17, 3, 2, 2, 2, 150, 151, 7, 55,
	2, 2, 151, 152, 7, 49, 2, 2, 152, 153, 5, 16, 9, 2, 153, 19, 3, 2, 2, 2,
	154, 155, 7, 11, 2, 2, 155, 156, 5, 8, 5, 2, 156, 157, 7, 9, 2, 2, 157,
	158, 5, 6, 4, 2, 158, 174, 3, 2, 2, 2, 159, 160, 7, 12, 2, 2, 160, 174,
	7, 13, 2, 2, 161, 162, 7, 22, 2, 2, 162, 163, 5, 8, 5, 2, 163, 164, 7,
	23, 2, 2, 164, 165, 7, 46, 2, 2, 165, 169, 7, 3, 2, 2, 166, 167, 5, 20,
	11, 2, 167, 168, 7, 3, 2, 2, 168, 170, 3, 2, 2, 2, 169, 166, 3, 2, 2, 2,
	170, 171, 3, 2, 2, 2, 171, 169, 3, 2, 2, 2, 171, 172, 3, 2, 2, 2, 172,
	174, 3, 2, 2, 2, 173, 154, 3, 2, 2, 2, 173, 159, 3, 2, 2, 2, 173, 161,
	3, 2, 2, 2, 174, 21, 3, 2, 2, 2, 175, 176, 7, 5, 2, 2, 176, 238, 5, 6,
	4, 2, 177, 238, 7, 6, 2, 2, 178, 179, 7, 7, 2, 2, 179, 180, 5, 6, 4, 2,
	180, 181, 5, 10, 6, 2, 181, 182, 7, 9, 2, 2, 182, 183, 5, 6, 4, 2, 183,
	238, 3, 2, 2, 2, 184, 185, 7, 7, 2, 2, 185, 186, 5, 6, 4, 2, 186, 187,
	5, 10, 6, 2, 187, 191, 7, 3, 2, 2, 188, 189, 5, 20, 11, 2, 189, 190, 7,
	3, 2, 2, 190, 192, 3, 2, 2, 2, 191, 188, 3, 2, 2, 2, 192, 193, 3, 2, 2,
	2, 193, 191, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 238, 3, 2, 2, 2, 195,
	196, 7, 7, 2, 2, 196, 197, 7, 14, 2, 2, 197, 198, 5, 2, 2, 2, 198, 199,
	5, 10, 6, 2, 199, 200, 7, 9, 2, 2, 200, 201, 5, 6, 4, 2, 201, 238, 3, 2,
	2, 2, 202, 203, 7, 20, 2, 2, 203, 204, 5, 2, 2, 2, 204, 205, 7, 21, 2,
	2, 205, 206, 5, 6, 4, 2, 206, 238, 3, 2, 2, 2, 207, 208, 7, 29, 2, 2, 208,
	209, 7, 30, 2, 2, 209, 210, 7, 31, 2, 2, 210, 211, 7, 55, 2, 2, 211, 212,
	7, 49, 2, 2, 212, 238, 5, 16, 9, 2, 213, 214, 7, 29, 2, 2, 214, 215, 7,
	30, 2, 2, 215, 216, 7, 31, 2, 2, 216, 217, 7, 43, 2, 2, 217, 221, 7, 3,
	2, 2, 218, 219, 5, 18, 10, 2, 219, 220, 7, 3, 2, 2, 220, 222, 3, 2, 2,
	2, 221, 218, 3, 2, 2, 2, 222, 223, 3, 2, 2, 2, 223, 221, 3, 2, 2, 2, 223,
	224, 3, 2, 2, 2, 224, 225, 3, 2, 2, 2, 225, 226, 7, 44, 2, 2, 226, 238,
	3, 2, 2, 2, 227, 228, 7, 29, 2, 2, 228, 229, 7, 32, 2, 2, 229, 230, 7,
	31, 2, 2, 230, 231, 7, 26, 2, 2, 231, 232, 5, 6, 4, 2, 232, 233, 7, 28,
	2, 2, 233, 234, 7, 55, 2, 2, 234, 235, 7, 49, 2, 2, 235, 236, 5, 16, 9,
	2, 236, 238, 3, 2, 2, 2, 237, 175, 3, 2, 2, 2, 237, 177, 3, 2, 2, 2, 237,
	178, 3, 2, 2, 2, 237, 184, 3, 2, 2, 2, 237, 195, 3, 2, 2, 2, 237, 202,
	3, 2, 2, 2, 237, 207, 3, 2, 2, 2, 237, 213, 3, 2, 2, 2, 237, 227, 3, 2,
	2, 2, 238, 23, 3, 2, 2, 2, 239, 241, 7, 3, 2, 2, 240, 239, 3, 2, 2, 2,
	241, 244, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243,
	253, 3, 2, 2, 2, 244, 242, 3, 2, 2, 2, 245, 247, 5, 14, 8, 2, 246, 248,
	7, 3, 2, 2, 247, 246, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 247, 3, 2,
	2, 2, 249, 250, 3, 2, 2, 2, 250, 252, 3, 2, 2, 2, 251, 245, 3, 2, 2, 2,
	252, 255, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2, 2, 2, 254,
	256, 3, 2, 2, 2, 255, 253, 3, 2, 2, 2, 256, 265, 5, 22, 12, 2, 257, 259,
	7, 3, 2, 2, 258, 257, 3, 2, 2, 2, 259, 260, 3, 2, 2, 2, 260, 258, 3, 2,
	2, 2, 260, 261, 3, 2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 264, 5, 22, 12,
	2, 263, 258, 3, 2, 2, 2, 264, 267, 3, 2, 2, 2, 265, 263, 3, 2, 2, 2, 265,
	266, 3, 2, 2, 2, 266, 269, 3, 2, 2, 2, 267, 265, 3, 2, 2, 2, 268, 270,
	7, 3, 2, 2, 269, 268, 3, 2, 2, 2, 269, 270, 3, 2, 2, 2, 270, 271, 3, 2,
	2, 2, 271, 272, 7, 2, 2, 3, 272, 25, 3, 2, 2, 2, 22, 41, 47, 52, 59, 66,
	108, 115, 144, 148, 171, 173, 193, 223, 237, 242, 249, 253, 260, 265, 269,
}
var literalNames = []string{
	"", "", "", "'print'", "'fail'", "'transfer'", "'from'", "'to'", "'then'",
	"'send'", "'keep'", "'remaining'", "'all'", "'take'", "'limit'", "'allow'",
	"'overdraft'", "'up'", "'reserve'", "'in'", "'split'", "'as'", "'var'",
	"'balance'", "'of'", "'meta'", "'key'", "'set'", "'transaction'", "'metadata'",
	"'account'", "'asset'", "'number'", "'monetary'", "'portion'", "'string'",
	"'+'", "'-'", "'%'", "'['", "']'", "'{'", "'}'", "'.'", "':'", "'/'", "'*'",
	"'='",
}
var symbolicNames = []string{
	"", "NEWLINE", "WHITESPACE", "PRINT", "FAIL", "TRANSFER", "FROM", "TO",
	"THEN", "SEND", "KEEP", "REMAINING", "ALL", "TAKE", "LIMIT", "ALLOW", "OVERDRAFT",
	"UP", "RESERVE", "IN", "SPLIT", "AS", "VAR", "BALANCE", "OF", "META", "KEY",
	"SET", "TRANSACTION", "METADATA", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER",
	"TY_MONETARY", "TY_PORTION", "TY_STRING", "OP_ADD", "OP_SUB", "PERCENT",
	"LBRACK", "RBRACK", "LBRACE", "RBRACE", "DOT", "COLON", "SLASH", "STAR",
	"EQ", "RATIO", "VARIABLE_NAME", "ACCOUNT", "ASSET", "NUMBER", "STRING",
}

var ruleNames = []string{
	"monetary", "literal", "expression", "portion", "source", "type_", "varDecl",
	"metadataValue", "metadataEntry", "sendClause", "statement", "script",
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
	FaRlParserEOF           = antlr.TokenEOF
	FaRlParserNEWLINE       = 1
	FaRlParserWHITESPACE    = 2
	FaRlParserPRINT         = 3
	FaRlParserFAIL          = 4
	FaRlParserTRANSFER      = 5
	FaRlParserFROM          = 6
	FaRlParserTO            = 7
	FaRlParserTHEN          = 8
	FaRlParserSEND          = 9
	FaRlParserKEEP          = 10
	FaRlParserREMAINING     = 11
	FaRlParserALL           = 12
	FaRlParserTAKE          = 13
	FaRlParserLIMIT         = 14
	FaRlParserALLOW         = 15
	FaRlParserOVERDRAFT     = 16
	FaRlParserUP            = 17
	FaRlParserRESERVE       = 18
	FaRlParserIN            = 19
	FaRlParserSPLIT         = 20
	FaRlParserAS            = 21
	FaRlParserVAR           = 22
	FaRlParserBALANCE       = 23
	FaRlParserOF            = 24
	FaRlParserMETA          = 25
	FaRlParserKEY           = 26
	FaRlParserSET           = 27
	FaRlParserTRANSACTION   = 28
	FaRlParserMETADATA      = 29
	FaRlParserTY_ACCOUNT    = 30
	FaRlParserTY_ASSET      = 31
	FaRlParserTY_NUMBER     = 32
	FaRlParserTY_MONETARY   = 33
	FaRlParserTY_PORTION    = 34
	FaRlParserTY_STRING     = 35
	FaRlParserOP_ADD        = 36
	FaRlParserOP_SUB        = 37
	FaRlParserPERCENT       = 38
	FaRlParserLBRACK        = 39
	FaRlParserRBRACK        = 40
	FaRlParserLBRACE        = 41
	FaRlParserRBRACE        = 42
	FaRlParserDOT           = 43
	FaRlParserCOLON         = 44
	FaRlParserSLASH         = 45
	FaRlParserSTAR          = 46
	FaRlParserEQ            = 47
	FaRlParserRATIO         = 48
	FaRlParserVARIABLE_NAME = 49
	FaRlParserACCOUNT       = 50
	FaRlParserASSET         = 51
	FaRlParserNUMBER        = 52
	FaRlParserSTRING        = 53
)

// FaRlParser rules.
const (
	FaRlParserRULE_monetary      = 0
	FaRlParserRULE_literal       = 1
	FaRlParserRULE_expression    = 2
	FaRlParserRULE_portion       = 3
	FaRlParserRULE_source        = 4
	FaRlParserRULE_type_         = 5
	FaRlParserRULE_varDecl       = 6
	FaRlParserRULE_metadataValue = 7
	FaRlParserRULE_metadataEntry = 8
	FaRlParserRULE_sendClause    = 9
	FaRlParserRULE_statement     = 10
	FaRlParserRULE_script        = 11
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

type MonetaryAllContext struct {
	*MonetaryContext
	asset     antlr.Token
	precision antlr.Token
}

func NewMonetaryAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MonetaryAllContext {
	var p = new(MonetaryAllContext)

	p.MonetaryContext = NewEmptyMonetaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MonetaryContext))

	return p
}

func (s *MonetaryAllContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryAllContext) GetPrecision() antlr.Token { return s.precision }

func (s *MonetaryAllContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryAllContext) SetPrecision(v antlr.Token) { s.precision = v }

func (s *MonetaryAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryAllContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACK, 0)
}

func (s *MonetaryAllContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaRlParserDOT, 0)
}

func (s *MonetaryAllContext) STAR() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTAR, 0)
}

func (s *MonetaryAllContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACK, 0)
}

func (s *MonetaryAllContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *MonetaryAllContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *MonetaryAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMonetaryAll(s)
	}
}

func (s *MonetaryAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMonetaryAll(s)
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

type MonetaryAssetOnlyContext struct {
	*MonetaryContext
	asset antlr.Token
}

func NewMonetaryAssetOnlyContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MonetaryAssetOnlyContext {
	var p = new(MonetaryAssetOnlyContext)

	p.MonetaryContext = NewEmptyMonetaryContext()
	p.parser = parser
	p.CopyFrom(ctx.(*MonetaryContext))

	return p
}

func (s *MonetaryAssetOnlyContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryAssetOnlyContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryAssetOnlyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryAssetOnlyContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACK, 0)
}

func (s *MonetaryAssetOnlyContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACK, 0)
}

func (s *MonetaryAssetOnlyContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *MonetaryAssetOnlyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMonetaryAssetOnly(s)
	}
}

func (s *MonetaryAssetOnlyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMonetaryAssetOnly(s)
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

	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) {
	case 1:
		localctx = NewMonetaryLitContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(24)
			p.Match(FaRlParserLBRACK)
		}
		{
			p.SetState(25)

			var _m = p.Match(FaRlParserASSET)

			localctx.(*MonetaryLitContext).asset = _m
		}
		{
			p.SetState(26)
			p.Match(FaRlParserDOT)
		}
		{
			p.SetState(27)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*MonetaryLitContext).precision = _m
		}
		{
			p.SetState(28)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*MonetaryLitContext).amount = _m
		}
		{
			p.SetState(29)
			p.Match(FaRlParserRBRACK)
		}

	case 2:
		localctx = NewMonetaryAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(30)
			p.Match(FaRlParserLBRACK)
		}
		{
			p.SetState(31)

			var _m = p.Match(FaRlParserASSET)

			localctx.(*MonetaryAllContext).asset = _m
		}
		{
			p.SetState(32)
			p.Match(FaRlParserDOT)
		}
		{
			p.SetState(33)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*MonetaryAllContext).precision = _m
		}
		{
			p.SetState(34)
			p.Match(FaRlParserSTAR)
		}
		{
			p.SetState(35)
			p.Match(FaRlParserRBRACK)
		}

	case 3:
		localctx = NewMonetaryAssetOnlyContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.Match(FaRlParserLBRACK)
		}
		{
			p.SetState(37)

			var _m = p.Match(FaRlParserASSET)

			localctx.(*MonetaryAssetOnlyContext).asset = _m
		}
		{
			p.SetState(38)
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
	p.EnterRule(localctx, 2, FaRlParserRULE_literal)

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

	p.SetState(45)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(41)
			p.Match(FaRlParserACCOUNT)
		}

	case FaRlParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(42)
			p.Match(FaRlParserASSET)
		}

	case FaRlParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(43)
			p.Match(FaRlParserNUMBER)
		}

	case FaRlParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(44)
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
	_startState := 4
	p.EnterRecursionRule(localctx, 4, FaRlParserRULE_expression, _p)
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
	p.SetState(50)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserLBRACK, FaRlParserACCOUNT, FaRlParserASSET, FaRlParserNUMBER:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(48)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case FaRlParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(49)

			var _m = p.Match(FaRlParserVARIABLE_NAME)

			localctx.(*ExprVariableContext).variable = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*ExprAddSubContext).lhs = _prevctx

			p.PushNewRecursionContext(localctx, _startState, FaRlParserRULE_expression)
			p.SetState(52)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(53)

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
				p.SetState(54)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext())
	}

	return localctx
}

// IPortionContext is an interface to support dynamic dispatch.
type IPortionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPortionContext differentiates from other interfaces.
	IsPortionContext()
}

type PortionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPortionContext() *PortionContext {
	var p = new(PortionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_portion
	return p
}

func (*PortionContext) IsPortionContext() {}

func NewPortionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PortionContext {
	var p = new(PortionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_portion

	return p
}

func (s *PortionContext) GetParser() antlr.Parser { return s.parser }

func (s *PortionContext) CopyFrom(ctx *PortionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PortionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PortionRemainingContext struct {
	*PortionContext
}

func NewPortionRemainingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PortionRemainingContext {
	var p = new(PortionRemainingContext)

	p.PortionContext = NewEmptyPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PortionContext))

	return p
}

func (s *PortionRemainingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionRemainingContext) REMAINING() antlr.TerminalNode {
	return s.GetToken(FaRlParserREMAINING, 0)
}

func (s *PortionRemainingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterPortionRemaining(s)
	}
}

func (s *PortionRemainingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitPortionRemaining(s)
	}
}

type PortionPercentContext struct {
	*PortionContext
	p antlr.Token
}

func NewPortionPercentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PortionPercentContext {
	var p = new(PortionPercentContext)

	p.PortionContext = NewEmptyPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PortionContext))

	return p
}

func (s *PortionPercentContext) GetP() antlr.Token { return s.p }

func (s *PortionPercentContext) SetP(v antlr.Token) { s.p = v }

func (s *PortionPercentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionPercentContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(FaRlParserPERCENT, 0)
}

func (s *PortionPercentContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *PortionPercentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterPortionPercent(s)
	}
}

func (s *PortionPercentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitPortionPercent(s)
	}
}

type PortionRatioContext struct {
	*PortionContext
	r antlr.Token
}

func NewPortionRatioContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PortionRatioContext {
	var p = new(PortionRatioContext)

	p.PortionContext = NewEmptyPortionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PortionContext))

	return p
}

func (s *PortionRatioContext) GetR() antlr.Token { return s.r }

func (s *PortionRatioContext) SetR(v antlr.Token) { s.r = v }

func (s *PortionRatioContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PortionRatioContext) RATIO() antlr.TerminalNode {
	return s.GetToken(FaRlParserRATIO, 0)
}

func (s *PortionRatioContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterPortionRatio(s)
	}
}

func (s *PortionRatioContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitPortionRatio(s)
	}
}

func (p *FaRlParser) Portion() (localctx IPortionContext) {
	localctx = NewPortionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, FaRlParserRULE_portion)

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

	p.SetState(64)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserNUMBER:
		localctx = NewPortionPercentContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(60)

			var _m = p.Match(FaRlParserNUMBER)

			localctx.(*PortionPercentContext).p = _m
		}
		{
			p.SetState(61)
			p.Match(FaRlParserPERCENT)
		}

	case FaRlParserRATIO:
		localctx = NewPortionRatioContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(62)

			var _m = p.Match(FaRlParserRATIO)

			localctx.(*PortionRatioContext).r = _m
		}

	case FaRlParserREMAINING:
		localctx = NewPortionRemainingContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(63)
			p.Match(FaRlParserREMAINING)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

type SrcRemainingContext struct {
	*SourceContext
}

func NewSrcRemainingContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcRemainingContext {
	var p = new(SrcRemainingContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcRemainingContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcRemainingContext) TAKE() antlr.TerminalNode {
	return s.GetToken(FaRlParserTAKE, 0)
}

func (s *SrcRemainingContext) REMAINING() antlr.TerminalNode {
	return s.GetToken(FaRlParserREMAINING, 0)
}

func (s *SrcRemainingContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SrcRemainingContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcRemainingContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcRemaining(s)
	}
}

func (s *SrcRemainingContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcRemaining(s)
	}
}

type SrcPercentLimitContext struct {
	*SourceContext
}

func NewSrcPercentLimitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcPercentLimitContext {
	var p = new(SrcPercentLimitContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcPercentLimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcPercentLimitContext) TAKE() antlr.TerminalNode {
	return s.GetToken(FaRlParserTAKE, 0)
}

func (s *SrcPercentLimitContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *SrcPercentLimitContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(FaRlParserPERCENT, 0)
}

func (s *SrcPercentLimitContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SrcPercentLimitContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcPercentLimitContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(FaRlParserLIMIT, 0)
}

func (s *SrcPercentLimitContext) Monetary() IMonetaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryContext)
}

func (s *SrcPercentLimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcPercentLimit(s)
	}
}

func (s *SrcPercentLimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcPercentLimit(s)
	}
}

type SrcCascadeContext struct {
	*SourceContext
}

func NewSrcCascadeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcCascadeContext {
	var p = new(SrcCascadeContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcCascadeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcCascadeContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SrcCascadeContext) THEN() antlr.TerminalNode {
	return s.GetToken(FaRlParserTHEN, 0)
}

func (s *SrcCascadeContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcCascadeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcCascade(s)
	}
}

func (s *SrcCascadeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcCascade(s)
	}
}

type SrcSimpleContext struct {
	*SourceContext
}

func NewSrcSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcSimpleContext {
	var p = new(SrcSimpleContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcSimpleContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SrcSimpleContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcSimple(s)
	}
}

func (s *SrcSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcSimple(s)
	}
}

type SrcOverdraftContext struct {
	*SourceContext
}

func NewSrcOverdraftContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcOverdraftContext {
	var p = new(SrcOverdraftContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcOverdraftContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcOverdraftContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SrcOverdraftContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcOverdraftContext) ALLOW() antlr.TerminalNode {
	return s.GetToken(FaRlParserALLOW, 0)
}

func (s *SrcOverdraftContext) OVERDRAFT() antlr.TerminalNode {
	return s.GetToken(FaRlParserOVERDRAFT, 0)
}

func (s *SrcOverdraftContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcOverdraft(s)
	}
}

func (s *SrcOverdraftContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcOverdraft(s)
	}
}

type SrcOverdraftCappedContext struct {
	*SourceContext
}

func NewSrcOverdraftCappedContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcOverdraftCappedContext {
	var p = new(SrcOverdraftCappedContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcOverdraftCappedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcOverdraftCappedContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SrcOverdraftCappedContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcOverdraftCappedContext) ALLOW() antlr.TerminalNode {
	return s.GetToken(FaRlParserALLOW, 0)
}

func (s *SrcOverdraftCappedContext) OVERDRAFT() antlr.TerminalNode {
	return s.GetToken(FaRlParserOVERDRAFT, 0)
}

func (s *SrcOverdraftCappedContext) UP() antlr.TerminalNode {
	return s.GetToken(FaRlParserUP, 0)
}

func (s *SrcOverdraftCappedContext) TO() antlr.TerminalNode {
	return s.GetToken(FaRlParserTO, 0)
}

func (s *SrcOverdraftCappedContext) Monetary() IMonetaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryContext)
}

func (s *SrcOverdraftCappedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcOverdraftCapped(s)
	}
}

func (s *SrcOverdraftCappedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcOverdraftCapped(s)
	}
}

type SrcLimitContext struct {
	*SourceContext
}

func NewSrcLimitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcLimitContext {
	var p = new(SrcLimitContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcLimitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcLimitContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SrcLimitContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcLimitContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(FaRlParserLIMIT, 0)
}

func (s *SrcLimitContext) Monetary() IMonetaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryContext)
}

func (s *SrcLimitContext) THEN() antlr.TerminalNode {
	return s.GetToken(FaRlParserTHEN, 0)
}

func (s *SrcLimitContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *SrcLimitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcLimit(s)
	}
}

func (s *SrcLimitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcLimit(s)
	}
}

type SrcPercentContext struct {
	*SourceContext
}

func NewSrcPercentContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SrcPercentContext {
	var p = new(SrcPercentContext)

	p.SourceContext = NewEmptySourceContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SourceContext))

	return p
}

func (s *SrcPercentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SrcPercentContext) TAKE() antlr.TerminalNode {
	return s.GetToken(FaRlParserTAKE, 0)
}

func (s *SrcPercentContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *SrcPercentContext) PERCENT() antlr.TerminalNode {
	return s.GetToken(FaRlParserPERCENT, 0)
}

func (s *SrcPercentContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *SrcPercentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SrcPercentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSrcPercent(s)
	}
}

func (s *SrcPercentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSrcPercent(s)
	}
}

func (p *FaRlParser) Source() (localctx ISourceContext) {
	return p.source(0)
}

func (p *FaRlParser) source(_p int) (localctx ISourceContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewSourceContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ISourceContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 8
	p.EnterRecursionRule(localctx, 8, FaRlParserRULE_source, _p)

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
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSrcSimpleContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(67)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(68)
			p.expression(0)
		}

	case 2:
		localctx = NewSrcOverdraftContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(69)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(70)
			p.expression(0)
		}
		{
			p.SetState(71)
			p.Match(FaRlParserALLOW)
		}
		{
			p.SetState(72)
			p.Match(FaRlParserOVERDRAFT)
		}

	case 3:
		localctx = NewSrcOverdraftCappedContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(74)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(75)
			p.expression(0)
		}
		{
			p.SetState(76)
			p.Match(FaRlParserALLOW)
		}
		{
			p.SetState(77)
			p.Match(FaRlParserOVERDRAFT)
		}
		{
			p.SetState(78)
			p.Match(FaRlParserUP)
		}
		{
			p.SetState(79)
			p.Match(FaRlParserTO)
		}
		{
			p.SetState(80)
			p.Monetary()
		}

	case 4:
		localctx = NewSrcLimitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(82)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(83)
			p.expression(0)
		}
		{
			p.SetState(84)
			p.Match(FaRlParserLIMIT)
		}
		{
			p.SetState(85)
			p.Monetary()
		}
		{
			p.SetState(86)
			p.Match(FaRlParserTHEN)
		}
		{
			p.SetState(87)
			p.source(5)
		}

	case 5:
		localctx = NewSrcPercentContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(89)
			p.Match(FaRlParserTAKE)
		}
		{
			p.SetState(90)
			p.Match(FaRlParserNUMBER)
		}
		{
			p.SetState(91)
			p.Match(FaRlParserPERCENT)
		}
		{
			p.SetState(92)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(93)
			p.expression(0)
		}

	case 6:
		localctx = NewSrcPercentLimitContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(94)
			p.Match(FaRlParserTAKE)
		}
		{
			p.SetState(95)
			p.Match(FaRlParserNUMBER)
		}
		{
			p.SetState(96)
			p.Match(FaRlParserPERCENT)
		}
		{
			p.SetState(97)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(98)
			p.expression(0)
		}
		{
			p.SetState(99)
			p.Match(FaRlParserLIMIT)
		}
		{
			p.SetState(100)
			p.Monetary()
		}

	case 7:
		localctx = NewSrcRemainingContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)
			p.Match(FaRlParserTAKE)
		}
		{
			p.SetState(103)
			p.Match(FaRlParserREMAINING)
		}
		{
			p.SetState(104)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(105)
			p.expression(0)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(113)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewSrcCascadeContext(p, NewSourceContext(p, _parentctx, _parentState))
			p.PushNewRecursionContext(localctx, _startState, FaRlParserRULE_source)
			p.SetState(108)

			if !(p.Precpred(p.GetParserRuleContext(), 4)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
			}
			{
				p.SetState(109)
				p.Match(FaRlParserTHEN)
			}
			{
				p.SetState(110)
				p.expression(0)
			}

		}
		p.SetState(115)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 10, FaRlParserRULE_type_)
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
		p.SetState(116)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-30)&-(0x1f+1)) == 0 && ((1<<uint((_la-30)))&((1<<(FaRlParserTY_ACCOUNT-30))|(1<<(FaRlParserTY_ASSET-30))|(1<<(FaRlParserTY_NUMBER-30))|(1<<(FaRlParserTY_MONETARY-30))|(1<<(FaRlParserTY_PORTION-30))|(1<<(FaRlParserTY_STRING-30)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
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

func (s *VarTypedContext) SetTy(v IType_Context) { s.ty = v }

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

type VarBalanceContext struct {
	*VarDeclContext
	name antlr.Token
}

func NewVarBalanceContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarBalanceContext {
	var p = new(VarBalanceContext)

	p.VarDeclContext = NewEmptyVarDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VarDeclContext))

	return p
}

func (s *VarBalanceContext) GetName() antlr.Token { return s.name }

func (s *VarBalanceContext) SetName(v antlr.Token) { s.name = v }

func (s *VarBalanceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarBalanceContext) VAR() antlr.TerminalNode {
	return s.GetToken(FaRlParserVAR, 0)
}

func (s *VarBalanceContext) EQ() antlr.TerminalNode {
	return s.GetToken(FaRlParserEQ, 0)
}

func (s *VarBalanceContext) BALANCE() antlr.TerminalNode {
	return s.GetToken(FaRlParserBALANCE, 0)
}

func (s *VarBalanceContext) OF() antlr.TerminalNode {
	return s.GetToken(FaRlParserOF, 0)
}

func (s *VarBalanceContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarBalanceContext) IN() antlr.TerminalNode {
	return s.GetToken(FaRlParserIN, 0)
}

func (s *VarBalanceContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *VarBalanceContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaRlParserDOT, 0)
}

func (s *VarBalanceContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *VarBalanceContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(FaRlParserVARIABLE_NAME, 0)
}

func (s *VarBalanceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterVarBalance(s)
	}
}

func (s *VarBalanceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitVarBalance(s)
	}
}

type VarMetaContext struct {
	*VarDeclContext
	name antlr.Token
}

func NewVarMetaContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VarMetaContext {
	var p = new(VarMetaContext)

	p.VarDeclContext = NewEmptyVarDeclContext()
	p.parser = parser
	p.CopyFrom(ctx.(*VarDeclContext))

	return p
}

func (s *VarMetaContext) GetName() antlr.Token { return s.name }

func (s *VarMetaContext) SetName(v antlr.Token) { s.name = v }

func (s *VarMetaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarMetaContext) VAR() antlr.TerminalNode {
	return s.GetToken(FaRlParserVAR, 0)
}

func (s *VarMetaContext) EQ() antlr.TerminalNode {
	return s.GetToken(FaRlParserEQ, 0)
}

func (s *VarMetaContext) META() antlr.TerminalNode {
	return s.GetToken(FaRlParserMETA, 0)
}

func (s *VarMetaContext) OF() antlr.TerminalNode {
	return s.GetToken(FaRlParserOF, 0)
}

func (s *VarMetaContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VarMetaContext) KEY() antlr.TerminalNode {
	return s.GetToken(FaRlParserKEY, 0)
}

func (s *VarMetaContext) STRING() antlr.TerminalNode {
	return s.GetToken(FaRlParserSTRING, 0)
}

func (s *VarMetaContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(FaRlParserVARIABLE_NAME, 0)
}

func (s *VarMetaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterVarMeta(s)
	}
}

func (s *VarMetaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitVarMeta(s)
	}
}

func (p *FaRlParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FaRlParserRULE_varDecl)

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

	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVarTypedContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(118)
			p.Match(FaRlParserVAR)
		}
		{
			p.SetState(119)

			var _m = p.Match(FaRlParserVARIABLE_NAME)

			localctx.(*VarTypedContext).name = _m
		}
		{
			p.SetState(120)
			p.Match(FaRlParserCOLON)
		}
		{
			p.SetState(121)

			var _x = p.Type_()

			localctx.(*VarTypedContext).ty = _x
		}

	case 2:
		localctx = NewVarBalanceContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(122)
			p.Match(FaRlParserVAR)
		}
		{
			p.SetState(123)

			var _m = p.Match(FaRlParserVARIABLE_NAME)

			localctx.(*VarBalanceContext).name = _m
		}
		{
			p.SetState(124)
			p.Match(FaRlParserEQ)
		}
		{
			p.SetState(125)
			p.Match(FaRlParserBALANCE)
		}
		{
			p.SetState(126)
			p.Match(FaRlParserOF)
		}
		{
			p.SetState(127)
			p.expression(0)
		}
		{
			p.SetState(128)
			p.Match(FaRlParserIN)
		}
		{
			p.SetState(129)
			p.Match(FaRlParserASSET)
		}
		{
			p.SetState(130)
			p.Match(FaRlParserDOT)
		}
		{
			p.SetState(131)
			p.Match(FaRlParserNUMBER)
		}

	case 3:
		localctx = NewVarMetaContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(133)
			p.Match(FaRlParserVAR)
		}
		{
			p.SetState(134)

			var _m = p.Match(FaRlParserVARIABLE_NAME)

			localctx.(*VarMetaContext).name = _m
		}
		{
			p.SetState(135)
			p.Match(FaRlParserEQ)
		}
		{
			p.SetState(136)
			p.Match(FaRlParserMETA)
		}
		{
			p.SetState(137)
			p.Match(FaRlParserOF)
		}
		{
			p.SetState(138)
			p.expression(0)
		}
		{
			p.SetState(139)
			p.Match(FaRlParserKEY)
		}
		{
			p.SetState(140)
			p.Match(FaRlParserSTRING)
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
	p.EnterRule(localctx, 14, FaRlParserRULE_metadataValue)

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

	p.SetState(146)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserLBRACK, FaRlParserVARIABLE_NAME, FaRlParserACCOUNT, FaRlParserASSET, FaRlParserNUMBER:
		localctx = NewMetaValueExprContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(144)
			p.expression(0)
		}

	case FaRlParserRATIO:
		localctx = NewMetaValueRatioContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)
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
	p.EnterRule(localctx, 16, FaRlParserRULE_metadataEntry)

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
		p.SetState(148)
		p.Match(FaRlParserSTRING)
	}
	{
		p.SetState(149)
		p.Match(FaRlParserEQ)
	}
	{
		p.SetState(150)
		p.MetadataValue()
	}

	return localctx
}

// ISendClauseContext is an interface to support dynamic dispatch.
type ISendClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSendClauseContext differentiates from other interfaces.
	IsSendClauseContext()
}

type SendClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySendClauseContext() *SendClauseContext {
	var p = new(SendClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_sendClause
	return p
}

func (*SendClauseContext) IsSendClauseContext() {}

func NewSendClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SendClauseContext {
	var p = new(SendClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_sendClause

	return p
}

func (s *SendClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *SendClauseContext) CopyFrom(ctx *SendClauseContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SendClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SendToContext struct {
	*SendClauseContext
}

func NewSendToContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendToContext {
	var p = new(SendToContext)

	p.SendClauseContext = NewEmptySendClauseContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SendClauseContext))

	return p
}

func (s *SendToContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendToContext) SEND() antlr.TerminalNode {
	return s.GetToken(FaRlParserSEND, 0)
}

func (s *SendToContext) Portion() IPortionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortionContext)
}

func (s *SendToContext) TO() antlr.TerminalNode {
	return s.GetToken(FaRlParserTO, 0)
}

func (s *SendToContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SendToContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSendTo(s)
	}
}

func (s *SendToContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSendTo(s)
	}
}

type SendSplitContext struct {
	*SendClauseContext
}

func NewSendSplitContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendSplitContext {
	var p = new(SendSplitContext)

	p.SendClauseContext = NewEmptySendClauseContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SendClauseContext))

	return p
}

func (s *SendSplitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendSplitContext) SPLIT() antlr.TerminalNode {
	return s.GetToken(FaRlParserSPLIT, 0)
}

func (s *SendSplitContext) Portion() IPortionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPortionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPortionContext)
}

func (s *SendSplitContext) AS() antlr.TerminalNode {
	return s.GetToken(FaRlParserAS, 0)
}

func (s *SendSplitContext) COLON() antlr.TerminalNode {
	return s.GetToken(FaRlParserCOLON, 0)
}

func (s *SendSplitContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *SendSplitContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *SendSplitContext) AllSendClause() []ISendClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISendClauseContext)(nil)).Elem())
	var tst = make([]ISendClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISendClauseContext)
		}
	}

	return tst
}

func (s *SendSplitContext) SendClause(i int) ISendClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISendClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISendClauseContext)
}

func (s *SendSplitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSendSplit(s)
	}
}

func (s *SendSplitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSendSplit(s)
	}
}

type SendKeepContext struct {
	*SendClauseContext
}

func NewSendKeepContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SendKeepContext {
	var p = new(SendKeepContext)

	p.SendClauseContext = NewEmptySendClauseContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SendClauseContext))

	return p
}

func (s *SendKeepContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SendKeepContext) KEEP() antlr.TerminalNode {
	return s.GetToken(FaRlParserKEEP, 0)
}

func (s *SendKeepContext) REMAINING() antlr.TerminalNode {
	return s.GetToken(FaRlParserREMAINING, 0)
}

func (s *SendKeepContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterSendKeep(s)
	}
}

func (s *SendKeepContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitSendKeep(s)
	}
}

func (p *FaRlParser) SendClause() (localctx ISendClauseContext) {
	localctx = NewSendClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, FaRlParserRULE_sendClause)
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

	p.SetState(171)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserSEND:
		localctx = NewSendToContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(FaRlParserSEND)
		}
		{
			p.SetState(153)
			p.Portion()
		}
		{
			p.SetState(154)
			p.Match(FaRlParserTO)
		}
		{
			p.SetState(155)
			p.expression(0)
		}

	case FaRlParserKEEP:
		localctx = NewSendKeepContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(157)
			p.Match(FaRlParserKEEP)
		}
		{
			p.SetState(158)
			p.Match(FaRlParserREMAINING)
		}

	case FaRlParserSPLIT:
		localctx = NewSendSplitContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(159)
			p.Match(FaRlParserSPLIT)
		}
		{
			p.SetState(160)
			p.Portion()
		}
		{
			p.SetState(161)
			p.Match(FaRlParserAS)
		}
		{
			p.SetState(162)
			p.Match(FaRlParserCOLON)
		}
		{
			p.SetState(163)
			p.Match(FaRlParserNEWLINE)
		}
		p.SetState(167)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaRlParserSEND)|(1<<FaRlParserKEEP)|(1<<FaRlParserSPLIT))) != 0) {
			{
				p.SetState(164)
				p.SendClause()
			}
			{
				p.SetState(165)
				p.Match(FaRlParserNEWLINE)
			}

			p.SetState(169)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
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

type ReserveContext struct {
	*StatementContext
}

func NewReserveContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReserveContext {
	var p = new(ReserveContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReserveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReserveContext) RESERVE() antlr.TerminalNode {
	return s.GetToken(FaRlParserRESERVE, 0)
}

func (s *ReserveContext) Monetary() IMonetaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryContext)
}

func (s *ReserveContext) IN() antlr.TerminalNode {
	return s.GetToken(FaRlParserIN, 0)
}

func (s *ReserveContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReserveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterReserve(s)
	}
}

func (s *ReserveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitReserve(s)
	}
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

type TransferSimpleContext struct {
	*StatementContext
	amount IExpressionContext
	src    ISourceContext
	dest   IExpressionContext
}

func NewTransferSimpleContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransferSimpleContext {
	var p = new(TransferSimpleContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TransferSimpleContext) GetAmount() IExpressionContext { return s.amount }

func (s *TransferSimpleContext) GetSrc() ISourceContext { return s.src }

func (s *TransferSimpleContext) GetDest() IExpressionContext { return s.dest }

func (s *TransferSimpleContext) SetAmount(v IExpressionContext) { s.amount = v }

func (s *TransferSimpleContext) SetSrc(v ISourceContext) { s.src = v }

func (s *TransferSimpleContext) SetDest(v IExpressionContext) { s.dest = v }

func (s *TransferSimpleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferSimpleContext) TRANSFER() antlr.TerminalNode {
	return s.GetToken(FaRlParserTRANSFER, 0)
}

func (s *TransferSimpleContext) TO() antlr.TerminalNode {
	return s.GetToken(FaRlParserTO, 0)
}

func (s *TransferSimpleContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TransferSimpleContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TransferSimpleContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *TransferSimpleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterTransferSimple(s)
	}
}

func (s *TransferSimpleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitTransferSimple(s)
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

type TransferWithDestContext struct {
	*StatementContext
	amount      IExpressionContext
	src         ISourceContext
	_sendClause ISendClauseContext
	sends       []ISendClauseContext
}

func NewTransferWithDestContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransferWithDestContext {
	var p = new(TransferWithDestContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TransferWithDestContext) GetAmount() IExpressionContext { return s.amount }

func (s *TransferWithDestContext) GetSrc() ISourceContext { return s.src }

func (s *TransferWithDestContext) Get_sendClause() ISendClauseContext { return s._sendClause }

func (s *TransferWithDestContext) SetAmount(v IExpressionContext) { s.amount = v }

func (s *TransferWithDestContext) SetSrc(v ISourceContext) { s.src = v }

func (s *TransferWithDestContext) Set_sendClause(v ISendClauseContext) { s._sendClause = v }

func (s *TransferWithDestContext) GetSends() []ISendClauseContext { return s.sends }

func (s *TransferWithDestContext) SetSends(v []ISendClauseContext) { s.sends = v }

func (s *TransferWithDestContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferWithDestContext) TRANSFER() antlr.TerminalNode {
	return s.GetToken(FaRlParserTRANSFER, 0)
}

func (s *TransferWithDestContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *TransferWithDestContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *TransferWithDestContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TransferWithDestContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *TransferWithDestContext) AllSendClause() []ISendClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISendClauseContext)(nil)).Elem())
	var tst = make([]ISendClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISendClauseContext)
		}
	}

	return tst
}

func (s *TransferWithDestContext) SendClause(i int) ISendClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISendClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISendClauseContext)
}

func (s *TransferWithDestContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterTransferWithDest(s)
	}
}

func (s *TransferWithDestContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitTransferWithDest(s)
	}
}

type TransferAllContext struct {
	*StatementContext
	src  ISourceContext
	dest IExpressionContext
}

func NewTransferAllContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransferAllContext {
	var p = new(TransferAllContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TransferAllContext) GetSrc() ISourceContext { return s.src }

func (s *TransferAllContext) GetDest() IExpressionContext { return s.dest }

func (s *TransferAllContext) SetSrc(v ISourceContext) { s.src = v }

func (s *TransferAllContext) SetDest(v IExpressionContext) { s.dest = v }

func (s *TransferAllContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferAllContext) TRANSFER() antlr.TerminalNode {
	return s.GetToken(FaRlParserTRANSFER, 0)
}

func (s *TransferAllContext) ALL() antlr.TerminalNode {
	return s.GetToken(FaRlParserALL, 0)
}

func (s *TransferAllContext) Monetary() IMonetaryContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonetaryContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonetaryContext)
}

func (s *TransferAllContext) TO() antlr.TerminalNode {
	return s.GetToken(FaRlParserTO, 0)
}

func (s *TransferAllContext) Source() ISourceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISourceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISourceContext)
}

func (s *TransferAllContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TransferAllContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterTransferAll(s)
	}
}

func (s *TransferAllContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitTransferAll(s)
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
	p.EnterRule(localctx, 20, FaRlParserRULE_statement)
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

	p.SetState(235)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(173)
			p.Match(FaRlParserPRINT)
		}
		{
			p.SetState(174)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case 2:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(175)
			p.Match(FaRlParserFAIL)
		}

	case 3:
		localctx = NewTransferSimpleContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(176)
			p.Match(FaRlParserTRANSFER)
		}
		{
			p.SetState(177)

			var _x = p.expression(0)

			localctx.(*TransferSimpleContext).amount = _x
		}
		{
			p.SetState(178)

			var _x = p.source(0)

			localctx.(*TransferSimpleContext).src = _x
		}
		{
			p.SetState(179)
			p.Match(FaRlParserTO)
		}
		{
			p.SetState(180)

			var _x = p.expression(0)

			localctx.(*TransferSimpleContext).dest = _x
		}

	case 4:
		localctx = NewTransferWithDestContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(182)
			p.Match(FaRlParserTRANSFER)
		}
		{
			p.SetState(183)

			var _x = p.expression(0)

			localctx.(*TransferWithDestContext).amount = _x
		}
		{
			p.SetState(184)

			var _x = p.source(0)

			localctx.(*TransferWithDestContext).src = _x
		}
		{
			p.SetState(185)
			p.Match(FaRlParserNEWLINE)
		}
		p.SetState(189)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaRlParserSEND)|(1<<FaRlParserKEEP)|(1<<FaRlParserSPLIT))) != 0) {
			{
				p.SetState(186)

				var _x = p.SendClause()

				localctx.(*TransferWithDestContext)._sendClause = _x
			}
			localctx.(*TransferWithDestContext).sends = append(localctx.(*TransferWithDestContext).sends, localctx.(*TransferWithDestContext)._sendClause)
			{
				p.SetState(187)
				p.Match(FaRlParserNEWLINE)
			}

			p.SetState(191)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 5:
		localctx = NewTransferAllContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(193)
			p.Match(FaRlParserTRANSFER)
		}
		{
			p.SetState(194)
			p.Match(FaRlParserALL)
		}
		{
			p.SetState(195)
			p.Monetary()
		}
		{
			p.SetState(196)

			var _x = p.source(0)

			localctx.(*TransferAllContext).src = _x
		}
		{
			p.SetState(197)
			p.Match(FaRlParserTO)
		}
		{
			p.SetState(198)

			var _x = p.expression(0)

			localctx.(*TransferAllContext).dest = _x
		}

	case 6:
		localctx = NewReserveContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(200)
			p.Match(FaRlParserRESERVE)
		}
		{
			p.SetState(201)
			p.Monetary()
		}
		{
			p.SetState(202)
			p.Match(FaRlParserIN)
		}
		{
			p.SetState(203)
			p.expression(0)
		}

	case 7:
		localctx = NewSetTxMetaContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(205)
			p.Match(FaRlParserSET)
		}
		{
			p.SetState(206)
			p.Match(FaRlParserTRANSACTION)
		}
		{
			p.SetState(207)
			p.Match(FaRlParserMETADATA)
		}
		{
			p.SetState(208)
			p.Match(FaRlParserSTRING)
		}
		{
			p.SetState(209)
			p.Match(FaRlParserEQ)
		}
		{
			p.SetState(210)
			p.MetadataValue()
		}

	case 8:
		localctx = NewSetTxMetaBlockContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(211)
			p.Match(FaRlParserSET)
		}
		{
			p.SetState(212)
			p.Match(FaRlParserTRANSACTION)
		}
		{
			p.SetState(213)
			p.Match(FaRlParserMETADATA)
		}
		{
			p.SetState(214)
			p.Match(FaRlParserLBRACE)
		}
		{
			p.SetState(215)
			p.Match(FaRlParserNEWLINE)
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FaRlParserSTRING {
			{
				p.SetState(216)
				p.MetadataEntry()
			}
			{
				p.SetState(217)
				p.Match(FaRlParserNEWLINE)
			}

			p.SetState(221)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(223)
			p.Match(FaRlParserRBRACE)
		}

	case 9:
		localctx = NewSetAccountMetaContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(225)
			p.Match(FaRlParserSET)
		}
		{
			p.SetState(226)
			p.Match(FaRlParserTY_ACCOUNT)
		}
		{
			p.SetState(227)
			p.Match(FaRlParserMETADATA)
		}
		{
			p.SetState(228)
			p.Match(FaRlParserOF)
		}
		{
			p.SetState(229)
			p.expression(0)
		}
		{
			p.SetState(230)
			p.Match(FaRlParserKEY)
		}
		{
			p.SetState(231)
			p.Match(FaRlParserSTRING)
		}
		{
			p.SetState(232)
			p.Match(FaRlParserEQ)
		}
		{
			p.SetState(233)
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
	p.EnterRule(localctx, 22, FaRlParserRULE_script)
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
	p.SetState(240)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaRlParserNEWLINE {
		{
			p.SetState(237)
			p.Match(FaRlParserNEWLINE)
		}

		p.SetState(242)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(251)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == FaRlParserVAR {
		{
			p.SetState(243)
			p.VarDecl()
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == FaRlParserNEWLINE {
			{
				p.SetState(244)
				p.Match(FaRlParserNEWLINE)
			}

			p.SetState(247)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

		p.SetState(253)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(254)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(263)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(256)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			for ok := true; ok; ok = _la == FaRlParserNEWLINE {
				{
					p.SetState(255)
					p.Match(FaRlParserNEWLINE)
				}

				p.SetState(258)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)
			}
			{
				p.SetState(260)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext())
	}
	p.SetState(267)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaRlParserNEWLINE {
		{
			p.SetState(266)
			p.Match(FaRlParserNEWLINE)
		}

	}
	{
		p.SetState(269)
		p.Match(FaRlParserEOF)
	}

	return localctx
}

func (p *FaRlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 2:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 4:
		var t *SourceContext = nil
		if localctx != nil {
			t = localctx.(*SourceContext)
		}
		return p.Source_Sempred(t, predIndex)

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

func (p *FaRlParser) Source_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
