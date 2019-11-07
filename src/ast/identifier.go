// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Identifier struct {
	line  uint32
	token string
}

func (this Identifier) GetType() string {
	return IDENTIFIER
}

func (this Identifier) Debug() string {
	return this.token
}

func (this Identifier) GetLine() uint32 {
	return this.line
}

func (this Identifier) GetValue() string {
	return this.token
}

func (this Identifier) Init(line uint32, token string) Identifier {
	this.line = line
	this.token = token
	return this
}
