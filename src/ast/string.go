// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type String struct {
	line  uint32
	value string
}

func (this String) GetType() string {
	return STRING
}

func (this String) Debug() string {
	return "\"" + this.value + "\""
}

func (this String) GetLine() uint32 {
	return this.line
}

func (this String) GetValue() string {
	return this.value
}

func (this String) Init(line uint32, value string) String {
	this.line = line
	this.value = value
	return this
}
