// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseProgram() ast.Node {
	program := ast.Program{}.Init(p.peek.Line, []ast.Node{})

	for p.peek.Type != types.EOF {
		if p.peek.Type == types.EOL {
			p.next()
			continue
		}
		expr := p.parseExpression(LOWEST)
		if expr.GetType() != ast.ASSIGN {
			console.ThrowParsingError(1, constants.PARSER_EXPECTED, expr.GetLine(), "declaration")
		}
		program.Append(expr)
	}

	return program
}
