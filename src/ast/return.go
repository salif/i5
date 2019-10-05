// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Return struct {
	Value string
	Body  Expression
}

func (rs Return) String() string {
	var out bytes.Buffer
	out.WriteString(rs.Value)
	out.WriteString(" ")
	if rs.Body != nil {
		out.WriteString(rs.Body.String())
	}
	return out.String()
}

func (rs Return) statement() {}
