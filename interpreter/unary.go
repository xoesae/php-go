package interpreter

import (
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
	"github.com/xoesae/php-go/utils"
)

func (i *Interpreter) VisitUnaryExpression(ctx *parser.UnaryExpressionContext) interface{} {
	logrus.Debug("Visiting UnaryExpression")
	value := i.Visit(ctx.Expression())
	if ctx.T_NOT() != nil {
		return !utils.ToBool(value)
	} else if ctx.T_MINUS() != nil {
		return -utils.ToFloat(value)
	}
	return value
}
