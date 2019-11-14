// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/object"
)

const FATAL int = 2

func Errf(err error) error {
	if err == nil {
		return err
	}
	return fmt.Errorf("%v%v\n", colors.Red("error: "), err.Error())
}

func Nil(line uint32) object.Error {
	return newError(false, line, constants.ERROR_NIL, "nil")
}

func newError(isFatal bool, line uint32, number int64, text string, format ...interface{}) object.Error {
	var message object.String = object.String{Value: fmt.Sprintf(text, format...)}
	return object.Error{IsFatal: isFatal, Line: line, Number: object.Integer{Value: number}, Message: message}
}

// if err is:
// not error -> return 0;
// nil -> return 1;
// fatal error -> return 2;
// other error -> return 3;
func ErrorType(err object.Object) int {
	if err, ok := err.(object.Error); ok {
		if err.IsFatal {
			return 2
		} else if err.Number.Value == constants.ERROR_NIL {
			return 1
		} else {
			return 3
		}
	} else {
		return 0
	}
}
