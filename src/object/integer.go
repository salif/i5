// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/io/console"

type Integer struct {
	Value int64
}

func (i *Integer) Type() TYPE {
	return INTEGER
}

func (i *Integer) StringValue() string {
	return console.Format("%v", i.Value)
}
