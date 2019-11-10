// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/types"
)

// TODO
func (p *Parser) parseSwitch() (ast.Node, error) {
	node := ast.Switch{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	e, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	node.SetCondition(e)
	err = p.require(p.peek.Type, types.LBRACE)
	if err != nil {
		return nil, err
	}
	p.next()
	err = p.require(p.peek.Type, types.EOL)
	if err != nil {
		return nil, err
	}
	p.next()

	cases := []ast.Case{}
	for p.peek.Type != types.RBRACE {
		_case, err := p.parseCase()
		if err != nil {
			return nil, err
		}
		cases = append(cases, _case)
	}
	node.SetCases(cases)
	node.SetElse(ast.Block{}.Init(p.peek.Line))

	err = p.require(p.peek.Type, types.RBRACE)
	if err != nil {
		return nil, err
	}
	p.next()
	err = p.require(p.peek.Type, types.EOL)
	if err != nil {
		return nil, err
	}
	p.next()

	return node, nil
}

func (p *Parser) parseCase() (ast.Case, error) {
	node := ast.Case{}.Init(p.peek.Line)

	err := p.require(p.peek.Type, types.CASE)
	if err != nil {
		return ast.Case{}, err
	}
	p.next() // 'case'

	e, err := p.parseExpression(LOWEST)
	if err != nil {
		return ast.Case{}, err
	}
	node.Append(e)

	for p.peek.Type == types.COMMA {
		p.next() // ','

		err := p.require(p.peek.Type, types.EOL)
		if err != nil {
			return ast.Case{}, err
		}
		p.next() // 'EOL'

		err = p.require(p.peek.Type, types.CASE)
		if err != nil {
			return ast.Case{}, err
		}
		p.next() // 'case'
		e, err = p.parseExpression(LOWEST)
		if err != nil {
			return ast.Case{}, err
		}
		node.Append(e)
	}

	e, err = p.parseBlock()
	if err != nil {
		return ast.Case{}, err
	}
	block := e.(ast.Block)
	node.SetBody(block)

	err = p.require(p.peek.Type, types.EOL)
	if err != nil {
		return ast.Case{}, err
	}
	p.next()

	return node, nil
}
