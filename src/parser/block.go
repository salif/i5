// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseBlock() (ast.Block, error) {
	node := ast.Block{}.Init(p.peek.Line)

	err := p.require(p.peek.Type, constants.TOKEN_LBRACE)
	if err != nil {
		return node, err
	}
	p.next() // '{'

	if p.peek.Type == constants.TOKEN_RBRACE {
		p.next()
		return node, nil
	}

	err = p.require(p.peek.Type, constants.TOKEN_EOL)
	if err != nil {
		return node, err
	}

	p.next() // 'EOL'

	stmts := []ast.Node{}
	for p.peek.Type != constants.TOKEN_RBRACE {
		if p.peek.Type == constants.TOKEN_EOL {
			p.next() // skip empty line
			continue
		}
		stmt, err := p.parseStatement()
		if err != nil {
			return node, err
		}
		stmts = append(stmts, stmt)
	}
	node.SetBody(stmts)

	err = p.require(p.peek.Type, constants.TOKEN_RBRACE)
	if err != nil {
		return node, err
	}
	p.next() // '}'
	return node, nil
}
