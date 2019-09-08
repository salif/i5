package ast

import (
	"bytes"
	"fmt"

	"github.com/i5/i5/src/types"
)

type Import struct {
	Token types.Token
	Val   Expression
}

func (i Import) Value() string {
	return i.Token.Value
}

func (i Import) String() string {
	var out bytes.Buffer
	out.WriteString(i.Value())
	out.WriteString("(")
	out.WriteString(fmt.Sprintf("\"%s\"", i.Val))
	out.WriteString(")")
	return out.String()
}

func (i Import) statement() {}
