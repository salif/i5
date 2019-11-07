// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parsePostfix(left ast.Node) (ast.Node, error) {
	node := ast.Postfix{}.Init(p.peek.Line, p.peek.Value, left)
	p.next()
	return node, nil
}
