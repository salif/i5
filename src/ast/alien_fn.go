// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import (
	"bytes"
)

type AlienFn struct {
	Alien    Expression
	Function Expression
}

func (a AlienFn) StringValue() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(a.Alien.StringValue())
	out.WriteString(".")
	out.WriteString(a.Function.StringValue())
	out.WriteString(")")

	return out.String()
}

func (a AlienFn) expression() {}
