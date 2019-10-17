// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseParams() []ast.Identifier {
	var identifiers []ast.Identifier
	p.require(types.LPAREN)
	p.next() // skip '('

	if p.peek.Value == types.RPAREN {
		p.next()
		return identifiers
	}

	for p.peek.Type != types.RPAREN {
		p.require(types.IDENT)
		ident := ast.Identifier{}.Init(p.peek.Line, p.peek.Value)
		p.next()
		identifiers = append(identifiers, ident)
	}

	p.require(types.RPAREN)
	p.next() // skip ')'
	return identifiers
}
