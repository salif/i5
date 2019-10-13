// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type String struct {
	line  int
	value string
}

func (this String) GetType() int {
	return STRING
}

func (this String) Print() {
	console.Print("\"" + this.value + "\"")
}

func (this String) GetLine() int {
	return this.line
}

func (this String) GetValue() string {
	return this.value
}

func (this String) Init(line int, value string) String {
	this.line = line
	this.value = value
	return this
}
