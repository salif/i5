// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseStatement() ast.Node {
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
	case types.TRY:
		return p.parseTry()
	default:
		return p.parseExprStatement()
	}
}
