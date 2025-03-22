// Code generated from PhpParser.g4 by ANTLR 4.13.0. DO NOT EDIT.

package parser // PhpParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type PhpParser struct {
	*antlr.BaseParser
}

var PhpParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func phpparserParserInit() {
	staticData := &PhpParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'?>'", "'function'", "'class'", "'public'", "'private'", "'protected'",
		"'var'", "'return'", "'if'", "'else'", "'while'", "'for'", "'foreach'",
		"'switch'", "'case'", "'default'", "'break'", "'continue'", "'echo'",
		"'print'", "'new'", "'extends'", "'implements'", "'namespace'", "'use'",
		"'try'", "'catch'", "'finally'", "'throw'", "'static'", "'abstract'",
		"'final'", "'as'", "", "", "", "'='", "'+'", "'-'", "'*'", "'/'", "'%'",
		"'**'", "'&&'", "'||'", "'!'", "'=='", "", "'>'", "'<'", "'>='", "'<='",
		"'.'", "':'", "';'", "','", "'('", "')'", "'{'", "'}'",
	}
	staticData.SymbolicNames = []string{
		"", "T_OPEN_TAG", "T_CLOSE_TAG", "T_FUNCTION", "T_CLASS", "T_PUBLIC",
		"T_PRIVATE", "T_PROTECTED", "T_VAR", "T_RETURN", "T_IF", "T_ELSE", "T_WHILE",
		"T_FOR", "T_FOREACH", "T_SWITCH", "T_CASE", "T_DEFAULT", "T_BREAK",
		"T_CONTINUE", "T_ECHO", "T_PRINT", "T_NEW", "T_EXTENDS", "T_IMPLEMENTS",
		"T_NAMESPACE", "T_USE", "T_TRY", "T_CATCH", "T_FINALLY", "T_THROW",
		"T_STATIC", "T_ABSTRACT", "T_FINAL", "T_AS", "T_BOOLEAN", "T_VARIABLE",
		"T_IDENTIFIER", "T_ASSIGN", "T_PLUS", "T_MINUS", "T_MULT", "T_DIV",
		"T_MOD", "T_POW", "T_AND", "T_OR", "T_NOT", "T_EQUAL", "T_NOT_EQUAL",
		"T_GREATER", "T_LESS", "T_GREATER_EQUAL", "T_LESS_EQUAL", "T_CONCAT",
		"T_COLON", "T_SEMICOLON", "T_COMMA", "T_L_PARENTHESIS", "T_R_PARENTHESIS",
		"T_L_BRACE", "T_R_BRACE", "T_INT", "T_FLOAT", "T_STRING", "T_COMMENT",
		"T_COMMENT_BLOCK", "T_WHITESPACE",
	}
	staticData.RuleNames = []string{
		"program", "statement", "expressionStatement", "expression", "assignmentExpression",
		"unaryExpression", "primaryExpression", "functionDeclaration", "parameterList",
		"parameter", "classDeclaration", "ifStatement", "whileStatement", "forStatement",
		"foreachStatement", "switchStatement", "caseStatement", "returnStatement",
		"echoStatement", "expressionList", "block",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 234, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 1,
		0, 5, 0, 44, 8, 0, 10, 0, 12, 0, 47, 9, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 61, 8, 1, 1, 2, 1, 2, 1,
		2, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 70, 8, 3, 1, 3, 1, 3, 1, 3, 5, 3, 75,
		8, 3, 10, 3, 12, 3, 78, 9, 3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5,
		1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 1, 6, 3, 6, 96, 8, 6, 1,
		7, 1, 7, 1, 7, 1, 7, 3, 7, 102, 8, 7, 1, 7, 1, 7, 1, 7, 1, 8, 1, 8, 1,
		8, 5, 8, 110, 8, 8, 10, 8, 12, 8, 113, 9, 8, 1, 9, 1, 9, 1, 10, 1, 10,
		1, 10, 1, 10, 3, 10, 121, 8, 10, 1, 10, 1, 10, 1, 10, 1, 10, 5, 10, 127,
		8, 10, 10, 10, 12, 10, 130, 9, 10, 3, 10, 132, 8, 10, 1, 10, 1, 10, 1,
		11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 1, 11, 3, 11, 143, 8, 11, 1, 12,
		1, 12, 1, 12, 1, 12, 1, 12, 1, 12, 1, 13, 1, 13, 1, 13, 3, 13, 154, 8,
		13, 1, 13, 1, 13, 3, 13, 158, 8, 13, 1, 13, 1, 13, 3, 13, 162, 8, 13, 1,
		13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14, 1, 14,
		1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 1, 15, 5, 15, 181, 8, 15, 10, 15, 12,
		15, 184, 9, 15, 1, 15, 1, 15, 1, 16, 1, 16, 1, 16, 1, 16, 5, 16, 192, 8,
		16, 10, 16, 12, 16, 195, 9, 16, 1, 16, 1, 16, 1, 16, 5, 16, 200, 8, 16,
		10, 16, 12, 16, 203, 9, 16, 3, 16, 205, 8, 16, 1, 17, 1, 17, 3, 17, 209,
		8, 17, 1, 17, 1, 17, 1, 18, 1, 18, 1, 18, 1, 18, 1, 19, 1, 19, 1, 19, 5,
		19, 220, 8, 19, 10, 19, 12, 19, 223, 9, 19, 1, 20, 1, 20, 5, 20, 227, 8,
		20, 10, 20, 12, 20, 230, 9, 20, 1, 20, 1, 20, 1, 20, 0, 1, 6, 21, 0, 2,
		4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40,
		0, 2, 3, 0, 39, 43, 48, 49, 54, 54, 2, 0, 40, 40, 47, 47, 246, 0, 45, 1,
		0, 0, 0, 2, 60, 1, 0, 0, 0, 4, 62, 1, 0, 0, 0, 6, 69, 1, 0, 0, 0, 8, 79,
		1, 0, 0, 0, 10, 83, 1, 0, 0, 0, 12, 95, 1, 0, 0, 0, 14, 97, 1, 0, 0, 0,
		16, 106, 1, 0, 0, 0, 18, 114, 1, 0, 0, 0, 20, 116, 1, 0, 0, 0, 22, 135,
		1, 0, 0, 0, 24, 144, 1, 0, 0, 0, 26, 150, 1, 0, 0, 0, 28, 166, 1, 0, 0,
		0, 30, 174, 1, 0, 0, 0, 32, 204, 1, 0, 0, 0, 34, 206, 1, 0, 0, 0, 36, 212,
		1, 0, 0, 0, 38, 216, 1, 0, 0, 0, 40, 224, 1, 0, 0, 0, 42, 44, 3, 2, 1,
		0, 43, 42, 1, 0, 0, 0, 44, 47, 1, 0, 0, 0, 45, 43, 1, 0, 0, 0, 45, 46,
		1, 0, 0, 0, 46, 48, 1, 0, 0, 0, 47, 45, 1, 0, 0, 0, 48, 49, 5, 0, 0, 1,
		49, 1, 1, 0, 0, 0, 50, 61, 3, 4, 2, 0, 51, 61, 3, 14, 7, 0, 52, 61, 3,
		20, 10, 0, 53, 61, 3, 22, 11, 0, 54, 61, 3, 24, 12, 0, 55, 61, 3, 26, 13,
		0, 56, 61, 3, 28, 14, 0, 57, 61, 3, 30, 15, 0, 58, 61, 3, 34, 17, 0, 59,
		61, 3, 36, 18, 0, 60, 50, 1, 0, 0, 0, 60, 51, 1, 0, 0, 0, 60, 52, 1, 0,
		0, 0, 60, 53, 1, 0, 0, 0, 60, 54, 1, 0, 0, 0, 60, 55, 1, 0, 0, 0, 60, 56,
		1, 0, 0, 0, 60, 57, 1, 0, 0, 0, 60, 58, 1, 0, 0, 0, 60, 59, 1, 0, 0, 0,
		61, 3, 1, 0, 0, 0, 62, 63, 3, 6, 3, 0, 63, 64, 5, 56, 0, 0, 64, 5, 1, 0,
		0, 0, 65, 66, 6, 3, -1, 0, 66, 70, 3, 8, 4, 0, 67, 70, 3, 10, 5, 0, 68,
		70, 3, 12, 6, 0, 69, 65, 1, 0, 0, 0, 69, 67, 1, 0, 0, 0, 69, 68, 1, 0,
		0, 0, 70, 76, 1, 0, 0, 0, 71, 72, 10, 3, 0, 0, 72, 73, 7, 0, 0, 0, 73,
		75, 3, 6, 3, 4, 74, 71, 1, 0, 0, 0, 75, 78, 1, 0, 0, 0, 76, 74, 1, 0, 0,
		0, 76, 77, 1, 0, 0, 0, 77, 7, 1, 0, 0, 0, 78, 76, 1, 0, 0, 0, 79, 80, 5,
		36, 0, 0, 80, 81, 5, 38, 0, 0, 81, 82, 3, 6, 3, 0, 82, 9, 1, 0, 0, 0, 83,
		84, 7, 1, 0, 0, 84, 85, 3, 6, 3, 0, 85, 11, 1, 0, 0, 0, 86, 96, 5, 36,
		0, 0, 87, 96, 5, 62, 0, 0, 88, 96, 5, 63, 0, 0, 89, 96, 5, 64, 0, 0, 90,
		96, 5, 35, 0, 0, 91, 92, 5, 58, 0, 0, 92, 93, 3, 6, 3, 0, 93, 94, 5, 59,
		0, 0, 94, 96, 1, 0, 0, 0, 95, 86, 1, 0, 0, 0, 95, 87, 1, 0, 0, 0, 95, 88,
		1, 0, 0, 0, 95, 89, 1, 0, 0, 0, 95, 90, 1, 0, 0, 0, 95, 91, 1, 0, 0, 0,
		96, 13, 1, 0, 0, 0, 97, 98, 5, 3, 0, 0, 98, 99, 5, 37, 0, 0, 99, 101, 5,
		58, 0, 0, 100, 102, 3, 16, 8, 0, 101, 100, 1, 0, 0, 0, 101, 102, 1, 0,
		0, 0, 102, 103, 1, 0, 0, 0, 103, 104, 5, 59, 0, 0, 104, 105, 3, 40, 20,
		0, 105, 15, 1, 0, 0, 0, 106, 111, 3, 18, 9, 0, 107, 108, 5, 57, 0, 0, 108,
		110, 3, 18, 9, 0, 109, 107, 1, 0, 0, 0, 110, 113, 1, 0, 0, 0, 111, 109,
		1, 0, 0, 0, 111, 112, 1, 0, 0, 0, 112, 17, 1, 0, 0, 0, 113, 111, 1, 0,
		0, 0, 114, 115, 5, 36, 0, 0, 115, 19, 1, 0, 0, 0, 116, 117, 5, 4, 0, 0,
		117, 120, 5, 37, 0, 0, 118, 119, 5, 23, 0, 0, 119, 121, 5, 37, 0, 0, 120,
		118, 1, 0, 0, 0, 120, 121, 1, 0, 0, 0, 121, 131, 1, 0, 0, 0, 122, 123,
		5, 24, 0, 0, 123, 128, 5, 37, 0, 0, 124, 125, 5, 57, 0, 0, 125, 127, 5,
		37, 0, 0, 126, 124, 1, 0, 0, 0, 127, 130, 1, 0, 0, 0, 128, 126, 1, 0, 0,
		0, 128, 129, 1, 0, 0, 0, 129, 132, 1, 0, 0, 0, 130, 128, 1, 0, 0, 0, 131,
		122, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0, 133, 134,
		3, 40, 20, 0, 134, 21, 1, 0, 0, 0, 135, 136, 5, 10, 0, 0, 136, 137, 5,
		58, 0, 0, 137, 138, 3, 6, 3, 0, 138, 139, 5, 59, 0, 0, 139, 142, 3, 40,
		20, 0, 140, 141, 5, 11, 0, 0, 141, 143, 3, 40, 20, 0, 142, 140, 1, 0, 0,
		0, 142, 143, 1, 0, 0, 0, 143, 23, 1, 0, 0, 0, 144, 145, 5, 12, 0, 0, 145,
		146, 5, 58, 0, 0, 146, 147, 3, 6, 3, 0, 147, 148, 5, 59, 0, 0, 148, 149,
		3, 40, 20, 0, 149, 25, 1, 0, 0, 0, 150, 151, 5, 13, 0, 0, 151, 153, 5,
		58, 0, 0, 152, 154, 3, 6, 3, 0, 153, 152, 1, 0, 0, 0, 153, 154, 1, 0, 0,
		0, 154, 155, 1, 0, 0, 0, 155, 157, 5, 56, 0, 0, 156, 158, 3, 6, 3, 0, 157,
		156, 1, 0, 0, 0, 157, 158, 1, 0, 0, 0, 158, 159, 1, 0, 0, 0, 159, 161,
		5, 56, 0, 0, 160, 162, 3, 6, 3, 0, 161, 160, 1, 0, 0, 0, 161, 162, 1, 0,
		0, 0, 162, 163, 1, 0, 0, 0, 163, 164, 5, 59, 0, 0, 164, 165, 3, 40, 20,
		0, 165, 27, 1, 0, 0, 0, 166, 167, 5, 14, 0, 0, 167, 168, 5, 58, 0, 0, 168,
		169, 5, 36, 0, 0, 169, 170, 5, 34, 0, 0, 170, 171, 5, 36, 0, 0, 171, 172,
		5, 59, 0, 0, 172, 173, 3, 40, 20, 0, 173, 29, 1, 0, 0, 0, 174, 175, 5,
		15, 0, 0, 175, 176, 5, 58, 0, 0, 176, 177, 3, 6, 3, 0, 177, 178, 5, 59,
		0, 0, 178, 182, 5, 60, 0, 0, 179, 181, 3, 32, 16, 0, 180, 179, 1, 0, 0,
		0, 181, 184, 1, 0, 0, 0, 182, 180, 1, 0, 0, 0, 182, 183, 1, 0, 0, 0, 183,
		185, 1, 0, 0, 0, 184, 182, 1, 0, 0, 0, 185, 186, 5, 61, 0, 0, 186, 31,
		1, 0, 0, 0, 187, 188, 5, 16, 0, 0, 188, 189, 3, 6, 3, 0, 189, 193, 5, 55,
		0, 0, 190, 192, 3, 2, 1, 0, 191, 190, 1, 0, 0, 0, 192, 195, 1, 0, 0, 0,
		193, 191, 1, 0, 0, 0, 193, 194, 1, 0, 0, 0, 194, 205, 1, 0, 0, 0, 195,
		193, 1, 0, 0, 0, 196, 197, 5, 17, 0, 0, 197, 201, 5, 55, 0, 0, 198, 200,
		3, 2, 1, 0, 199, 198, 1, 0, 0, 0, 200, 203, 1, 0, 0, 0, 201, 199, 1, 0,
		0, 0, 201, 202, 1, 0, 0, 0, 202, 205, 1, 0, 0, 0, 203, 201, 1, 0, 0, 0,
		204, 187, 1, 0, 0, 0, 204, 196, 1, 0, 0, 0, 205, 33, 1, 0, 0, 0, 206, 208,
		5, 9, 0, 0, 207, 209, 3, 6, 3, 0, 208, 207, 1, 0, 0, 0, 208, 209, 1, 0,
		0, 0, 209, 210, 1, 0, 0, 0, 210, 211, 5, 56, 0, 0, 211, 35, 1, 0, 0, 0,
		212, 213, 5, 20, 0, 0, 213, 214, 3, 38, 19, 0, 214, 215, 5, 56, 0, 0, 215,
		37, 1, 0, 0, 0, 216, 221, 3, 6, 3, 0, 217, 218, 5, 57, 0, 0, 218, 220,
		3, 6, 3, 0, 219, 217, 1, 0, 0, 0, 220, 223, 1, 0, 0, 0, 221, 219, 1, 0,
		0, 0, 221, 222, 1, 0, 0, 0, 222, 39, 1, 0, 0, 0, 223, 221, 1, 0, 0, 0,
		224, 228, 5, 60, 0, 0, 225, 227, 3, 2, 1, 0, 226, 225, 1, 0, 0, 0, 227,
		230, 1, 0, 0, 0, 228, 226, 1, 0, 0, 0, 228, 229, 1, 0, 0, 0, 229, 231,
		1, 0, 0, 0, 230, 228, 1, 0, 0, 0, 231, 232, 5, 61, 0, 0, 232, 41, 1, 0,
		0, 0, 21, 45, 60, 69, 76, 95, 101, 111, 120, 128, 131, 142, 153, 157, 161,
		182, 193, 201, 204, 208, 221, 228,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// PhpParserInit initializes any static state used to implement PhpParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewPhpParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func PhpParserInit() {
	staticData := &PhpParserParserStaticData
	staticData.once.Do(phpparserParserInit)
}

