// SPDX-License-Identifier: GPL-3.0-or-later
package interpreter

import (
	"github.com/i5/i5/src/ast"
	"github.com/i5/i5/src/object"
)

func evalBlock(node ast.Block, env *object.Env) error {
	for _, statement := range node.GetBody() {
		_, err := Eval(statement, env)
		if err != nil {
			return err
		}
	}
	return nil
}
