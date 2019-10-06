// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
	"strings"
)

type Call struct {
	Caller    Expression
	Arguments []Expression
}

func (c Call) StringValue() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range c.Arguments {
		args = append(args, a.StringValue())
	}
	out.WriteString(c.Caller.StringValue())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

func (c Call) expression() {}
