// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

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
