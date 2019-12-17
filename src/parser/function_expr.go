// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseFunctionExpr() (ast.Node, error) {
	node := ast.FunctionExpr{}.Init(p.peek.Line, p.peek.Type)
	p.next() // 'fn'

	err := p.require(p.peek.Type, constants.TOKEN_COLON)
	if err != nil {
		return nil, err
	}
	p.next()

	err = p.require(p.peek.Type, constants.TOKEN_LPAREN)
	if err != nil {
		return nil, err
	}
	p.next()

	var params []ast.Identifier = make([]ast.Identifier, 0)
	for p.peek.Type == constants.TOKEN_IDENTIFIER {
		ident, err := p.parseIdentifier()
		if err != nil {
			return nil, err
		}
		params = append(params, ident.(ast.Identifier))
	}
	node.SetParams(params)

	err = p.require(p.peek.Type, constants.TOKEN_RPAREN)
	if err != nil {
		return nil, err
	}
	p.next()

	err = p.require(p.peek.Type, constants.TOKEN_EQGT)
	if err != nil {
		return nil, err
	}
	p.next()

	if p.peek.Type == constants.TOKEN_LBRACE {
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
