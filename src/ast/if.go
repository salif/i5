// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type If struct {
	Token       types.Token
	Condition   Expression
	Consequence *Block
	Alternative *Block
}

func (i If) Value() string {
	return i.Token.Value
}

func (i If) String() string {
	var out bytes.Buffer
	out.WriteString("if ")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())
	if i.Alternative.Body != nil {
		out.WriteString("else ")
		out.WriteString(i.Alternative.String())
	}
	return out.String()
}
func (i If) statement() {}
