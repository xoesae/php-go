package interpreter

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
	"github.com/xoesae/php-go/utils"
)

func (i *Interpreter) VisitExpression(ctx *parser.ExpressionContext) interface{} {
	logrus.Debug("VisitExpression ", ctx.GetChildCount())
	if ctx.GetChildCount() == 1 {
		if primaryCtx, ok := ctx.PrimaryExpression().(*parser.PrimaryExpressionContext); ok {
			return i.VisitPrimaryExpression(primaryCtx)
		}

		if assgnCtx, ok := ctx.AssignmentExpression().(*parser.AssignmentExpressionContext); ok {
			return i.VisitAssignmentExpression(assgnCtx)
		}

		if unaryCtx, ok := ctx.UnaryExpression().(*parser.UnaryExpressionContext); ok {
			return i.VisitUnaryExpression(unaryCtx)
		}

		return nil
	}

	left := i.Visit(ctx.Expression(0))
	right := i.Visit(ctx.Expression(1))
	op := ctx.GetChild(1).(antlr.TerminalNode).GetText()

	switch op {
	case "+":
		return utils.ToFloat(left) + utils.ToFloat(right)
	case "-":
		return utils.ToFloat(left) - utils.ToFloat(right)
	case "*":
		return utils.ToFloat(left) * utils.ToFloat(right)
	case "/":
		return utils.ToFloat(left) / utils.ToFloat(right)
	case "%":
		return int(utils.ToFloat(left)) % int(utils.ToFloat(right))
	case ".":
		return fmt.Sprintf("%v%v", left, right)
	case "==":
		if utils.IsNumber(left) && utils.IsNumber(right) {
			return utils.ToFloat(left) == utils.ToFloat(right)
		}

		return left == right
	case "!=":
		if utils.IsNumber(left) && utils.IsNumber(right) {
			return utils.ToFloat(left) != utils.ToFloat(right)
		}

		return left != right
	}
	return nil
}

func (i *Interpreter) VisitExpressionStatement(ctx *parser.ExpressionStatementContext) interface{} {
	logrus.Debug("Visit Expression Statement")
	var results []string

	if exprCtx, ok := ctx.Expression().(*parser.ExpressionContext); ok {
		value := i.VisitExpression(exprCtx)
		results = append(results, fmt.Sprintf("%v", value))
	}

	return nil
}
