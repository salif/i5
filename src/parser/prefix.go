// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parsePrefix() (ast.Node, error) {
	node := ast.Prefix{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	e, err := p.parseExpression(PREFIX)
	if err != nil {
		return nil, err
	}
	node.SetRight(e)
	return node, nil
}
