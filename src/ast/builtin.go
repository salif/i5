// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Builtin struct {
	line  int
	value string
}

func (this Builtin) GetType() string {
	return BUILTIN
}

func (this Builtin) Print() {
	console.Print("$" + this.value)
}

func (this Builtin) GetLine() int {
	return this.line
}

func (this Builtin) GetValue() string {
	return this.value
}

func (this Builtin) Init(line int, value string) Builtin {
	this.line = line
	this.value = value
	return this
}
