// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseIdentifier() ast.Node {
	p.require(types.IDENT)
	expr := ast.Identifier{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return expr
}
