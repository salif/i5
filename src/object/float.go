// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/io/console"

type Float struct {
	Value float64
}

func (this Float) Type() TYPE {
	return FLOAT
}

func (this Float) StringValue() string {
	return console.Format("%v", this.Value)
}

func (this Float) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
