// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Throw struct {
	Value string
	Body  Expression
}

func (t Throw) String() string {
	var out bytes.Buffer
	out.WriteString(t.Value)
	out.WriteString("(")
	out.WriteString(t.Body.String())
	out.WriteString(")")
	return out.String()
}

func (t Throw) statement() {}
