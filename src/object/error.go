// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Error struct {
	Message string
}

func (e *Error) Type() TYPE {
	return ERROR
}

func (e *Error) StringValue() string {
	return e.Message
}

func (e *Error) Clone() Object {
	return &Error{Message: e.Message}
}

type Throw struct {
	Value Object
}

func (t *Throw) Type() TYPE {
	return THROW
}

func (t *Throw) StringValue() string {
	return t.Value.StringValue()
}
