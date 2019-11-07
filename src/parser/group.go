// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseGroup() (ast.Node, error) {
	err := p.require(p.peek.Type, types.LPAREN)
	if err != nil {
		return nil, err
	}
	p.next() // '('
	var node ast.Node
	if p.peek.Type == types.RPAREN {
		p.next()
		return ast.Identifiers{}.Init(p.peek.Line), nil
	}
	node, err = p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	err = p.require(p.peek.Type, types.RPAREN)
	if err != nil {
		return nil, err
	}
	p.next() // ')'
	return node, nil
}
