// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Key struct {
	Type  TYPE
	Value Object
}

func (this Key) GetValue() Object {
	return this.Value
}
