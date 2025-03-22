package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xoesae/php-go/parser"
	"testing"
)

func parseAssignmentExpression(input string) *parser.AssignmentExpressionContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	return p.AssignmentExpression().(*parser.AssignmentExpressionContext)
}

func TestVisitAssignmentExpressionInteger(t *testing.T) {
	tree := parseAssignmentExpression("$x = 42")
	visitor := &Interpreter{
		variables: make(map[string]interface{}),
	}
	result := visitor.VisitAssignmentExpression(tree)

	assert.Equal(t, 42, visitor.variables["$x"])
	assert.Equal(t, 42, result)
}

func TestVisitAssignmentExpressionString(t *testing.T) {
	tree := parseAssignmentExpression("$name = \"Carlos\"")
	visitor := NewInterpreter()

	result := visitor.VisitAssignmentExpression(tree)

	assert.Equal(t, "Carlos", visitor.variables["$name"])
	assert.Equal(t, "Carlos", result)
}

func TestVisitAssignmentExpressionBoolean(t *testing.T) {
	tree := parseAssignmentExpression("$isActive = true")
	visitor := NewInterpreter()

	result := visitor.VisitAssignmentExpression(tree)

	assert.Equal(t, true, visitor.variables["$isActive"])
	assert.Equal(t, true, result)
}

func TestVisitAssignmentExpressionExpression(t *testing.T) {
	tree := parseAssignmentExpression("$sum = 10 + 5")
	visitor := NewInterpreter()

	result := visitor.VisitAssignmentExpression(tree)

	assert.Equal(t, 15.0, visitor.variables["$sum"])
	assert.Equal(t, 15.0, result)
}
