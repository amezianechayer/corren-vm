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

// EnterMonetaryLit is called when production MonetaryLit is entered.
func (s *BaseFaRlListener) EnterMonetaryLit(ctx *MonetaryLitContext) {}

// ExitMonetaryLit is called when production MonetaryLit is exited.
func (s *BaseFaRlListener) ExitMonetaryLit(ctx *MonetaryLitContext) {}

// EnterMonetaryAll is called when production MonetaryAll is entered.
func (s *BaseFaRlListener) EnterMonetaryAll(ctx *MonetaryAllContext) {}

// ExitMonetaryAll is called when production MonetaryAll is exited.
func (s *BaseFaRlListener) ExitMonetaryAll(ctx *MonetaryAllContext) {}

// EnterMonetaryNoPrecision is called when production MonetaryNoPrecision is entered.
func (s *BaseFaRlListener) EnterMonetaryNoPrecision(ctx *MonetaryNoPrecisionContext) {}

// ExitMonetaryNoPrecision is called when production MonetaryNoPrecision is exited.
func (s *BaseFaRlListener) ExitMonetaryNoPrecision(ctx *MonetaryNoPrecisionContext) {}

// EnterMonetaryNoPrecisionAll is called when production MonetaryNoPrecisionAll is entered.
func (s *BaseFaRlListener) EnterMonetaryNoPrecisionAll(ctx *MonetaryNoPrecisionAllContext) {}

// ExitMonetaryNoPrecisionAll is called when production MonetaryNoPrecisionAll is exited.
func (s *BaseFaRlListener) ExitMonetaryNoPrecisionAll(ctx *MonetaryNoPrecisionAllContext) {}

// EnterMonetaryAssetOnly is called when production MonetaryAssetOnly is entered.
func (s *BaseFaRlListener) EnterMonetaryAssetOnly(ctx *MonetaryAssetOnlyContext) {}

// ExitMonetaryAssetOnly is called when production MonetaryAssetOnly is exited.
func (s *BaseFaRlListener) ExitMonetaryAssetOnly(ctx *MonetaryAssetOnlyContext) {}

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

// EnterPortionPercent is called when production PortionPercent is entered.
func (s *BaseFaRlListener) EnterPortionPercent(ctx *PortionPercentContext) {}

// ExitPortionPercent is called when production PortionPercent is exited.
func (s *BaseFaRlListener) ExitPortionPercent(ctx *PortionPercentContext) {}

// EnterPortionRatio is called when production PortionRatio is entered.
func (s *BaseFaRlListener) EnterPortionRatio(ctx *PortionRatioContext) {}

// ExitPortionRatio is called when production PortionRatio is exited.
func (s *BaseFaRlListener) ExitPortionRatio(ctx *PortionRatioContext) {}

// EnterPortionRemaining is called when production PortionRemaining is entered.
func (s *BaseFaRlListener) EnterPortionRemaining(ctx *PortionRemainingContext) {}

// ExitPortionRemaining is called when production PortionRemaining is exited.
func (s *BaseFaRlListener) ExitPortionRemaining(ctx *PortionRemainingContext) {}

// EnterPortionVar is called when production PortionVar is entered.
func (s *BaseFaRlListener) EnterPortionVar(ctx *PortionVarContext) {}

// ExitPortionVar is called when production PortionVar is exited.
func (s *BaseFaRlListener) ExitPortionVar(ctx *PortionVarContext) {}

// EnterSrcRemaining is called when production SrcRemaining is entered.
func (s *BaseFaRlListener) EnterSrcRemaining(ctx *SrcRemainingContext) {}

// ExitSrcRemaining is called when production SrcRemaining is exited.
func (s *BaseFaRlListener) ExitSrcRemaining(ctx *SrcRemainingContext) {}

// EnterSrcPercentLimit is called when production SrcPercentLimit is entered.
func (s *BaseFaRlListener) EnterSrcPercentLimit(ctx *SrcPercentLimitContext) {}

// ExitSrcPercentLimit is called when production SrcPercentLimit is exited.
func (s *BaseFaRlListener) ExitSrcPercentLimit(ctx *SrcPercentLimitContext) {}

