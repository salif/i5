// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseContinue() ast.Node {
	stmt := ast.Continue{}.Init(p.peek.Line, p.peek.Type)
	p.next() // skip 'continue'
	p.require(types.EOL)
	p.next() // skip EOL
	return stmt
}
