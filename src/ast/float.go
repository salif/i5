// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "fmt"

type Float struct {
	Value float64
}

func (f Float) StringValue() string {
	return fmt.Sprintf("%v", f.Value)
}

func (f Float) expression() {}
