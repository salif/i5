// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalThrow(node ast.Throw, env *object.Env) error {
	evRight, err := Eval(node.GetBody(), env)
	if err != nil {
		return err
	}
	if evRight.Type() == constants.TYPE_STRING {
		exc := evRight.(object.String)
		return constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: exc.StringValue()}
	} else {
		return constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_IS_NOT_A_STRING, evRight.Type())}
	}
}
