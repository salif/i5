// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "fmt"

type Integer struct {
	line  uint32
	value int64
}

func (this Integer) GetType() string {
	return INTEGER
}

func (this Integer) Debug() string {
	return fmt.Sprint(this.value)
}

func (this Integer) GetLine() uint32 {
	return this.line
}

func (this Integer) GetValue() int64 {
	return this.value
}

func (this Integer) Init(line uint32, value int64) Integer {
	this.line = line
	this.value = value
	return this
}
