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
VAR      : 'var' ;

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
COLON  : ':' ;

// Tokens
VARIABLE_NAME : '$' [a-z_] [a-z0-9_]* ;
ACCOUNT       : '@' [a-z_] [a-z0-9_:]* ;
ASSET         : [A-Z] [A-Z0-9]* ;
NUMBER        : [0-9]+ ;  // supprimé PRECISION

// ─── PARSER ──────────────────────────────────────────────────────────────────
monetary
    : LBRACK asset=ASSET DOT precision=NUMBER amount=NUMBER RBRACK  // precision=NUMBER
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
    : VAR name=VARIABLE_NAME COLON ty=type_
    ;

varListDecl
    : LBRACE NEWLINE+ (v+=varDecl NEWLINE+)+ RBRACE  // supprimé NEWLINE+ après RBRACE
    ;

statement
    : PRINT expr=expression                            # Print
    | FAIL                                             # Fail
    | TRANSFER NEWLINE* amount=expression NEWLINE*
      FROM NEWLINE* source=expression NEWLINE*
      TO NEWLINE* dest=expression                      # Transfer
    ;

script
    : NEWLINE*
      vars=varListDecl?
      NEWLINE*
      stmts+=statement
      (NEWLINE+ stmts+=statement)*
      NEWLINE*
      EOF
    ;