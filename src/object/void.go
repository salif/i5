// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/io/console"

type Void struct {
	Value string
}

func (this Void) Type() TYPE {
	return VOID
}

func (this Void) StringValue() string {
	if this.Value == "" {
		return VOID
	} else {
		return console.Format("'%v' is null", this.Value)
	}
}
