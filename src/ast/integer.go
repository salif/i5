// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Integer struct {
	line  int
	value int64
}

func (this Integer) GetType() int {
	return INTEGER
}

func (this Integer) Print() {
	console.Printf("%v", this.value)
}

func (this Integer) GetLine() int {
	return this.line
}

func (this Integer) GetValue() int64 {
	return this.value
}

func (this Integer) Init(line int, value int64) Integer {
	this.line = line
	this.value = value
	return this
}
