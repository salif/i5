// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type Float struct {
	Value float64
}

func (f *Float) Type() TYPE {
	return FLOAT
}

func (f *Float) StringValue() string {
	return fmt.Sprintf("%v", f.Value)
}
