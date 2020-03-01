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

func evalCall(node ast.Call, env *object.Env) (object.Object, error) {
	evCaller, err := Eval(node.GetCaller(), env)
	if err != nil {
		return nil, err
	}

	var evaluatedArguments []object.Object = []object.Object{}
	for _, e := range node.GetArguments() {
		ev, err := Eval(e, env)
		if err != nil {
			return nil, err
		}

		evaluatedArguments = append(evaluatedArguments, ev)
	}

	return callFunction(evCaller, evaluatedArguments, node.GetLine())
}

func callFunction(fn object.Object, args []object.Object, line uint32) (object.Object, error) {
	switch fn := fn.(type) {
	case object.Function:
		if len(args) < len(fn.Params) {
			return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: constants.IR_NOT_ENOUGH_ARGS}
		}
		env := extendFunctionEnv(fn, args)
		result, err := Eval(fn.Body, env)
		if err != nil {
			if er, ok := err.(constants.Error); ok {
				if er.Type == constants.ERROR_RETURN {
					return er.Value.(object.Object), nil
				}
			}
		}
		return result, err

	case object.BuiltinFunction:
		if len(args) < fn.MinParams {
			return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: constants.IR_NOT_ENOUGH_ARGS}
		}
		if result := fn.Function(args...); result != nil {
			return result, nil
		} else {
			return Null, nil
		}
	default:
		return nil, constants.Error{Line: line, Type: constants.ERROR_FATAL, Message: fmt.Sprintf(constants.IR_IS_NOT_A_FUNCTION, fn.Type())}
	}
}

func extendFunctionEnv(fn object.Function, args []object.Object) *object.Env {
	var env *object.Env = fn.Env.Clone()
	for paramIdx, param := range fn.Params {
		env.Set(param.GetValue(), args[paramIdx])
	}
	return env
}
