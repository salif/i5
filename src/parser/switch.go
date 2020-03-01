// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
)

func (p *Parser) parseSwitch() (ast.Node, error) {
	node := ast.Switch{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
	if err != nil {
		return nil, err
	}
	node.SetCondition(e)
	err = p.require(p.peek.Type, constants.TOKEN_LBRACE)
	if err != nil {
		return nil, err
	}
	p.next()
	err = p.require(p.peek.Type, constants.TOKEN_EOL)
	if err != nil {
		return nil, err
	}
	p.next()

	cases := []ast.Case{}
	for p.peek.Type != constants.TOKEN_RBRACE {
		casesToAppend := []ast.Case{}

		_case := ast.Case{}.Init(p.peek.Line)
		err := p.require(p.peek.Type, constants.TOKEN_CASE)
		if err != nil {
			return nil, err
		}
		p.next() // 'case'

		e, err := p.parseExpression(constants.PRECEDENCE_LOWEST)
		if err != nil {
			return nil, err
		}
		_case.SetCase(e)

		casesToAppend = append(casesToAppend, _case)

		for p.peek.Type == constants.TOKEN_COMMA {
			p.next() // ','

			err := p.require(p.peek.Type, constants.TOKEN_EOL)
			if err != nil {
				return nil, err
			}
			p.next() // 'EOL'

			_case = ast.Case{}.Init(p.peek.Line)
			err = p.require(p.peek.Type, constants.TOKEN_CASE)
			if err != nil {
				return nil, err
			}
			p.next() // 'case'
			e, err = p.parseExpression(constants.PRECEDENCE_LOWEST)
			if err != nil {
				return nil, err
			}

			_case.SetCase(e)
			casesToAppend = append(casesToAppend, _case)
		}

		err = p.require(p.peek.Type, constants.TOKEN_EQGT)
		if err != nil {
			return nil, err
		}
		p.next()

		e, err = p.parseBlock()
		if err != nil {
			return nil, err
		}
		block := e.(ast.Block)

		for _, c := range casesToAppend {
			c.SetBody(block)
			cases = append(cases, c)
		}

		err = p.require(p.peek.Type, constants.TOKEN_EOL)
		if err != nil {
			return nil, err
		}
		p.next()

	}
	node.SetCases(cases)

	err = p.require(p.peek.Type, constants.TOKEN_RBRACE)
	if err != nil {
		return nil, err
	}
	p.next()
	err = p.require(p.peek.Type, constants.TOKEN_EOL)
	if err != nil {
		return nil, err
	}
	p.next()

	return node, nil
}
