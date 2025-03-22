package interpreter

import (
	"bytes"
	"github.com/antlr4-go/antlr/v4"
	"github.com/xoesae/php-go/parser"
	"os"
	"testing"
)

func parseEcho(input string) *parser.EchoStatementContext {
	lexer := parser.NewPhpLexer(antlr.NewInputStream(input))
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)

	return p.EchoStatement().(*parser.EchoStatementContext)
}

func TestEcho(t *testing.T) {
	tree := parseEcho(`echo "Hello, World!";`)
	interpreter := NewInterpreter()

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	interpreter.VisitEchoStatement(tree)
	w.Close()
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("Expected: %q, received: %q", expected, output)
	}
}
