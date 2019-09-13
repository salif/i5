package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Throw struct {
	Token types.Token
	Val   Expression
}

func (t Throw) Value() string {
	return t.Token.Value
}

func (t Throw) String() string {
	var out bytes.Buffer
	out.WriteString(t.Value())
	out.WriteString("(")
	out.WriteString(t.Val.String())
	out.WriteString(")")
	return out.String()
}

func (t Throw) statement() {}
