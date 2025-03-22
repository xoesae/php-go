package interpreter

import (
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
	"strconv"
)

func (i *Interpreter) VisitPrimaryExpression(ctx *parser.PrimaryExpressionContext) interface{} {
	logrus.Debug("Visiting primary expression")

	if ctx.T_INT() != nil {
		val, _ := strconv.Atoi(ctx.T_INT().GetText())
		return val
	} else if ctx.T_FLOAT() != nil {
		val, _ := strconv.ParseFloat(ctx.T_FLOAT().GetText(), 64)
		return val
	} else if ctx.T_BOOLEAN() != nil {
		val, _ := strconv.ParseBool(ctx.T_BOOLEAN().GetText())
		return val
	} else if ctx.T_STRING() != nil {
		text := ctx.T_STRING().GetText()
		return text[1 : len(text)-1]
	} else if ctx.T_VARIABLE() != nil {
		varName := ctx.T_VARIABLE().GetText()
		if value, exists := i.variables[varName]; exists {
			return value
		}
		panic("Variable not defined: " + varName)
	} else if ctx.T_L_PARENTHESIS() != nil {
		return i.Visit(ctx.Expression())
	}
	return nil
}
