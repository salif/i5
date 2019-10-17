// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseBreak() ast.Node {
	stmt := ast.Break{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'break'
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
