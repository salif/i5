// SPDX-License-Identifier: GPL-3.0-or-later
package object

type String struct {
	Value string
}

func (this String) Type() TYPE {
	return STRING
}

func (this String) StringValue() string {
	return this.Value
}
