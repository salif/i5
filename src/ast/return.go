package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type Return struct {
	Token types.Token
	Body  Expression
}

func (rs Return) Value() string {
	return rs.Token.Value
}

func (rs Return) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Value() + " ")
	if rs.Body != nil {
		out.WriteString(rs.Body.String())
	}
	return out.String()
}

func (rs Return) statement() {}
