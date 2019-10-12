// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"strconv"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseExpression(precedence int) ast.Expression {
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

func (p *Parser) parseIdentifier() ast.Expression {
	p.require(types.IDENT)
	expr := &ast.Identifier{Line: p.peek.Line, Value: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseInteger() ast.Expression {
	p.require(types.INT)
	value, err := strconv.ParseInt(p.peek.Value, 0, 64)

	if err != nil {
		console.ThrowParsingError(1, constants.PARSER_NOT_INT, p.peek.Line, p.peek.Value)
	}

	expr := &ast.Integer{Line: p.peek.Line, Value: value}
	p.next()
	return expr
}

func (p *Parser) parseFloat() ast.Expression {
	p.require(types.FLOAT)
	value, err := strconv.ParseFloat(p.peek.Value, 64)

	if err != nil {
		console.ThrowParsingError(1, constants.PARSER_NOT_FLOAT, p.peek.Line, p.peek.Value)
	}

	expr := &ast.Float{Line: p.peek.Line, Value: value}
	p.next()
	return expr
}

func (p *Parser) parseString() ast.Expression {
	p.require(types.STRING)
	expr := &ast.String{Line: p.peek.Line, Value: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseBuiltin() ast.Expression {
	p.require(types.BUILTIN)
	expr := &ast.Builtin{Line: p.peek.Line, Value: p.peek.Value}
	p.next()
	return expr
}

func (p *Parser) parseBool() ast.Expression {
	expr := &ast.Bool{Line: p.peek.Line, Value: p.peek.Type == types.TRUE}
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
	expr := &ast.Call{Line: p.peek.Line, Caller: fn}
	expr.Arguments = p.parseList(types.RPAREN)
	//console.ThrowParsingError(1, console.PARSER_EXPECTED_ARG, p.peek.Value, p.peek.Line)
	p.require(types.RPAREN)
	p.next() // skip ')'
	return expr
}

func (p *Parser) parseList(end string) []ast.Expression {
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

func (p *Parser) parseAssign(left ast.Expression) ast.Expression {
	expr := &ast.Assign{Line: p.peek.Line, Value: types.EQ, Left: left}
	switch p.peek.Type {
	case types.EQ:
		p.next()
		expr.Right = p.parseExpression(LOWEST)
	case types.COLONEQ:
		p.next()
		expr.Right = &ast.Infix{Line: p.peek.Line, Left: left, Operator: types.COLON, Right: p.parseExpression(LOWEST)}
	case types.PLUSEQ:
		p.next()
		expr.Right = &ast.Infix{Line: p.peek.Line, Left: left, Operator: types.PLUS, Right: p.parseExpression(LOWEST)}
	case types.MINUSEQ:
		p.next()
		expr.Right = &ast.Infix{Line: p.peek.Line, Left: left, Operator: types.MINUS, Right: p.parseExpression(LOWEST)}
	case types.MULTIPLYEQ:
		p.next()
		expr.Right = &ast.Infix{Line: p.peek.Line, Left: left, Operator: types.MULTIPLY, Right: p.parseExpression(LOWEST)}
	case types.DIVIDEEQ:
		p.next()
		expr.Right = &ast.Infix{Line: p.peek.Line, Left: left, Operator: types.DIVIDE, Right: p.parseExpression(LOWEST)}
	case types.MODULOEQ:
		p.next()
		expr.Right = &ast.Infix{Line: p.peek.Line, Left: left, Operator: types.MODULO, Right: p.parseExpression(LOWEST)}
	default:
		console.ThrowParsingError(1, constants.PARSER_UNEXPECTED, p.peek.Line, p.peek.Type)
	}
	return expr
}

func (p *Parser) parseFn() ast.Expression {
	p.require(types.FN)
	fn := &ast.Function{Line: p.peek.Line, Value: p.peek.Type}
	var expr *ast.Assign
	p.next() // skip 'fn'
	if p.peek.Type != types.IDENT {
		fn.Anonymous = true
	} else {
		fn.Anonymous = false
		expr = &ast.Assign{Line: p.peek.Line, Value: types.EQ}
		expr.Left = p.parseIdentifier()
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
	expr := &ast.AlienFn{Line: p.peek.Line, Alien: alien}
	p.next()
	expr.Function = p.parseExpression(DOT)
	return expr
}

func (p *Parser) parseImportExpr() ast.Expression {
	expr := &ast.ImportExpr{Line: p.peek.Line, Value: p.peek.Type}
	p.next()
	expr.Body = p.parseExpression(LOWEST)
	return expr
}

func (p *Parser) parsePrefix() ast.Expression {
	expr := &ast.Prefix{Line: p.peek.Line, Operator: p.peek.Value}
	p.next()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	expr := &ast.Infix{Line: p.peek.Line, Operator: p.peek.Value, Left: left}
	precedence := p.precedence()
	p.next()
	expr.Right = p.parseExpression(precedence)
	return expr
}

func (p *Parser) parseSuffix(left ast.Expression) ast.Expression {
	expr := &ast.Suffix{Line: p.peek.Line, Operator: p.peek.Value, Left: left}
	p.next()
	return expr
}
