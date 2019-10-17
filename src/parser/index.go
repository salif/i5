// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parseIndex(left ast.Node) ast.Node {
	expr := ast.Index{}.Init(p.peek.Line, left, p.peek.Value)
	p.next()
	expr.SetRight(p.parseExpression(DOT))
	return expr
}
