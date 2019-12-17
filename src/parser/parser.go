// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/lexer"
)

const (
	_ int = iota
	LOWEST
	ASSIGN   // = := += -= *= /= %=
	IF       // ?? ::
	OR       // or
	AND      // and
	EQ       // == != < > <= >=
	CONCAT   // :
	QM       // ?
	BTWOR    // |
	BTWXOR   // ^
	BTWAND   // &
	BTWSHIFT // << >>
	SUM      // + -
	PRODUCT  // * / %
	PREFIX   // prefix
	POSTFIX  // postfix
	CALL     // (
	DOT      // .
)

var precedences = map[string]int{
	constants.TOKEN_EQ:         ASSIGN,
	constants.TOKEN_COLONEQ:    ASSIGN,
	constants.TOKEN_PLUSEQ:     ASSIGN,
	constants.TOKEN_MINUSEQ:    ASSIGN,
	constants.TOKEN_MULTIPLYEQ: ASSIGN,
	constants.TOKEN_DIVIDEEQ:   ASSIGN,
	constants.TOKEN_MODULOEQ:   ASSIGN,
	constants.TOKEN_QMQM:       IF,
	constants.TOKEN_OROR:       OR,
	constants.TOKEN_ANDAND:     AND,
	constants.TOKEN_EQEQ:       EQ,
	constants.TOKEN_NOTEQ:      EQ,
	constants.TOKEN_LT:         EQ,
	constants.TOKEN_GT:         EQ,
	constants.TOKEN_LTEQ:       EQ,
	constants.TOKEN_GTEQ:       EQ,
	constants.TOKEN_COLON:      CONCAT,
	constants.TOKEN_QM:         QM,
	constants.TOKEN_OR:         BTWOR,
	constants.TOKEN_XOR:        BTWXOR,
	constants.TOKEN_AND:        BTWAND,
	constants.TOKEN_LTLT:       BTWSHIFT,
	constants.TOKEN_GTGT:       BTWSHIFT,
	constants.TOKEN_PLUS:       SUM,
	constants.TOKEN_MINUS:      SUM,
	constants.TOKEN_MULTIPLY:   PRODUCT,
	constants.TOKEN_DIVIDE:     PRODUCT,
	constants.TOKEN_MODULO:     PRODUCT,
	constants.TOKEN_LPAREN:     CALL,
	constants.TOKEN_DOT:        DOT,
}

type (
	prefixFunction func() (ast.Node, error)
	infixFunction  func(ast.Node) (ast.Node, error)
)

type Parser struct {
	code            []byte
	tokens          []constants.Token
	position        int
	peek            constants.Token
	currentFile     string
	prefixFunctions map[string]prefixFunction
	infixFunctions  map[string]infixFunction
}

func ParseProgram(fileName string, code []byte) (ast.Node, error) {
	var p Parser
	p.Init(fileName, code)
	tokens, err := lexer.Run(p.currentFile, p.code)
	if err != nil {
		return nil, err
	}
	p.tokens = tokens
	p.next()
	return p.parseProgram()
}

func Parse(fileName string, code []byte) (ast.Node, error) {
	var p Parser
	p.Init(fileName, code)
	tokens, err := lexer.Run(p.currentFile, p.code)
	if err != nil {
		return nil, err
	}
	p.tokens = tokens
	p.next()
	return p.parseStatement()
}

