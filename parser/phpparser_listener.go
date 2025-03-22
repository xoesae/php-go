// Code generated from PhpParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // PhpParser

import "github.com/antlr4-go/antlr/v4"

// PhpParserListener is a complete listener for a parse tree produced by PhpParser.
type PhpParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterExpressionStatement is called when entering the expressionStatement production.
	EnterExpressionStatement(c *ExpressionStatementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignmentExpression is called when entering the assignmentExpression production.
	EnterAssignmentExpression(c *AssignmentExpressionContext)

	// EnterUnaryExpression is called when entering the unaryExpression production.
	EnterUnaryExpression(c *UnaryExpressionContext)

	// EnterPrimaryExpression is called when entering the primaryExpression production.
	EnterPrimaryExpression(c *PrimaryExpressionContext)

	// EnterFunctionDeclaration is called when entering the functionDeclaration production.
	EnterFunctionDeclaration(c *FunctionDeclarationContext)

	// EnterParameterList is called when entering the parameterList production.
	EnterParameterList(c *ParameterListContext)

	// EnterParameter is called when entering the parameter production.
	EnterParameter(c *ParameterContext)

	// EnterClassDeclaration is called when entering the classDeclaration production.
	EnterClassDeclaration(c *ClassDeclarationContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterForeachStatement is called when entering the foreachStatement production.
	EnterForeachStatement(c *ForeachStatementContext)

	// EnterSwitchStatement is called when entering the switchStatement production.
	EnterSwitchStatement(c *SwitchStatementContext)

	// EnterCaseStatement is called when entering the caseStatement production.
	EnterCaseStatement(c *CaseStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterEchoStatement is called when entering the echoStatement production.
	EnterEchoStatement(c *EchoStatementContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterBlock is called when entering the block production.
	EnterBlock(c *BlockContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitExpressionStatement is called when exiting the expressionStatement production.
	ExitExpressionStatement(c *ExpressionStatementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignmentExpression is called when exiting the assignmentExpression production.
	ExitAssignmentExpression(c *AssignmentExpressionContext)

	// ExitUnaryExpression is called when exiting the unaryExpression production.
	ExitUnaryExpression(c *UnaryExpressionContext)

	// ExitPrimaryExpression is called when exiting the primaryExpression production.
	ExitPrimaryExpression(c *PrimaryExpressionContext)

	// ExitFunctionDeclaration is called when exiting the functionDeclaration production.
	ExitFunctionDeclaration(c *FunctionDeclarationContext)

	// ExitParameterList is called when exiting the parameterList production.
	ExitParameterList(c *ParameterListContext)

	// ExitParameter is called when exiting the parameter production.
	ExitParameter(c *ParameterContext)

	// ExitClassDeclaration is called when exiting the classDeclaration production.
	ExitClassDeclaration(c *ClassDeclarationContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitForeachStatement is called when exiting the foreachStatement production.
	ExitForeachStatement(c *ForeachStatementContext)

	// ExitSwitchStatement is called when exiting the switchStatement production.
	ExitSwitchStatement(c *SwitchStatementContext)

	// ExitCaseStatement is called when exiting the caseStatement production.
	ExitCaseStatement(c *CaseStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitEchoStatement is called when exiting the echoStatement production.
	ExitEchoStatement(c *EchoStatementContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitBlock is called when exiting the block production.
	ExitBlock(c *BlockContext)
}
