// SPDX-License-Identifier: GPL-3.0-or-later
package parser

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/types"
)

// TODO
func (p *Parser) parseSwitch() (ast.Node, error) {
	stmt := ast.Switch{}.Init(p.peek.Line, p.peek.Type)
	p.next()
	e, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	stmt.SetCondition(e)
	var cases []ast.Case
	cs := ast.Case{}.Init(p.peek.Line)
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
	for p.peek.Type == types.CASE {
		p.next()
		expr, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		cs.Append(expr)
		if p.peek.Type == types.LBRACE {
			block, err := p.parseBlock()
			if block, ok := block.(ast.Block); ok {
				cs.SetBody(block)
			} else {
				return nil, p.Throw(block.GetLine(), constants.PARSER_EXPECTED, "block statement")
			}
			err = p.require(p.peek.Type, types.EOL)
			if err != nil {
				return nil, err
			}
			p.next()
			cases = append(cases, cs)
			cs = ast.Case{}.Init(p.peek.Line)
		} else {
			err = p.require(p.peek.Type, types.EOL)
			if err != nil {
				return nil, err
			}
			p.next()
		}
	}
	stmt.SetCases(cases)

	if p.peek.Type == types.ELSE {
		p.next()
		block, err := p.parseBlock()
		if err != nil {
			return nil, err
		}
		if block, ok := block.(ast.Block); ok {
			stmt.SetElse(block)
		} else {
			return nil, p.Throw(block.GetLine(), constants.PARSER_EXPECTED, "block statement")
		}
	}

	err = p.require(p.peek.Type, types.EOL)
	if err != nil {
		return nil, err
	}
	p.next()
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

	return stmt, nil
}
