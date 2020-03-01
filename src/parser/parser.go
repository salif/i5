// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/lexer"
)

type (
	prefixFunction func() (ast.Node, error)
	infixFunction  func(ast.Node) (ast.Node, error)
)

type Parser struct {
	lexer           lexer.Lexer
	position        int
	peek            constants.Token
	prefixFunctions map[string]prefixFunction
	infixFunctions  map[string]infixFunction
}

func (p *Parser) ParseProgram() (ast.Node, error) {
	err := p.lexer.Run()
	if err != nil {
		return nil, err
	}
	p.next()
	return p.parseProgram()
}

func (p *Parser) Parse() (ast.Node, error) {
	err := p.lexer.Run()
	if err != nil {
		return nil, err
	}
	p.next()
	return p.parseStatement()
}

func (p *Parser) Init(fileName string, code []byte) {
	var _lexer lexer.Lexer
	_lexer.Init(fileName, code)
	p.lexer = _lexer
	p.position = 0
	p.peek = constants.Token{}
	p.prefixFunctions = map[string]prefixFunction{
		constants.TOKEN_IDENTIFIER: p.parseIdentifier,
		constants.TOKEN_INTEGER:    p.parseInteger,
		constants.TOKEN_FLOAT:      p.parseFloat,
		constants.TOKEN_STRING:     p.parseString,
		constants.TOKEN_BUILTIN:    p.parseBuiltin,
		constants.TOKEN_FN:         p.parseFunctionExpr,
		constants.TOKEN_LPAREN:     p.parseGroup,
		constants.TOKEN_NOT:        p.parsePrefix,
		constants.TOKEN_BNOT:       p.parsePrefix,
		constants.TOKEN_MINUS:      p.parsePrefix,
	}
	p.infixFunctions = map[string]infixFunction{
		constants.TOKEN_EQ:         p.parseAssign,
		constants.TOKEN_COLONEQ:    p.parseAssign,
		constants.TOKEN_PLUSEQ:     p.parseAssign,
		constants.TOKEN_MINUSEQ:    p.parseAssign,
		constants.TOKEN_MULTIPLYEQ: p.parseAssign,
		constants.TOKEN_DIVIDEEQ:   p.parseAssign,
		constants.TOKEN_MODULOEQ:   p.parseAssign,
		constants.TOKEN_QMQM:       p.parseTernary,
		constants.TOKEN_OROR:       p.parseInfix,
		constants.TOKEN_ANDAND:     p.parseInfix,
		constants.TOKEN_EQEQ:       p.parseInfix,
		constants.TOKEN_NOTEQ:      p.parseInfix,
		constants.TOKEN_LT:         p.parseInfix,
		constants.TOKEN_GT:         p.parseInfix,
		constants.TOKEN_LTEQ:       p.parseInfix,
		constants.TOKEN_GTEQ:       p.parseInfix,
		constants.TOKEN_COLON:      p.parseInfix,
		constants.TOKEN_QM:         p.parseInfix,
		constants.TOKEN_OR:         p.parseInfix,
		constants.TOKEN_XOR:        p.parseInfix,
		constants.TOKEN_AND:        p.parseInfix,
		constants.TOKEN_LTLT:       p.parseInfix,
		constants.TOKEN_GTGT:       p.parseInfix,
		constants.TOKEN_PLUS:       p.parseInfix,
		constants.TOKEN_MINUS:      p.parseInfix,
		constants.TOKEN_MULTIPLY:   p.parseInfix,
		constants.TOKEN_DIVIDE:     p.parseInfix,
		constants.TOKEN_MODULO:     p.parseInfix,
		constants.TOKEN_LPAREN:     p.parseCall,
		constants.TOKEN_DOT:        p.parseIndex,
	}
}

func (p *Parser) next() {
	if p.position < len(p.lexer.Tokens) {
		p.peek = p.lexer.Tokens[p.position]
	} else {
		p.peek = constants.Token{Type: constants.TOKEN_EOF, Value: constants.TOKEN_EOF, Line: 0}
	}
	p.position++
}

func (p *Parser) Throw(line uint32, text string, format ...interface{}) error {
	return constants.SyntaxError{Message: fmt.Sprintf(text, format...), In: fmt.Sprintf("%v:%d", p.lexer.FileName, line)}
}

func (p *Parser) precedence() int {
	if p, ok := constants.PRECEDENCES[p.peek.Type]; ok {
		return p
	}
	return constants.PRECEDENCE_LOWEST
}

func (p *Parser) require(found string, expected string) error {
	if found == expected {
		return nil
	} else {
		return p.Throw(p.peek.Line, constants.SYNTAX_EXPECTED_FOUND, expected, found)
	}
}
