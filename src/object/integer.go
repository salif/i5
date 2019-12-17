// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Integer struct {
	Value int64
}

func (this Integer) Type() string {
	return constants.TYPE_INTEGER
}

func (this Integer) StringValue() string {
	return fmt.Sprint(this.Value)
}

func (this Integer) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
