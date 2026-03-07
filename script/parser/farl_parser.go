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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 28, 92, 4,
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7, 4,
	8, 9, 8, 4, 9, 9, 9, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 3, 3, 3,
	3, 3, 3, 3, 3, 5, 3, 30, 10, 3, 3, 4, 3, 4, 3, 4, 5, 4, 35, 10, 4, 3, 4,
	3, 4, 3, 4, 7, 4, 40, 10, 4, 12, 4, 14, 4, 43, 11, 4, 3, 5, 3, 5, 3, 6,
	3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 7, 7, 56, 10, 7, 12, 7,
	14, 7, 59, 11, 7, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 74, 10, 8, 3, 9, 5, 9, 77, 10, 9, 3, 9, 3,
	9, 3, 9, 7, 9, 82, 10, 9, 12, 9, 14, 9, 85, 11, 9, 3, 9, 5, 9, 88, 10,
	9, 3, 9, 3, 9, 3, 9, 2, 3, 6, 10, 2, 4, 6, 8, 10, 12, 14, 16, 2, 4, 3,
	2, 15, 16, 3, 2, 11, 14, 2, 94, 2, 18, 3, 2, 2, 2, 4, 29, 3, 2, 2, 2, 6,
	34, 3, 2, 2, 2, 8, 44, 3, 2, 2, 2, 10, 46, 3, 2, 2, 2, 12, 49, 3, 2, 2,
	2, 14, 73, 3, 2, 2, 2, 16, 76, 3, 2, 2, 2, 18, 19, 7, 17, 2, 2, 19, 20,
	7, 25, 2, 2, 20, 21, 7, 21, 2, 2, 21, 22, 7, 26, 2, 2, 22, 23, 7, 27, 2,
	2, 23, 24, 7, 18, 2, 2, 24, 3, 3, 2, 2, 2, 25, 30, 7, 24, 2, 2, 26, 30,
	7, 25, 2, 2, 27, 30, 7, 27, 2, 2, 28, 30, 5, 2, 2, 2, 29, 25, 3, 2, 2,
	2, 29, 26, 3, 2, 2, 2, 29, 27, 3, 2, 2, 2, 29, 28, 3, 2, 2, 2, 30, 5, 3,
	2, 2, 2, 31, 32, 8, 4, 1, 2, 32, 35, 5, 4, 3, 2, 33, 35, 7, 23, 2, 2, 34,
	31, 3, 2, 2, 2, 34, 33, 3, 2, 2, 2, 35, 41, 3, 2, 2, 2, 36, 37, 12, 5,
	2, 2, 37, 38, 9, 2, 2, 2, 38, 40, 5, 6, 4, 6, 39, 36, 3, 2, 2, 2, 40, 43,
	3, 2, 2, 2, 41, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2, 42, 7, 3, 2, 2, 2,
	43, 41, 3, 2, 2, 2, 44, 45, 9, 3, 2, 2, 45, 9, 3, 2, 2, 2, 46, 47, 5, 8,
	5, 2, 47, 48, 7, 23, 2, 2, 48, 11, 3, 2, 2, 2, 49, 50, 7, 10, 2, 2, 50,
	51, 7, 19, 2, 2, 51, 57, 7, 3, 2, 2, 52, 53, 5, 10, 6, 2, 53, 54, 7, 3,
	2, 2, 54, 56, 3, 2, 2, 2, 55, 52, 3, 2, 2, 2, 56, 59, 3, 2, 2, 2, 57, 55,
	3, 2, 2, 2, 57, 58, 3, 2, 2, 2, 58, 60, 3, 2, 2, 2, 59, 57, 3, 2, 2, 2,
	60, 61, 7, 20, 2, 2, 61, 62, 7, 3, 2, 2, 62, 13, 3, 2, 2, 2, 63, 64, 7,
	5, 2, 2, 64, 74, 5, 6, 4, 2, 65, 74, 7, 6, 2, 2, 66, 67, 7, 7, 2, 2, 67,
	68, 5, 6, 4, 2, 68, 69, 7, 8, 2, 2, 69, 70, 5, 6, 4, 2, 70, 71, 7, 9, 2,
	2, 71, 72, 5, 6, 4, 2, 72, 74, 3, 2, 2, 2, 73, 63, 3, 2, 2, 2, 73, 65,
	3, 2, 2, 2, 73, 66, 3, 2, 2, 2, 74, 15, 3, 2, 2, 2, 75, 77, 5, 12, 7, 2,
	76, 75, 3, 2, 2, 2, 76, 77, 3, 2, 2, 2, 77, 78, 3, 2, 2, 2, 78, 83, 5,
	14, 8, 2, 79, 80, 7, 3, 2, 2, 80, 82, 5, 14, 8, 2, 81, 79, 3, 2, 2, 2,
	82, 85, 3, 2, 2, 2, 83, 81, 3, 2, 2, 2, 83, 84, 3, 2, 2, 2, 84, 87, 3,
	2, 2, 2, 85, 83, 3, 2, 2, 2, 86, 88, 7, 3, 2, 2, 87, 86, 3, 2, 2, 2, 87,
	88, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2, 89, 90, 7, 2, 2, 3, 90, 17, 3, 2, 2,
	2, 10, 29, 34, 41, 57, 73, 76, 83, 87,
}
var literalNames = []string{
	"", "", "", "'print'", "'fail'", "'transfer'", "'from'", "'to'", "'vars'",
	"'account'", "'asset'", "'number'", "'monetary'", "'+'", "'-'", "'['",
	"']'", "'{'", "'}'", "'.'", "','",
}
var symbolicNames = []string{
	"", "NEWLINE", "WHITESPACE", "PRINT", "FAIL", "TRANSFER", "FROM", "TO",
	"VARS", "TY_ACCOUNT", "TY_ASSET", "TY_NUMBER", "TY_MONETARY", "OP_ADD",
	"OP_SUB", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "DOT", "COMMA", "VARIABLE_NAME",
	"ACCOUNT", "ASSET", "PRECISION", "NUMBER", "IDENTIFIER",
}

