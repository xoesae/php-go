package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
)

type Interpreter struct {
	*antlr.BaseParseTreeVisitor
	variables map[string]interface{}
}

func NewInterpreter() *Interpreter {
	return &Interpreter{
		variables: make(map[string]interface{}),
	}
}

func (i *Interpreter) Visit(tree antlr.ParseTree) interface{} {
	switch node := tree.(type) {
	case *parser.ProgramContext:
		return i.VisitProgram(node)
	case *parser.EchoStatementContext:
		return i.VisitEchoStatement(node)
	case *parser.AssignmentExpressionContext:
		return i.VisitAssignmentExpression(node)
	case *parser.ExpressionContext:
		return i.VisitExpression(node)
	case *parser.StatementContext:
		return i.VisitStatement(node)
	case *parser.BlockContext:
		return i.VisitBlock(node)
	default:
		logrus.Debug("Not found")
	}
	return nil
}

func (i *Interpreter) VisitProgram(ctx *parser.ProgramContext) interface{} {
	logrus.Debug("Visit Program")
	for _, stmt := range ctx.AllStatement() {
		i.VisitStatement(stmt.(*parser.StatementContext))
	}
	return nil
}
