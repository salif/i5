// Adapted from https://github.com/prologic/monkey-lang/blob/v1.3.5/parser/parser.go

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/errors"
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
	PRO
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
	types.DOT:        PRO,
}

func Run(tokens types.TokenList) ast.Program {
	var parser Parser = Parser{
		tokenlist:       types.TokenList{},
		position:        0,
		prefixFunctions: make(map[string]prefixFunction),
		infixFunctions:  make(map[string]infixFunction)}

	parser.Init(tokens)
	return parser.parseProgram()
}

func (p *Parser) Init(tokens types.TokenList) {
	p.tokenlist = tokens
	p.position = 0
	p.peek = types.Token{}

	p.prefixFunctions[types.IDENTIFIER] = p.parseIdentifier
	p.prefixFunctions[types.NUMBER] = p.parseNumber
	p.prefixFunctions[types.STRING] = p.parseString
	p.prefixFunctions[types.BUILTIN] = p.parseBuiltin
	p.prefixFunctions[types.TRUE] = p.parseBool
	p.prefixFunctions[types.FALSE] = p.parseBool
	p.prefixFunctions[types.NIL] = p.parseNil
	p.prefixFunctions[types.LPAREN] = p.parseGroup

	p.infixFunctions[types.OROR] = p.parseInfix
	p.infixFunctions[types.ANDAND] = p.parseInfix
	p.infixFunctions[types.EQ] = p.parseInfix
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

	p.next()
}

type (
	prefixFunction func() ast.Expression
	infixFunction  func(ast.Expression) ast.Expression

	Parser struct {
		tokenlist       types.TokenList
		position        int
		peek            types.Token
		prefixFunctions map[string]prefixFunction
		infixFunctions  map[string]infixFunction
	}
)

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
		errors.FatalError(errors.F(errors.PARSER_EXPECTED_FOUND, expected, p.peek.Value, p.peek.Line), 1)
	}
}

func (p *Parser) expect(b bool) {
	if b {
		p.next()
	}
}

func (p *Parser) parseProgram() ast.Program {
	program := ast.Program{}
	program.Body = []ast.Statement{}

	for p.peek.Type != types.EOF {
		if p.peek.Type == types.EOL {
			p.next()
			continue
		}
		stmt := p.parseStatement()
		if stmt != nil {
			program.Body = append(program.Body, stmt)
		}
	}

	return program
}

func (p *Parser) parseParams() []ast.Identifier {
	identifiers := []ast.Identifier{}
	p.require(types.LPAREN)
	p.next() // skip '('

	if p.peek.Value == types.RPAREN {
		p.next()
		return identifiers
	}

	p.require(types.IDENTIFIER)
	identifiers = append(identifiers, ast.Identifier{Token: p.peek, Val: p.peek.Value})
	p.next()

	for p.peek.Value == types.COMMA {
		p.next() // skip ','
		p.require(types.IDENTIFIER)
		identifiers = append(identifiers, ast.Identifier{Token: p.peek, Val: p.peek.Value})
		p.next() // skip argument
	}

	p.require(types.RPAREN)
	p.next() // skip ')'
	return identifiers
}

func (p *Parser) parseBlock() ast.Block {
	block := ast.Block{Token: p.peek}
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
	p.expect(p.peek.Type == types.EOL)
	return block
}
