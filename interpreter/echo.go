package interpreter

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
	"strings"
)

func (i *Interpreter) VisitEchoStatement(ctx *parser.EchoStatementContext) interface{} {
	logrus.Debug("Visiting echo statement")
	var results []string

	for _, exprCtx := range ctx.ExpressionList().AllExpression() {
		value := i.Visit(exprCtx)
		results = append(results, fmt.Sprintf("%v", value))
	}

	fmt.Println(strings.Join(results, " "))

	return nil
}
