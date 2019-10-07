// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

var (
	TRUE  = &object.Bool{Value: true}
	FALSE = &object.Bool{Value: false}
	NIL   = &object.Nil{}
)

func Run(program ast.Node) {
	Eval(program, object.InitEnv())
}

func nativeToBool(input bool) *object.Bool {
	if input {
		return TRUE
	}
	return FALSE
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR
	}
	return false
}

func isTrue(obj object.Object) bool {
	switch obj {
	case NIL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
