// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Float struct {
	line  int
	value float64
}

func (this Float) GetType() int {
	return FLOAT
}

func (this Float) Print() {
	console.Printf("%v", this.value)
}

func (this Float) GetLine() int {
	return this.line
}

func (this Float) GetValue() float64 {
	return this.value
}

func (this Float) Init(line int, value float64) Float {
	this.line = line
	this.value = value
	return this
}
