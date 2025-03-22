package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xoesae/php-go/parser"
	"testing"
)

func parseUnaryExpression(input string) *parser.UnaryExpressionContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	return p.UnaryExpression().(*parser.UnaryExpressionContext)
}

func TestVisitUnaryExpressionNegation(t *testing.T) {
	tree := parseUnaryExpression("!true")
	visitor := NewInterpreter()

	result := visitor.VisitUnaryExpression(tree)

	assert.Equal(t, false, result)
}

func TestVisitUnaryExpressionNot(t *testing.T) {
	tree := parseUnaryExpression("!false")
	visitor := NewInterpreter()

	result := visitor.VisitUnaryExpression(tree)

	assert.Equal(t, true, result)
}

func TestVisitUnaryExpressionUnaryMinus(t *testing.T) {
	tree := parseUnaryExpression("-10")
	visitor := NewInterpreter()

	result := visitor.VisitUnaryExpression(tree)

	assert.Equal(t, -10.0, result)
}

func TestVisitUnaryExpressionUnaryMinusPositive(t *testing.T) {
	tree := parseUnaryExpression("-(-5)")
	visitor := NewInterpreter()

	result := visitor.VisitUnaryExpression(tree)

	assert.Equal(t, 5.0, result)
}
