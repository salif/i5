// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseCall(fn ast.Node) ast.Node {
	p.require(types.LPAREN)
	p.next() // skip '('
	p.expect(p.peek.Type == types.EOL)
	expr := ast.Call{}.Init(p.peek.Line, fn)
	expr.SetArguments(p.parseList(types.RPAREN))
	p.expect(p.peek.Type == types.EOL)
	p.require(types.RPAREN)
	p.next() // skip ')'
	return expr
}