func (p *Parser) Init(fileName string, code []byte) {
	p.code = code
	p.tokens = make([]constants.Token, 0)
	p.position = 0
	p.peek = constants.Token{}
	p.currentFile = fileName
	p.prefixFunctions = make(map[string]prefixFunction)
	p.infixFunctions = make(map[string]infixFunction)
	p.prefixFunctions[constants.TOKEN_IDENTIFIER] = p.parseIdentifier
	p.prefixFunctions[constants.TOKEN_INTEGER] = p.parseInteger
	p.prefixFunctions[constants.TOKEN_FLOAT] = p.parseFloat
	p.prefixFunctions[constants.TOKEN_STRING] = p.parseString
	p.prefixFunctions[constants.TOKEN_BUILTIN] = p.parseBuiltin
	p.prefixFunctions[constants.TOKEN_FN] = p.parseFunctionExpr
	p.prefixFunctions[constants.TOKEN_LPAREN] = p.parseGroup
	p.prefixFunctions[constants.TOKEN_NOT] = p.parsePrefix
	p.prefixFunctions[constants.TOKEN_BNOT] = p.parsePrefix
	p.prefixFunctions[constants.TOKEN_MINUS] = p.parsePrefix
	p.infixFunctions[constants.TOKEN_EQ] = p.parseAssign
	p.infixFunctions[constants.TOKEN_COLONEQ] = p.parseAssign
	p.infixFunctions[constants.TOKEN_PLUSEQ] = p.parseAssign
	p.infixFunctions[constants.TOKEN_MINUSEQ] = p.parseAssign
	p.infixFunctions[constants.TOKEN_MULTIPLYEQ] = p.parseAssign
	p.infixFunctions[constants.TOKEN_DIVIDEEQ] = p.parseAssign
	p.infixFunctions[constants.TOKEN_MODULOEQ] = p.parseAssign
	p.infixFunctions[constants.TOKEN_QMQM] = p.parseTernary
	p.infixFunctions[constants.TOKEN_OROR] = p.parseInfix
	p.infixFunctions[constants.TOKEN_ANDAND] = p.parseInfix
	p.infixFunctions[constants.TOKEN_EQEQ] = p.parseInfix
	p.infixFunctions[constants.TOKEN_NOTEQ] = p.parseInfix
	p.infixFunctions[constants.TOKEN_LT] = p.parseInfix
	p.infixFunctions[constants.TOKEN_GT] = p.parseInfix
	p.infixFunctions[constants.TOKEN_LTEQ] = p.parseInfix
	p.infixFunctions[constants.TOKEN_GTEQ] = p.parseInfix
	p.infixFunctions[constants.TOKEN_COLON] = p.parseInfix
	p.infixFunctions[constants.TOKEN_QM] = p.parseInfix
	p.infixFunctions[constants.TOKEN_OR] = p.parseInfix
	p.infixFunctions[constants.TOKEN_XOR] = p.parseInfix
	p.infixFunctions[constants.TOKEN_AND] = p.parseInfix
	p.infixFunctions[constants.TOKEN_LTLT] = p.parseInfix
	p.infixFunctions[constants.TOKEN_GTGT] = p.parseInfix
	p.infixFunctions[constants.TOKEN_PLUS] = p.parseInfix
	p.infixFunctions[constants.TOKEN_MINUS] = p.parseInfix
	p.infixFunctions[constants.TOKEN_MULTIPLY] = p.parseInfix
	p.infixFunctions[constants.TOKEN_DIVIDE] = p.parseInfix
	p.infixFunctions[constants.TOKEN_MODULO] = p.parseInfix
	p.infixFunctions[constants.TOKEN_LPAREN] = p.parseCall
	p.infixFunctions[constants.TOKEN_DOT] = p.parseIndex
}

func (p *Parser) next() {
	if p.position < len(p.tokens) {
		p.peek = p.tokens[p.position]
	} else {
		p.peek = constants.Token{Type: constants.TOKEN_EOF, Value: constants.TOKEN_EOF, Line: 0}
	}
	p.position++
}

func (p *Parser) Throw(line uint32, text string, format ...interface{}) error {
	return fmt.Errorf("%v%v\n%v%v:%v\n", colors.Red("parsing error: "), fmt.Sprintf(text, format...), colors.Red("in: "), p.currentFile, line)
}

func (p *Parser) precedence() int {
	if p, ok := precedences[p.peek.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) require(found string, expected string) error {
	if found == expected {
		return nil
	} else {
		return p.Throw(p.peek.Line, constants.PARSER_EXPECTED_FOUND, expected, found)
	}
}