var ruleNames = []string{
	"monetary", "literal", "expression", "type_", "varDecl", "varListDecl",
	"statement", "script",
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
	FaRlParserVARS          = 8
	FaRlParserTY_ACCOUNT    = 9
	FaRlParserTY_ASSET      = 10
	FaRlParserTY_NUMBER     = 11
	FaRlParserTY_MONETARY   = 12
	FaRlParserOP_ADD        = 13
	FaRlParserOP_SUB        = 14
	FaRlParserLBRACK        = 15
	FaRlParserRBRACK        = 16
	FaRlParserLBRACE        = 17
	FaRlParserRBRACE        = 18
	FaRlParserDOT           = 19
	FaRlParserCOMMA         = 20
	FaRlParserVARIABLE_NAME = 21
	FaRlParserACCOUNT       = 22
	FaRlParserASSET         = 23
	FaRlParserPRECISION     = 24
	FaRlParserNUMBER        = 25
	FaRlParserIDENTIFIER    = 26
)

// FaRlParser rules.
const (
	FaRlParserRULE_monetary    = 0
	FaRlParserRULE_literal     = 1
	FaRlParserRULE_expression  = 2
	FaRlParserRULE_type_       = 3
	FaRlParserRULE_varDecl     = 4
	FaRlParserRULE_varListDecl = 5
	FaRlParserRULE_statement   = 6
	FaRlParserRULE_script      = 7
)

// IMonetaryContext is an interface to support dynamic dispatch.
type IMonetaryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAsset returns the asset token.
	GetAsset() antlr.Token

	// GetPrecision returns the precision token.
	GetPrecision() antlr.Token

	// GetAmount returns the amount token.
	GetAmount() antlr.Token

	// SetAsset sets the asset token.
	SetAsset(antlr.Token)

	// SetPrecision sets the precision token.
	SetPrecision(antlr.Token)

	// SetAmount sets the amount token.
	SetAmount(antlr.Token)

	// IsMonetaryContext differentiates from other interfaces.
	IsMonetaryContext()
}

type MonetaryContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	asset     antlr.Token
	precision antlr.Token
	amount    antlr.Token
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

func (s *MonetaryContext) GetAsset() antlr.Token { return s.asset }

func (s *MonetaryContext) GetPrecision() antlr.Token { return s.precision }

