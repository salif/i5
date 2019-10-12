// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Identifier struct {
	Line  int
	Value string
}

func (this Identifier) StringValue() string {
	return this.Value
}

func (this Identifier) GetLine() int {
	return this.Line
}

func (this Identifier) expression() {}
