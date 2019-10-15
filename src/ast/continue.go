// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Continue struct {
	line  int
	token string
}

func (this Continue) GetType() string {
	return CONTINUE
}

func (this Continue) Print() {
	console.Print(this.token)
}

func (this Continue) GetLine() int {
	return this.line
}

func (this Continue) Init(line int, token string) Continue {
	this.line = line
	this.token = token
	return this
}
