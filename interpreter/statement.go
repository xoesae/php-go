package interpreter

import (
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
)

func (i *Interpreter) VisitStatement(ctx *parser.StatementContext) interface{} {
	logrus.Debug("Visit Statement")

	if echoCtx, ok := ctx.EchoStatement().(*parser.EchoStatementContext); ok {
		return i.VisitEchoStatement(echoCtx)
	}

	if exprCtx, ok := ctx.ExpressionStatement().(*parser.ExpressionStatementContext); ok {
		return i.VisitExpressionStatement(exprCtx)
	}

	if ifCtx, ok := ctx.IfStatement().(*parser.IfStatementContext); ok {
		return i.VisitIfStatement(ifCtx)
	}

	return nil
}
