package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type For struct {
	Token     types.Token
	Condition Expression
	Body      Block
	Return    Expression
}

func (f For) Value() string {
	return f.Token.Value
}

func (f For) String() string {
	var out bytes.Buffer
	out.WriteString("for ")
	out.WriteString(f.Condition.String())
	out.WriteString(" ")
	out.WriteString(f.Body.String())
	return out.String()
}
func (f For) statement() {}
