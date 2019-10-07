// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
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
	INCDEC
	CALL
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
	types.PLUSPLUS:   INCDEC,
	types.MINUSMINUS: INCDEC,
	types.LPAREN:     CALL,
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
	p.prefixFunctions[types.IMPORT] = p.parseImportExpr
	p.prefixFunctions[types.IDENTIFIER] = p.parseIdentifier
	p.prefixFunctions[types.NUMBER] = p.parseNumber
	p.prefixFunctions[types.STRING] = p.parseString
	p.prefixFunctions[types.BUILTIN] = p.parseBuiltin
	p.prefixFunctions[types.TRUE] = p.parseBool
	p.prefixFunctions[types.FALSE] = p.parseBool
	p.prefixFunctions[types.NIL] = p.parseNil
	p.prefixFunctions[types.LPAREN] = p.parseGroup
	p.prefixFunctions[types.NOT] = p.parsePrefix
	p.prefixFunctions[types.PLUSPLUS] = p.parsePrefix
	p.prefixFunctions[types.MINUSMINUS] = p.parsePrefix
	p.prefixFunctions[types.BNOT] = p.parsePrefix
	p.prefixFunctions[types.MINUS] = p.parsePrefix

	p.infixFunctions[types.OROR] = p.parseInfix
	p.infixFunctions[types.ANDAND] = p.parseInfix
	p.infixFunctions[types.EQ] = p.parseAssign
	p.infixFunctions[types.COLONEQ] = p.parseInfix
	p.infixFunctions[types.PLUSEQ] = p.parseInfix
	p.infixFunctions[types.MINUSEQ] = p.parseInfix
	p.infixFunctions[types.MULTIPLYEQ] = p.parseInfix
	p.infixFunctions[types.DIVIDEEQ] = p.parseInfix
	p.infixFunctions[types.MODULOEQ] = p.parseInfix
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
	p.infixFunctions[types.PLUSPLUS] = p.parseSuffix
	p.infixFunctions[types.AND] = p.parseInfix
	p.infixFunctions[types.OR] = p.parseInfix
	p.infixFunctions[types.XOR] = p.parseInfix
	p.infixFunctions[types.LTLT] = p.parseInfix
	p.infixFunctions[types.GTGT] = p.parseInfix
	p.infixFunctions[types.DOT] = p.parseAlienFn
	p.infixFunctions[types.PLUSPLUS] = p.parseSuffix
	p.infixFunctions[types.MINUSMINUS] = p.parseSuffix

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
		console.ThrowParsingError(1, console.PARSER_EXPECTED_FOUND, expected, p.peek.Value, p.peek.Line)
	}
}

func (p *Parser) expect(b bool) {
	if b {
		p.next()
	}
}

func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{}
	program.Body = []ast.Expression{}

	for p.peek.Type != types.EOF {
		if p.peek.Type == types.EOL {
			p.next()
			continue
		}
		expr := p.parseExpression(LOWEST)
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
		p.require(types.IDENTIFIER)
		ident := &ast.Identifier{Value: p.peek.Value}
		p.next()
		identifiers = append(identifiers, ident)
	}

	p.require(types.RPAREN)
	p.next() // skip ')'
	return identifiers
}

func (p *Parser) parseBlock() *ast.Block {
	block := &ast.Block{}
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
