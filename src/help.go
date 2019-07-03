package src

import (
	"fmt"
)

func PrintHelp() int {
	fmt.Println(`
Usage:

	i5 [options] [file]

options:

	--help      print help
`)
	return 0
}
