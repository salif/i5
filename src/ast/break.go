// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Break struct {
	Line  int
	Value string
}

func (this Break) StringValue() string {
	return this.Value
}

func (this Break) GetLine() int {
	return this.Line
}

func (this Break) statement() {}
