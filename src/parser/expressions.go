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
	expr := ast.Identifier{Token: p.peek, Val: p.peek.Value, Strict: false}
	p.next()
	if p.peek.Type == types.META {
		expr.Type = p.peek
		expr.Strict = true
		p.next()
	}
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

func (p *Parser) parseMeta() ast.Expression {
	p.require(types.META)
	expr := ast.Meta{Token: p.peek, Val: p.peek.Value}
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

func (p *Parser) parseList(expr ast.Expression) ast.Expression {
	list := ast.ExprList{}
	list.Exprs = append(list.Exprs, expr)

	for p.peek.Type == types.COMMA {
		p.next() // skip ','
		list.Exprs = append(list.Exprs, p.parseExpression(LOWEST))
	}

	return list
}

func (p *Parser) parseAssign(left ast.Expression) ast.Expression {
	expr := ast.Assign{Token: p.peek}
	expr.Left = left
	p.require(types.EQ)
	p.next()
	expr.Right = p.parseExpression(LOWEST)
	return expr
}

func (p *Parser) parseFn() ast.Expression {
	p.require(types.FN)
	fn := ast.Function{Token: p.peek}
	var expr ast.Assign
	p.next() // skip 'fn'
	if p.peek.Type != types.IDENTIFIER {
		fn.Anonymous = true
	} else {
		fn.Anonymous = false
		expr = ast.Assign{Token: types.Token{Type: types.EQ, Value: types.EQ, Line: p.peek.Line}}
		exprs := ast.ExprList{}
		exprs.Exprs = append(exprs.Exprs, p.parseIdentifier())
		expr.Left = exprs
	}

	fn.Params = p.parseParams()
	if p.peek.Type == types.META {
		list := p.parseList(p.parseExpression(LOWEST))
		fn.Return = list
		fn.Strict = true
	}
	fn.Body = p.parseBlock()
	if fn.Anonymous {
		return fn
	}
	expr.Right = fn
	return expr
}

// func (p *Parser) parseFn() ast.Expression {
// 	p.require(types.FN)
// 	peek := p.peek
// 	p.next() // skip 'fn'
// 	if p.peek.Type != types.IDENTIFIER {
// 		fn := ast.Function{Token: peek}
// 		fn.Anonymous = true
// 		fn.Params = p.parseParams()
// 		if p.peek.Type == types.META {
// 			returns := []types.Token{}
// 			fn.Return = returns
// 			fn.Strict = true
// 			p.next()
// 		}
// 		fn.Body = p.parseBlock()
// 		return fn
// 	}
// 	expr := ast.Assign{Token: types.Token{Type: types.EQ, Value: types.EQ, Line: p.peek.Line}}
// 	fn := ast.Function{Token: peek}
// 	exprs := ast.ExprList{}
// 	exprs.Exprs = append(exprs.Exprs, p.parseIdentifier())
// 	expr.Left = exprs
// 	fn.Anonymous = false
// 	fn.Params = p.parseParams()
// 	if p.peek.Type == types.META {
// 		returns := []types.Token{}
// 		fn.Return = returns
// 		fn.Strict = true
// 		p.next()
// 	}
// 	fn.Body = p.parseBlock()
// 	expr.Right = fn
// 	return expr
// }

func (p *Parser) parseAlienFn(alien ast.Expression) ast.Expression {
	p.next()
	expr := ast.AlienFn{Token: p.peek}
	expr.Alien = alien
	expr.Function = p.parseExpression(DOT)
	return expr
}

func (p *Parser) parsePrefix() ast.Expression {
	expr := ast.Prefix{
		Token:    p.peek,
		Operator: p.peek.Value,
	}

	p.next()
	expr.Right = p.parseExpression(PREFIX)
	return expr
}

func (p *Parser) parseInfix(left ast.Expression) ast.Expression {
	expr := ast.Infix{
		Token:    p.peek,
		Operator: p.peek.Value,
		Left:     left,
	}

	precedence := p.precedence()
	p.next()
	expr.Right = p.parseExpression(precedence)
	return expr
}

func (p *Parser) parseSuffix(left ast.Expression) ast.Expression {
	expr := ast.Suffix{
		Token: p.peek,
	}

	expr.Left = left
	expr.Operator = p.peek.Value
	p.next()
	return expr
}
