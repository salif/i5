// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/builtins"
	"github.com/i5/i5/src/object"
)

func evalProgram(program *ast.Program, env *object.Env) object.Object {
	var result object.Object
	for _, expr := range program.Body {
		result = Eval(expr, env)
		switch result := result.(type) {
		case *object.Error:
			return result
		}
	}

	return result
}
func evalIf(ie *ast.If, env *object.Env) object.Object {
	condition := Eval(ie.Condition, env)

	if isError(condition) {
		return condition
	}

	if isTrue(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NIL
	}
}

func evalSwitch(s *ast.Switch, env *object.Env) object.Object {
	return newError("'switch' not implemented yet")
	// TODO
}

func evalWhile(w *ast.While, env *object.Env) object.Object {
	return newError("'while' not implemented yet")
	// TODO
}

func evalImportExpr(i *ast.ImportExpr, env *object.Env) object.Object {
	return newError("'import expression' not implemented yet")
	// TODO
}
func evalImportStatement(i *ast.ImportStatement, env *object.Env) object.Object {
	return newError("'import statement' not implemented yet")
	// TODO
}

func evalTry(t *ast.Try, env *object.Env) object.Object {
	return newError("'try' not implemented yet")
	// TODO
}

func evalIdentifier(node *ast.Identifier, env *object.Env) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	} else {
		return newError("identifier not found: " + node.Value)
	}
}

func evalBuiltin(node *ast.Builtin, env *object.Env) object.Object {
	if builtin, ok := builtins.Get(node.Value); ok {
		return builtin
	} else {
		return newError("buitin not found: " + node.Value)
	}
}

func callFunction(fn object.Object, args []object.Object) object.Object {
	switch fn := fn.(type) {

	case *object.Function:
		env := extendFunctionEnv(fn, args)
		return unwrapReturnValue(Eval(fn.Body, env))
	case *object.Builtin:
		if result := fn.Function(args...); result != nil {
			return result
		}
		return NIL
	default:
		return newError("not a function: %s", fn.Type())
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
	}

	return obj
}

func evalBlock(block *ast.Block, env *object.Env) object.Object {
	var result object.Object
	for _, statement := range block.Body {
		result = Eval(statement, env)
		if result != nil {
			rt := result.Type()
			if rt == object.RETURN || rt == object.ERROR {
				return result
			}
		}
	}
	return result
}

func evalExpressions(exps []ast.Expression, env *object.Env) []object.Object {
	var result []object.Object
	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}
		result = append(result, evaluated)
	}
	return result
}
