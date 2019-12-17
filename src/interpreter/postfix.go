// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalPostfixNode(node ast.Postfix, env *object.Env) (object.Object, error) {
	evLeft, err := Eval(node.GetLeft(), env)
	if err != nil {
		return nil, err
	}
	return evalPostfix(node.GetOperator(), evLeft, env, node.GetLine())
}

func evalPostfix(operator string, right object.Object, env *object.Env, line uint32) (object.Object, error) {
	switch operator {
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_POSTFIX, right.Type(), operator)}
	}
}
