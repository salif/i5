// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type AlienFn struct {
	Line     int
	Alien    Expression
	Function Expression
}

func (this AlienFn) StringValue() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(this.Alien.StringValue())
	out.WriteString(".")
	out.WriteString(this.Function.StringValue())
	out.WriteString(")")

	return out.String()
}

func (this AlienFn) GetLine() int {
	return this.Line
}

func (this AlienFn) expression() {}
