// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "github.com/i5/i5/src/io/console"

type Float struct {
	Value float64
}

func (f *Float) Type() TYPE {
	return FLOAT
}

func (f *Float) StringValue() string {
	return console.Format("%v", f.Value)
}
