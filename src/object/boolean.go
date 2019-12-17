// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Boolean struct {
	Value bool
}

func (this Boolean) Type() string {
	return constants.TYPE_BOOLEAN
}

func (this Boolean) StringValue() string {
	return fmt.Sprint(this.Value)
}
