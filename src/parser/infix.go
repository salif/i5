// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parseInfix(left ast.Node) ast.Node {
	expr := ast.Infix{}.Init(p.peek.Line, p.peek.Value, left)
	precedence := p.precedence()
	p.next()
	expr.SetRight(p.parseExpression(precedence))
	return expr
}
