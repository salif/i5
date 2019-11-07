// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseBlock() (ast.Node, error) {
	node := ast.Block{}.Init(p.peek.Line)
	p.next() // '{'

	if p.peek.Type == types.RBRACE {
		p.next()
		return node, nil
	}

	err := p.require(p.peek.Type, types.EOL)
	if err != nil {
		return nil, err
	}

	p.next() // 'EOL'

	for p.peek.Type != types.RBRACE {
		if p.peek.Type == types.EOL {
			p.next() // skip empty line
			continue
		}
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		node.Append(stmt)
	}

	err = p.require(p.peek.Type, types.RBRACE)
	if err != nil {
		return nil, err
	}
	p.next() // '}'
	return node, nil
}
