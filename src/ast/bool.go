// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Bool struct {
	Line  int
	Value bool
}

func (this Bool) StringValue() string {
	return console.Format("%v", this.Value)
}

func (this Bool) GetLine() int {
	return this.Line
}

func (this Bool) expression() {}
