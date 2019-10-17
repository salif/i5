// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseIf() ast.Node {
	expression := ast.If{}.Init(p.peek.Line, p.peek.Type)

	p.next() // skip 'if' or 'elif'

	expression.SetCondition(p.parseExpression(LOWEST))

	expression.SetConsequence(p.parseBlock())

	if p.peek.Type == types.ELIF {
		expression.SetAlternative(ast.Block{}.Set(p.peek.Line, []ast.Node{p.parseIf()}))
	} else if p.peek.Type == types.ELSE {
		p.next() // skip 'else'
		expression.SetAlternative(p.parseBlock())
	}

	return expression
}
