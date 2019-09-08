package ast

import "github.com/i5/i5/src/types"

type Nil struct {
	Token types.Token
}

func (n Nil) Value() string {
	return n.Token.Value
}

func (n Nil) String() string {
	return n.Token.Value
}

func (n Nil) expression() {}
