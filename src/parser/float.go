// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"strconv"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseFloat() (ast.Node, error) {
	err := p.require(p.peek.Type, constants.TOKEN_FLOAT)
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
