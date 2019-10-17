// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseWhile() ast.Node {
	stmt := ast.While{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'while'
	stmt.SetCondition(p.parseExpression(LOWEST))
	stmt.SetBody(p.parseBlock())
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
