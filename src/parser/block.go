// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseBlock() ast.Block {
	block := ast.Block{}.Init(p.peek.Line)
	p.require(types.LBRACE)
	p.next() // skip '{'
	p.require(types.EOL)
	p.next() // skip EOL

	for p.peek.Value != types.RBRACE {
		if p.peek.Type == types.EOL {
			p.next() // skip empty line
			continue
		}
		stmt := p.parseStatement()
		block.Append(stmt)
	}

	p.require(types.RBRACE)
	p.next() // skip '}'
	return block
}
