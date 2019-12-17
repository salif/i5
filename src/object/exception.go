// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/constants"
)

type Exception struct {
	Name    String
	Message String
}

func (this Exception) Type() string {
	return constants.TYPE_EXCEPTION
}

func (this Exception) StringValue() string {
	return fmt.Sprintf("%v exception: %v", this.Name.StringValue(), this.Message.StringValue())
}
