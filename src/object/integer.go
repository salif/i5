// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type Integer struct {
	Value int64
}

func (this Integer) Type() TYPE {
	return INTEGER
}

func (this Integer) StringValue() string {
	return fmt.Sprint(this.Value)
}

func (this Integer) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
