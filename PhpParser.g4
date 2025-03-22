parser grammar PhpParser;

options { tokenVocab=PhpLexer; }

program: statement* EOF;

statement
    : expressionStatement
    | functionDeclaration
    | classDeclaration
    | ifStatement
    | whileStatement
    | forStatement
    | foreachStatement
    | switchStatement
    | returnStatement
    | echoStatement
    ;

expressionStatement: expression T_SEMICOLON ;

expression
    : assignmentExpression
    | expression (T_PLUS | T_MINUS | T_MULT | T_DIV | T_MOD | T_CONCAT | T_EQUAL | T_NOT_EQUAL) expression
    | unaryExpression
    | primaryExpression
    ;

assignmentExpression: T_VARIABLE T_ASSIGN expression;

unaryExpression: (T_NOT | T_MINUS) expression;

primaryExpression
    : T_VARIABLE
    | T_INT
    | T_FLOAT
    | T_STRING
    | T_BOOLEAN
    | T_L_PARENTHESIS expression T_R_PARENTHESIS
    ;

functionDeclaration: T_FUNCTION T_IDENTIFIER T_L_PARENTHESIS parameterList? T_R_PARENTHESIS block;

parameterList: parameter (T_COMMA parameter)*;
parameter: T_VARIABLE;

classDeclaration: T_CLASS T_IDENTIFIER ('extends' T_IDENTIFIER)? ('implements' T_IDENTIFIER (T_COMMA T_IDENTIFIER)*)? block;

ifStatement: T_IF T_L_PARENTHESIS expression T_R_PARENTHESIS block (T_ELSE block)?;

whileStatement: T_WHILE T_L_PARENTHESIS expression T_R_PARENTHESIS block;

forStatement: T_FOR T_L_PARENTHESIS expression? T_SEMICOLON expression? T_SEMICOLON expression? T_R_PARENTHESIS block;

foreachStatement: T_FOREACH T_L_PARENTHESIS T_VARIABLE T_AS T_VARIABLE T_R_PARENTHESIS block;

switchStatement: T_SWITCH T_L_PARENTHESIS expression T_R_PARENTHESIS T_L_BRACE caseStatement* T_R_BRACE;

caseStatement: T_CASE expression T_COLON statement* | T_DEFAULT T_COLON statement*;

returnStatement: T_RETURN expression? T_SEMICOLON;

echoStatement: T_ECHO expressionList T_SEMICOLON;

expressionList: expression (T_COMMA expression)*;

block: T_L_BRACE statement* T_R_BRACE;
