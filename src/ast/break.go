// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Break struct {
	line  int
	token string
}

func (this Break) GetType() string {
	return BREAK
}

func (this Break) Print() {
	console.Print(this.token)
}

func (this Break) GetLine() int {
	return this.line
}

func (this Break) Init(line int, token string) Break {
	this.line = line
	this.token = token
	return this
}
