// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseIdentifier() (ast.Node, error) {
	err := p.require(p.peek.Type, types.IDENT)
	if err != nil {
		return nil, err
	}
	node := ast.Identifier{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	if p.peek.Type == types.IDENT {
		result := ast.Identifiers{}.Init(node.GetLine())
		result.Append(node)
		for p.peek.Type == types.IDENT {
			result.Append(ast.Identifier{}.Init(p.peek.Line, p.peek.Value))
			p.next()
		}
		return result, nil
	} else {
		return node, nil
	}
}
