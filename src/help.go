package src

import (
	"fmt"
)

func PrintHelp() int {
	fmt.Println(`
Usage:

	iota5 [options] [file]

options:

	--help      print help
`)
	return 0
}
