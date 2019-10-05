// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Error struct {
	Message string
}

func (e *Error) Type() TYPE {
	return ERROR
}

func (e *Error) StringValue() string {
	return "error: " + e.Message
}

func (e *Error) Clone() Object {
	return &Error{Message: e.Message}
}
