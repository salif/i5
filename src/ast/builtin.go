// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Builtin struct {
	Line  int
	Value string
}

func (this Builtin) StringValue() string {
	return "$" + this.Value
}

func (this Builtin) GetLine() int {
	return this.Line
}

func (this Builtin) expression() {}
