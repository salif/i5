// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseList(end string) ([]ast.Node, error) {
	var list []ast.Node = []ast.Node{}

	if p.peek.Type == end {
		return list, nil
	}

	e, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	list = append(list, e)

	for p.peek.Type == constants.TOKEN_COMMA {
		p.next() // ','
		if p.peek.Type == constants.TOKEN_EOL {
			p.next() // 'EOL'
		}
		e, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		list = append(list, e)
	}

	return list, nil
}
