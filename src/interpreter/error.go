// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/i5/colors"
	"github.com/i5/i5/src/object"
)

var Null = object.Exception{Name: object.String{Value: constants.EXCEPTION_NULL}, Message: object.String{Value: constants.EXCEPTION_NULL}}

func Errf(err error) error {
	if err == nil {
		return err
	}
	return fmt.Errorf("%v%v\n", colors.Red("error: "), err.Error())
}

func newException(name string, text string, format ...interface{}) object.Exception {
	return object.Exception{
		Name:    object.String{Value: name},
		Message: object.String{Value: fmt.Sprintf(text, format...)}}
}