// NewPhpParser produces a new parser instance for the optional input antlr.TokenStream.
func NewPhpParser(input antlr.TokenStream) *PhpParser {
	PhpParserInit()
	this := new(PhpParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &PhpParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "PhpParser.g4"

	return this
}

// PhpParser tokens.
const (
	PhpParserEOF             = antlr.TokenEOF
	PhpParserT_OPEN_TAG      = 1
	PhpParserT_CLOSE_TAG     = 2
	PhpParserT_FUNCTION      = 3
	PhpParserT_CLASS         = 4
	PhpParserT_PUBLIC        = 5
	PhpParserT_PRIVATE       = 6
	PhpParserT_PROTECTED     = 7
	PhpParserT_VAR           = 8
	PhpParserT_RETURN        = 9
	PhpParserT_IF            = 10
	PhpParserT_ELSE          = 11
	PhpParserT_WHILE         = 12
	PhpParserT_FOR           = 13
	PhpParserT_FOREACH       = 14
	PhpParserT_SWITCH        = 15
	PhpParserT_CASE          = 16
	PhpParserT_DEFAULT       = 17
	PhpParserT_BREAK         = 18
	PhpParserT_CONTINUE      = 19
	PhpParserT_ECHO          = 20
	PhpParserT_PRINT         = 21
	PhpParserT_NEW           = 22
	PhpParserT_EXTENDS       = 23
	PhpParserT_IMPLEMENTS    = 24
	PhpParserT_NAMESPACE     = 25
	PhpParserT_USE           = 26
	PhpParserT_TRY           = 27
	PhpParserT_CATCH         = 28
	PhpParserT_FINALLY       = 29
	PhpParserT_THROW         = 30
	PhpParserT_STATIC        = 31
	PhpParserT_ABSTRACT      = 32
	PhpParserT_FINAL         = 33
	PhpParserT_AS            = 34
	PhpParserT_BOOLEAN       = 35
	PhpParserT_VARIABLE      = 36
	PhpParserT_IDENTIFIER    = 37
	PhpParserT_ASSIGN        = 38
	PhpParserT_PLUS          = 39
	PhpParserT_MINUS         = 40
	PhpParserT_MULT          = 41
	PhpParserT_DIV           = 42
	PhpParserT_MOD           = 43
	PhpParserT_POW           = 44
	PhpParserT_AND           = 45
	PhpParserT_OR            = 46
	PhpParserT_NOT           = 47
	PhpParserT_EQUAL         = 48
	PhpParserT_NOT_EQUAL     = 49
	PhpParserT_GREATER       = 50
	PhpParserT_LESS          = 51
	PhpParserT_GREATER_EQUAL = 52
	PhpParserT_LESS_EQUAL    = 53
	PhpParserT_CONCAT        = 54
	PhpParserT_COLON         = 55
	PhpParserT_SEMICOLON     = 56
	PhpParserT_COMMA         = 57
	PhpParserT_L_PARENTHESIS = 58
	PhpParserT_R_PARENTHESIS = 59
	PhpParserT_L_BRACE       = 60
	PhpParserT_R_BRACE       = 61
	PhpParserT_INT           = 62
	PhpParserT_FLOAT         = 63
	PhpParserT_STRING        = 64
	PhpParserT_COMMENT       = 65
	PhpParserT_COMMENT_BLOCK = 66
	PhpParserT_WHITESPACE    = 67
)

// PhpParser rules.
const (
	PhpParserRULE_program              = 0
	PhpParserRULE_statement            = 1
	PhpParserRULE_expressionStatement  = 2
	PhpParserRULE_expression           = 3
	PhpParserRULE_assignmentExpression = 4
	PhpParserRULE_unaryExpression      = 5
	PhpParserRULE_primaryExpression    = 6
	PhpParserRULE_functionDeclaration  = 7
	PhpParserRULE_parameterList        = 8
	PhpParserRULE_parameter            = 9
	PhpParserRULE_classDeclaration     = 10
	PhpParserRULE_ifStatement          = 11
	PhpParserRULE_whileStatement       = 12
	PhpParserRULE_forStatement         = 13
	PhpParserRULE_foreachStatement     = 14
	PhpParserRULE_switchStatement      = 15
	PhpParserRULE_caseStatement        = 16
	PhpParserRULE_returnStatement      = 17
	PhpParserRULE_echoStatement        = 18
	PhpParserRULE_expressionList       = 19
	PhpParserRULE_block                = 20
)

// IProgramContext is an interface to support dynamic dispatch.
type IProgramContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	EOF() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsProgramContext differentiates from other interfaces.
	IsProgramContext()
}

type ProgramContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramContext() *ProgramContext {
	var p = new(ProgramContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_program
	return p
}

func InitEmptyProgramContext(p *ProgramContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_program
}

func (*ProgramContext) IsProgramContext() {}

func NewProgramContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramContext {
	var p = new(ProgramContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_program

	return p
}

func (s *ProgramContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramContext) EOF() antlr.TerminalNode {
	return s.GetToken(PhpParserEOF, 0)
}

func (s *ProgramContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *ProgramContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *ProgramContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterProgram(s)
	}
}

func (s *ProgramContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitProgram(s)
	}
}

func (p *PhpParser) Program() (localctx IProgramContext) {
	localctx = NewProgramContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, PhpParserRULE_program)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(45)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&4071271805652967107) != 0 {
		{
			p.SetState(42)
			p.Statement()
		}

		p.SetState(47)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(48)
		p.Match(PhpParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ExpressionStatement() IExpressionStatementContext
	FunctionDeclaration() IFunctionDeclarationContext
	ClassDeclaration() IClassDeclarationContext
	IfStatement() IIfStatementContext
	WhileStatement() IWhileStatementContext
	ForStatement() IForStatementContext
	ForeachStatement() IForeachStatementContext
	SwitchStatement() ISwitchStatementContext
	ReturnStatement() IReturnStatementContext
	EchoStatement() IEchoStatementContext

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_statement
	return p
}

func InitEmptyStatementContext(p *StatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_statement
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) ExpressionStatement() IExpressionStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionStatementContext)
}

