// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
	"strings"
)

type Call struct {
	Line      int
	Caller    Expression
	Arguments []Expression
}

func (this Call) StringValue() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range this.Arguments {
		args = append(args, a.StringValue())
	}
	out.WriteString(this.Caller.StringValue())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

func (this Call) GetLine() int {
	return this.Line
}

func (this Call) expression() {}