func (s *MonetaryContext) GetAmount() antlr.Token { return s.amount }

func (s *MonetaryContext) SetAsset(v antlr.Token) { s.asset = v }

func (s *MonetaryContext) SetPrecision(v antlr.Token) { s.precision = v }

func (s *MonetaryContext) SetAmount(v antlr.Token) { s.amount = v }

func (s *MonetaryContext) LBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACK, 0)
}

func (s *MonetaryContext) DOT() antlr.TerminalNode {
	return s.GetToken(FaRlParserDOT, 0)
}

func (s *MonetaryContext) RBRACK() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACK, 0)
}

func (s *MonetaryContext) ASSET() antlr.TerminalNode {
	return s.GetToken(FaRlParserASSET, 0)
}

func (s *MonetaryContext) PRECISION() antlr.TerminalNode {
	return s.GetToken(FaRlParserPRECISION, 0)
}

func (s *MonetaryContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(FaRlParserNUMBER, 0)
}

func (s *MonetaryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonetaryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonetaryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterMonetary(s)
	}
}

func (s *MonetaryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitMonetary(s)
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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.Match(FaRlParserLBRACK)
	}
	{
		p.SetState(17)

		var _m = p.Match(FaRlParserASSET)

		localctx.(*MonetaryContext).asset = _m
	}
	{
		p.SetState(18)
		p.Match(FaRlParserDOT)
	}
	{
		p.SetState(19)

		var _m = p.Match(FaRlParserPRECISION)

		localctx.(*MonetaryContext).precision = _m
	}
	{
		p.SetState(20)

		var _m = p.Match(FaRlParserNUMBER)

		localctx.(*MonetaryContext).amount = _m
	}
	{
		p.SetState(21)
		p.Match(FaRlParserRBRACK)
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

	p.SetState(27)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserACCOUNT:
		localctx = NewLitAccountContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(23)
			p.Match(FaRlParserACCOUNT)
		}

	case FaRlParserASSET:
		localctx = NewLitAssetContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(24)
			p.Match(FaRlParserASSET)
		}

	case FaRlParserNUMBER:
		localctx = NewLitNumberContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(25)
			p.Match(FaRlParserNUMBER)
		}

	case FaRlParserLBRACK:
		localctx = NewLitMonetaryContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(26)
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
	p.SetState(32)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserLBRACK, FaRlParserACCOUNT, FaRlParserASSET, FaRlParserNUMBER:
		localctx = NewExprLiteralContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(30)

			var _x = p.Literal()

			localctx.(*ExprLiteralContext).lit = _x
		}

	case FaRlParserVARIABLE_NAME:
		localctx = NewExprVariableContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)

			var _m = p.Match(FaRlParserVARIABLE_NAME)

			localctx.(*ExprVariableContext).variable = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(39)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExprAddSubContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*ExprAddSubContext).lhs = _prevctx

			p.PushNewRecursionContext(localctx, _startState, FaRlParserRULE_expression)
			p.SetState(34)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
			}
			{
				p.SetState(35)

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
				p.SetState(36)

				var _x = p.expression(4)

				localctx.(*ExprAddSubContext).rhs = _x
			}

		}
		p.SetState(41)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 6, FaRlParserRULE_type_)
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
		p.SetState(42)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaRlParserTY_ACCOUNT)|(1<<FaRlParserTY_ASSET)|(1<<FaRlParserTY_NUMBER)|(1<<FaRlParserTY_MONETARY))) != 0) {
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

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTy returns the ty rule contexts.
	GetTy() IType_Context

	// SetTy sets the ty rule contexts.
	SetTy(IType_Context)

	// IsVarDeclContext differentiates from other interfaces.
	IsVarDeclContext()
}

type VarDeclContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	ty     IType_Context
	name   antlr.Token
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

func (s *VarDeclContext) GetName() antlr.Token { return s.name }

func (s *VarDeclContext) SetName(v antlr.Token) { s.name = v }

func (s *VarDeclContext) GetTy() IType_Context { return s.ty }

func (s *VarDeclContext) SetTy(v IType_Context) { s.ty = v }

