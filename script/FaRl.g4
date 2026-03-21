grammar FaRl;

NEWLINE    : [\r\n]+ ;
WHITESPACE : [ \t]+ -> skip ;

PRINT       : 'print' ;
FAIL        : 'fail' ;
TRANSFER    : 'transfer' ;
FROM        : 'from' ;
TO          : 'to' ;
THEN        : 'then' ;
SEND        : 'send' ;
KEEP        : 'keep' ;
REMAINING   : 'remaining' ;
ALL         : 'all' ;
TAKE        : 'take' ;
LIMIT       : 'limit' ;
ALLOW       : 'allow' ;
OVERDRAFT   : 'overdraft' ;
UP          : 'up' ;
RESERVE     : 'reserve' ;
IN          : 'in' ;
SPLIT       : 'split' ;
AS          : 'as' ;
VAR         : 'var' ;
BALANCE     : 'balance' ;
OF          : 'of' ;
META        : 'meta' ;
KEY         : 'key' ;
SET         : 'set' ;
TRANSACTION : 'transaction' ;
METADATA    : 'metadata' ;

TY_ACCOUNT  : 'account' ;
TY_ASSET    : 'asset' ;
TY_NUMBER   : 'number' ;
TY_MONETARY : 'monetary' ;
TY_PORTION  : 'portion' ;
TY_STRING   : 'string' ;

OP_ADD : '+' ;
OP_SUB : '-' ;
PERCENT: '%' ;

LBRACK : '[' ;
RBRACK : ']' ;
LBRACE : '{' ;
RBRACE : '}' ;
DOT    : '.' ;
COLON  : ':' ;
STAR   : '*' ;
EQ     : '=' ;

RATIO         : [0-9]+ '/' [0-9]+ ;
VARIABLE_NAME : '$' [a-z_] [a-z0-9_]* ;
ACCOUNT       : '@' [a-z_] [a-z0-9_:]* ;
ASSET         : [A-Z] [A-Z0-9]* ;
NUMBER        : [0-9]+ ;
STRING        : '"' (~["\r\n])* '"' ;

monetary
    : LBRACK asset=ASSET DOT precision=NUMBER amount=NUMBER RBRACK  # MonetaryLit
    | LBRACK asset=ASSET DOT precision=NUMBER STAR RBRACK           # MonetaryAll
    | LBRACK asset=ASSET amount=NUMBER RBRACK                       # MonetaryNoPrecision
    | LBRACK asset=ASSET STAR RBRACK                                # MonetaryNoPrecisionAll
    | LBRACK asset=ASSET RBRACK                                     # MonetaryAssetOnly
    ;

literal
    : ACCOUNT   # LitAccount
    | ASSET     # LitAsset
    | NUMBER    # LitNumber
    | monetary  # LitMonetary
    ;

expression
    : lhs=expression op=(OP_ADD | OP_SUB) rhs=expression  # ExprAddSub
    | lit=literal                                          # ExprLiteral
    | variable=VARIABLE_NAME                               # ExprVariable
    ;

portion
    : p=NUMBER PERCENT    # PortionPercent
    | r=RATIO             # PortionRatio
    | REMAINING           # PortionRemaining
    ;

source
    : FROM expression                                                # SrcSimple
    | FROM expression ALLOW OVERDRAFT                               # SrcOverdraft
    | FROM expression ALLOW OVERDRAFT UP TO monetary                # SrcOverdraftCapped
    | FROM expression LIMIT monetary THEN source                    # SrcLimit
    | source THEN expression                                        # SrcCascade
    | TAKE NUMBER PERCENT FROM expression                           # SrcPercent
    | TAKE NUMBER PERCENT FROM expression LIMIT monetary            # SrcPercentLimit
    | TAKE REMAINING FROM expression                                # SrcRemaining
    ;

type_
    : TY_ACCOUNT
    | TY_ASSET
    | TY_NUMBER
    | TY_MONETARY
    | TY_PORTION
    | TY_STRING
    ;

varDecl
    : VAR name=VARIABLE_NAME COLON ty=type_                                 # VarTyped
    | VAR name=VARIABLE_NAME EQ BALANCE OF expression IN ASSET DOT NUMBER  # VarBalance
    | VAR name=VARIABLE_NAME EQ META OF expression KEY STRING               # VarMeta
    ;

metadataValue
    : expression    # MetaValueExpr
    | RATIO         # MetaValueRatio
    ;

metadataEntry
    : STRING EQ metadataValue
    ;

sendClause
    : SEND portion TO expression                         # SendTo
    | KEEP REMAINING                                     # SendKeep
    | SPLIT portion AS COLON NEWLINE
        (sendClause NEWLINE)+                            # SendSplit
    ;

statement
    : PRINT expr=expression
        # Print

    | FAIL
        # Fail

    | TRANSFER amount=expression src=source TO dest=expression
        # TransferSimple

    | TRANSFER amount=expression src=source NEWLINE
        (sends+=sendClause NEWLINE)+
        # TransferWithDest

    | TRANSFER ALL monetary src=source TO dest=expression
        # TransferAll

    | RESERVE monetary IN expression
        # Reserve

    | SET TRANSACTION METADATA STRING EQ metadataValue
        # SetTxMeta

    | SET TRANSACTION METADATA LBRACE NEWLINE
        (metadataEntry NEWLINE)+
      RBRACE
        # SetTxMetaBlock

    | SET TY_ACCOUNT METADATA OF expression KEY STRING EQ metadataValue
        # SetAccountMeta
    ;

script
    : NEWLINE*
      (varDecl NEWLINE+)*
      stmts+=statement (NEWLINE+ stmts+=statement)*
      NEWLINE? EOF
    ;
