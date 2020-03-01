// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseExpression(precedence int) (ast.Node, error) {
	prefix := p.prefixFunctions[p.peek.Type]
	if prefix == nil {
		return nil, p.Throw(p.peek.Line, constants.SYNTAX_UNEXPECTED, p.peek.Value)
	}
	leftExpression, err := prefix()
	if err != nil {
		return nil, err
	}

	for p.peek.Type != constants.TOKEN_EOL && precedence < p.precedence() {
		infix := p.infixFunctions[p.peek.Type]
		if infix == nil {
			return nil, p.Throw(p.peek.Line, constants.SYNTAX_UNEXPECTED, p.peek.Value)
		}
		leftExpression, err = infix(leftExpression)
		if err != nil {
			return nil, err
		}
	}

	return leftExpression, nil
}
