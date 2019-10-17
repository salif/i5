// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

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
