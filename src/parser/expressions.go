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
	expr := ast.Identifier{Token: p.peek, Val: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseNumber() ast.Expression {
	p.require(types.NUMBER)
	value, err := strconv.ParseInt(p.peek.Value, 0, 64)

	if err != nil {
		errors.FatalError(errors.F("could not parse %q as number", p.peek.Value), 1)
	}

	expr := ast.Number{Token: p.peek, Val: value}
	p.next()
	return expr
}

func (p *Parser) parseString() ast.Expression {
	p.require(types.STRING)
	expr := ast.String{Token: p.peek, Val: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseBuiltin() ast.Expression {
	p.require(types.BUILTIN)
	expr := ast.Builtin{Token: p.peek, Val: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseBool() ast.Expression {
	expr := ast.Bool{Token: p.peek, Val: p.peek.Type == types.TRUE}
	p.next()
	return expr
}

func (p *Parser) parseNil() ast.Expression {
	p.require(types.NIL)
	expr := ast.Nil{Token: p.peek}
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
	expr := ast.Call{Token: p.peek, Function: fn}
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

func (p *Parser) parsePrefix() ast.Expression {
	expression := ast.Prefix{
		Token:    p.peek,
		Operator: p.peek.Value,
	}

	p.next()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	expression := ast.Infix{
		Token:    p.peek,
		Operator: p.peek.Value,
		Left:     left,
	}

	precedence := p.precedence()
	p.next()
	expression.Right = p.parseExpression(precedence)
	return expression
}
