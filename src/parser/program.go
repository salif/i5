// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseProgram() (ast.Node, error) {
	node := ast.Program{}.Init(p.peek.Line, []ast.Assign{})

	body := []ast.Assign{}
	for p.peek.Type != types.EOF {
		if p.peek.Type == types.EOL {
			p.next()
			continue
		}
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		} else if e, ok := e.(ast.Assign); ok {
			body = append(body, e)
		} else {
			return nil, p.Throw(e.GetLine(), constants.PARSER_EXPECTED, "declaration")
		}
	}
	node.SetBody(body)

	return node, nil
}