// EnterSrcCascade is called when production SrcCascade is entered.
func (s *BaseFaRlListener) EnterSrcCascade(ctx *SrcCascadeContext) {}

// ExitSrcCascade is called when production SrcCascade is exited.
func (s *BaseFaRlListener) ExitSrcCascade(ctx *SrcCascadeContext) {}

// EnterSrcSimple is called when production SrcSimple is entered.
func (s *BaseFaRlListener) EnterSrcSimple(ctx *SrcSimpleContext) {}

// ExitSrcSimple is called when production SrcSimple is exited.
func (s *BaseFaRlListener) ExitSrcSimple(ctx *SrcSimpleContext) {}

// EnterSrcOverdraft is called when production SrcOverdraft is entered.
func (s *BaseFaRlListener) EnterSrcOverdraft(ctx *SrcOverdraftContext) {}

// ExitSrcOverdraft is called when production SrcOverdraft is exited.
func (s *BaseFaRlListener) ExitSrcOverdraft(ctx *SrcOverdraftContext) {}

// EnterSrcOverdraftCapped is called when production SrcOverdraftCapped is entered.
func (s *BaseFaRlListener) EnterSrcOverdraftCapped(ctx *SrcOverdraftCappedContext) {}

// ExitSrcOverdraftCapped is called when production SrcOverdraftCapped is exited.
func (s *BaseFaRlListener) ExitSrcOverdraftCapped(ctx *SrcOverdraftCappedContext) {}

// EnterSrcLimit is called when production SrcLimit is entered.
func (s *BaseFaRlListener) EnterSrcLimit(ctx *SrcLimitContext) {}

// ExitSrcLimit is called when production SrcLimit is exited.
func (s *BaseFaRlListener) ExitSrcLimit(ctx *SrcLimitContext) {}

// EnterSrcPercent is called when production SrcPercent is entered.
func (s *BaseFaRlListener) EnterSrcPercent(ctx *SrcPercentContext) {}

// ExitSrcPercent is called when production SrcPercent is exited.
func (s *BaseFaRlListener) ExitSrcPercent(ctx *SrcPercentContext) {}

// EnterType_ is called when production type_ is entered.
func (s *BaseFaRlListener) EnterType_(ctx *Type_Context) {}

// ExitType_ is called when production type_ is exited.
func (s *BaseFaRlListener) ExitType_(ctx *Type_Context) {}

// EnterOrigin is called when production origin is entered.
func (s *BaseFaRlListener) EnterOrigin(ctx *OriginContext) {}

// ExitOrigin is called when production origin is exited.
func (s *BaseFaRlListener) ExitOrigin(ctx *OriginContext) {}

// EnterVarTyped is called when production VarTyped is entered.
func (s *BaseFaRlListener) EnterVarTyped(ctx *VarTypedContext) {}

// ExitVarTyped is called when production VarTyped is exited.
func (s *BaseFaRlListener) ExitVarTyped(ctx *VarTypedContext) {}

// EnterMetaValueExpr is called when production MetaValueExpr is entered.
func (s *BaseFaRlListener) EnterMetaValueExpr(ctx *MetaValueExprContext) {}

// ExitMetaValueExpr is called when production MetaValueExpr is exited.
func (s *BaseFaRlListener) ExitMetaValueExpr(ctx *MetaValueExprContext) {}

// EnterMetaValueRatio is called when production MetaValueRatio is entered.
func (s *BaseFaRlListener) EnterMetaValueRatio(ctx *MetaValueRatioContext) {}

// ExitMetaValueRatio is called when production MetaValueRatio is exited.
func (s *BaseFaRlListener) ExitMetaValueRatio(ctx *MetaValueRatioContext) {}

// EnterMetadataEntry is called when production metadataEntry is entered.
func (s *BaseFaRlListener) EnterMetadataEntry(ctx *MetadataEntryContext) {}

// ExitMetadataEntry is called when production metadataEntry is exited.
func (s *BaseFaRlListener) ExitMetadataEntry(ctx *MetadataEntryContext) {}

