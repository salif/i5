// SPDX-License-Identifier: GPL-3.0-or-later
package ast

import "github.com/i5/i5/src/io/console"

type Try struct {
	line    int
	token   string
	body    Block
	errn    bool
	err     Identifier
	catch   Block
	finally Block
}

func (this Try) GetType() int {
	return TRY
}

func (this Try) Print() {
	console.Print(this.token)
	console.Print(" ")
	this.body.Print()
	if this.catch.body != nil {
		console.Print(" catch ")
		if this.errn {
			this.err.Print()
			console.Print(" ")
		}
		this.catch.Print()
		if this.finally.body != nil {
			console.Print(" finally ")
			this.finally.Print()
		}
	}
}

func (this Try) GetLine() int {
	return this.line
}

func (this Try) GetBody() Block {
	return this.body
}

func (this Try) HaveCatch() bool {
	return this.catch.body != nil
}

func (this Try) GetCatch() Block {
	return this.catch
}

func (this Try) HaveErr() bool {
	return this.errn
}

func (this Try) GetErr() Identifier {
	return this.err
}

func (this Try) HaveFinally() bool {
	return this.finally.body != nil
}

func (this Try) GetFinally() Block {
	return this.finally
}

func (this Try) Init(line int, token string) Try {
	this.line = line
	this.token = token
	this.errn = false
	return this
}

func (this *Try) SetBody(body Block) {
	this.body = body
}

func (this *Try) SetErr(err Identifier) {
	this.err = err
	this.errn = true
}

func (this *Try) SetCatch(catch Block) {
	this.catch = catch
}

func (this *Try) SetFinally(finally Block) {
	this.finally = finally
}
