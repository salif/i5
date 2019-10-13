// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Bool struct {
	line  int
	value bool
}

func (this Bool) GetType() int {
	return BOOL
}

func (this Bool) Print() {
	console.Printf("%v", this.value)
}

func (this Bool) GetLine() int {
	return this.line
}

func (this Bool) GetValue() bool {
	return this.value
}

func (this Bool) Init(line int, value bool) Bool {
	this.line = line
	this.value = value
	return this
}
