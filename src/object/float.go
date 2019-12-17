// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Float struct {
	Value float64
}

func (this Float) Type() string {
	return constants.TYPE_FLOAT
}

func (this Float) StringValue() string {
	return fmt.Sprint(this.Value)
}

func (this Float) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
