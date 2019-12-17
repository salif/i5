// SPDX-License-Identifier: GPL-3.0-or-later
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
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(e)
	case constants.TOKEN_COLONEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_COLON, e))
	case constants.TOKEN_PLUSEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_PLUS, e))
	case constants.TOKEN_MINUSEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_MINUS, e))
	case constants.TOKEN_MULTIPLYEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_MULTIPLY, e))
	case constants.TOKEN_DIVIDEEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_DIVIDE, e))
	case constants.TOKEN_MODULOEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, constants.TOKEN_MODULO, e))
	default:
		return nil, p.Throw(p.peek.Line, constants.PARSER_UNEXPECTED, p.peek.Type)
	}
	return node, nil
}
