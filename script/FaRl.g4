grammar FaRl;

// ─── LEXER ───────────────────────────────────────────────────────────────────
WHITESPACE : [ \t\r\n]+ -> skip ;  // changé : ajout \r\n

// Keywords
PRINT    : 'print' ;
FAIL     : 'fail' ;
TRANSFER : 'transfer' ;
FROM     : 'from' ;
TO       : 'to' ;
VAR      : 'var' ;  // changé : VARS → VAR

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
COLON  : ':' ;  // ajouté

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
    : VAR name=VARIABLE_NAME COLON ty=type_  // changé : var $name: type
    ;

varListDecl
    : LBRACE v+=varDecl+ RBRACE  // changé : plus de VARS, plus de NEWLINE
    ;

statement
    : PRINT expr=expression                                                # Print
    | FAIL                                                                 # Fail
    | TRANSFER amount=expression FROM source=expression TO dest=expression  # Transfer
    ;

script
    : vars=varListDecl?
      stmts+=statement+
      EOF
    ;