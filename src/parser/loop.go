// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseLoop() (ast.Node, error) {
	node := ast.Loop{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	e, err := p.parseBlock()
	if err != nil {
		return nil, err
	}
	node.SetBody(e)
	err = p.require(p.peek.Type, types.EOL)
	if err != nil {
		return nil, err
	}
	p.next()
	return node, nil
}
