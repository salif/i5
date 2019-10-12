// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/builtins"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
)

func evalProgram(program *ast.Program, env *object.Env, line int) object.Object {
	var result object.Object
	for _, expr := range program.Body {
		result = Eval(expr, env)
		if isVoid(result) {
			return &object.Error{Message: result.StringValue(), Line: line}
		}
		switch result := result.(type) {
		case *object.Return:
			return result.Value
		case *object.Error:
			return result
		}
	}

	if result == nil {
		return &object.Void{}
	} else {
		return result
	}
}
func evalIf(ie *ast.If, env *object.Env, line int) object.Object {
	condition := Eval(ie.Condition, env)

	if isError(condition) {
		return condition
	}

	if isVoid(condition) {
		return &object.Error{Message: condition.StringValue(), Line: line}
	}

	if condition.Type() != object.BOOL {
		return &object.Error{Message: console.Format(constants.IR_NON_BOOL, condition.Type(), "if"), Line: line}
	}
	if isTrue(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return &object.Void{}
	}
}

func evalSwitch(s *ast.Switch, env *object.Env, line int) object.Object {
	return &object.Error{Message: console.Format(constants.IR_NOT_IMPLEMENTED, "switch"), Line: line}
	// TODO
}

func evalWhile(w *ast.While, env *object.Env, line int) object.Object {
	for {
		condition := Eval(w.Condition, env)
		if isError(condition) {
			return condition
		}

		if condition.Type() != object.BOOL {
			return &object.Error{Message: console.Format(constants.IR_NON_BOOL, condition.Type(), "while"), Line: line}
		}

		if isTrue(condition) {
			result := Eval(w.Body, env)
			if result.Type() == object.BREAK {
				break
			} else if result.Type() == object.CONTINUE {
				continue
			} else if result.Type() == object.ERROR {
				return result
			} else if result.Type() == object.RETURN {
				return result
			}
		} else {
			break
		}
	}

	return &object.Void{}
}

func evalImport(i *ast.Import, env *object.Env, line int) object.Object {
	return &object.Error{Message: console.Format(constants.IR_NOT_IMPLEMENTED, "import"), Line: line}
	// TODO
}

func evalTry(t *ast.Try, env *object.Env, line int) object.Object {
	result := Eval(t.Body, env)
	if isError(result) {
		if t.Catch == nil {
			return &object.Void{}
		}

		if t.Err != nil {
			env.Set(t.Err.Value, result)
			// TODO make err usable in catch
		}
		catchResult := Eval(t.Catch, env)

		switch catchResult.Type() {
		case object.ERROR:
			fallthrough
		case object.RETURN:
			fallthrough
		case object.BREAK:
			fallthrough
		case object.CONTINUE:
			return catchResult
		}

		if t.Finally == nil {
			return &object.Void{}
		}

		return Eval(t.Finally, env)

	} else {
		return result
	}
}

func evalIdentifier(node *ast.Identifier, env *object.Env, line int) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	} else {
		return &object.Error{Message: "identifier not found: " + node.Value, Line: line}
	}
}

func evalBuiltin(node *ast.Builtin, env *object.Env, line int) object.Object {
	if builtin, ok := builtins.Get(node.Value); ok {
		return builtin
	} else {
		return &object.Error{Message: "buitin not found: " + node.Value, Line: line}
	}
}

func callFunction(fn object.Object, args []object.Object, line int) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		if len(args) < len(fn.Params) {
			return &object.Error{Message: constants.IR_NOT_ENOUGH_ARGS, Line: line}
		} else {
			env := extendFunctionEnv(fn, args)
			result := Eval(fn.Body, env)
			switch result.Type() {
			case object.BREAK:
				fallthrough
			case object.CONTINUE:
				return &object.Void{}
			}
			return unwrapReturnValue(result)
		}
	case *object.Builtin:
		if result := fn.Function(args...); result != nil {
			return result
		} else {
			return &object.Error{Message: console.Format(constants.IR_BUILTIN_NOT_FOUND, fn.Name), Line: line}
		}
	default:
		return &object.Error{Message: console.Format(constants.IR_INVALID_CALL, fn.Type()), Line: line}
	}
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Env {
	env := fn.Env.Clone()
	for paramIdx, param := range fn.Params {
		env.Set(param.Value, args[paramIdx])
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.Return); ok {
		return returnValue.Value
	} else {
		return obj
	}
}

func evalBlock(block *ast.Block, env *object.Env, line int) object.Object {
	var result object.Object
	for _, statement := range block.Body {
		result = Eval(statement, env)
		switch result.Type() {
		case object.RETURN:
			fallthrough
		case object.ERROR:
			fallthrough
		case object.BREAK:
			fallthrough
		case object.CONTINUE:
			return result
		}
	}
	return &object.Void{}
}

func evalExpressions(exps []ast.Expression, env *object.Env, line int) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		} else if isVoid(evaluated) {
			return []object.Object{&object.Error{Message: evaluated.StringValue(), Line: line}}
		} else {
			result = append(result, evaluated)
		}
	}
	return result
}
