// SPDX-License-Identifier: GPL-3.0-or-later
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
