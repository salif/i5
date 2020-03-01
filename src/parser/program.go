// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseProgram() (ast.Node, error) {
	node := ast.Program{}.Init(p.peek.Line)

	body := make([]ast.Function, 0)
	for p.peek.Type != constants.TOKEN_EOF {
		if p.peek.Type == constants.TOKEN_EOL {
			p.next()
			continue
		}
		e, err := p.parseStatement()
		if err != nil {
			return nil, err
		} else if e, ok := e.(ast.Function); ok {
			body = append(body, e)
		} else {
			return nil, p.Throw(e.GetLine(), constants.SYNTAX_EXPECTED, "function")
		}
	}
	node.SetBody(body)

	return node, nil
}
