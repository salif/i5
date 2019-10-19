// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseReturn() ast.Node {
	stmt := ast.Return{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'return'
	if p.peek.Type == ast.RETURN {
		stmt.SetBody(ast.Return{}.Init(p.peek.Line, p.peek.Type))
		p.next()
	} else {
		stmt.SetBody(p.parseExpression(LOWEST))
	}
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
