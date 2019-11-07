// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Builtin struct {
	line  uint32
	value string
}

func (this Builtin) GetType() string {
	return BUILTIN
}

func (this Builtin) Debug() string {
	return "$" + this.value
}

func (this Builtin) GetLine() uint32 {
	return this.line
}

func (this Builtin) GetValue() string {
	return this.value
}

func (this Builtin) Init(line uint32, value string) Builtin {
	this.line = line
	this.value = value
	return this
}
