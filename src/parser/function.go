// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseFunction(left ast.Node) (ast.Node, error) {
	node := ast.Function{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	if nident, ok := left.(ast.Identifiers); ok {
		node.SetParams(nident)
	} else {
		idents := ast.Identifiers{}.Init(p.peek.Line)
		if nident, ok := left.(ast.Identifier); ok {
			idents.Append(nident)
		} else {
			return nil, p.Throw(p.peek.Line, constants.PARSER_UNEXPECTED, left.GetType())
		}
		node.SetParams(idents)
	}
	if p.peek.Type == types.LBRACE {
		expr, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		node.SetBody(expr)
	} else {
		expr, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		node.SetBody(expr)
	}
	return node, nil
}
