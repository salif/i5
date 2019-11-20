// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseLambda() (ast.Node, error) {
	node := ast.Lambda{}.Init(p.peek.Line, p.peek.Type)
	p.next() // 'lambda'

	err := p.require(p.peek.Type, types.COLON)
	if err != nil {
		return nil, err
	}
	p.next()

	err = p.require(p.peek.Type, types.LPAREN)
	if err != nil {
		return nil, err
	}
	p.next()

	var params []ast.Identifier = make([]ast.Identifier, 0)
	for p.peek.Type == types.IDENT {
		ident, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		params = append(params, ident.(ast.Identifier))
	}
	node.SetParams(params)

	err = p.require(p.peek.Type, types.RPAREN)
	if err != nil {
		return nil, err
	}
	p.next()

	err = p.require(p.peek.Type, types.EQGT)
	if err != nil {
		return nil, err
	}
	p.next()

	if p.peek.Type == types.LBRACE {
		body, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		node.SetBody(body)
	} else {
		body, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetBody(body)
	}

	return node, nil
}
