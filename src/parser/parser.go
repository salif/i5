// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/types"
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
	types.EQ:         ASSIGN,
	types.COLONEQ:    ASSIGN,
	types.PLUSEQ:     ASSIGN,
	types.MINUSEQ:    ASSIGN,
	types.MULTIPLYEQ: ASSIGN,
	types.DIVIDEEQ:   ASSIGN,
	types.MODULOEQ:   ASSIGN,
	types.QMQM:       IF,
	types.OROR:       OR,
	types.ANDAND:     AND,
	types.EQEQ:       EQ,
	types.NOTEQ:      EQ,
	types.LT:         EQ,
	types.GT:         EQ,
	types.LTEQ:       EQ,
	types.GTEQ:       EQ,
	types.COLON:      CONCAT,
	types.QM:         QM,
	types.OR:         BTWOR,
	types.XOR:        BTWXOR,
	types.AND:        BTWAND,
	types.LTLT:       BTWSHIFT,
	types.GTGT:       BTWSHIFT,
	types.PLUS:       SUM,
	types.MINUS:      SUM,
	types.MULTIPLY:   PRODUCT,
	types.DIVIDE:     PRODUCT,
	types.MODULO:     PRODUCT,
	types.LPAREN:     CALL,
	types.DOT:        DOT,
}

func Run(fileName string, tokens types.TokenList) (ast.Node, error) {
	var parser Parser
	parser.Init(fileName, tokens)
	return parser.parseProgram()
}

type Parser struct {
	tokenlist       types.TokenList
	position        int
	peek            types.Token
	currentFile     string
	prefixFunctions map[string]prefixFunction
	infixFunctions  map[string]infixFunction
}

type (
	prefixFunction func() (ast.Node, error)
	infixFunction  func(ast.Node) (ast.Node, error)
)

func (p *Parser) Init(fileName string, tokens types.TokenList) {
	p.tokenlist = tokens
	p.position = 0
	p.peek = types.Token{}
	p.currentFile = fileName
	p.prefixFunctions = make(map[string]prefixFunction)
	p.infixFunctions = make(map[string]infixFunction)

	p.prefixFunctions[types.IDENT] = p.parseIdentifier
	p.prefixFunctions[types.INT] = p.parseInteger
	p.prefixFunctions[types.FLOAT] = p.parseFloat
	p.prefixFunctions[types.STRING] = p.parseString
	p.prefixFunctions[types.BUILTIN] = p.parseBuiltin
	p.prefixFunctions[types.LAMBDA] = p.parseLambda
	p.prefixFunctions[types.LPAREN] = p.parseGroup
	p.prefixFunctions[types.NOT] = p.parsePrefix
	p.prefixFunctions[types.BNOT] = p.parsePrefix
	p.prefixFunctions[types.MINUS] = p.parsePrefix

	p.infixFunctions[types.EQ] = p.parseAssign
	p.infixFunctions[types.COLONEQ] = p.parseAssign
	p.infixFunctions[types.PLUSEQ] = p.parseAssign
	p.infixFunctions[types.MINUSEQ] = p.parseAssign
	p.infixFunctions[types.MULTIPLYEQ] = p.parseAssign
	p.infixFunctions[types.DIVIDEEQ] = p.parseAssign
	p.infixFunctions[types.MODULOEQ] = p.parseAssign
	p.infixFunctions[types.QMQM] = p.parseTernary
	p.infixFunctions[types.OROR] = p.parseInfix
	p.infixFunctions[types.ANDAND] = p.parseInfix
	p.infixFunctions[types.EQEQ] = p.parseInfix
	p.infixFunctions[types.NOTEQ] = p.parseInfix
	p.infixFunctions[types.LT] = p.parseInfix
	p.infixFunctions[types.GT] = p.parseInfix
	p.infixFunctions[types.LTEQ] = p.parseInfix
	p.infixFunctions[types.GTEQ] = p.parseInfix
	p.infixFunctions[types.COLON] = p.parseInfix
	p.infixFunctions[types.QM] = p.parseInfix
	p.infixFunctions[types.OR] = p.parseInfix
	p.infixFunctions[types.XOR] = p.parseInfix
	p.infixFunctions[types.AND] = p.parseInfix
	p.infixFunctions[types.LTLT] = p.parseInfix
	p.infixFunctions[types.GTGT] = p.parseInfix
	p.infixFunctions[types.PLUS] = p.parseInfix
	p.infixFunctions[types.MINUS] = p.parseInfix
	p.infixFunctions[types.MULTIPLY] = p.parseInfix
	p.infixFunctions[types.DIVIDE] = p.parseInfix
	p.infixFunctions[types.MODULO] = p.parseInfix
	p.infixFunctions[types.LPAREN] = p.parseCall
	p.infixFunctions[types.DOT] = p.parseIndex

	p.next()
}

func (p *Parser) next() {
	p.peek = p.tokenlist.Get(p.position)
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
