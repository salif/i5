// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
	"github.com/i5/i5/src/types"
)

func Eval(nodei ast.Node, env *object.Env) object.Object {
	switch node := nodei.(type) {

	case ast.Program:
		return evalProgram(node, env, node.GetLine())

	case ast.Expression:
		return Eval(node.GetBody(), env)

	case ast.Block:
		return evalBlock(node, env, node.GetLine())

	case ast.Return:
		val := Eval(node.GetBody(), env)
		if isError(val) {
			return val
		}
		return object.Return{Value: val}

	case ast.Assign:
		left := node.GetLeft()
		switch left := left.(type) {
		case ast.Identifier:
			right := Eval(node.GetRight(), env)
			if isError(right) {
				return right
			}
			if isVoid(right) {
				right = object.Void{Value: left.GetValue()}
			}
			env.Set(left.GetValue(), right)
			return right
		case ast.Index:
			right := Eval(node.GetRight(), env)
			if isError(right) {
				return right
			}
			if isVoid(right) {
				return object.Error{Message: right.StringValue(), Line: node.GetLine()}
			}
			leftIndex := Eval(left.GetLeft(), env)
			if isError(leftIndex) {
				return leftIndex
			}
			if isVoid(leftIndex) {
				return object.Error{Message: leftIndex.StringValue(), Line: node.GetLine()}
			}
			if leftIndex.Type() == object.MAP {
				switch rightIndex := left.GetRight().(type) {
				case ast.Identifier:
					_map := leftIndex.(object.Map)
					return nativeToBool(_map.Set(object.String{Value: rightIndex.GetValue()}, right))
				case ast.Integer:
					_map := leftIndex.(object.Map)
					return nativeToBool(_map.Set(object.Integer{Value: rightIndex.GetValue()}, right))
				default:
					return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, leftIndex.Type(), left.GetOperator(), rightIndex.GetType()), Line: node.GetLine()}
				}
			} else {
				return object.Error{Message: console.Format(constants.IR_INVALID_POSTFIX, leftIndex.Type(), left.GetOperator()), Line: node.GetLine()}
			}
		default:
			return object.Error{Message: console.Format(constants.IR_CANNOT_ASSIGN, left.GetType()), Line: node.GetLine()}
		}
	case ast.Call:
		function := Eval(node.GetCaller(), env)
		if isError(function) {
			return function
		}
		if isVoid(function) {
			return object.Error{Message: function.StringValue(), Line: node.GetLine()}
		}
		args := evalExpressions(node.GetArguments(), env, node.GetLine())
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}
		if len(args) == 1 && isVoid(args[0]) {
			return object.Error{Message: args[0].StringValue(), Line: node.GetLine()}
		}
		return callFunction(function, args, node.GetLine())
	case ast.Function:
		return object.Function{Params: node.GetParams(), Body: node.GetBody(), Env: env}
	case ast.Identifier:
		return evalIdentifier(node, env, node.GetLine())
	case ast.Builtin:
		return evalBuiltin(node, env, node.GetLine())
	case ast.Integer:
		return object.Integer{Value: node.GetValue()}
	case ast.Float:
		return object.Float{Value: node.GetValue()}
	case ast.String:
		return object.String{Value: node.GetValue()}
	case ast.Bool:
		return nativeToBool(node.GetValue())
	case ast.Throw:
		val := Eval(node.GetBody(), env)
		if isError(val) {
			return val
		}
		if isVoid(val) {
			return object.Error{Message: val.StringValue(), Line: node.GetLine()}
		}
		return object.Error{Message: val.StringValue(), Line: node.GetLine()}
	case ast.Prefix:
		right := Eval(node.GetRight(), env)
		if isError(right) {
			return right
		}
		if isVoid(right) {
			return object.Error{Message: right.StringValue(), Line: node.GetLine()}
		}
		return evalPrefix(node.GetOperator(), right, env, node.GetLine())
	case ast.Infix:
		left := Eval(node.GetLeft(), env)
		if isError(left) {
			return left
		}
		if isVoid(left) {
			return object.Error{Message: left.StringValue(), Line: node.GetLine()}
		}
		right := Eval(node.GetRight(), env)
		if isError(right) {
			return right
		}
		if isVoid(right) {
			return object.Error{Message: right.StringValue(), Line: node.GetLine()}
		}
		return evalInfix(node.GetOperator(), left, right, env, node.GetLine())
	case ast.Postfix:
		left := Eval(node.GetLeft(), env)
		if isError(left) {
			return left
		}
		if node.GetOperator() == types.QM {
			return nativeToBool(left.Type() != object.VOID)
		}
		if isVoid(left) {
			return object.Error{Message: left.StringValue(), Line: node.GetLine()}
		}
		return evalPostfix(node.GetOperator(), left, env, node.GetLine())
	case ast.Index:
		left := Eval(node.GetLeft(), env)
		if isError(left) {
			return left
		}
		if isVoid(left) {
			return object.Error{Message: left.StringValue(), Line: node.GetLine()}
		}
		if left.Type() == object.MAP {
			switch rnode := node.GetRight().(type) {
			case ast.Identifier:
				_map := left.(object.Map)
				obj := _map.Get(object.String{Value: rnode.GetValue()})
				if isVoid(obj) {
					return object.Error{Message: console.Format(constants.IR_MAP_KEY_NOT_FOUND, rnode.GetValue()), Line: node.GetLine()}
				} else {
					return obj
				}
			case ast.Integer:
				_map := left.(object.Map)
				obj := _map.Get(object.Integer{Value: rnode.GetValue()})
				if isVoid(obj) {
					return object.Error{Message: console.Format(constants.IR_MAP_KEY_NOT_FOUND, rnode.GetValue()), Line: node.GetLine()}
				} else {
					return obj
				}
			default:
				return object.Error{Message: console.Format(constants.IR_INVALID_INFIX, left.Type(), node.GetOperator(), rnode.GetType()), Line: node.GetLine()}
			}
		} else {
			return object.Error{Message: console.Format(constants.IR_INVALID_POSTFIX, left.Type(), node.GetOperator()), Line: node.GetLine()}
		}
	case ast.If:
		return evalIf(node, env, node.GetLine())
	case ast.Switch:
		return evalSwitch(node, env, node.GetLine())
	case ast.Import:
		return evalImport(node, env, node.GetLine())
	case ast.Try:
		return evalTry(node, env, node.GetLine())
	default:
		return object.Error{Message: console.Format(constants.IR_INVALID_EVAL, node.GetType()), Line: node.GetLine()}
	}
}
