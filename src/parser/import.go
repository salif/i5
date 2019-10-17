// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parseImport() ast.Node {
	expr := ast.Import{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	expr.SetBody(p.parseExpression(LOWEST))
	return expr
}
