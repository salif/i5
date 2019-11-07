// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"strconv"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseFloat() (ast.Node, error) {
	err := p.require(p.peek.Type, types.FLOAT)
	if err != nil {
		return nil, err
	}

	value, err := strconv.ParseFloat(p.peek.Value, 64)
	if err != nil {
		return nil, p.Throw(p.peek.Line, constants.PARSER_NOT_FLOAT, p.peek.Value)
	}

	node := ast.Float{}.Init(p.peek.Line, value)
	p.next()
	return node, nil
}
