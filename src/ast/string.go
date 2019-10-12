// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type String struct {
	Line  int
	Value string
}

func (this String) StringValue() string {
	return "\"" + this.Value + "\""
}

func (this String) GetLine() int {
	return this.Line
}

func (this String) expression() {}
