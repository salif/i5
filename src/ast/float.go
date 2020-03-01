// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package ast

import "fmt"

type Float struct {
	line  uint32
	value float64
}

func (this Float) GetType() string {
	return FLOAT
}

func (this Float) Debug() string {
	return fmt.Sprint(this.value)
}

func (this Float) GetLine() uint32 {
	return this.line
}

func (this Float) GetValue() float64 {
	return this.value
}

func (this Float) Init(line uint32, value float64) Float {
	this.line = line
	this.value = value
	return this
}
