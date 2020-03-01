// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseAssign(left ast.Node) (ast.Node, error) {
	node := ast.Assign{}.Init(p.peek.Line, constants.TOKEN_EQ, left)
	switch p.peek.Type {

	case constants.TOKEN_EQ:
		p.next()
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(e)
	case constants.TOKEN_COLONEQ:
		p.next()
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_COLON, e))
	case constants.TOKEN_PLUSEQ:
		p.next()
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_PLUS, e))
	case constants.TOKEN_MINUSEQ:
		p.next()
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_MINUS, e))
	case constants.TOKEN_MULTIPLYEQ:
		p.next()
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_MULTIPLY, e))
	case constants.TOKEN_DIVIDEEQ:
		p.next()
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_DIVIDE, e))
	case constants.TOKEN_MODULOEQ:
		p.next()
		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_MODULO, e))
	default:
		return nil, p.Throw(p.peek.Line, constants.SYNTAX_UNEXPECTED, p.peek.Type)
	}
	return node, nil
}
