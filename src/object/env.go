// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package object

type Env struct {
	Stack  map[string]Object
	parent *Env
}

func InitEnv() *Env {
	return &Env{Stack: map[string]Object{}}
}

func (e *Env) Contains(str string) bool {
	_, err := e.Stack[str]
	return err
}

func (e *Env) Get(str string) (Object, bool) {
	obj, ok := e.Stack[str]
	if !ok && e.parent != nil {
		obj, ok = e.parent.Get(str)
	}
	return obj, ok
}

func (e *Env) Set(str string, obj Object) bool {
	e.Stack[str] = obj
	return true
}

func (e *Env) Clone() *Env {
	env := InitEnv()
	env.parent = e
	return env
}
