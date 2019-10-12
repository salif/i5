// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/io/console"

type Bool struct {
	Value bool
}

func (b *Bool) Type() TYPE {
	return BOOL
}

func (b *Bool) StringValue() string {
	return console.Format("%t", b.Value)
}
