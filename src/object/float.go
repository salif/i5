// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type Float struct {
	Value float64
}

func (this Float) Type() TYPE {
	return FLOAT
}

func (this Float) StringValue() string {
	return fmt.Sprint(this.Value)
}

func (this Float) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