func (s *VarDeclContext) Type_() IType_Context {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IType_Context)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IType_Context)
}

func (s *VarDeclContext) VARIABLE_NAME() antlr.TerminalNode {
	return s.GetToken(FaRlParserVARIABLE_NAME, 0)
}

func (s *VarDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterVarDecl(s)
	}
}

func (s *VarDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitVarDecl(s)
	}
}

func (p *FaRlParser) VarDecl() (localctx IVarDeclContext) {
	localctx = NewVarDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, FaRlParserRULE_varDecl)

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
		p.SetState(44)

		var _x = p.Type_()

		localctx.(*VarDeclContext).ty = _x
	}
	{
		p.SetState(45)

		var _m = p.Match(FaRlParserVARIABLE_NAME)

		localctx.(*VarDeclContext).name = _m
	}

	return localctx
}

// IVarListDeclContext is an interface to support dynamic dispatch.
type IVarListDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_varDecl returns the _varDecl rule contexts.
	Get_varDecl() IVarDeclContext

	// Set_varDecl sets the _varDecl rule contexts.
	Set_varDecl(IVarDeclContext)

	// GetV returns the v rule context list.
	GetV() []IVarDeclContext

	// SetV sets the v rule context list.
	SetV([]IVarDeclContext)

	// IsVarListDeclContext differentiates from other interfaces.
	IsVarListDeclContext()
}

type VarListDeclContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	_varDecl IVarDeclContext
	v        []IVarDeclContext
}

func NewEmptyVarListDeclContext() *VarListDeclContext {
	var p = new(VarListDeclContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = FaRlParserRULE_varListDecl
	return p
}

func (*VarListDeclContext) IsVarListDeclContext() {}

func NewVarListDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarListDeclContext {
	var p = new(VarListDeclContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = FaRlParserRULE_varListDecl

	return p
}

func (s *VarListDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *VarListDeclContext) Get_varDecl() IVarDeclContext { return s._varDecl }

func (s *VarListDeclContext) Set_varDecl(v IVarDeclContext) { s._varDecl = v }

func (s *VarListDeclContext) GetV() []IVarDeclContext { return s.v }

func (s *VarListDeclContext) SetV(v []IVarDeclContext) { s.v = v }

func (s *VarListDeclContext) VARS() antlr.TerminalNode {
	return s.GetToken(FaRlParserVARS, 0)
}

func (s *VarListDeclContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserLBRACE, 0)
}

func (s *VarListDeclContext) AllNEWLINE() []antlr.TerminalNode {
	return s.GetTokens(FaRlParserNEWLINE)
}

func (s *VarListDeclContext) NEWLINE(i int) antlr.TerminalNode {
	return s.GetToken(FaRlParserNEWLINE, i)
}

func (s *VarListDeclContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(FaRlParserRBRACE, 0)
}

func (s *VarListDeclContext) AllVarDecl() []IVarDeclContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVarDeclContext)(nil)).Elem())
	var tst = make([]IVarDeclContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVarDeclContext)
		}
	}

	return tst
}

func (s *VarListDeclContext) VarDecl(i int) IVarDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarDeclContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVarDeclContext)
}

func (s *VarListDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarListDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarListDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.EnterVarListDecl(s)
	}
}

func (s *VarListDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(FaRlListener); ok {
		listenerT.ExitVarListDecl(s)
	}
}

func (p *FaRlParser) VarListDecl() (localctx IVarListDeclContext) {
	localctx = NewVarListDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, FaRlParserRULE_varListDecl)
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
		p.SetState(47)
		p.Match(FaRlParserVARS)
	}
	{
		p.SetState(48)
		p.Match(FaRlParserLBRACE)
	}
	{
		p.SetState(49)
		p.Match(FaRlParserNEWLINE)
	}
	p.SetState(55)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<FaRlParserTY_ACCOUNT)|(1<<FaRlParserTY_ASSET)|(1<<FaRlParserTY_NUMBER)|(1<<FaRlParserTY_MONETARY))) != 0 {
		{
			p.SetState(50)

			var _x = p.VarDecl()

			localctx.(*VarListDeclContext)._varDecl = _x
		}
		localctx.(*VarListDeclContext).v = append(localctx.(*VarListDeclContext).v, localctx.(*VarListDeclContext)._varDecl)
		{
			p.SetState(51)
			p.Match(FaRlParserNEWLINE)
		}

		p.SetState(57)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(58)
		p.Match(FaRlParserRBRACE)
	}
	{
		p.SetState(59)
		p.Match(FaRlParserNEWLINE)
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
	amount IExpressionContext
	source IExpressionContext
	dest   IExpressionContext
}

func NewTransferContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TransferContext {
	var p = new(TransferContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *TransferContext) GetAmount() IExpressionContext { return s.amount }

func (s *TransferContext) GetSource() IExpressionContext { return s.source }

func (s *TransferContext) GetDest() IExpressionContext { return s.dest }

func (s *TransferContext) SetAmount(v IExpressionContext) { s.amount = v }

func (s *TransferContext) SetSource(v IExpressionContext) { s.source = v }

func (s *TransferContext) SetDest(v IExpressionContext) { s.dest = v }

func (s *TransferContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TransferContext) TRANSFER() antlr.TerminalNode {
	return s.GetToken(FaRlParserTRANSFER, 0)
}

func (s *TransferContext) FROM() antlr.TerminalNode {
	return s.GetToken(FaRlParserFROM, 0)
}

func (s *TransferContext) TO() antlr.TerminalNode {
	return s.GetToken(FaRlParserTO, 0)
}

func (s *TransferContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TransferContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
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

func (p *FaRlParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, FaRlParserRULE_statement)

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

	p.SetState(71)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case FaRlParserPRINT:
		localctx = NewPrintContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(61)
			p.Match(FaRlParserPRINT)
		}
		{
			p.SetState(62)

			var _x = p.expression(0)

			localctx.(*PrintContext).expr = _x
		}

	case FaRlParserFAIL:
		localctx = NewFailContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(63)
			p.Match(FaRlParserFAIL)
		}

	case FaRlParserTRANSFER:
		localctx = NewTransferContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(64)
			p.Match(FaRlParserTRANSFER)
		}
		{
			p.SetState(65)

			var _x = p.expression(0)

			localctx.(*TransferContext).amount = _x
		}
		{
			p.SetState(66)
			p.Match(FaRlParserFROM)
		}
		{
			p.SetState(67)

			var _x = p.expression(0)

			localctx.(*TransferContext).source = _x
		}
		{
			p.SetState(68)
			p.Match(FaRlParserTO)
		}
		{
			p.SetState(69)

			var _x = p.expression(0)

			localctx.(*TransferContext).dest = _x
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IScriptContext is an interface to support dynamic dispatch.
type IScriptContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetVars returns the vars rule contexts.
	GetVars() IVarListDeclContext

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// SetVars sets the vars rule contexts.
	SetVars(IVarListDeclContext)

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
	vars       IVarListDeclContext
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

func (s *ScriptContext) GetVars() IVarListDeclContext { return s.vars }

func (s *ScriptContext) Get_statement() IStatementContext { return s._statement }

func (s *ScriptContext) SetVars(v IVarListDeclContext) { s.vars = v }

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

func (s *ScriptContext) VarListDecl() IVarListDeclContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVarListDeclContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVarListDeclContext)
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
	p.EnterRule(localctx, 14, FaRlParserRULE_script)
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
	p.SetState(74)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaRlParserVARS {
		{
			p.SetState(73)

			var _x = p.VarListDecl()

			localctx.(*ScriptContext).vars = _x
		}

	}
	{
		p.SetState(76)

		var _x = p.Statement()

		localctx.(*ScriptContext)._statement = _x
	}
	localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)
	p.SetState(81)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(77)
				p.Match(FaRlParserNEWLINE)
			}
			{
				p.SetState(78)

				var _x = p.Statement()

				localctx.(*ScriptContext)._statement = _x
			}
			localctx.(*ScriptContext).stmts = append(localctx.(*ScriptContext).stmts, localctx.(*ScriptContext)._statement)

		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext())
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == FaRlParserNEWLINE {
		{
			p.SetState(84)
			p.Match(FaRlParserNEWLINE)
		}

	}
	{
		p.SetState(87)
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
