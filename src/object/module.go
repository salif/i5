// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Module struct {
	Env *Env
}

func (this Module) Type() string {
	return constants.TYPE_MODULE
}

func (this Module) StringValue() string {
	return fmt.Sprintf("[type: %v]", this.Type())
}
