// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseExprStatement() ast.Node {
	stmt := ast.Expression{}.Init(p.peek.Line, p.parseExpression(LOWEST))
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
