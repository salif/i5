// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Integer struct {
	Line  int
	Value int64
}

func (this Integer) StringValue() string {
	return console.Format("%v", this.Value)
}

func (this Integer) GetLine() int {
	return this.Line
}

func (this Integer) expression() {}
