// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"strconv"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseExpression(precedence int) ast.Node {
	prefix := p.prefixFunctions[p.peek.Type]
	if prefix == nil {
		console.ThrowParsingError(1, constants.PARSER_UNEXPECTED, p.peek.Line, p.peek.Value)
	}
	leftExpression := prefix()

	for p.peek.Type != types.EOL && precedence < p.precedence() {
		infix := p.infixFunctions[p.peek.Type]
		if infix == nil {
			console.ThrowParsingError(1, constants.PARSER_UNEXPECTED, p.peek.Line, p.peek.Value)
		}
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (p *Parser) parseIdentifier() ast.Node {
	p.require(types.IDENT)
	expr := ast.Identifier{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return expr
}

func (p *Parser) parseInteger() ast.Node {
	p.require(types.INT)
	value, err := strconv.ParseInt(p.peek.Value, 0, 64)

	if err != nil {
		console.ThrowParsingError(1, constants.PARSER_NOT_INT, p.peek.Line, p.peek.Value)
	}

	expr := ast.Integer{}.Init(p.peek.Line, value)
	p.next()
	return expr
}

func (p *Parser) parseFloat() ast.Node {
	p.require(types.FLOAT)
	value, err := strconv.ParseFloat(p.peek.Value, 64)

	if err != nil {
		console.ThrowParsingError(1, constants.PARSER_NOT_FLOAT, p.peek.Line, p.peek.Value)
	}

	expr := ast.Float{}.Init(p.peek.Line, value)
	p.next()
	return expr
}

func (p *Parser) parseString() ast.Node {
	p.require(types.STRING)
	expr := ast.String{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return expr
}

func (p *Parser) parseBuiltin() ast.Node {
	p.require(types.BUILTIN)
	expr := ast.Builtin{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return expr
}

func (p *Parser) parseBool() ast.Node {
	expr := ast.Bool{}.Init(p.peek.Line, p.peek.Type == types.TRUE)
	p.next()
	return expr
}

func (p *Parser) parseGroup() ast.Node {
	p.require(types.LPAREN)
	p.next() // skip '('
	expr := p.parseExpression(LOWEST)
	p.require(types.RPAREN)
	p.next() // skip ')'
	return expr
}

func (p *Parser) parseCall(fn ast.Node) ast.Node {
	p.require(types.LPAREN)
	p.next() // skip '('
	expr := ast.Call{}.Init(p.peek.Line, fn)
	expr.SetArguments(p.parseList(types.RPAREN))
	//console.ThrowParsingError(1, console.PARSER_EXPECTED_ARG, p.peek.Value, p.peek.Line)
	p.require(types.RPAREN)
	p.next() // skip ')'
	return expr
}

func (p *Parser) parseList(end string) []ast.Node {
	var list []ast.Node

	if p.peek.Type == end {
		return list
	}

	list = append(list, p.parseExpression(LOWEST))

	for p.peek.Type == types.COMMA {
		p.next() // skip ','
		list = append(list, p.parseExpression(LOWEST))
	}

	return list
}

func (p *Parser) parseAssign(left ast.Node) ast.Node {
	expr := ast.Assign{}.Init(p.peek.Line, types.EQ, left)
	switch p.peek.Type {
	case types.EQ:
		p.next()
		expr.SetRight(p.parseExpression(LOWEST))
	case types.COLONEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.COLON, p.parseExpression(LOWEST)))
	case types.PLUSEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.PLUS, p.parseExpression(LOWEST)))
	case types.MINUSEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MINUS, p.parseExpression(LOWEST)))
	case types.MULTIPLYEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MULTIPLY, p.parseExpression(LOWEST)))
	case types.DIVIDEEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.DIVIDE, p.parseExpression(LOWEST)))
	case types.MODULOEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MODULO, p.parseExpression(LOWEST)))
	default:
		console.ThrowParsingError(1, constants.PARSER_UNEXPECTED, p.peek.Line, p.peek.Type)
	}
	return expr
}

func (p *Parser) parseFn() ast.Node {
	p.require(types.FN)
	fn := ast.Function{}.Init(p.peek.Line, p.peek.Type)
	var expr ast.Assign
	p.next() // skip 'fn'
	if p.peek.Type != types.IDENT {
		fn.SetAnonymous(true)
	} else {
		fn.SetAnonymous(false)
		expr = ast.Assign{}.Init(p.peek.Line, types.EQ, p.parseIdentifier())
	}

	fn.SetParams(p.parseParams())
	fn.SetBody(p.parseBlock())
	if fn.GetAnonymous() {
		return fn
	}
	expr.SetRight(fn)
	return expr
}

func (p *Parser) parseIndex(left ast.Node) ast.Node {
	expr := ast.Index{}.Init(p.peek.Line, left, p.peek.Value)
	p.next()
	expr.SetRight(p.parseExpression(DOT))
	return expr
}

func (p *Parser) parseImport() ast.Node {
	expr := ast.Import{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	expr.SetBody(p.parseExpression(LOWEST))
	return expr
}

func (p *Parser) parsePrefix() ast.Node {
	expr := ast.Prefix{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	expr.SetRight(p.parseExpression(PREFIX))
	return expr
}

func (p *Parser) parseInfix(left ast.Node) ast.Node {
	expr := ast.Infix{}.Init(p.peek.Line, p.peek.Value, left)
	precedence := p.precedence()
	p.next()
	expr.SetRight(p.parseExpression(precedence))
	return expr
}

func (p *Parser) parsePostfix(left ast.Node) ast.Node {
	expr := ast.Postfix{}.Init(p.peek.Line, p.peek.Value, left)
	p.next()
	return expr
}
