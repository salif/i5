// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
