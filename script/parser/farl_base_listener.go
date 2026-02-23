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

// EnterMonetary is called when production monetary is entered.
func (s *BaseFaRlListener) EnterMonetary(ctx *MonetaryContext) {}

// ExitMonetary is called when production monetary is exited.
func (s *BaseFaRlListener) ExitMonetary(ctx *MonetaryContext) {}

// EnterLitAccount is called when production LitAccount is entered.
func (s *BaseFaRlListener) EnterLitAccount(ctx *LitAccountContext) {}

// ExitLitAccount is called when production LitAccount is exited.
func (s *BaseFaRlListener) ExitLitAccount(ctx *LitAccountContext) {}

// EnterLitAsset is called when production LitAsset is entered.
func (s *BaseFaRlListener) EnterLitAsset(ctx *LitAssetContext) {}

// ExitLitAsset is called when production LitAsset is exited.
func (s *BaseFaRlListener) ExitLitAsset(ctx *LitAssetContext) {}

// EnterLitNumber is called when production LitNumber is entered.
func (s *BaseFaRlListener) EnterLitNumber(ctx *LitNumberContext) {}

// ExitLitNumber is called when production LitNumber is exited.
func (s *BaseFaRlListener) ExitLitNumber(ctx *LitNumberContext) {}

// EnterLitMonetary is called when production LitMonetary is entered.
func (s *BaseFaRlListener) EnterLitMonetary(ctx *LitMonetaryContext) {}

// ExitLitMonetary is called when production LitMonetary is exited.
func (s *BaseFaRlListener) ExitLitMonetary(ctx *LitMonetaryContext) {}

// EnterExprAddSub is called when production ExprAddSub is entered.
func (s *BaseFaRlListener) EnterExprAddSub(ctx *ExprAddSubContext) {}

// ExitExprAddSub is called when production ExprAddSub is exited.
func (s *BaseFaRlListener) ExitExprAddSub(ctx *ExprAddSubContext) {}

// EnterExprLiteral is called when production ExprLiteral is entered.
func (s *BaseFaRlListener) EnterExprLiteral(ctx *ExprLiteralContext) {}

// ExitExprLiteral is called when production ExprLiteral is exited.
func (s *BaseFaRlListener) ExitExprLiteral(ctx *ExprLiteralContext) {}

// EnterExprVariable is called when production ExprVariable is entered.
func (s *BaseFaRlListener) EnterExprVariable(ctx *ExprVariableContext) {}

// ExitExprVariable is called when production ExprVariable is exited.
func (s *BaseFaRlListener) ExitExprVariable(ctx *ExprVariableContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseFaRlListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseFaRlListener) ExitType_(ctx *Type_Context) {}

// EnterVarDecl is called when production varDecl is entered.
func (s *BaseFaRlListener) EnterVarDecl(ctx *VarDeclContext) {}

// ExitVarDecl is called when production varDecl is exited.
func (s *BaseFaRlListener) ExitVarDecl(ctx *VarDeclContext) {}

// EnterVarListDecl is called when production varListDecl is entered.
func (s *BaseFaRlListener) EnterVarListDecl(ctx *VarListDeclContext) {}

// ExitVarListDecl is called when production varListDecl is exited.
func (s *BaseFaRlListener) ExitVarListDecl(ctx *VarListDeclContext) {}

// EnterPrint is called when production Print is entered.
func (s *BaseFaRlListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production Print is exited.
func (s *BaseFaRlListener) ExitPrint(ctx *PrintContext) {}

// EnterFail is called when production Fail is entered.
func (s *BaseFaRlListener) EnterFail(ctx *FailContext) {}

// ExitFail is called when production Fail is exited.
func (s *BaseFaRlListener) ExitFail(ctx *FailContext) {}

// EnterTransfer is called when production Transfer is entered.
func (s *BaseFaRlListener) EnterTransfer(ctx *TransferContext) {}

// ExitTransfer is called when production Transfer is exited.
func (s *BaseFaRlListener) ExitTransfer(ctx *TransferContext) {}

// EnterScript is called when production script is entered.
func (s *BaseFaRlListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseFaRlListener) ExitScript(ctx *ScriptContext) {}
