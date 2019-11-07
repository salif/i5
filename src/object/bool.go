// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type Bool struct {
	Value bool
}

func (this Bool) Type() TYPE {
	return BOOL
}

func (this Bool) StringValue() string {
	return fmt.Sprintf("%t", this.Value)
}
