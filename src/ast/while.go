// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"

	"github.com/i5/i5/src/types"
)

type While struct {
	Token     types.Token
	Condition Expression
	Body      *Block
}

func (w While) Value() string {
	return w.Token.Value
}

func (w While) String() string {
	var out bytes.Buffer
	out.WriteString("while ")
	out.WriteString(w.Condition.String())
	out.WriteString(" ")
	out.WriteString(w.Body.String())
	return out.String()
}
func (w While) statement() {}
