grammar FaRl;

NEWLINE    : [\r\n]+ ;
WHITESPACE : [ \t]+ -> skip ;

TRANSFER   : 'transfer' ;
PRINT      : 'print' ;
FAIL       : 'fail' ;
FROM       : 'from' ;
TO         : 'to' ;

OP_ADD : '+' ;
OP_SUB : '-' ;

LBRACK : '[' ;
RBRACK : ']' ;
DOT    : '.' ;

IDENTIFIER : '@' [a-z_] [a-z0-9_:]* ;
ASSET      : [A-Z] [A-Z0-9]* ;
PRECISION  : [0-9] ;
NUMBER     : [0-9]+ ;

monetary
    : LBRACK asset=ASSET DOT precision=PRECISION amount=NUMBER RBRACK
    ;

literal
    : IDENTIFIER # LitAddress
    | ASSET      # LitAsset
    | NUMBER     # LitNumber
    | monetary   # LitMonetary
    ;

expression
    : lhs=expression op=(OP_ADD | OP_SUB) rhs=expression # ExprAddSub
    | lit=literal                                         # ExprLiteral
    ;

statement
    : PRINT expr=expression                                            # Print
    | FAIL                                                             # Fail
    | TRANSFER amount=monetary FROM source=literal TO dest=literal     # Transfer
    ;

script
    : stmts+=statement
      (NEWLINE stmts+=statement)*
      EOF
    ;