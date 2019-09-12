package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Case struct {
	Token types.Token
	Cases []Expression
	Body  Block
}

func (c Case) String() string {
	var out bytes.Buffer
	for _, i := range c.Cases {
		out.WriteString("case " + i.String() + ";")
	}
	out.WriteString(c.Body.String())
	return out.String()
}
