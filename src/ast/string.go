package ast

import "github.com/i5/i5/src/types"

type String struct {
	Token types.Token
	Val   string
}

func (s String) Value() string {
	return s.Token.Value
}

func (s String) String() string {
	return s.Token.Value
}

func (s String) expression() {}
