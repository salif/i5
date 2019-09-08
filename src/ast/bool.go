package ast

import "github.com/i5/i5/src/types"

type Bool struct {
	Token types.Token
	Val   bool
}

func (b Bool) Value() string {
	return b.Token.Value
}

func (b Bool) String() string {
	return b.Token.Value
}

func (b Bool) expression() {}
