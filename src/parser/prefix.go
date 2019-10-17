// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parsePrefix() ast.Node {
	expr := ast.Prefix{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	expr.SetRight(p.parseExpression(PREFIX))
	return expr
}
