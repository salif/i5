// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type Integer struct {
	Value int64
}

func (i *Integer) Type() TYPE {
	return INTEGER
}

func (i *Integer) StringValue() string {
	return fmt.Sprintf("%v", i.Value)
}
