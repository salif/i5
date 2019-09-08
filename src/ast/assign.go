package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Assign struct {
	Token types.Token
	Left  Expression
	Val   Expression
}

func (a Assign) Value() string {
	return a.Token.Value
}

func (a Assign) String() string {
	var out bytes.Buffer
	out.WriteString(a.Left.String())
	out.WriteString(a.Value())
	out.WriteString(a.Val.String())
	return out.String()
}

func (a Assign) expression() {}
