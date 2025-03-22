package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xoesae/php-go/parser"
	"testing"
)

func parseIfStatement(input string) *parser.IfStatementContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	return p.IfStatement().(*parser.IfStatementContext)
}

func TestVisitIfStatementTrueCondition(t *testing.T) {
	tree := parseIfStatement("if (5 == 5) { $x = 10; } else { $x = 20; }")
	visitor := NewInterpreter()

	visitor.VisitIfStatement(tree)

	assert.Equal(t, 10, visitor.variables["$x"])
}

func TestVisitIfStatementFalseCondition(t *testing.T) {
	tree := parseIfStatement("if (5 != 5) { $x = 10; } else { $x = 20; }")
	visitor := NewInterpreter()

	visitor.VisitIfStatement(tree)

	assert.Equal(t, 20, visitor.variables["$x"])
}

func TestVisitIfStatementNoElse(t *testing.T) {
	tree := parseIfStatement("if (5 == 5) { $x = 10; }")
	visitor := NewInterpreter()

	visitor.VisitIfStatement(tree)

	assert.Equal(t, 10, visitor.variables["$x"])
}

func TestVisitIfStatementElseConditionNotExecuted(t *testing.T) {
	tree := parseIfStatement("if (5 != 5) { $x = 10; } else { $x = 20; }")
	visitor := NewInterpreter()

	visitor.VisitIfStatement(tree)

	assert.Equal(t, 20, visitor.variables["$x"])
}
