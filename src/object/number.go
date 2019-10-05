// SPDX-License-Identifier: GPL-3.0-or-later
package object

import "fmt"

type Number struct {
	Value int64
}

func (n *Number) Type() TYPE {
	return NUMBER
}

func (n *Number) StringValue() string {
	return fmt.Sprintf("%d", n.Value)
}
