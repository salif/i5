// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/i5/colors"
)

type Error struct {
	IsFatal bool
	Line    uint32
	Number  Integer
	Message String
}

func (this Error) Type() TYPE {
	return ERROR
}

func (this Error) StringValue() string {
	return fmt.Sprintf("%s", this.Message.StringValue())
}

func (this Error) NativeError(fileName string) error {
	return fmt.Errorf("%s%s\n%s%s:%v\n", colors.Red("error: "), this.StringValue(), colors.Red("in: "), fileName, this.Line)
}
