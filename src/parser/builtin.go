// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseBuiltin() ast.Node {
	p.require(types.BUILTIN)
	expr := ast.Builtin{}.Init(p.peek.Line, p.peek.Value)
	p.next()
	return expr
}
