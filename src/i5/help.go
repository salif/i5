package i5

import (
	"fmt"
	"github.com/i5-lang/i5/src/error"
)

func PrintHelp() {
	fmt.Println(`
Usage:

	i5 [options] [file]

options:

	--help      print help
	`)
	error.Exit(0)
}
