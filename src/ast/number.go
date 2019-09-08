package ast

import "github.com/i5/i5/src/types"

type Number struct {
	Token types.Token
	Val   int64
}

func (n Number) Value() string {
	return n.Token.Value
}

func (n Number) String() string {
	return n.Token.Value
}

func (n Number) expression() {}
