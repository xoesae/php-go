package interpreter

import (
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xoesae/php-go/parser"
	"testing"
)

func parseBlock(input string) *parser.BlockContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	return p.Block().(*parser.BlockContext)
}

func TestVisitBlockEmpty(t *testing.T) {
	tree := parseBlock("{}")
	visitor := NewInterpreter()

	result := visitor.VisitBlock(tree)

	assert.Nil(t, result)
}

func TestVisitBlockSingleStatement(t *testing.T) {
	tree := parseBlock("{ $x = 10; }")
	visitor := NewInterpreter()

	result := visitor.VisitBlock(tree)

	assert.Equal(t, 10, visitor.variables["$x"])
	assert.Nil(t, result)
}

func TestVisitBlockMultipleStatements(t *testing.T) {
	tree := parseBlock("{ $x = 5; $y = $x + 3; }")
	visitor := NewInterpreter()

	result := visitor.VisitBlock(tree)

	assert.Equal(t, 5, visitor.variables["$x"])
	assert.Equal(t, 8.0, visitor.variables["$y"])
	assert.Nil(t, result)
}

func TestVisitBlockStatementsWithExpressions(t *testing.T) {
	tree := parseBlock("{ $x = 2 + 3; $y = $x * 2; }")
	visitor := NewInterpreter()

	result := visitor.VisitBlock(tree)

	assert.Equal(t, 5.0, visitor.variables["$x"])
	assert.Equal(t, 10.0, visitor.variables["$y"])
	assert.Nil(t, result)
}
