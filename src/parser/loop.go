// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseLoop() ast.Node {
	stmt := ast.Loop{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	stmt.SetBody(p.parseBlock())
	p.require(types.EOL)
	p.next()
	return stmt
}
