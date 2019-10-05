// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Try struct {
	Value string
	Body  *Block
}

func (t Try) String() string {
	var out bytes.Buffer
	out.WriteString(t.Value)
	out.WriteString(" ")
	out.WriteString(t.Body.String())
	return out.String()
}
func (t Try) statement() {}
