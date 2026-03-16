// Code generated from FaRl.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FaRl

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FaRlListener is a complete listener for a parse tree produced by FaRlParser.
type FaRlListener interface {
	antlr.ParseTreeListener

	// EnterMonetaryLit is called when entering the MonetaryLit production.
	EnterMonetaryLit(c *MonetaryLitContext)

	// EnterMonetaryAll is called when entering the MonetaryAll production.
	EnterMonetaryAll(c *MonetaryAllContext)

	// EnterMonetaryAssetOnly is called when entering the MonetaryAssetOnly production.
	EnterMonetaryAssetOnly(c *MonetaryAssetOnlyContext)

	// EnterLitAccount is called when entering the LitAccount production.
	EnterLitAccount(c *LitAccountContext)

	// EnterLitAsset is called when entering the LitAsset production.
	EnterLitAsset(c *LitAssetContext)

	// EnterLitNumber is called when entering the LitNumber production.
	EnterLitNumber(c *LitNumberContext)

	// EnterLitMonetary is called when entering the LitMonetary production.
	EnterLitMonetary(c *LitMonetaryContext)

	// EnterExprAddSub is called when entering the ExprAddSub production.
	EnterExprAddSub(c *ExprAddSubContext)

	// EnterExprLiteral is called when entering the ExprLiteral production.
	EnterExprLiteral(c *ExprLiteralContext)

	// EnterExprVariable is called when entering the ExprVariable production.
	EnterExprVariable(c *ExprVariableContext)

	// EnterPortionPercent is called when entering the PortionPercent production.
	EnterPortionPercent(c *PortionPercentContext)

	// EnterPortionRatio is called when entering the PortionRatio production.
	EnterPortionRatio(c *PortionRatioContext)

	// EnterPortionRemaining is called when entering the PortionRemaining production.
	EnterPortionRemaining(c *PortionRemainingContext)

	// EnterSrcRemaining is called when entering the SrcRemaining production.
	EnterSrcRemaining(c *SrcRemainingContext)

	// EnterSrcPercentLimit is called when entering the SrcPercentLimit production.
	EnterSrcPercentLimit(c *SrcPercentLimitContext)

	// EnterSrcCascade is called when entering the SrcCascade production.
	EnterSrcCascade(c *SrcCascadeContext)

	// EnterSrcSimple is called when entering the SrcSimple production.
	EnterSrcSimple(c *SrcSimpleContext)

	// EnterSrcOverdraft is called when entering the SrcOverdraft production.
	EnterSrcOverdraft(c *SrcOverdraftContext)

	// EnterSrcOverdraftCapped is called when entering the SrcOverdraftCapped production.
	EnterSrcOverdraftCapped(c *SrcOverdraftCappedContext)

	// EnterSrcLimit is called when entering the SrcLimit production.
	EnterSrcLimit(c *SrcLimitContext)

	// EnterSrcPercent is called when entering the SrcPercent production.
	EnterSrcPercent(c *SrcPercentContext)

	// EnterType_ is called when entering the type_ production.
	EnterType_(c *Type_Context)

	// EnterVarTyped is called when entering the VarTyped production.
	EnterVarTyped(c *VarTypedContext)

	// EnterVarBalance is called when entering the VarBalance production.
	EnterVarBalance(c *VarBalanceContext)

	// EnterVarMeta is called when entering the VarMeta production.
	EnterVarMeta(c *VarMetaContext)

	// EnterMetaValueExpr is called when entering the MetaValueExpr production.
	EnterMetaValueExpr(c *MetaValueExprContext)

	// EnterMetaValueRatio is called when entering the MetaValueRatio production.
	EnterMetaValueRatio(c *MetaValueRatioContext)

	// EnterMetadataEntry is called when entering the metadataEntry production.
	EnterMetadataEntry(c *MetadataEntryContext)

	// EnterSendTo is called when entering the SendTo production.
	EnterSendTo(c *SendToContext)

	// EnterSendKeep is called when entering the SendKeep production.
	EnterSendKeep(c *SendKeepContext)

	// EnterSendSplit is called when entering the SendSplit production.
	EnterSendSplit(c *SendSplitContext)

	// EnterPrint is called when entering the Print production.
	EnterPrint(c *PrintContext)

	// EnterFail is called when entering the Fail production.
	EnterFail(c *FailContext)

	// EnterTransferSimple is called when entering the TransferSimple production.
	EnterTransferSimple(c *TransferSimpleContext)

	// EnterTransferWithDest is called when entering the TransferWithDest production.
	EnterTransferWithDest(c *TransferWithDestContext)

	// EnterTransferAll is called when entering the TransferAll production.
	EnterTransferAll(c *TransferAllContext)

	// EnterReserve is called when entering the Reserve production.
	EnterReserve(c *ReserveContext)

	// EnterSetTxMeta is called when entering the SetTxMeta production.
	EnterSetTxMeta(c *SetTxMetaContext)

	// EnterSetTxMetaBlock is called when entering the SetTxMetaBlock production.
	EnterSetTxMetaBlock(c *SetTxMetaBlockContext)

	// EnterSetAccountMeta is called when entering the SetAccountMeta production.
	EnterSetAccountMeta(c *SetAccountMetaContext)

	// EnterScript is called when entering the script production.
	EnterScript(c *ScriptContext)

	// ExitMonetaryLit is called when exiting the MonetaryLit production.
	ExitMonetaryLit(c *MonetaryLitContext)

	// ExitMonetaryAll is called when exiting the MonetaryAll production.
	ExitMonetaryAll(c *MonetaryAllContext)

	// ExitMonetaryAssetOnly is called when exiting the MonetaryAssetOnly production.
	ExitMonetaryAssetOnly(c *MonetaryAssetOnlyContext)

	// ExitLitAccount is called when exiting the LitAccount production.
	ExitLitAccount(c *LitAccountContext)

	// ExitLitAsset is called when exiting the LitAsset production.
	ExitLitAsset(c *LitAssetContext)

	// ExitLitNumber is called when exiting the LitNumber production.
	ExitLitNumber(c *LitNumberContext)

	// ExitLitMonetary is called when exiting the LitMonetary production.
	ExitLitMonetary(c *LitMonetaryContext)

	// ExitExprAddSub is called when exiting the ExprAddSub production.
	ExitExprAddSub(c *ExprAddSubContext)

	// ExitExprLiteral is called when exiting the ExprLiteral production.
	ExitExprLiteral(c *ExprLiteralContext)

	// ExitExprVariable is called when exiting the ExprVariable production.
	ExitExprVariable(c *ExprVariableContext)

	// ExitPortionPercent is called when exiting the PortionPercent production.
	ExitPortionPercent(c *PortionPercentContext)

	// ExitPortionRatio is called when exiting the PortionRatio production.
	ExitPortionRatio(c *PortionRatioContext)

	// ExitPortionRemaining is called when exiting the PortionRemaining production.
	ExitPortionRemaining(c *PortionRemainingContext)

	// ExitSrcRemaining is called when exiting the SrcRemaining production.
	ExitSrcRemaining(c *SrcRemainingContext)

	// ExitSrcPercentLimit is called when exiting the SrcPercentLimit production.
	ExitSrcPercentLimit(c *SrcPercentLimitContext)

	// ExitSrcCascade is called when exiting the SrcCascade production.
	ExitSrcCascade(c *SrcCascadeContext)

	// ExitSrcSimple is called when exiting the SrcSimple production.
	ExitSrcSimple(c *SrcSimpleContext)

	// ExitSrcOverdraft is called when exiting the SrcOverdraft production.
	ExitSrcOverdraft(c *SrcOverdraftContext)

	// ExitSrcOverdraftCapped is called when exiting the SrcOverdraftCapped production.
	ExitSrcOverdraftCapped(c *SrcOverdraftCappedContext)

	// ExitSrcLimit is called when exiting the SrcLimit production.
	ExitSrcLimit(c *SrcLimitContext)

	// ExitSrcPercent is called when exiting the SrcPercent production.
	ExitSrcPercent(c *SrcPercentContext)

	// ExitType_ is called when exiting the type_ production.
	ExitType_(c *Type_Context)

	// ExitVarTyped is called when exiting the VarTyped production.
	ExitVarTyped(c *VarTypedContext)

	// ExitVarBalance is called when exiting the VarBalance production.
	ExitVarBalance(c *VarBalanceContext)

	// ExitVarMeta is called when exiting the VarMeta production.
	ExitVarMeta(c *VarMetaContext)

	// ExitMetaValueExpr is called when exiting the MetaValueExpr production.
	ExitMetaValueExpr(c *MetaValueExprContext)

	// ExitMetaValueRatio is called when exiting the MetaValueRatio production.
	ExitMetaValueRatio(c *MetaValueRatioContext)

	// ExitMetadataEntry is called when exiting the metadataEntry production.
	ExitMetadataEntry(c *MetadataEntryContext)

	// ExitSendTo is called when exiting the SendTo production.
	ExitSendTo(c *SendToContext)

	// ExitSendKeep is called when exiting the SendKeep production.
	ExitSendKeep(c *SendKeepContext)

	// ExitSendSplit is called when exiting the SendSplit production.
	ExitSendSplit(c *SendSplitContext)

	// ExitPrint is called when exiting the Print production.
	ExitPrint(c *PrintContext)

	// ExitFail is called when exiting the Fail production.
	ExitFail(c *FailContext)

	// ExitTransferSimple is called when exiting the TransferSimple production.
	ExitTransferSimple(c *TransferSimpleContext)

	// ExitTransferWithDest is called when exiting the TransferWithDest production.
	ExitTransferWithDest(c *TransferWithDestContext)

	// ExitTransferAll is called when exiting the TransferAll production.
	ExitTransferAll(c *TransferAllContext)

	// ExitReserve is called when exiting the Reserve production.
	ExitReserve(c *ReserveContext)

	// ExitSetTxMeta is called when exiting the SetTxMeta production.
	ExitSetTxMeta(c *SetTxMetaContext)

	// ExitSetTxMetaBlock is called when exiting the SetTxMetaBlock production.
	ExitSetTxMetaBlock(c *SetTxMetaBlockContext)

	// ExitSetAccountMeta is called when exiting the SetAccountMeta production.
	ExitSetAccountMeta(c *SetAccountMetaContext)

	// ExitScript is called when exiting the script production.
	ExitScript(c *ScriptContext)
}
