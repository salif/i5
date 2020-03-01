// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import "fmt"

type Integer struct {
	line  uint32
	value int64
}

func (this Integer) GetType() string {
	return INTEGER
}

func (this Integer) Debug() string {
	return fmt.Sprint(this.value)
}

func (this Integer) GetLine() uint32 {
	return this.line
}

func (this Integer) GetValue() int64 {
	return this.value
}

func (this Integer) Init(line uint32, value int64) Integer {
	this.line = line
	this.value = value
	return this
}
