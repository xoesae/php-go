package interpreter

import (
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
)

func (i *Interpreter) VisitIfStatement(ctx *parser.IfStatementContext) interface{} {
	logrus.Debug("Visit IfStatement")
	conditionExpr := ctx.Expression().(*parser.ExpressionContext)
	condition := i.VisitExpression(conditionExpr)
	b, ok := condition.(bool)

	logrus.Debug("IfStatement ", condition)

	if ok {
		if b {
			block := ctx.Block(0)
			i.Visit(block)
		} else {
			if ctx.Block(1) != nil {
				logrus.Debug("Else Block")
				i.Visit(ctx.Block(1))
			}
		}
	}

	return nil
}
