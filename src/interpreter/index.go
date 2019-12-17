// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalIndex(node ast.Index, env *object.Env) (object.Object, error) {
	evLeft, err := Eval(node.GetLeft(), env)

	if err != nil {
		return nil, err
	}

	if evLeft.Type() == constants.TYPE_MODULE {
		switch rnode := node.GetRight().(type) {
		case ast.Identifier:
			_module := evLeft.(object.Module)
			obj, _ := _module.Env.Get(rnode.GetValue())
			return obj, nil

		default:
			return nil, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_INFIX, evLeft.Type(), node.GetOperator(), rnode.GetType())}
		}
	} else {
		return nil, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_POSTFIX, evLeft.Type(), node.GetOperator())}
	}
}
