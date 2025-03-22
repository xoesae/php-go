package main

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/xoesae/php-go/interpreter"
	"github.com/xoesae/php-go/parser"
	"log"
	"os"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	level := os.Getenv("LOG_LEVEL")
	if level == "" {
		level = "error"
	}

	ll, err := logrus.ParseLevel(level)
	if err != nil {
		ll = logrus.DebugLevel
	}

	logrus.SetLevel(ll)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter the file to be executed.")
		return
	}

	filepath := os.Args[1]
	content, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("Error on read file:", err)
		return
	}

	is := antlr.NewInputStream(string(content))
	lexer := parser.NewPhpLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewPhpParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Program()
	visitor := interpreter.NewInterpreter()

	logrus.Debug("AST", tree.ToStringTree(nil, p))

	visitor.Visit(tree)
}
