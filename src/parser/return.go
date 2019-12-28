// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseReturn() (ast.Node, error) {
	node := ast.Return{}.Init(p.peek.Line, p.peek.Type)
	p.next() // 'return'
	e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
	if err != nil {
		return nil, err
	}
	node.SetBody(e)
	err = p.require(p.peek.Type, constants.TOKEN_EOL)
	if err != nil {
		return nil, err
	}
	p.next() // 'EOL'
	return node, nil
}
