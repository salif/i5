// SPDX-License-Identifier: GPL-3.0-or-later
package ast

type Break struct {
	line  uint32
	token string
}

func (this Break) GetType() string {
	return BREAK
}

func (this Break) Debug() string {
	return this.token
}

func (this Break) GetLine() uint32 {
	return this.line
}

func (this Break) Init(line uint32, token string) Break {
	this.line = line
	this.token = token
	return this
}
