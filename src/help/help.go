package help

import (
	"fmt"
)

func Print() int {
	fmt.Println(`
	Usage:

		iota5 [options] [file]

	options:

		--help         print help
	`)
	return 0
}
