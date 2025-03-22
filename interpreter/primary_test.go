package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xoesae/php-go/parser"
	"testing"
)

func parsePrimaryExpression(input string) *parser.PrimaryExpressionContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	return p.PrimaryExpression().(*parser.PrimaryExpressionContext)
}

func TestVisitPrimaryExpressionVariable(t *testing.T) {
	tree := parsePrimaryExpression("$variable")
	visitor := &Interpreter{
		variables: make(map[string]interface{}),
	}
	visitor.variables["$variable"] = 1

	result := visitor.VisitPrimaryExpression(tree)

	assert.Equal(t, 1, result)
}

func TestVisitPrimaryExpressionInt(t *testing.T) {
	tree := parsePrimaryExpression("42")
	visitor := NewInterpreter()

	result := visitor.VisitPrimaryExpression(tree)
	assert.Equal(t, 42, result)
}

func TestVisitPrimaryExpressionFloat(t *testing.T) {
	tree := parsePrimaryExpression("3.14")
	visitor := NewInterpreter()

	result := visitor.VisitPrimaryExpression(tree)
	assert.Equal(t, 3.14, result)
}

func TestVisitPrimaryExpressionString(t *testing.T) {
	tree := parsePrimaryExpression("\"Hello, World!\"")
	visitor := NewInterpreter()

	result := visitor.VisitPrimaryExpression(tree)
	assert.Equal(t, "Hello, World!", result)
}

func TestVisitPrimaryExpressionBooleanTrue(t *testing.T) {
	tree := parsePrimaryExpression("true")
	visitor := NewInterpreter()

	result := visitor.VisitPrimaryExpression(tree)
	assert.Equal(t, true, result)
}

func TestVisitPrimaryExpressionBooleanFalse(t *testing.T) {
	tree := parsePrimaryExpression("false")
	visitor := NewInterpreter()

	result := visitor.VisitPrimaryExpression(tree)
	assert.Equal(t, false, result)
}

func TestVisitPrimaryExpressionParentheses(t *testing.T) {
	tree := parsePrimaryExpression("(1 + 2)")
	visitor := NewInterpreter()

	result := visitor.VisitPrimaryExpression(tree)
	assert.Equal(t, 3.0, result)
}
