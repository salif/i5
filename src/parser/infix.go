// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import "github.com/i5/i5/src/ast"

func (p *Parser) parseInfix(left ast.Node) (ast.Node, error) {
	expr := ast.Infix{}.Init(p.peek.Line, p.peek.Value, left)
	precedence := p.precedence()
	p.next()
	e, err := p.parseExpression(precedence)
	if err != nil {
		return nil, err
	}
	expr.SetRight(e)
	return expr, nil
}
