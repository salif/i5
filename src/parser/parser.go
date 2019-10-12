// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

const (
	_ int = iota
	LOWEST
	OR
	AND
	NOT
	ASSIGN
	EQ
	CONCAT
	LTGT
	BTWOR
	BTWXOR
	BTWAND
	BTWSHIFT
	SUM
	PRODUCT
	PREFIX
	CALL
	QM
	DOT
)

var precedences = map[string]int{
	types.OROR:       OR,
	types.ANDAND:     AND,
	types.NOT:        NOT,
	types.EQ:         ASSIGN,
	types.COLONEQ:    ASSIGN,
	types.PLUSEQ:     ASSIGN,
	types.MINUSEQ:    ASSIGN,
	types.MULTIPLYEQ: ASSIGN,
	types.DIVIDEEQ:   ASSIGN,
	types.MODULOEQ:   ASSIGN,
	types.EQEQ:       EQ,
	types.NOTEQ:      EQ,
	types.COLON:      CONCAT,
	types.LT:         LTGT,
	types.GT:         LTGT,
	types.LTEQ:       LTGT,
	types.GTEQ:       LTGT,
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
	types.QM:         QM,
	types.DOT:        DOT,
}

func Run(tokens types.TokenList) *ast.Program {
	parser := Parser{
		tokenlist:       types.TokenList{},
		position:        0,
		prefixFunctions: make(map[string]prefixFunction),
		infixFunctions:  make(map[string]infixFunction)}

	parser.Init(tokens)
	return parser.parseProgram()
}

type Parser struct {
	tokenlist       types.TokenList
	position        int
	peek            types.Token
	prefixFunctions map[string]prefixFunction
	infixFunctions  map[string]infixFunction
}

type (
	prefixFunction func() ast.Expression
	infixFunction  func(ast.Expression) ast.Expression
)

func (p *Parser) Init(tokens types.TokenList) {
	p.tokenlist = tokens
	p.position = 0
	p.peek = types.Token{}

	p.prefixFunctions[types.FN] = p.parseFn
	p.prefixFunctions[types.IMPORT] = p.parseImport
	p.prefixFunctions[types.IDENT] = p.parseIdentifier
	p.prefixFunctions[types.INT] = p.parseInteger
	p.prefixFunctions[types.FLOAT] = p.parseFloat
	p.prefixFunctions[types.STRING] = p.parseString
	p.prefixFunctions[types.BUILTIN] = p.parseBuiltin
	p.prefixFunctions[types.TRUE] = p.parseBool
	p.prefixFunctions[types.FALSE] = p.parseBool
	p.prefixFunctions[types.LPAREN] = p.parseGroup
	p.prefixFunctions[types.NOT] = p.parsePrefix
	p.prefixFunctions[types.BNOT] = p.parsePrefix
	p.prefixFunctions[types.MINUS] = p.parsePrefix

	p.infixFunctions[types.OROR] = p.parseInfix
	p.infixFunctions[types.ANDAND] = p.parseInfix
	p.infixFunctions[types.EQ] = p.parseAssign
	p.infixFunctions[types.COLONEQ] = p.parseAssign
	p.infixFunctions[types.PLUSEQ] = p.parseAssign
	p.infixFunctions[types.MINUSEQ] = p.parseAssign
	p.infixFunctions[types.MULTIPLYEQ] = p.parseAssign
	p.infixFunctions[types.DIVIDEEQ] = p.parseAssign
	p.infixFunctions[types.MODULOEQ] = p.parseAssign
	p.infixFunctions[types.EQEQ] = p.parseInfix
	p.infixFunctions[types.NOTEQ] = p.parseInfix
	p.infixFunctions[types.COLON] = p.parseInfix
	p.infixFunctions[types.LT] = p.parseInfix
	p.infixFunctions[types.GT] = p.parseInfix
	p.infixFunctions[types.LTEQ] = p.parseInfix
	p.infixFunctions[types.GTEQ] = p.parseInfix
	p.infixFunctions[types.PLUS] = p.parseInfix
	p.infixFunctions[types.MINUS] = p.parseInfix
	p.infixFunctions[types.MULTIPLY] = p.parseInfix
	p.infixFunctions[types.DIVIDE] = p.parseInfix
	p.infixFunctions[types.MODULO] = p.parseInfix
	p.infixFunctions[types.LPAREN] = p.parseCall
	p.infixFunctions[types.AND] = p.parseInfix
	p.infixFunctions[types.OR] = p.parseInfix
	p.infixFunctions[types.XOR] = p.parseInfix
	p.infixFunctions[types.LTLT] = p.parseInfix
	p.infixFunctions[types.GTGT] = p.parseInfix
	p.infixFunctions[types.DOT] = p.parseAlienFn
	p.infixFunctions[types.QM] = p.parseSuffix

	p.next()
}

func (p *Parser) next() {
	p.peek = p.tokenlist.Get(p.position)
	p.position++
}

func (p *Parser) precedence() int {
	if p, ok := precedences[p.peek.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) require(expected string) {
	if expected != p.peek.Type {
		console.ThrowParsingError(1, constants.PARSER_EXPECTED_FOUND, p.peek.Line, expected, p.peek.Value)
	}
}

func (p *Parser) expect(b bool) {
	if b {
		p.next()
	}
}

func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{Line: p.peek.Line}
	program.Body = []ast.Expression{}

	for p.peek.Type != types.EOF {
		if p.peek.Type == types.EOL {
			p.next()
			continue
		}
		expr := p.parseExpression(LOWEST)
		if _, ok := expr.(*ast.Assign); !ok {
			console.ThrowParsingError(1, constants.PARSER_EXPECTED, expr.GetLine(), "declaration")
		}
		program.Body = append(program.Body, expr)
	}

	return program
}

func (p *Parser) parseParams() []*ast.Identifier {
	identifiers := []*ast.Identifier{}
	p.require(types.LPAREN)
	p.next() // skip '('

	if p.peek.Value == types.RPAREN {
		p.next()
		return identifiers
	}

	for p.peek.Type != types.RPAREN {
		p.require(types.IDENT)
		ident := &ast.Identifier{Line: p.peek.Line, Value: p.peek.Value}
		p.next()
		identifiers = append(identifiers, ident)
	}

	p.require(types.RPAREN)
	p.next() // skip ')'
	return identifiers
}

func (p *Parser) parseBlock() *ast.Block {
	block := &ast.Block{Line: p.peek.Line}
	p.require(types.LBRACE)
	p.next() // skip '{'
	p.require(types.EOL)
	p.next() // skip EOL

	for p.peek.Value != types.RBRACE {
		if p.peek.Type == types.EOL {
			p.next() // skip empty line
			continue
		}
		stmt := p.parseStatement()
		block.Body = append(block.Body, stmt)
	}

	p.require(types.RBRACE)
	p.next() // skip '}'
	return block
}
