// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parseIndex(left ast.Node) (ast.Node, error) {
	node := ast.Index{}.Init(p.peek.Line, left, p.peek.Value)
	p.next()
	e, err := p.parseExpression(DOT)
	if err != nil {
		return nil, err
	}
	node.SetRight(e)
	return node, nil
}
