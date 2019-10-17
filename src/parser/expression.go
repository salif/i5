// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseExpression(precedence int) ast.Node {
	prefix := p.prefixFunctions[p.peek.Type]
	if prefix == nil {
		console.ThrowParsingError(1, constants.PARSER_UNEXPECTED, p.peek.Line, p.peek.Value)
	}
	leftExpression := prefix()

	for p.peek.Type != types.EOL && precedence < p.precedence() {
		infix := p.infixFunctions[p.peek.Type]
		if infix == nil {
			console.ThrowParsingError(1, constants.PARSER_UNEXPECTED, p.peek.Line, p.peek.Value)
		}
		leftExpression = infix(leftExpression)
	}

	return leftExpression
}
