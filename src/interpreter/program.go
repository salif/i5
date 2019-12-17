// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalProgram(node ast.Program, env *object.Env) error {
	for _, fn := range node.GetBody() {
		_, err := Eval(fn, env)
		if err != nil {
			return err
		}
	}
	return nil
}
