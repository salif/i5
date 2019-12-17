// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalAssign(node ast.Assign, env *object.Env) (object.Object, error) {
	var left ast.Node = node.GetLeft()
	evRight, err := Eval(node.GetRight(), env)

	if err != nil {
		return nil, err
	}

	switch left := left.(type) {

	case ast.Identifier:
		env.Set(left.GetValue(), evRight)
		return evRight, nil

	default:
		return nil, constants.Error{Line: node.GetLine(), Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_CANNOT_ASSIGN, left.GetType())}
	}
}
