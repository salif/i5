// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/io/console"

type Error struct {
	Line    int
	Message string
}

func (e *Error) Type() TYPE {
	return ERROR
}

func (e *Error) StringValue() string {
	return console.Format("line %d: %v", e.Line, e.Message)
}

func (e *Error) Clone() Object {
	return &Error{Message: e.Message, Line: e.Line}
}
