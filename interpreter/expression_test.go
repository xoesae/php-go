package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xoesae/php-go/parser"
	"testing"
)

func parseExpression(input string) *parser.ExpressionContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	return p.Expression().(*parser.ExpressionContext)
}

func TestVisitExpressionAddition(t *testing.T) {
	tree := parseExpression("1 + 2")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, 3.0, result)
}

func TestVisitExpressionSubtraction(t *testing.T) {
	tree := parseExpression("10 - 3")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, 7.0, result)
}

func TestVisitExpressionMultiplication(t *testing.T) {
	tree := parseExpression("4 * 5")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, 20.0, result)
}

func TestVisitExpressionDivision(t *testing.T) {
	tree := parseExpression("10 / 2")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, 5.0, result)
}

func TestVisitExpressionMod(t *testing.T) {
	tree := parseExpression("10 % 3")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, 1, result)
}

func TestVisitExpressionConcatenation(t *testing.T) {
	tree := parseExpression("\"Hello, \" . \"World!\"")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, "Hello, World!", result)
}

func TestVisitExpressionEquality(t *testing.T) {
	tree := parseExpression("5 == 5")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, true, result)
}

func TestVisitExpressionInequality(t *testing.T) {
	tree := parseExpression("5 != 3")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, true, result)
}

func TestVisitExpressionUnaryMinus(t *testing.T) {
	tree := parseExpression("-7")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, -7.0, result)
}

func TestVisitExpressionParentheses(t *testing.T) {
	tree := parseExpression("(2 + 3) * 4")
	visitor := NewInterpreter()

	result := visitor.VisitExpression(tree)
	assert.Equal(t, 20.0, result)
}
