// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseSwitch() ast.Node {
	stmt := ast.Switch{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	stmt.SetCondition(p.parseExpression(LOWEST))
	var cases []ast.Case
	cs := ast.Case{}.Init(p.peek.Line)
	p.require(types.EOL)
	p.next()

	for p.peek.Type == types.CASE {
		p.next()
		expr := p.parseExpression(LOWEST)
		cs.Append(expr)
		if p.peek.Type == types.LBRACE {
			cs.SetBody(p.parseBlock())
			p.require(types.EOL)
			p.next()
			cases = append(cases, cs)
			cs = ast.Case{}.Init(p.peek.Line)
		} else {
			p.require(types.EOL)
			p.next()
		}
	}
	stmt.SetCases(cases)

	if p.peek.Type == types.ELSE {
		p.next()
		stmt.SetElse(p.parseBlock())
	}

	p.require(types.EOL)
	p.next() // skip EOL

	return stmt
}
