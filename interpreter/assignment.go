package interpreter

import (
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
)

func (i *Interpreter) VisitAssignmentExpression(ctx *parser.AssignmentExpressionContext) interface{} {
	logrus.Debug("Visit Assignment Expression")
	varName := ctx.T_VARIABLE().GetText()
	value := i.Visit(ctx.Expression())

	logrus.Debug("Assignment Expression", varName, value)

	i.variables[varName] = value
	return value
}
