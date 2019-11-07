// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

// All statements ends with EOL
func (p *Parser) parseStatement() (ast.Node, error) {
	switch p.peek.Type {
	case types.IF:
		return p.parseIf()
	case types.SWITCH:
		return p.parseSwitch()
	case types.LOOP:
		return p.parseLoop()
	case types.RETURN:
		return p.parseReturn()
	case types.THROW:
		return p.parseThrow()
	default:
		node, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		err = p.require(p.peek.Type, types.EOL)
		if err != nil {
			return nil, err
		}
		p.next() // 'EOL'
		return node, nil
	}
}
