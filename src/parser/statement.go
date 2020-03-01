// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

// All statements ends with EOL
func (p *Parser) parseStatement() (ast.Node, error) {
	switch p.peek.Type {
	case constants.TOKEN_FN:
		return p.parseFunction()
	case constants.TOKEN_RETURN:
		return p.parseReturn()
	case constants.TOKEN_IF:
		return p.parseIf()
	case constants.TOKEN_SWITCH:
		return p.parseSwitch()
	case constants.TOKEN_LOOP:
		return p.parseLoop()
	case constants.TOKEN_BREAK:
		return p.parseBreak()
	case constants.TOKEN_THROW:
		return p.parseThrow()
	case constants.TOKEN_IMPORT:
		return p.parseImport()
	default:
		node, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		err = p.require(p.peek.Type, constants.TOKEN_EOL)
		if err != nil {
			return nil, err
		}
		p.next() // 'EOL'
		return node, nil
	}
}
