// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseList(end string) ([]ast.Node, error) {
	var list []ast.Node = []ast.Node{}

	if p.peek.Type == end {
		return list, nil
	}

	e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
	if err != nil {
		return nil, err
	}
	list = append(list, e)

	for p.peek.Type == constants.TOKEN_COMMA {
		p.next() // ','
		if p.peek.Type == constants.TOKEN_EOL {
			p.next() // 'EOL'
		}
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}

	return list, nil
}
