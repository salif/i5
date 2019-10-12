// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/io/console"

type Void struct {
	Value string
}

func (v *Void) Type() TYPE {
	return VOID
}

func (v *Void) StringValue() string {
	if v.Value == "" {
		return VOID
	} else {
		return console.Format("undefined '%v'", v.Value)
	}
}
