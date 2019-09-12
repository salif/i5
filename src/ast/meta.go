package ast

import "github.com/i5/i5/src/types"

type Meta struct {
	Token types.Token
	Val   string
}

func (m Meta) Value() string {
	return m.Token.Value
}

func (m Meta) String() string {
	return m.Val
}

func (m Meta) expression() {}
