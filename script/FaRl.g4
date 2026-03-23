grammar FaRl;

NEWLINE           : [\r\n]+ ;
WHITESPACE        : [ \t]+ -> skip ;
MULTILINE_COMMENT : '/*' (MULTILINE_COMMENT|.)*? '*/' -> skip ;
LINE_COMMENT      : '//' .*? NEWLINE -> skip ;

PRINT       : 'print' ;
FAIL        : 'fail' ;
TRANSFER    : 'transfer' ;
FROM        : 'from' ;
TO          : 'to' ;
MAX         : 'max' ;
KEEP        : 'keep' ;
ALL         : 'all' ;
VAR         : 'var' ;
SET         : 'set' ;
TRANSACTION : 'transaction' ;
METADATA    : 'metadata' ;
OF          : 'of' ;
KEY         : 'key' ;
META        : 'meta' ;

TY_ACCOUNT  : 'account' ;
TY_ASSET    : 'asset' ;
TY_NUMBER   : 'number' ;
TY_MONETARY : 'monetary' ;
TY_PORTION  : 'portion' ;
TY_STRING   : 'string' ;

OP_ADD : '+' ;
OP_SUB : '-' ;
PERCENT           : '%' ;
PORTION_REMAINING : 'remaining' ;

LBRACK : '[' ;
RBRACK : ']' ;
LBRACE : '{' ;
RBRACE : '}' ;
LPAREN : '(' ;
RPAREN : ')' ;
DOT    : '.' ;
COLON  : ':' ;
STAR   : '*' ;
EQ     : '=' ;

STRING        : '"' [a-zA-Z0-9_/.]* '"' ;
RATIO         : [0-9]+ '/' [0-9]+ ;
VARIABLE_NAME : '$' [a-z_] [a-z0-9_]* ;
ACCOUNT       : '@' [a-z_] [a-z0-9_:]* ;
ASSET         : [A-Z] [A-Z0-9]* ;
NUMBER        : [0-9]+ ;

monetary
    : LBRACK asset=ASSET DOT precision=NUMBER amount=NUMBER RBRACK  # MonetaryLit
    | LBRACK asset=ASSET amount=NUMBER RBRACK                       # MonetaryNoPrecision
    ;

monetaryAll
    : LBRACK asset=ASSET DOT precision=NUMBER STAR RBRACK           # MonetaryAllPrec
    | LBRACK asset=ASSET STAR RBRACK                                # MonetaryAllNoPrec
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

allotmentPortion
    : RATIO                    # allotmentPortionConst
    | NUMBER ('.' NUMBER)? PERCENT  # allotmentPortionConstPercent
    | por=VARIABLE_NAME        # allotmentPortionVar
    | PORTION_REMAINING        # allotmentPortionRemaining
    ;

destinationAllotment
    : LBRACE NEWLINE (portions+=allotmentPortion TO dests+=expression NEWLINE)+ RBRACE
    ;

destination
    : expression               # DestAccount
    | destinationAllotment     # DestAllotment
    ;

sourceAllotment
    : LBRACE NEWLINE (portions+=allotmentPortion FROM sources+=source NEWLINE)+ RBRACE
    ;

sourceMaxed
    : MAX max=expression FROM src=source
    ;

sourceInOrder
    : LBRACE NEWLINE (sources+=source NEWLINE)+ RBRACE
    ;

source
    : expression               # SrcAccount
    | sourceMaxed              # SrcMaxed
    | sourceInOrder            # SrcInOrder
    ;

valueAwareSource
    : source                   # Src
    | sourceAllotment          # SrcAllotment
    ;

type_
    : TY_ACCOUNT
    | TY_ASSET
    | TY_NUMBER
    | TY_MONETARY
    | TY_PORTION
    | TY_STRING
    ;

origin
    : META LPAREN acc=expression ',' key=STRING RPAREN
    ;

varDecl
    : VAR name=VARIABLE_NAME COLON ty=type_ (EQ orig=origin)?  # VarTyped
    ;

metadataValue
    : expression    # MetaValueExpr
    | RATIO         # MetaValueRatio
    ;

metadataEntry
    : STRING EQ metadataValue
    ;

statement
    : PRINT expr=expression
        # Print

    | FAIL
        # Fail

    | TRANSFER (mon=expression | monAll=monetaryAll) LPAREN NEWLINE
        ( FROM src=valueAwareSource NEWLINE TO dest=destination
        | TO dest=destination NEWLINE FROM src=valueAwareSource) NEWLINE RPAREN
        # Transfer

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
