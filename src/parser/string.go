// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseString() (ast.Node, error) {
	err := p.require(p.peek.Type, types.STRING)
	if err != nil {
		return nil, err
	}
	node := ast.String{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return node, nil
}