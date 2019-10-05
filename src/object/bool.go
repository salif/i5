// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type Bool struct {
	Value bool
}

func (b *Bool) Type() TYPE {
	return BOOL
}

func (b *Bool) StringValue() string {
	return fmt.Sprintf("%t", b.Value)
}
