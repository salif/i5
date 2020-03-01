// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseImport() (ast.Node, error) {
	node := ast.Import{}.Init(p.peek.Line, p.peek.Type)
	p.next() // 'import'
	e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
	if err != nil {
		return nil, err
	}
	node.SetBody(e)

	err = p.require(p.peek.Type, constants.TOKEN_AS)
	if err != nil {
		return nil, err
	}
	p.next() // 'as'

	e, err = p.parseIdentifier()
	if err != nil {
		return nil, err
	}
	node.SetAs(e.(ast.Identifier))

	err = p.require(p.peek.Type, constants.TOKEN_EOL)
	if err != nil {
		return nil, err
	}
	p.next() // 'EOL'
	return node, nil
}
