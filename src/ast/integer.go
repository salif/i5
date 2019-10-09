// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "fmt"

type Integer struct {
	Value int64
}

func (i Integer) StringValue() string {
	return fmt.Sprintf("%v", i.Value)
}

func (i Integer) expression() {}
