// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Break struct {
}

func (this Break) Type() TYPE {
	return BREAK
}

func (this Break) StringValue() string {
	return "[type: break]"
}

type Continue struct {
}

func (this Continue) Type() TYPE {
	return CONTINUE
}

func (this Continue) StringValue() string {
	return "[type: continue]"
}
