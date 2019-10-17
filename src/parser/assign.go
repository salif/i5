// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseAssign(left ast.Node) ast.Node {
	expr := ast.Assign{}.Init(p.peek.Line, types.EQ, left)
	switch p.peek.Type {
	case types.EQ:
		p.next()
		expr.SetRight(p.parseExpression(LOWEST))
	case types.COLONEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.COLON, p.parseExpression(LOWEST)))
	case types.PLUSEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.PLUS, p.parseExpression(LOWEST)))
	case types.MINUSEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MINUS, p.parseExpression(LOWEST)))
	case types.MULTIPLYEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MULTIPLY, p.parseExpression(LOWEST)))
	case types.DIVIDEEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.DIVIDE, p.parseExpression(LOWEST)))
	case types.MODULOEQ:
		p.next()
		expr.SetRight(ast.Infix{}.Set(p.peek.Line, left, types.MODULO, p.parseExpression(LOWEST)))
	default:
		console.ThrowParsingError(1, constants.PARSER_UNEXPECTED, p.peek.Line, p.peek.Type)
	}
	return expr
}
