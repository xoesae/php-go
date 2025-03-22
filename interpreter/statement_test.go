package interpreter

import (
	"bytes"
	"github.com/antlr4-go/antlr/v4"
	"github.com/stretchr/testify/assert"
	"github.com/xoesae/php-go/parser"
	"os"
	"testing"
)

func parseStatement(input string) *parser.StatementContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	return p.Statement().(*parser.StatementContext)
}

func captureOutput(f func()) string {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	return buf.String()
}

func TestVisitExpressionStatement(t *testing.T) {
	tree := parseStatement("$x = 10")
	visitor := NewInterpreter()

	result := visitor.VisitStatement(tree)

	assert.Equal(t, 10, visitor.variables["$x"])
	assert.Nil(t, result)
}

func TestVisitIfStatement(t *testing.T) {
	tree := parseStatement("if (5 == 5) { $x = 10 } else { $x = 20 }")
	visitor := NewInterpreter()

	visitor.VisitStatement(tree)

	assert.Equal(t, 10, visitor.variables["$x"])
}

func TestVisitEchoStatement(t *testing.T) {
	tree := parseStatement("echo \"Hello, World!\"")
	visitor := NewInterpreter()

	output := captureOutput(func() {
		visitor.VisitStatement(tree)
	})

	assert.Equal(t, "Hello, World!\n", output)
}
