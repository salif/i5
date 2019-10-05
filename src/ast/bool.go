// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "fmt"

type Bool struct {
	Value bool
}

func (b Bool) String() string {
	return fmt.Sprintf("%v", b.Value)
}

func (b Bool) expression() {}
