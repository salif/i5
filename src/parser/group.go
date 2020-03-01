// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseGroup() (ast.Node, error) {
	err := p.require(p.peek.Type, constants.TOKEN_LPAREN)
	if err != nil {
		return nil, err
	}
	p.next() // '('
	node, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
	if err != nil {
		return nil, err
	}
	err = p.require(p.peek.Type, constants.TOKEN_RPAREN)
	if err != nil {
		return nil, err
	}
	p.next() // ')'
	return node, nil
}