func (s *StatementContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *StatementContext) ClassDeclaration() IClassDeclarationContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IClassDeclarationContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IClassDeclarationContext)
}

func (s *StatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *StatementContext) WhileStatement() IWhileStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhileStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhileStatementContext)
}

func (s *StatementContext) ForStatement() IForStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForStatementContext)
}

func (s *StatementContext) ForeachStatement() IForeachStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IForeachStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IForeachStatementContext)
}

func (s *StatementContext) SwitchStatement() ISwitchStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISwitchStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISwitchStatementContext)
}

func (s *StatementContext) ReturnStatement() IReturnStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IReturnStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IReturnStatementContext)
}

func (s *StatementContext) EchoStatement() IEchoStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEchoStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IEchoStatementContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (p *PhpParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, PhpParserRULE_statement)
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PhpParserT_BOOLEAN, PhpParserT_VARIABLE, PhpParserT_MINUS, PhpParserT_NOT, PhpParserT_L_PARENTHESIS, PhpParserT_INT, PhpParserT_FLOAT, PhpParserT_STRING:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(50)
			p.ExpressionStatement()
		}

	case PhpParserT_FUNCTION:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(51)
			p.FunctionDeclaration()
		}

	case PhpParserT_CLASS:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(52)
			p.ClassDeclaration()
		}

	case PhpParserT_IF:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(53)
			p.IfStatement()
		}

	case PhpParserT_WHILE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(54)
			p.WhileStatement()
		}

	case PhpParserT_FOR:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(55)
			p.ForStatement()
		}

	case PhpParserT_FOREACH:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(56)
			p.ForeachStatement()
		}

	case PhpParserT_SWITCH:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(57)
			p.SwitchStatement()
		}

	case PhpParserT_RETURN:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(58)
			p.ReturnStatement()
		}

	case PhpParserT_ECHO:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(59)
			p.EchoStatement()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionStatementContext is an interface to support dynamic dispatch.
type IExpressionStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	T_SEMICOLON() antlr.TerminalNode

	// IsExpressionStatementContext differentiates from other interfaces.
	IsExpressionStatementContext()
}

type ExpressionStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionStatementContext() *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_expressionStatement
	return p
}

func InitEmptyExpressionStatementContext(p *ExpressionStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_expressionStatement
}

func (*ExpressionStatementContext) IsExpressionStatementContext() {}

func NewExpressionStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_expressionStatement

	return p
}

func (s *ExpressionStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) T_SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PhpParserT_SEMICOLON, 0)
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (p *PhpParser) ExpressionStatement() (localctx IExpressionStatementContext) {
	localctx = NewExpressionStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, PhpParserRULE_expressionStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(62)
		p.expression(0)
	}
	{
		p.SetState(63)
		p.Match(PhpParserT_SEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AssignmentExpression() IAssignmentExpressionContext
	UnaryExpression() IUnaryExpressionContext
	PrimaryExpression() IPrimaryExpressionContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	T_PLUS() antlr.TerminalNode
	T_MINUS() antlr.TerminalNode
	T_MULT() antlr.TerminalNode
	T_DIV() antlr.TerminalNode
	T_MOD() antlr.TerminalNode
	T_CONCAT() antlr.TerminalNode
	T_EQUAL() antlr.TerminalNode
	T_NOT_EQUAL() antlr.TerminalNode

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) AssignmentExpression() IAssignmentExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAssignmentExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAssignmentExpressionContext)
}

func (s *ExpressionContext) UnaryExpression() IUnaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *ExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) T_PLUS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_PLUS, 0)
}

func (s *ExpressionContext) T_MINUS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_MINUS, 0)
}

func (s *ExpressionContext) T_MULT() antlr.TerminalNode {
	return s.GetToken(PhpParserT_MULT, 0)
}

func (s *ExpressionContext) T_DIV() antlr.TerminalNode {
	return s.GetToken(PhpParserT_DIV, 0)
}

func (s *ExpressionContext) T_MOD() antlr.TerminalNode {
	return s.GetToken(PhpParserT_MOD, 0)
}

func (s *ExpressionContext) T_CONCAT() antlr.TerminalNode {
	return s.GetToken(PhpParserT_CONCAT, 0)
}

func (s *ExpressionContext) T_EQUAL() antlr.TerminalNode {
	return s.GetToken(PhpParserT_EQUAL, 0)
}

func (s *ExpressionContext) T_NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(PhpParserT_NOT_EQUAL, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *PhpParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *PhpParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 6
	p.EnterRecursionRule(localctx, 6, PhpParserRULE_expression, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(66)
			p.AssignmentExpression()
		}

	case 2:
		{
			p.SetState(67)
			p.UnaryExpression()
		}

	case 3:
		{
			p.SetState(68)
			p.PrimaryExpression()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewExpressionContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, PhpParserRULE_expression)
			p.SetState(71)

			if !(p.Precpred(p.GetParserRuleContext(), 3)) {
				p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				goto errorExit
			}
			{
				p.SetState(72)
				_la = p.GetTokenStream().LA(1)

				if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&18875865869844480) != 0) {
					p.GetErrorHandler().RecoverInline(p)
				} else {
					p.GetErrorHandler().ReportMatch(p)
					p.Consume()
				}
			}
			{
				p.SetState(73)
				p.expression(4)
			}

		}
		p.SetState(78)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAssignmentExpressionContext is an interface to support dynamic dispatch.
type IAssignmentExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_VARIABLE() antlr.TerminalNode
	T_ASSIGN() antlr.TerminalNode
	Expression() IExpressionContext

	// IsAssignmentExpressionContext differentiates from other interfaces.
	IsAssignmentExpressionContext()
}

type AssignmentExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAssignmentExpressionContext() *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_assignmentExpression
	return p
}

func InitEmptyAssignmentExpressionContext(p *AssignmentExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_assignmentExpression
}

func (*AssignmentExpressionContext) IsAssignmentExpressionContext() {}

func NewAssignmentExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AssignmentExpressionContext {
	var p = new(AssignmentExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_assignmentExpression

	return p
}

func (s *AssignmentExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AssignmentExpressionContext) T_VARIABLE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_VARIABLE, 0)
}

func (s *AssignmentExpressionContext) T_ASSIGN() antlr.TerminalNode {
	return s.GetToken(PhpParserT_ASSIGN, 0)
}

func (s *AssignmentExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AssignmentExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterAssignmentExpression(s)
	}
}

func (s *AssignmentExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitAssignmentExpression(s)
	}
}

func (p *PhpParser) AssignmentExpression() (localctx IAssignmentExpressionContext) {
	localctx = NewAssignmentExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, PhpParserRULE_assignmentExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(79)
		p.Match(PhpParserT_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(80)
		p.Match(PhpParserT_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(81)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	T_NOT() antlr.TerminalNode
	T_MINUS() antlr.TerminalNode

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) T_NOT() antlr.TerminalNode {
	return s.GetToken(PhpParserT_NOT, 0)
}

func (s *UnaryExpressionContext) T_MINUS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_MINUS, 0)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (p *PhpParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, PhpParserRULE_unaryExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(83)
		_la = p.GetTokenStream().LA(1)

		if !(_la == PhpParserT_MINUS || _la == PhpParserT_NOT) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}
	{
		p.SetState(84)
		p.expression(0)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_VARIABLE() antlr.TerminalNode
	T_INT() antlr.TerminalNode
	T_FLOAT() antlr.TerminalNode
	T_STRING() antlr.TerminalNode
	T_BOOLEAN() antlr.TerminalNode
	T_L_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	T_R_PARENTHESIS() antlr.TerminalNode

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) T_VARIABLE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_VARIABLE, 0)
}

func (s *PrimaryExpressionContext) T_INT() antlr.TerminalNode {
	return s.GetToken(PhpParserT_INT, 0)
}

func (s *PrimaryExpressionContext) T_FLOAT() antlr.TerminalNode {
	return s.GetToken(PhpParserT_FLOAT, 0)
}

func (s *PrimaryExpressionContext) T_STRING() antlr.TerminalNode {
	return s.GetToken(PhpParserT_STRING, 0)
}

func (s *PrimaryExpressionContext) T_BOOLEAN() antlr.TerminalNode {
	return s.GetToken(PhpParserT_BOOLEAN, 0)
}

func (s *PrimaryExpressionContext) T_L_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_PARENTHESIS, 0)
}

func (s *PrimaryExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *PrimaryExpressionContext) T_R_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_PARENTHESIS, 0)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (p *PhpParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, PhpParserRULE_primaryExpression)
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PhpParserT_VARIABLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(86)
			p.Match(PhpParserT_VARIABLE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PhpParserT_INT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(87)
			p.Match(PhpParserT_INT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PhpParserT_FLOAT:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(88)
			p.Match(PhpParserT_FLOAT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PhpParserT_STRING:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(89)
			p.Match(PhpParserT_STRING)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PhpParserT_BOOLEAN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(90)
			p.Match(PhpParserT_BOOLEAN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case PhpParserT_L_PARENTHESIS:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(91)
			p.Match(PhpParserT_L_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.expression(0)
		}
		{
			p.SetState(93)
			p.Match(PhpParserT_R_PARENTHESIS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_FUNCTION() antlr.TerminalNode
	T_IDENTIFIER() antlr.TerminalNode
	T_L_PARENTHESIS() antlr.TerminalNode
	T_R_PARENTHESIS() antlr.TerminalNode
	Block() IBlockContext
	ParameterList() IParameterListContext

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_functionDeclaration
	return p
}

func InitEmptyFunctionDeclarationContext(p *FunctionDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_functionDeclaration
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) T_FUNCTION() antlr.TerminalNode {
	return s.GetToken(PhpParserT_FUNCTION, 0)
}

func (s *FunctionDeclarationContext) T_IDENTIFIER() antlr.TerminalNode {
	return s.GetToken(PhpParserT_IDENTIFIER, 0)
}

func (s *FunctionDeclarationContext) T_L_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_PARENTHESIS, 0)
}

func (s *FunctionDeclarationContext) T_R_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_PARENTHESIS, 0)
}

func (s *FunctionDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *FunctionDeclarationContext) ParameterList() IParameterListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterListContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (p *PhpParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, PhpParserRULE_functionDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(97)
		p.Match(PhpParserT_FUNCTION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(98)
		p.Match(PhpParserT_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(99)
		p.Match(PhpParserT_L_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PhpParserT_VARIABLE {
		{
			p.SetState(100)
			p.ParameterList()
		}

	}
	{
		p.SetState(103)
		p.Match(PhpParserT_R_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(104)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterListContext is an interface to support dynamic dispatch.
type IParameterListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllParameter() []IParameterContext
	Parameter(i int) IParameterContext
	AllT_COMMA() []antlr.TerminalNode
	T_COMMA(i int) antlr.TerminalNode

	// IsParameterListContext differentiates from other interfaces.
	IsParameterListContext()
}

type ParameterListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterListContext() *ParameterListContext {
	var p = new(ParameterListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_parameterList
	return p
}

func InitEmptyParameterListContext(p *ParameterListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_parameterList
}

func (*ParameterListContext) IsParameterListContext() {}

func NewParameterListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterListContext {
	var p = new(ParameterListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_parameterList

	return p
}

func (s *ParameterListContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterListContext) AllParameter() []IParameterContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IParameterContext); ok {
			len++
		}
	}

	tst := make([]IParameterContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IParameterContext); ok {
			tst[i] = t.(IParameterContext)
			i++
		}
	}

	return tst
}

func (s *ParameterListContext) Parameter(i int) IParameterContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IParameterContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IParameterContext)
}

func (s *ParameterListContext) AllT_COMMA() []antlr.TerminalNode {
	return s.GetTokens(PhpParserT_COMMA)
}

func (s *ParameterListContext) T_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PhpParserT_COMMA, i)
}

func (s *ParameterListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterParameterList(s)
	}
}

func (s *ParameterListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitParameterList(s)
	}
}

