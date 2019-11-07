// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseAssign(left ast.Node) (ast.Node, error) {
	node := ast.Assign{}.Init(p.peek.Line, types.EQ, left)
	switch p.peek.Type {
	case types.EQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(e)
	case types.COLONEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.COLON, e))
	case types.PLUSEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.PLUS, e))
	case types.MINUSEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MINUS, e))
	case types.MULTIPLYEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MULTIPLY, e))
	case types.DIVIDEEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.DIVIDE, e))
	case types.MODULOEQ:
		p.next()
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MODULO, e))
	default:
		return nil, p.Throw(p.peek.Line, constants.PARSER_UNEXPECTED, p.peek.Type)
	}
	return node, nil
}
