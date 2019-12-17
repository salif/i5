// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/object"
)

func evalLoop(node ast.Loop, env *object.Env) error {
	for {
		_, err := Eval(node.GetBody(), env)
		if err != nil {
			if er, ok := err.(constants.Error); ok {
				if er.Type == constants.ERROR_BREAK {
					break
				}
			}
			return err
		}
	}
	return nil
}