func (p *PhpParser) ParameterList() (localctx IParameterListContext) {
	localctx = NewParameterListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, PhpParserRULE_parameterList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(106)
		p.Parameter()
	}
	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PhpParserT_COMMA {
		{
			p.SetState(107)
			p.Match(PhpParserT_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(108)
			p.Parameter()
		}

		p.SetState(113)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IParameterContext is an interface to support dynamic dispatch.
type IParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_VARIABLE() antlr.TerminalNode

	// IsParameterContext differentiates from other interfaces.
	IsParameterContext()
}

type ParameterContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParameterContext() *ParameterContext {
	var p = new(ParameterContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_parameter
	return p
}

func InitEmptyParameterContext(p *ParameterContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_parameter
}

func (*ParameterContext) IsParameterContext() {}

func NewParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParameterContext {
	var p = new(ParameterContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_parameter

	return p
}

func (s *ParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *ParameterContext) T_VARIABLE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_VARIABLE, 0)
}

func (s *ParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterParameter(s)
	}
}

func (s *ParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitParameter(s)
	}
}

func (p *PhpParser) Parameter() (localctx IParameterContext) {
	localctx = NewParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, PhpParserRULE_parameter)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(114)
		p.Match(PhpParserT_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IClassDeclarationContext is an interface to support dynamic dispatch.
type IClassDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_CLASS() antlr.TerminalNode
	AllT_IDENTIFIER() []antlr.TerminalNode
	T_IDENTIFIER(i int) antlr.TerminalNode
	Block() IBlockContext
	T_EXTENDS() antlr.TerminalNode
	T_IMPLEMENTS() antlr.TerminalNode
	AllT_COMMA() []antlr.TerminalNode
	T_COMMA(i int) antlr.TerminalNode

	// IsClassDeclarationContext differentiates from other interfaces.
	IsClassDeclarationContext()
}

type ClassDeclarationContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyClassDeclarationContext() *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_classDeclaration
	return p
}

func InitEmptyClassDeclarationContext(p *ClassDeclarationContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_classDeclaration
}

func (*ClassDeclarationContext) IsClassDeclarationContext() {}

func NewClassDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ClassDeclarationContext {
	var p = new(ClassDeclarationContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_classDeclaration

	return p
}

func (s *ClassDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ClassDeclarationContext) T_CLASS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_CLASS, 0)
}

func (s *ClassDeclarationContext) AllT_IDENTIFIER() []antlr.TerminalNode {
	return s.GetTokens(PhpParserT_IDENTIFIER)
}

func (s *ClassDeclarationContext) T_IDENTIFIER(i int) antlr.TerminalNode {
	return s.GetToken(PhpParserT_IDENTIFIER, i)
}

func (s *ClassDeclarationContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ClassDeclarationContext) T_EXTENDS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_EXTENDS, 0)
}

func (s *ClassDeclarationContext) T_IMPLEMENTS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_IMPLEMENTS, 0)
}

func (s *ClassDeclarationContext) AllT_COMMA() []antlr.TerminalNode {
	return s.GetTokens(PhpParserT_COMMA)
}

func (s *ClassDeclarationContext) T_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PhpParserT_COMMA, i)
}

func (s *ClassDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ClassDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ClassDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterClassDeclaration(s)
	}
}

func (s *ClassDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitClassDeclaration(s)
	}
}

