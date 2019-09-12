package ast

import (
	"bytes"
	"strings"

	"github.com/i5/i5/src/types"
)

type ExprList struct {
	Token types.Token
	Exprs []Expression
}

func (e ExprList) Value() string {
	return e.Token.Value
}

func (e ExprList) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range e.Exprs {
		args = append(args, a.String())
	}
	out.WriteString(strings.Join(args, ", "))
	return out.String()
}

func (e ExprList) expression() {}
