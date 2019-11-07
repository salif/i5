// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

func (p *Parser) parseLoop() (ast.Node, error) {
	node := ast.Loop{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	e, err := p.parseBlock()
	if err != nil {
		return nil, err
	}
	if e, ok := e.(ast.Block); ok {
		node.SetBody(e)
	} else {
		return nil, p.Throw(e.GetLine(), constants.PARSER_EXPECTED, "block statement")
	}
	err = p.require(p.peek.Type, types.EOL)
	if err != nil {
		return nil, err
	}
	p.next()
	return node, nil
}
