// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

var Null = object.Exception{Name: object.String{Value: constants.EXCEPTION_NULL}, Message: object.String{Value: constants.EXCEPTION_NULL}}

func Errf(err error) error {
	if err == nil {
		return err
	}
	return fmt.Errorf("%v%v\n", "error: ", err.Error())
}

func newException(name string, text string, format ...interface{}) object.Exception {
	return object.Exception{
		Name:    object.String{Value: name},
		Message: object.String{Value: fmt.Sprintf(text, format...)}}
}
