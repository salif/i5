// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"strconv"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseInteger() (ast.Node, error) {
	err := p.require(p.peek.Type, constants.TOKEN_INTEGER)
	if err != nil {
		return nil, err
	}
	value, err := strconv.ParseInt(p.peek.Value, 0, 64)

	if err != nil {
		return nil, p.Throw(p.peek.Line, constants.PARSER_NOT_INT, p.peek.Value)
	}

	expr := ast.Integer{}.Init(p.peek.Line, value)
	p.next()
	return expr, nil
}
