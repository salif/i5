// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"fmt"

	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/object"
)

var (
	TRUE  = &object.Bool{Value: true}
	FALSE = &object.Bool{Value: false}
)

func RunPackage(dir string, arguments []string) {
	console.ThrowError(1, "not implemented yet")
	// TODO
}

func RunModule(module string, arguments []string) {
	console.ThrowError(1, "not implemented yet")
	// TODO
}

func RunFile(program ast.Node, arguments []string) {
	err := Eval(program, object.InitEnv())
	if err.Type() == object.ERROR {
		console.ThrowError(1, err.StringValue())
	}
}

func nativeToBool(input bool) *object.Bool {
	if input {
		return TRUE
	}
	return FALSE
}

func isError(obj object.Object) bool {
	return obj.Type() == object.ERROR
}

func isTrue(obj object.Object) bool {
	if obj == TRUE {
		return true
	} else {
		return false
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}
