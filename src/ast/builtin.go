package ast

import "github.com/i5/i5/src/types"

type Builtin struct {
	Token types.Token
	Val   string
}

func (b Builtin) Value() string {
	return b.Token.Value
}

func (b Builtin) String() string {
	return b.Val
}

func (b Builtin) expression() {}
