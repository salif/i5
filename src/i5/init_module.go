package i5

import (
	"github.com/i5/i5/src/errors"
	"github.com/i5/i5/src/io/file"
)

func InitModule() {
	var err string = file.Write("module.json", newModuleTemplate(), 0644)
	if err != "" {
		errors.FatalError(err, 1)
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
	"url": "",
	"main": "main.i5",
	"dependencies": {
		"@i5": ">0.0.1"
	},
	"author": {
		"name": "",
		"email": "",
		"url": ""
	},
	"license": ""
}
`
}
