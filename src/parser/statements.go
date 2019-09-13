package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.peek.Type {
	case types.IF:
		return p.parseIf()
	case types.SWITCH:
		return p.parseSwitch()
	case types.FOR:
		return p.parseFor()
	case types.RETURN:
		return p.parseReturn()
	case types.THROW:
		return p.parseThrow()
	default:
		return p.parseExprStatement()
	}
}

func (p *Parser) parseExprStatement() ast.Expr {
	stmt := ast.Expr{Token: p.peek}
	stmt.Expr = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseIf() ast.Statement {
	expression := ast.If{Token: p.peek}

	p.next() // skip 'if' or 'elif'

	expression.Condition = p.parseExpression(LOWEST)

	expression.Consequence = p.parseBlock()

	if p.peek.Type == types.ELIF {
		expression.Alternative = ast.Block{Body: []ast.Statement{p.parseIf()}}
	} else if p.peek.Type == types.ELSE {
		p.next() // skip 'else'
		expression.Alternative = p.parseBlock()
	}

	return expression
}

func (p *Parser) parseSwitch() ast.Statement {
	stmt := ast.Switch{Token: p.peek}
	p.next()
	stmt.Condition = p.parseExpression(LOWEST)
	cases := []ast.Case{}
	cs := ast.Case{}
	p.require(types.EOL)
	p.next()

	for p.peek.Type == types.CASE {
		p.next()
		expr := p.parseExpression(LOWEST)
		cs.Cases = append(cs.Cases, expr)
		if p.peek.Type == types.LBRACE {
			cs.Body = p.parseBlock()
			cases = append(cases, cs)
			cs = ast.Case{}
		} else {
			p.require(types.EOL)
			p.next()
		}
	}
	stmt.Cases = cases

	if p.peek.Type == types.ELSE {
		p.next()
		stmt.Else = p.parseBlock()
	}
	return stmt
}

func (p *Parser) parseFor() ast.Statement {
	stmt := ast.For{Token: p.peek}
	p.next() // skip 'for'
	stmt.Condition = p.parseExpression(LOWEST)
	stmt.Body = p.parseBlock()
	return stmt
}

func (p *Parser) parseReturn() ast.Statement {
	stmt := ast.Return{Token: p.peek}
	p.next() // skip 'return'
	stmt.Body = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseThrow() ast.Statement {
	stmt := ast.Throw{Token: p.peek}
	p.next() // skip 'throw'
	stmt.Val = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
