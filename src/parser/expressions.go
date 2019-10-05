// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"strconv"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixFunctions[p.peek.Type]
	if prefix == nil {
		errors.FatalError(errors.F("unexpected '%v' at line %v", p.peek.Value, p.peek.Line), 1)
	}
	leftExpression := prefix()

	for p.peek.Type != types.EOL && precedence < p.precedence() {
		infix := p.infixFunctions[p.peek.Type]
		if infix == nil {
			errors.FatalError(errors.F("unexpected '%v' at line %v", p.peek.Value, p.peek.Line), 1)
		}
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}

func (p *Parser) parseIdentifier() ast.Expression {
	p.require(types.IDENTIFIER)
	expr := &ast.Identifier{Value: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseNumber() ast.Expression {
	p.require(types.NUMBER)
	value, err := strconv.ParseInt(p.peek.Value, 0, 64)

	if err != nil {
		errors.FatalError(errors.F("could not parse %q as number", p.peek.Value), 1)
	}

	expr := &ast.Number{Value: value}
	p.next()
	return expr
}

func (p *Parser) parseString() ast.Expression {
	p.require(types.STRING)
	expr := &ast.String{Value: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseBuiltin() ast.Expression {
	p.require(types.BUILTIN)
	expr := &ast.Builtin{Value: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseBool() ast.Expression {
	expr := &ast.Bool{Value: p.peek.Type == types.TRUE}
	p.next()
	return expr
}

func (p *Parser) parseNil() ast.Expression {
	p.require(types.NIL)
	expr := &ast.Nil{Value: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseGroup() ast.Expression {
	p.require(types.LPAREN)
	p.next() // skip '('
	expr := p.parseExpression(LOWEST)
	p.require(types.RPAREN)
	p.next() // skip ')'
	return expr
}

func (p *Parser) parseCall(fn ast.Expression) ast.Expression {
	p.require(types.LPAREN)
	p.next() // skip '('
	expr := &ast.Call{Caller: fn}
	expr.Arguments = p.parseExpressionList(types.RPAREN)
	p.require(types.RPAREN)
	p.next() // skip ')'
	return expr
}

func (p *Parser) parseExpressionList(end string) []ast.Expression {
	list := []ast.Expression{}

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

func (p *Parser) parseList(expr ast.Expression) ast.Expression {
	list := &ast.ExprList{}
	list.Body = append(list.Body, expr)

	for p.peek.Type == types.COMMA {
		p.next() // skip ','
		list.Body = append(list.Body, p.parseExpression(LOWEST))
	}

	return list
}

func (p *Parser) parseAssign(left ast.Expression) ast.Expression {
	expr := &ast.Assign{Value: p.peek.Type}
	expr.Left = left
	p.require(types.EQ)
	p.next()
	expr.Right = p.parseExpression(LOWEST)
	return expr
}

func (p *Parser) parseFn() ast.Expression {
	p.require(types.FN)
	fn := &ast.Function{Value: p.peek.Type}
	var expr *ast.Assign
	p.next() // skip 'fn'
	if p.peek.Type != types.IDENTIFIER {
		fn.Anonymous = true
	} else {
		fn.Anonymous = false
		expr = &ast.Assign{Value: types.EQ}
		exprs := &ast.ExprList{}
		exprs.Body = append(exprs.Body, p.parseIdentifier())
		expr.Left = exprs
	}

	fn.Params = p.parseParams()
	fn.Body = p.parseBlock()
	if fn.Anonymous {
		return fn
	}
	expr.Right = fn
	return expr
}

func (p *Parser) parseAlienFn(alien ast.Expression) ast.Expression {
	expr := &ast.AlienFn{}
	p.next()
	expr.Alien = alien
	expr.Function = p.parseExpression(DOT)
	return expr
}

func (p *Parser) parseImport() ast.Expression {
	expr := &ast.Import{Value: p.peek.Type}
	p.next()
	expr.Body = p.parseExpression(LOWEST)
	return expr
}

func (p *Parser) parsePrefix() ast.Expression {
	expr := &ast.Prefix{Operator: p.peek.Value}
	p.next()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	expr := &ast.Infix{Operator: p.peek.Value, Left: left}
	precedence := p.precedence()
	p.next()
	expr.Right = p.parseExpression(precedence)
	return expr
}

func (p *Parser) parseSuffix(left ast.Expression) ast.Expression {
	expr := &ast.Suffix{Operator: p.peek.Value, Left: left}
	p.next()
	return expr
}
