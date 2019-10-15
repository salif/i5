// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"github.com/i5/i5/src/io/console"
)

type Error struct {
	Line    int
	Message string
}

func (this Error) Type() TYPE {
	return ERROR
}

func (this Error) StringValue() string {
	if this.Line > 0 {
		return console.Format("line %d: %v", this.Line, this.Message)
	} else {
		return console.Format("%v", this.Message)
	}
}

func (this Error) GetMessage() Object {
	return String{Value: this.Message}
}

func (this Error) Clone() Object {
	return Error{Message: this.Message, Line: this.Line}
}