func (p *PhpParser) ClassDeclaration() (localctx IClassDeclarationContext) {
	localctx = NewClassDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, PhpParserRULE_classDeclaration)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(116)
		p.Match(PhpParserT_CLASS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(117)
		p.Match(PhpParserT_IDENTIFIER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PhpParserT_EXTENDS {
		{
			p.SetState(118)
			p.Match(PhpParserT_EXTENDS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(119)
			p.Match(PhpParserT_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	p.SetState(131)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PhpParserT_IMPLEMENTS {
		{
			p.SetState(122)
			p.Match(PhpParserT_IMPLEMENTS)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(123)
			p.Match(PhpParserT_IDENTIFIER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(128)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == PhpParserT_COMMA {
			{
				p.SetState(124)
				p.Match(PhpParserT_COMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(125)
				p.Match(PhpParserT_IDENTIFIER)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

			p.SetState(130)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(133)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_IF() antlr.TerminalNode
	T_L_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	T_R_PARENTHESIS() antlr.TerminalNode
	AllBlock() []IBlockContext
	Block(i int) IBlockContext
	T_ELSE() antlr.TerminalNode

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) T_IF() antlr.TerminalNode {
	return s.GetToken(PhpParserT_IF, 0)
}

func (s *IfStatementContext) T_L_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_PARENTHESIS, 0)
}

func (s *IfStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfStatementContext) T_R_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_PARENTHESIS, 0)
}

func (s *IfStatementContext) AllBlock() []IBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBlockContext); ok {
			len++
		}
	}

	tst := make([]IBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBlockContext); ok {
			tst[i] = t.(IBlockContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Block(i int) IBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *IfStatementContext) T_ELSE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_ELSE, 0)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (p *PhpParser) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, PhpParserRULE_ifStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(135)
		p.Match(PhpParserT_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(136)
		p.Match(PhpParserT_L_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(137)
		p.expression(0)
	}
	{
		p.SetState(138)
		p.Match(PhpParserT_R_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(139)
		p.Block()
	}
	p.SetState(142)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == PhpParserT_ELSE {
		{
			p.SetState(140)
			p.Match(PhpParserT_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(141)
			p.Block()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhileStatementContext is an interface to support dynamic dispatch.
type IWhileStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_WHILE() antlr.TerminalNode
	T_L_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	T_R_PARENTHESIS() antlr.TerminalNode
	Block() IBlockContext

	// IsWhileStatementContext differentiates from other interfaces.
	IsWhileStatementContext()
}

type WhileStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhileStatementContext() *WhileStatementContext {
	var p = new(WhileStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_whileStatement
	return p
}

func InitEmptyWhileStatementContext(p *WhileStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_whileStatement
}

func (*WhileStatementContext) IsWhileStatementContext() {}

func NewWhileStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhileStatementContext {
	var p = new(WhileStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_whileStatement

	return p
}

func (s *WhileStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *WhileStatementContext) T_WHILE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_WHILE, 0)
}

func (s *WhileStatementContext) T_L_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_PARENTHESIS, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) T_R_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_PARENTHESIS, 0)
}

func (s *WhileStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (p *PhpParser) WhileStatement() (localctx IWhileStatementContext) {
	localctx = NewWhileStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, PhpParserRULE_whileStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(144)
		p.Match(PhpParserT_WHILE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(145)
		p.Match(PhpParserT_L_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(146)
		p.expression(0)
	}
	{
		p.SetState(147)
		p.Match(PhpParserT_R_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(148)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForStatementContext is an interface to support dynamic dispatch.
type IForStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_FOR() antlr.TerminalNode
	T_L_PARENTHESIS() antlr.TerminalNode
	AllT_SEMICOLON() []antlr.TerminalNode
	T_SEMICOLON(i int) antlr.TerminalNode
	T_R_PARENTHESIS() antlr.TerminalNode
	Block() IBlockContext
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext

	// IsForStatementContext differentiates from other interfaces.
	IsForStatementContext()
}

type ForStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForStatementContext() *ForStatementContext {
	var p = new(ForStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_forStatement
	return p
}

func InitEmptyForStatementContext(p *ForStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_forStatement
}

func (*ForStatementContext) IsForStatementContext() {}

func NewForStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForStatementContext {
	var p = new(ForStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_forStatement

	return p
}

func (s *ForStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForStatementContext) T_FOR() antlr.TerminalNode {
	return s.GetToken(PhpParserT_FOR, 0)
}

func (s *ForStatementContext) T_L_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_PARENTHESIS, 0)
}

func (s *ForStatementContext) AllT_SEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(PhpParserT_SEMICOLON)
}

func (s *ForStatementContext) T_SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(PhpParserT_SEMICOLON, i)
}

func (s *ForStatementContext) T_R_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_PARENTHESIS, 0)
}

func (s *ForStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForStatementContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ForStatementContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ForStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterForStatement(s)
	}
}

func (s *ForStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitForStatement(s)
	}
}

func (p *PhpParser) ForStatement() (localctx IForStatementContext) {
	localctx = NewForStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, PhpParserRULE_forStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(150)
		p.Match(PhpParserT_FOR)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(151)
		p.Match(PhpParserT_L_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(153)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&947916835) != 0 {
		{
			p.SetState(152)
			p.expression(0)
		}

	}
	{
		p.SetState(155)
		p.Match(PhpParserT_SEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(157)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&947916835) != 0 {
		{
			p.SetState(156)
			p.expression(0)
		}

	}
	{
		p.SetState(159)
		p.Match(PhpParserT_SEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(161)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&947916835) != 0 {
		{
			p.SetState(160)
			p.expression(0)
		}

	}
	{
		p.SetState(163)
		p.Match(PhpParserT_R_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(164)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IForeachStatementContext is an interface to support dynamic dispatch.
type IForeachStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_FOREACH() antlr.TerminalNode
	T_L_PARENTHESIS() antlr.TerminalNode
	AllT_VARIABLE() []antlr.TerminalNode
	T_VARIABLE(i int) antlr.TerminalNode
	T_AS() antlr.TerminalNode
	T_R_PARENTHESIS() antlr.TerminalNode
	Block() IBlockContext

	// IsForeachStatementContext differentiates from other interfaces.
	IsForeachStatementContext()
}

type ForeachStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyForeachStatementContext() *ForeachStatementContext {
	var p = new(ForeachStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_foreachStatement
	return p
}

func InitEmptyForeachStatementContext(p *ForeachStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_foreachStatement
}

func (*ForeachStatementContext) IsForeachStatementContext() {}

func NewForeachStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ForeachStatementContext {
	var p = new(ForeachStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_foreachStatement

	return p
}

func (s *ForeachStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ForeachStatementContext) T_FOREACH() antlr.TerminalNode {
	return s.GetToken(PhpParserT_FOREACH, 0)
}

func (s *ForeachStatementContext) T_L_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_PARENTHESIS, 0)
}

func (s *ForeachStatementContext) AllT_VARIABLE() []antlr.TerminalNode {
	return s.GetTokens(PhpParserT_VARIABLE)
}

func (s *ForeachStatementContext) T_VARIABLE(i int) antlr.TerminalNode {
	return s.GetToken(PhpParserT_VARIABLE, i)
}

func (s *ForeachStatementContext) T_AS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_AS, 0)
}

func (s *ForeachStatementContext) T_R_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_PARENTHESIS, 0)
}

func (s *ForeachStatementContext) Block() IBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *ForeachStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ForeachStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ForeachStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterForeachStatement(s)
	}
}

func (s *ForeachStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitForeachStatement(s)
	}
}

func (p *PhpParser) ForeachStatement() (localctx IForeachStatementContext) {
	localctx = NewForeachStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, PhpParserRULE_foreachStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(166)
		p.Match(PhpParserT_FOREACH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(167)
		p.Match(PhpParserT_L_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(168)
		p.Match(PhpParserT_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(169)
		p.Match(PhpParserT_AS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(170)
		p.Match(PhpParserT_VARIABLE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(171)
		p.Match(PhpParserT_R_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(172)
		p.Block()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISwitchStatementContext is an interface to support dynamic dispatch.
type ISwitchStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_SWITCH() antlr.TerminalNode
	T_L_PARENTHESIS() antlr.TerminalNode
	Expression() IExpressionContext
	T_R_PARENTHESIS() antlr.TerminalNode
	T_L_BRACE() antlr.TerminalNode
	T_R_BRACE() antlr.TerminalNode
	AllCaseStatement() []ICaseStatementContext
	CaseStatement(i int) ICaseStatementContext

	// IsSwitchStatementContext differentiates from other interfaces.
	IsSwitchStatementContext()
}

type SwitchStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchStatementContext() *SwitchStatementContext {
	var p = new(SwitchStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_switchStatement
	return p
}

func InitEmptySwitchStatementContext(p *SwitchStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_switchStatement
}

func (*SwitchStatementContext) IsSwitchStatementContext() {}

func NewSwitchStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_switchStatement

	return p
}

func (s *SwitchStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchStatementContext) T_SWITCH() antlr.TerminalNode {
	return s.GetToken(PhpParserT_SWITCH, 0)
}

func (s *SwitchStatementContext) T_L_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_PARENTHESIS, 0)
}

func (s *SwitchStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStatementContext) T_R_PARENTHESIS() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_PARENTHESIS, 0)
}

func (s *SwitchStatementContext) T_L_BRACE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_BRACE, 0)
}

func (s *SwitchStatementContext) T_R_BRACE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_BRACE, 0)
}

func (s *SwitchStatementContext) AllCaseStatement() []ICaseStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ICaseStatementContext); ok {
			len++
		}
	}

	tst := make([]ICaseStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ICaseStatementContext); ok {
			tst[i] = t.(ICaseStatementContext)
			i++
		}
	}

	return tst
}

