// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type Return struct {
	Value string
	Body  Expression
}

func (rs Return) StringValue() string {
	var out bytes.Buffer
	out.WriteString(rs.Value)
	out.WriteString(" ")
	if rs.Body != nil {
		out.WriteString(rs.Body.StringValue())
	}
	return out.String()
}

func (rs Return) statement() {}
