// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Float struct {
	Line  int
	Value float64
}

func (this Float) StringValue() string {
	return console.Format("%v", this.Value)
}

func (this Float) GetLine() int {
	return this.Line
}

func (this Float) expression() {}
