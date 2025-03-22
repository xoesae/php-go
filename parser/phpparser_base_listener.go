// Code generated from PhpParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // PhpParser

import "github.com/antlr4-go/antlr/v4"

// BasePhpParserListener is a complete listener for a parse tree produced by PhpParser.
type BasePhpParserListener struct{}

var _ PhpParserListener = &BasePhpParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasePhpParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasePhpParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasePhpParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasePhpParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BasePhpParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BasePhpParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterStatement is called when production statement is entered.
func (s *BasePhpParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BasePhpParserListener) ExitStatement(ctx *StatementContext) {}

// EnterExpressionStatement is called when production expressionStatement is entered.
func (s *BasePhpParserListener) EnterExpressionStatement(ctx *ExpressionStatementContext) {}

// ExitExpressionStatement is called when production expressionStatement is exited.
func (s *BasePhpParserListener) ExitExpressionStatement(ctx *ExpressionStatementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BasePhpParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BasePhpParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignmentExpression is called when production assignmentExpression is entered.
func (s *BasePhpParserListener) EnterAssignmentExpression(ctx *AssignmentExpressionContext) {}

// ExitAssignmentExpression is called when production assignmentExpression is exited.
func (s *BasePhpParserListener) ExitAssignmentExpression(ctx *AssignmentExpressionContext) {}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BasePhpParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BasePhpParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BasePhpParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BasePhpParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterFunctionDeclaration is called when production functionDeclaration is entered.
func (s *BasePhpParserListener) EnterFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// ExitFunctionDeclaration is called when production functionDeclaration is exited.
func (s *BasePhpParserListener) ExitFunctionDeclaration(ctx *FunctionDeclarationContext) {}

// EnterParameterList is called when production parameterList is entered.
func (s *BasePhpParserListener) EnterParameterList(ctx *ParameterListContext) {}

// ExitParameterList is called when production parameterList is exited.
func (s *BasePhpParserListener) ExitParameterList(ctx *ParameterListContext) {}

// EnterParameter is called when production parameter is entered.
func (s *BasePhpParserListener) EnterParameter(ctx *ParameterContext) {}

// ExitParameter is called when production parameter is exited.
func (s *BasePhpParserListener) ExitParameter(ctx *ParameterContext) {}

// EnterClassDeclaration is called when production classDeclaration is entered.
func (s *BasePhpParserListener) EnterClassDeclaration(ctx *ClassDeclarationContext) {}

// ExitClassDeclaration is called when production classDeclaration is exited.
func (s *BasePhpParserListener) ExitClassDeclaration(ctx *ClassDeclarationContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BasePhpParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BasePhpParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BasePhpParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BasePhpParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BasePhpParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BasePhpParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterForeachStatement is called when production foreachStatement is entered.
func (s *BasePhpParserListener) EnterForeachStatement(ctx *ForeachStatementContext) {}

// ExitForeachStatement is called when production foreachStatement is exited.
func (s *BasePhpParserListener) ExitForeachStatement(ctx *ForeachStatementContext) {}

// EnterSwitchStatement is called when production switchStatement is entered.
func (s *BasePhpParserListener) EnterSwitchStatement(ctx *SwitchStatementContext) {}

// ExitSwitchStatement is called when production switchStatement is exited.
func (s *BasePhpParserListener) ExitSwitchStatement(ctx *SwitchStatementContext) {}

// EnterCaseStatement is called when production caseStatement is entered.
func (s *BasePhpParserListener) EnterCaseStatement(ctx *CaseStatementContext) {}

// ExitCaseStatement is called when production caseStatement is exited.
func (s *BasePhpParserListener) ExitCaseStatement(ctx *CaseStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BasePhpParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BasePhpParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterEchoStatement is called when production echoStatement is entered.
func (s *BasePhpParserListener) EnterEchoStatement(ctx *EchoStatementContext) {}

// ExitEchoStatement is called when production echoStatement is exited.
func (s *BasePhpParserListener) ExitEchoStatement(ctx *EchoStatementContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BasePhpParserListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BasePhpParserListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterBlock is called when production block is entered.
func (s *BasePhpParserListener) EnterBlock(ctx *BlockContext) {}

// ExitBlock is called when production block is exited.
func (s *BasePhpParserListener) ExitBlock(ctx *BlockContext) {}
