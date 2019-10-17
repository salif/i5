// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseThrow() ast.Node {
	stmt := ast.Throw{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'throw'
	stmt.SetBody(p.parseExpression(LOWEST))
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
