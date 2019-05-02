package help

import (
	"fmt"
)

func Print() int {
	fmt.Println(`
	Usage:

		iota9 [options] [file]

	options:

		--help         print help
	`)
	return 0
}
