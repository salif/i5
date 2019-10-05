// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type While struct {
	Value     string
	Condition Expression
	Body      *Block
}

func (w While) String() string {
	var out bytes.Buffer
	out.WriteString(w.Value)
	out.WriteString(" ")
	out.WriteString(w.Condition.String())
	out.WriteString(" ")
	out.WriteString(w.Body.String())
	return out.String()
}
func (w While) statement() {}
