// SPDX-License-Identifier: GPL-3.0-or-later
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
	case types.WHILE:
		return p.parseWhile()
	case types.RETURN:
		return p.parseReturn()
	case types.IMPORT:
		return p.parseImportStatement()
	case types.THROW:
		return p.parseThrow()
	case types.TRY:
		return p.parseTry()
	default:
		return p.parseExprStatement()
	}
}

func (p *Parser) parseExprStatement() *ast.Expr {
	stmt := &ast.Expr{}
	stmt.Body = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseIf() ast.Statement {
	expression := &ast.If{Value: p.peek.Type}

	p.next() // skip 'if' or 'elif'

	expression.Condition = p.parseExpression(LOWEST)

	expression.Consequence = p.parseBlock()

	if p.peek.Type == types.ELIF {
		expression.Alternative = &ast.Block{Body: []ast.Statement{p.parseIf()}}
	} else if p.peek.Type == types.ELSE {
		p.next() // skip 'else'
		expression.Alternative = p.parseBlock()
	}

	return expression
}

func (p *Parser) parseSwitch() ast.Statement {
	stmt := &ast.Switch{Value: p.peek.Type}
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
			p.require(types.EOL)
			p.next()
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

	p.require(types.EOL)
	p.next() // skip EOL

	return stmt
}

func (p *Parser) parseWhile() ast.Statement {
	stmt := &ast.While{Value: p.peek.Type}
	p.next() // skip 'while'
	stmt.Condition = p.parseExpression(LOWEST)
	stmt.Body = p.parseBlock()
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseReturn() ast.Statement {
	stmt := &ast.Return{Value: p.peek.Type}
	p.next() // skip 'return'
	stmt.Body = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseImportStatement() ast.Statement {
	stmt := &ast.ImportStatement{Value: p.peek.Type}
	p.next() // skip 'import'
	stmt.Body = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseThrow() ast.Statement {
	stmt := &ast.Throw{Value: p.peek.Type}
	p.next() // skip 'throw'
	stmt.Body = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseTry() ast.Statement {
	stmt := &ast.Try{Value: p.peek.Type}
	p.next() // skip 'try'
	stmt.Body = p.parseBlock()
	if p.peek.Type == types.CATCH {
		p.next() // skip 'catch'
		if p.peek.Type == types.IDENT {
			stmt.Err = &ast.Identifier{Value: p.peek.Value}
			p.next()
		}
		stmt.Catch = p.parseBlock()
	}
	if p.peek.Type == types.FINALLY {
		p.next() // skip 'finally'
		stmt.Finally = p.parseBlock()
	}
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
