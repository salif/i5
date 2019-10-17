// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseFn() ast.Node {
	p.require(types.FN)
	fn := ast.Function{}.Init(p.peek.Line, p.peek.Type)
	var expr ast.Assign
	p.next() // skip 'fn'
	if p.peek.Type != types.IDENT {
		fn.SetAnonymous(true)
	} else {
		fn.SetAnonymous(false)
		expr = ast.Assign{}.Init(p.peek.Line, types.EQ, p.parseIdentifier())
	}

	fn.SetParams(p.parseParams())
	fn.SetBody(p.parseBlock())
	if fn.GetAnonymous() {
		return fn
	}
	expr.SetRight(fn)
	return expr
}
