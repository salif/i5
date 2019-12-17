// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalPrefixNode(node ast.Prefix, env *object.Env) (object.Object, error) {
	evRight, err := Eval(node.GetRight(), env)
	if err != nil {
		return nil, err
	}
	return evalPrefix(node.GetOperator(), evRight, env, node.GetLine())
}

func evalPrefix(operator string, right object.Object, env *object.Env, line uint32) (object.Object, error) {
	switch operator {
	case constants.TOKEN_NOT:
		if right.Type() == constants.TYPE_BOOLEAN {
			switch right {
			case TRUE:
				return FALSE, nil
			case FALSE:
				return TRUE, nil
			default:
				return FALSE, nil
			}
		} else {
			return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_PREFIX, operator, right.Type())}
		}
	case constants.TOKEN_BNOT:
		if right.Type() == constants.TYPE_INTEGER {
			value := right.(object.Integer).Value
			return object.Integer{Value: ^value}, nil
		} else {
			return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_PREFIX, operator, right.Type())}
		}
	case constants.TOKEN_MINUS:
		if right.Type() == constants.TYPE_INTEGER {
			value := right.(object.Integer).Value
			return object.Integer{Value: -value}, nil
		} else if right.Type() == constants.TYPE_FLOAT {
			value := right.(object.Float).Value
			return object.Float{Value: -value}, nil
		} else {
			return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_PREFIX, operator, right.Type())}
		}
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_INVALID_PREFIX, operator, right.Type())}
	}
}
