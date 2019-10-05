// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "fmt"

type Number struct {
	Value int64
}

func (n Number) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n Number) expression() {}
