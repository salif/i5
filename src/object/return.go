// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Return struct {
	Value Object
}

func (this Return) Type() TYPE {
	return RETURN
}

func (this Return) StringValue() string {
	return this.Value.StringValue()
}
