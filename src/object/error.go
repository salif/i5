// SPDX-License-Identifier: GPL-3.0-or-later
package object

import (
	"fmt"

	"github.com/i5/i5/src/i5/colors"
)

type Error struct {
	isFatal bool
	line    uint32
	number  Integer
	message String
}

func (this Error) Type() TYPE {
	return ERROR
}

func (this Error) Init(isFatal bool, line uint32, number Integer, message String) Error {
	return Error{isFatal: isFatal, line: line, number: number, message: message}
}

func (this Error) GetIsFatal() bool {
	return this.isFatal
}

func (this Error) GetNumber() Integer {
	return this.number
}

func (this Error) GetMessage() String {
	return this.message
}

func (this *Error) SetIsFatal(isFatal bool) {
	this.isFatal = isFatal
}

func (this Error) StringValue() string {
	return fmt.Sprintf("%s", this.message.StringValue())
}

func (this Error) NativeError(fileName string) error {
	return fmt.Errorf("%s%s\n%s%s:%v\n", colors.Red("error: "), this.StringValue(), colors.Red("in: "), fileName, this.line)
}
