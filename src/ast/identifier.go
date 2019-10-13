// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Identifier struct {
	line  int
	token string
}

func (this Identifier) GetType() int {
	return IDENTIFIER
}

func (this Identifier) Print() {
	console.Print(this.token)
}

func (this Identifier) GetLine() int {
	return this.line
}

func (this Identifier) GetValue() string {
	return this.token
}

func (this Identifier) Init(line int, token string) Identifier {
	this.line = line
	this.token = token
	return this
}
