package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.peek.Type {
	case types.IF:
		return p.parseIf()
	case types.FN:
		return p.parseFn()
	case types.RETURN:
		return p.parseReturn()
	case types.IMPORT:
		return p.parseImport()
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

func (p *Parser) parseFn() ast.Statement {
	fn := ast.Function{Token: p.peek}
	p.next() // skip 'fn'
	p.require(types.IDENTIFIER)
	fn.Function = p.peek.Value
	p.next() // skip function name
	fn.Params = p.parseParams()
	fn.Body = p.parseBlock()

	return fn
}

func (p *Parser) parseReturn() ast.Statement {
	stmt := ast.Return{Token: p.peek}
	p.next() // skip 'return'
	stmt.Body = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}

func (p *Parser) parseImport() ast.Statement {
	stmt := ast.Import{Token: p.peek}
	p.next() // skip 'import'
	stmt.Val = p.parseExpression(LOWEST)
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
