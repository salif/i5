// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalCall(node ast.Call, env *object.Env) object.Object {
	var evaluatedFunctionCaller object.Object = Eval(node.GetCaller(), env)
	var evaluatedArguments []object.Object = []object.Object{}
	for _, e := range node.GetArguments() {
		var evaluatedArgument object.Object = Eval(e, env)
		if ErrorType(evaluatedArgument) == FATAL {
			return evaluatedArgument
		}
		evaluatedArguments = append(evaluatedArguments, evaluatedArgument)
	}
	return callFunction(evaluatedFunctionCaller, evaluatedArguments, node.GetLine())
}

func callFunction(fn object.Object, args []object.Object, line uint32) object.Object {
	switch fn := fn.(type) {
	case object.Function:
		if len(args) < len(fn.Params) {
			return newError(true, line, constants.ERROR_RANGE, constants.IR_NOT_ENOUGH_ARGS)
		}
		env := extendFunctionEnv(fn, args)
		var result object.Object = Eval(fn.Body, env)
		if result.Type() == object.RETURN {
			var forReturn object.Return = result.(object.Return)
			return forReturn.Value
		} else {
			return result
		}

	case object.BuiltinFunction:
		if len(args) < fn.MinParams {
			return newError(true, line, constants.ERROR_RANGE, constants.IR_NOT_ENOUGH_ARGS)
		}
		if result := fn.Function(args...); result != nil {
			return result
		} else {
			return Nil(line)
		}
	default:
		return newError(true, line, constants.ERROR_TYPE, constants.IR_IS_NOT_A_FUNCTION, fn.Type())
	}
}

func extendFunctionEnv(fn object.Function, args []object.Object) *object.Env {
	var env *object.Env = fn.Env.Clone()
	for paramIdx, param := range fn.Params {
		env.Set(param.GetValue(), args[paramIdx])
	}
	return env
}
