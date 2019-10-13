// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseStatement() ast.Node {
	switch p.peek.Type {
	case types.IF:
		return p.parseIf()
	case types.SWITCH:
		return p.parseSwitch()
	case types.WHILE:
		return p.parseWhile()
	case types.RETURN:
		return p.parseReturn()
	case types.THROW:
		return p.parseThrow()
	case types.TRY:
		return p.parseTry()
	case types.BREAK:
		return p.parseBreak()
	case types.CONTINUE:
		return p.parseContinue()
	default:
		return p.parseExprStatement()
	}
}

func (p *Parser) parseExprStatement() ast.Node {
	stmt := ast.Expression{}.Init(p.peek.Line, p.parseExpression(LOWEST))
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseIf() ast.Node {
	expression := ast.If{}.Init(p.peek.Line, p.peek.Type)

	p.next() // skip 'if' or 'elif'

	expression.SetCondition(p.parseExpression(LOWEST))

	expression.SetConsequence(p.parseBlock())

	if p.peek.Type == types.ELIF {
		expression.SetAlternative(ast.Block{}.Set(p.peek.Line, []ast.Node{p.parseIf()}))
	} else if p.peek.Type == types.ELSE {
		p.next() // skip 'else'
		expression.SetAlternative(p.parseBlock())
	}

	return expression
}

func (p *Parser) parseSwitch() ast.Node {
	stmt := ast.Switch{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	stmt.SetCondition(p.parseExpression(LOWEST))
	var cases []ast.Case
	cs := ast.Case{}.Init(p.peek.Line)
	p.require(types.EOL)
	p.next()

	for p.peek.Type == types.CASE {
		p.next()
		expr := p.parseExpression(LOWEST)
		cs.Append(expr)
		if p.peek.Type == types.LBRACE {
			cs.SetBody(p.parseBlock())
			p.require(types.EOL)
			p.next()
			cases = append(cases, cs)
			cs = ast.Case{}.Init(p.peek.Line)
		} else {
			p.require(types.EOL)
			p.next()
		}
	}
	stmt.SetCases(cases)

	if p.peek.Type == types.ELSE {
		p.next()
		stmt.SetElse(p.parseBlock())
	}

	p.require(types.EOL)
	p.next() // skip EOL

	return stmt
}

func (p *Parser) parseWhile() ast.Node {
	stmt := ast.While{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'while'
	stmt.SetCondition(p.parseExpression(LOWEST))
	stmt.SetBody(p.parseBlock())
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseReturn() ast.Node {
	stmt := ast.Return{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'return'
	stmt.SetBody(p.parseExpression(LOWEST))
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseThrow() ast.Node {
	stmt := ast.Throw{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'throw'
	stmt.SetBody(p.parseExpression(LOWEST))
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseTry() ast.Node {
	stmt := ast.Try{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'try'
	stmt.SetBody(p.parseBlock())
	if p.peek.Type == types.CATCH {
		p.next() // skip 'catch'
		if p.peek.Type == types.IDENT {
			stmt.SetErr(ast.Identifier{}.Init(p.peek.Line, p.peek.Value))
			p.next()
		}
		stmt.SetCatch(p.parseBlock())
	}
	if p.peek.Type == types.FINALLY {
		p.next() // skip 'finally'
		stmt.SetFinally(p.parseBlock())
	}
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseBreak() ast.Node {
	stmt := ast.Break{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'break'
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseContinue() ast.Node {
	stmt := ast.Continue{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'continue'
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
