// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseBool() ast.Node {
	expr := ast.Bool{}.Init(p.peek.Line, p.peek.Type == types.TRUE)
	p.next()
	return expr
}
