package interpreter

import (
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/parser"
)

func (i *Interpreter) VisitBlock(ctx *parser.BlockContext) interface{} {
	logrus.Debug("Visit Block")

	for _, stmt := range ctx.AllStatement() {
		i.Visit(stmt)
	}

	logrus.Debug("Exiting Block")

	return nil
}
