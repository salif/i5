// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"strconv"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseInteger() ast.Node {
	p.require(types.INT)
	value, err := strconv.ParseInt(p.peek.Value, 0, 64)

	if err != nil {
		console.ThrowParsingError(1, constants.PARSER_NOT_INT, p.peek.Line, p.peek.Value)
	}

	expr := ast.Integer{}.Init(p.peek.Line, value)
	p.next()
	return expr
}
