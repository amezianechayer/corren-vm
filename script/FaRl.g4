grammar FaRl;

// ─── LEXER ───────────────────────────────────────────────────────────────────

NEWLINE    : [\r\n]+ ;
WHITESPACE : [ \t]+ -> skip ;

// Keywords
PRINT    : 'print' ;
FAIL     : 'fail' ;
TRANSFER : 'transfer' ;
FROM     : 'from' ;
TO       : 'to' ;
VARS     : 'vars' ;

// Types
TY_ACCOUNT  : 'account' ;
TY_ASSET    : 'asset' ;
TY_NUMBER   : 'number' ;
TY_MONETARY : 'monetary' ;

// Operators
OP_ADD : '+' ;
OP_SUB : '-' ;

// Symbols
LBRACK : '[' ;
RBRACK : ']' ;
LBRACE : '{' ;
RBRACE : '}' ;
DOT    : '.' ;
COMMA  : ',' ;

// Tokens
VARIABLE_NAME : '$' [a-z_] [a-z0-9_]* ;
ACCOUNT       : '@' [a-z_] [a-z0-9_:]* ;
ASSET         : [A-Z] [A-Z0-9]* ;
PRECISION     : [0-9] ;
NUMBER        : [0-9]+ ;
IDENTIFIER    : [a-z_] [a-z0-9_:]* ;

// ─── PARSER ──────────────────────────────────────────────────────────────────

monetary
    : LBRACK asset=ASSET DOT precision=PRECISION amount=NUMBER RBRACK
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

type_
    : TY_ACCOUNT
    | TY_ASSET
    | TY_NUMBER
    | TY_MONETARY
    ;

varDecl
    : ty=type_ name=VARIABLE_NAME
    ;

varListDecl
    : VARS LBRACE NEWLINE (v+=varDecl NEWLINE)* RBRACE NEWLINE
    ;

statement
    : PRINT expr=expression                                                # Print
    | FAIL                                                                 # Fail
    | TRANSFER amount=monetary FROM source=expression TO dest=expression   # Transfer
    ;

script
    : vars=varListDecl?
      stmts+=statement (NEWLINE stmts+=statement)*
      NEWLINE? EOF
    ;