func (s *SwitchStatementContext) CaseStatement(i int) ICaseStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICaseStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (p *PhpParser) SwitchStatement() (localctx ISwitchStatementContext) {
	localctx = NewSwitchStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, PhpParserRULE_switchStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(174)
		p.Match(PhpParserT_SWITCH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(175)
		p.Match(PhpParserT_L_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.expression(0)
	}
	{
		p.SetState(177)
		p.Match(PhpParserT_R_PARENTHESIS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(178)
		p.Match(PhpParserT_L_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(182)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PhpParserT_CASE || _la == PhpParserT_DEFAULT {
		{
			p.SetState(179)
			p.CaseStatement()
		}

		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(185)
		p.Match(PhpParserT_R_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICaseStatementContext is an interface to support dynamic dispatch.
type ICaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_CASE() antlr.TerminalNode
	Expression() IExpressionContext
	T_COLON() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext
	T_DEFAULT() antlr.TerminalNode

	// IsCaseStatementContext differentiates from other interfaces.
	IsCaseStatementContext()
}

type CaseStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCaseStatementContext() *CaseStatementContext {
	var p = new(CaseStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_caseStatement
	return p
}

func InitEmptyCaseStatementContext(p *CaseStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_caseStatement
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) T_CASE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_CASE, 0)
}

func (s *CaseStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseStatementContext) T_COLON() antlr.TerminalNode {
	return s.GetToken(PhpParserT_COLON, 0)
}

func (s *CaseStatementContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *CaseStatementContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseStatementContext) T_DEFAULT() antlr.TerminalNode {
	return s.GetToken(PhpParserT_DEFAULT, 0)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterCaseStatement(s)
	}
}

func (s *CaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitCaseStatement(s)
	}
}

func (p *PhpParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, PhpParserRULE_caseStatement)
	var _la int

	p.SetState(204)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case PhpParserT_CASE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(187)
			p.Match(PhpParserT_CASE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(188)
			p.expression(0)
		}
		{
			p.SetState(189)
			p.Match(PhpParserT_COLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(193)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&4071271805652967107) != 0 {
			{
				p.SetState(190)
				p.Statement()
			}

			p.SetState(195)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case PhpParserT_DEFAULT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(196)
			p.Match(PhpParserT_DEFAULT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(197)
			p.Match(PhpParserT_COLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(201)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&4071271805652967107) != 0 {
			{
				p.SetState(198)
				p.Statement()
			}

			p.SetState(203)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IReturnStatementContext is an interface to support dynamic dispatch.
type IReturnStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_RETURN() antlr.TerminalNode
	T_SEMICOLON() antlr.TerminalNode
	Expression() IExpressionContext

	// IsReturnStatementContext differentiates from other interfaces.
	IsReturnStatementContext()
}

type ReturnStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyReturnStatementContext() *ReturnStatementContext {
	var p = new(ReturnStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_returnStatement
	return p
}

func InitEmptyReturnStatementContext(p *ReturnStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_returnStatement
}

func (*ReturnStatementContext) IsReturnStatementContext() {}

func NewReturnStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_returnStatement

	return p
}

func (s *ReturnStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ReturnStatementContext) T_RETURN() antlr.TerminalNode {
	return s.GetToken(PhpParserT_RETURN, 0)
}

func (s *ReturnStatementContext) T_SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PhpParserT_SEMICOLON, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (p *PhpParser) ReturnStatement() (localctx IReturnStatementContext) {
	localctx = NewReturnStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, PhpParserRULE_returnStatement)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(206)
		p.Match(PhpParserT_RETURN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-35)) & ^0x3f) == 0 && ((int64(1)<<(_la-35))&947916835) != 0 {
		{
			p.SetState(207)
			p.expression(0)
		}

	}
	{
		p.SetState(210)
		p.Match(PhpParserT_SEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IEchoStatementContext is an interface to support dynamic dispatch.
type IEchoStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_ECHO() antlr.TerminalNode
	ExpressionList() IExpressionListContext
	T_SEMICOLON() antlr.TerminalNode

	// IsEchoStatementContext differentiates from other interfaces.
	IsEchoStatementContext()
}

type EchoStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEchoStatementContext() *EchoStatementContext {
	var p = new(EchoStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_echoStatement
	return p
}

func InitEmptyEchoStatementContext(p *EchoStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_echoStatement
}

func (*EchoStatementContext) IsEchoStatementContext() {}

func NewEchoStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EchoStatementContext {
	var p = new(EchoStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_echoStatement

	return p
}

func (s *EchoStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *EchoStatementContext) T_ECHO() antlr.TerminalNode {
	return s.GetToken(PhpParserT_ECHO, 0)
}

func (s *EchoStatementContext) ExpressionList() IExpressionListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionListContext)
}

func (s *EchoStatementContext) T_SEMICOLON() antlr.TerminalNode {
	return s.GetToken(PhpParserT_SEMICOLON, 0)
}

func (s *EchoStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EchoStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EchoStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterEchoStatement(s)
	}
}

func (s *EchoStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitEchoStatement(s)
	}
}

func (p *PhpParser) EchoStatement() (localctx IEchoStatementContext) {
	localctx = NewEchoStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, PhpParserRULE_echoStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(212)
		p.Match(PhpParserT_ECHO)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(213)
		p.ExpressionList()
	}
	{
		p.SetState(214)
		p.Match(PhpParserT_SEMICOLON)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionListContext is an interface to support dynamic dispatch.
type IExpressionListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllT_COMMA() []antlr.TerminalNode
	T_COMMA(i int) antlr.TerminalNode

	// IsExpressionListContext differentiates from other interfaces.
	IsExpressionListContext()
}

type ExpressionListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionListContext() *ExpressionListContext {
	var p = new(ExpressionListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_expressionList
	return p
}

func InitEmptyExpressionListContext(p *ExpressionListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_expressionList
}

func (*ExpressionListContext) IsExpressionListContext() {}

func NewExpressionListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionListContext {
	var p = new(ExpressionListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_expressionList

	return p
}

func (s *ExpressionListContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ExpressionListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionListContext) AllT_COMMA() []antlr.TerminalNode {
	return s.GetTokens(PhpParserT_COMMA)
}

func (s *ExpressionListContext) T_COMMA(i int) antlr.TerminalNode {
	return s.GetToken(PhpParserT_COMMA, i)
}

func (s *ExpressionListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterExpressionList(s)
	}
}

func (s *ExpressionListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitExpressionList(s)
	}
}

func (p *PhpParser) ExpressionList() (localctx IExpressionListContext) {
	localctx = NewExpressionListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, PhpParserRULE_expressionList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.expression(0)
	}
	p.SetState(221)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == PhpParserT_COMMA {
		{
			p.SetState(217)
			p.Match(PhpParserT_COMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(218)
			p.expression(0)
		}

		p.SetState(223)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	T_L_BRACE() antlr.TerminalNode
	T_R_BRACE() antlr.TerminalNode
	AllStatement() []IStatementContext
	Statement(i int) IStatementContext

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_block
	return p
}

func InitEmptyBlockContext(p *BlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = PhpParserRULE_block
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = PhpParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) T_L_BRACE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_L_BRACE, 0)
}

func (s *BlockContext) T_R_BRACE() antlr.TerminalNode {
	return s.GetToken(PhpParserT_R_BRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IStatementContext); ok {
			len++
		}
	}

	tst := make([]IStatementContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IStatementContext); ok {
			tst[i] = t.(IStatementContext)
			i++
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatementContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(PhpParserListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (p *PhpParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, PhpParserRULE_block)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Match(PhpParserT_L_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(228)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&4071271805652967107) != 0 {
		{
			p.SetState(225)
			p.Statement()
		}

		p.SetState(230)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(231)
		p.Match(PhpParserT_R_BRACE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *PhpParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 3:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *PhpParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
