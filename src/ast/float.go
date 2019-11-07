// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "fmt"

type Float struct {
	line  uint32
	value float64
}

func (this Float) GetType() string {
	return FLOAT
}

func (this Float) Debug() string {
	return fmt.Sprint(this.value)
}

func (this Float) GetLine() uint32 {
	return this.line
}

func (this Float) GetValue() float64 {
	return this.value
}

func (this Float) Init(line uint32, value float64) Float {
	this.line = line
	this.value = value
	return this
}
