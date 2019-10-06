// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/io/file"
)

func InitModule() {
	var err string = file.Write("module.json", newModuleTemplate(), 0644)
	if err != "" {
		console.ThrowError(1, err)
	}
	file.Write("main.i5", newEmptyTemplate(), 0744)
}

func newEmptyTemplate() string {
	return `#!/usr/bin/env i5

fn main() {
}

`
}

func newModuleTemplate() string {
	return `{
    "name": "",
    "version": "",
    "description": "",
    "main": "main.i5",
    "author": {
        "name": "",
        "email": ""
    },
    "license": ""
}
`
}
