// SPDX-License-Identifier: GPL-3.0-or-later
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
