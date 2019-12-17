// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseThrow() (ast.Node, error) {
	node := ast.Throw{}.Init(p.peek.Line, p.peek.Type)
	p.next() // 'throw'
	e, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	node.SetBody(e)
	p.require(p.peek.Type, constants.TOKEN_EOL)
	p.next() // 'EOL'
	return node, nil
}
