package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Suffix struct {
	Token    types.Token
	Left     Expression
	Operator string
}

func (s Suffix) Value() string {
	return s.Token.Value
}

func (s Suffix) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(s.Left.String())
	out.WriteString(s.Operator)
	out.WriteString(")")
	return out.String()
}
func (s Suffix) expression() {}