// EnterSendTo is called when production SendTo is entered.
func (s *BaseFaRlListener) EnterSendTo(ctx *SendToContext) {}

// ExitSendTo is called when production SendTo is exited.
func (s *BaseFaRlListener) ExitSendTo(ctx *SendToContext) {}

// EnterSendKeep is called when production SendKeep is entered.
func (s *BaseFaRlListener) EnterSendKeep(ctx *SendKeepContext) {}

// ExitSendKeep is called when production SendKeep is exited.
func (s *BaseFaRlListener) ExitSendKeep(ctx *SendKeepContext) {}

// EnterSendSplit is called when production SendSplit is entered.
func (s *BaseFaRlListener) EnterSendSplit(ctx *SendSplitContext) {}

// ExitSendSplit is called when production SendSplit is exited.
func (s *BaseFaRlListener) ExitSendSplit(ctx *SendSplitContext) {}

// EnterPrint is called when production Print is entered.
func (s *BaseFaRlListener) EnterPrint(ctx *PrintContext) {}

// ExitPrint is called when production Print is exited.
func (s *BaseFaRlListener) ExitPrint(ctx *PrintContext) {}

// EnterFail is called when production Fail is entered.
func (s *BaseFaRlListener) EnterFail(ctx *FailContext) {}

// ExitFail is called when production Fail is exited.
func (s *BaseFaRlListener) ExitFail(ctx *FailContext) {}

// EnterTransferSimple is called when production TransferSimple is entered.
func (s *BaseFaRlListener) EnterTransferSimple(ctx *TransferSimpleContext) {}

// ExitTransferSimple is called when production TransferSimple is exited.
func (s *BaseFaRlListener) ExitTransferSimple(ctx *TransferSimpleContext) {}

// EnterTransferWithDest is called when production TransferWithDest is entered.
func (s *BaseFaRlListener) EnterTransferWithDest(ctx *TransferWithDestContext) {}

// ExitTransferWithDest is called when production TransferWithDest is exited.
func (s *BaseFaRlListener) ExitTransferWithDest(ctx *TransferWithDestContext) {}

// EnterTransferAll is called when production TransferAll is entered.
func (s *BaseFaRlListener) EnterTransferAll(ctx *TransferAllContext) {}

// ExitTransferAll is called when production TransferAll is exited.
func (s *BaseFaRlListener) ExitTransferAll(ctx *TransferAllContext) {}

// EnterReserve is called when production Reserve is entered.
func (s *BaseFaRlListener) EnterReserve(ctx *ReserveContext) {}

// ExitReserve is called when production Reserve is exited.
func (s *BaseFaRlListener) ExitReserve(ctx *ReserveContext) {}

// EnterSetTxMeta is called when production SetTxMeta is entered.
func (s *BaseFaRlListener) EnterSetTxMeta(ctx *SetTxMetaContext) {}

// ExitSetTxMeta is called when production SetTxMeta is exited.
func (s *BaseFaRlListener) ExitSetTxMeta(ctx *SetTxMetaContext) {}

// EnterSetTxMetaBlock is called when production SetTxMetaBlock is entered.
func (s *BaseFaRlListener) EnterSetTxMetaBlock(ctx *SetTxMetaBlockContext) {}

// ExitSetTxMetaBlock is called when production SetTxMetaBlock is exited.
func (s *BaseFaRlListener) ExitSetTxMetaBlock(ctx *SetTxMetaBlockContext) {}

// EnterSetAccountMeta is called when production SetAccountMeta is entered.
func (s *BaseFaRlListener) EnterSetAccountMeta(ctx *SetAccountMetaContext) {}

// ExitSetAccountMeta is called when production SetAccountMeta is exited.
func (s *BaseFaRlListener) ExitSetAccountMeta(ctx *SetAccountMetaContext) {}

// EnterScript is called when production script is entered.
func (s *BaseFaRlListener) EnterScript(ctx *ScriptContext) {}

// ExitScript is called when production script is exited.
func (s *BaseFaRlListener) ExitScript(ctx *ScriptContext) {}
