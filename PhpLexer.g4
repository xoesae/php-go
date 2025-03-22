lexer grammar PhpLexer;

// Definição de Tokens
T_OPEN_TAG : '<?php' | '<?=';
T_CLOSE_TAG : '?>';

T_FUNCTION : 'function';
T_CLASS : 'class';
T_PUBLIC : 'public';
T_PRIVATE : 'private';
T_PROTECTED : 'protected';
T_VAR : 'var';
T_RETURN : 'return';
T_IF : 'if';
T_ELSE : 'else';
T_WHILE : 'while';
T_FOR : 'for';
T_FOREACH : 'foreach';
T_SWITCH : 'switch';
T_CASE : 'case';
T_DEFAULT : 'default';
T_BREAK : 'break';
T_CONTINUE : 'continue';
T_ECHO : 'echo';
T_PRINT : 'print';
T_NEW : 'new';
T_EXTENDS : 'extends';
T_IMPLEMENTS : 'implements';
T_NAMESPACE : 'namespace';
T_USE : 'use';
T_TRY : 'try';
T_CATCH : 'catch';
T_FINALLY : 'finally';
T_THROW : 'throw';
T_STATIC : 'static';
T_ABSTRACT : 'abstract';
T_FINAL : 'final';
T_AS : 'as' ;

// Boolean
T_BOOLEAN : 'true' | 'TRUE' | 'false' | 'FALSE';

// Identificadores e variáveis
T_VARIABLE : '$' [a-zA-Z_][a-zA-Z0-9_]*;
T_IDENTIFIER : [a-zA-Z_][a-zA-Z0-9_]*;

// Operadores
T_ASSIGN : '=';
T_PLUS : '+';
T_MINUS : '-';
T_MULT : '*';
T_DIV : '/';
T_MOD : '%';
T_POW : '**';
T_AND : '&&';
T_OR : '||';
T_NOT : '!';
T_EQUAL : '==';
T_NOT_EQUAL : '!=';
T_GREATER : '>';
T_LESS : '<';
T_GREATER_EQUAL : '>=';
T_LESS_EQUAL : '<=';
T_CONCAT : '.';
T_COLON : ':' ;
T_SEMICOLON : ';' ;
T_COMMA : ',' ;
T_L_PARENTHESIS : '(' ;
T_R_PARENTHESIS : ')' ;
T_L_BRACE : '{' ;
T_R_BRACE : '}' ;

// Números
T_INT : [0-9]+;
T_FLOAT : [0-9]+ '.' [0-9]+;

// Strings
T_STRING : '"' ( ~('"' | '\\') | '\\' . )* '"' | '\'' ( ~('\'' | '\\') | '\\' . )* '\'';

// Comentários
T_COMMENT : '//' ~[\r\n]* -> skip;
T_COMMENT_BLOCK : '/*' .*? '*/' -> skip;

// Espaços em branco
T_WHITESPACE : [ \t\r\n]+ -> skip;
