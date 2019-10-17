// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseList(end string) []ast.Node {
	var list []ast.Node

	if p.peek.Type == end {
		return list
	}

	list = append(list, p.parseExpression(LOWEST))

	for p.peek.Type == types.COMMA {
		p.next() // skip ','
		p.expect(p.peek.Type == types.EOL)
		list = append(list, p.parseExpression(LOWEST))
	}

	return list
}
