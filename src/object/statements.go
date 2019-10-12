// SPDX-License-Identifier: GPL-3.0-or-later
package object

type Break struct {
}

func (b *Break) Type() TYPE {
	return BREAK
}

func (b *Break) StringValue() string {
	return "[type: break]"
}

type Continue struct {
}

func (c *Continue) Type() TYPE {
	return CONTINUE
}

func (c *Continue) StringValue() string {
	return "[type: continue]"
}
