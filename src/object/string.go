// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/constants"

type String struct {
	Value string
}

func (this String) Type() string {
	return constants.TYPE_STRING
}

func (this String) StringValue() string {
	return this.Value
}

func (this String) GenKey() Key {
	return Key{Type: this.Type(), Value: this}
}
