//
//
//

grammar Bund;

expressions
  : ( integer
    | oct_integer
    | hex_integer
    | bin_integer
    | float_term
    | string_term
    | false_term
    | true_term
    | name_term
    | duplicate_term
    | drop_term
    | begin
    | end
    | pair
    | block
    | directive
    | cmd
    | lambda )*
  ;

term: ( integer
  | oct_integer | hex_integer | bin_integer
  | float_term | string_term | false_term | true_term
  | name_term | duplicate_term | drop_term
  | begin | end | pair | block | directive | cmd ) ;
//
// Common terms
//
true_term:    value=TRUE ;
false_term:   value=FALSE ;
string_term:  value=STRING ;
float_term:   value=FLOAT_NUMBER ;
bin_integer:  value=BIN_INTEGER ;
hex_integer:  value=HEX_INTEGER ;
oct_integer:  value=OCT_INTEGER ;
integer:      value=INTEGER ;


//
// Name term
//
name_term:    value=NAME ;
//
// Stack management terms
//
duplicate_term: value=DUPLICATE ;
drop_term:      value=DROP ;
begin:          value=TOBEGIN;
end:            value=TOEND ;

block
  : '|' (values+=term)* '|'
  ;

pair
  : '(' head=NAME tail=term ')'
  ;

directive
  : op=DIR name=NAME
  ;

cmd
  : command=CMD
  ;

lambda
  : '[' name=NAME ']' (body+=term)* '.'
  ;

INTEGER
  : (SIGN)? DECIMAL_INTEGER
  ;

DECIMAL_INTEGER
  : NON_ZERO_DIGIT DIGIT*
  | '0'+
  ;

OCT_INTEGER
  : '0' [oO] [0-7]+
  ;

HEX_INTEGER
  : '0' [xX] [0-9a-fA-F]+
  ;

BIN_INTEGER
  : '0' [bB] [01]+
  ;

FLOAT_NUMBER
  : EXPONENT_OR_POINT_FLOAT
  ;

STRING
  : SHORT_STRING
  | LONG_STRING
  ;

TRUE
  : 'true'
  | 'TRUE'
  | '#T'
  ;

FALSE
  : 'false'
  | 'FALSE'
  | '#F'
  ;

NAME
  : ID_START ID_CONTINUE*
  ;



TOBEGIN: ':' ;
TOEND:   ';' ;
SLASH:   '/' ;

DIR
  : ('+'|'-'|'*'|'`'|'~')+
  ;

CMD
  : SPACES ('@'|'$'|'&'|'='|','|'%'|'<'|'>')+ SPACES
  ;

DUPLICATE
  : SPACES '^' SPACES
  ;

DROP
  : SPACES '_' SPACES
  ;

SKIP_
  : ( SPACES | COMMENT | LINE_JOINING | CCOMMENT | BLOCK_COMMENT ) -> skip
  ;


//
// Fragments generic term fragments
//
fragment NON_ZERO_DIGIT
  : [1-9]
  ;

fragment DIGIT
  : [0-9]
  ;

fragment EXPONENT_OR_POINT_FLOAT
  : (SIGN)? ([0-9]+ | POINT_FLOAT) [eE] [+-]? [0-9]+
  | (SIGN)? POINT_FLOAT
  ;

fragment POINT_FLOAT
  : [0-9]* '.' [0-9]+
  | [0-9]+ '.'
  ;

fragment SIGN
  : ('+' | '-')
  ;

fragment RN
  : '\r'? '\n'
  ;

fragment SHORT_STRING
  : '\'' ('\\' (RN | .) | ~[\\\r\n'])* '\''
  | '"'  ('\\' (RN | .) | ~[\\\r\n"])* '"'
  ;

fragment LONG_STRING
  : '\'\'\'' LONG_STRING_ITEM*? '\'\'\''
  | '"""' LONG_STRING_ITEM*? '"""'
  ;

fragment LONG_STRING_ITEM
  : ~'\\'
  | '\\' (RN | .)
  ;

fragment ID_START
 : ([A-Z]|[a-z]|SLASH)
 | [a-z]
 ;

fragment ID_CONTINUE
 : ID_START
 | [0-9]
 | SLASH
 ;

//
// SKIP fragments
//

fragment LINE_JOINING
  : '\\' SPACES? ( '\r'? '\n' | '\r' | '\f' )
  ;

fragment COMMENT
  : '##' ~[\r\n\f]*
  ;

fragment BLOCK_COMMENT
  :   '/*' .*? '*/'
  ;

fragment CCOMMENT
  :   '//' ~[\r\n]*
  ;

fragment SPACES
  : [ \t]+
  ;
