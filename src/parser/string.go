// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseString() ast.Node {
	p.require(types.STRING)
	expr := ast.String{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return expr
}
