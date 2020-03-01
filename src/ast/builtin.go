// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

type Builtin struct {
	line  uint32
	value string
}

func (this Builtin) GetType() string {
	return BUILTIN
}

func (this Builtin) Debug() string {
	return "$" + this.value
}

func (this Builtin) GetLine() uint32 {
	return this.line
}

func (this Builtin) GetValue() string {
	return this.value
}

func (this Builtin) Init(line uint32, value string) Builtin {
	this.line = line
	this.value = value
	return this
}
