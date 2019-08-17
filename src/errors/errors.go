package errors

import (
	"github.com/i5/i5/src/io/console"
	"github.com/i5/i5/src/types"
	"os"
)

func FatalError(text interface{}, status int) types.Node {
	console.Println(text)
	Exit(status)
	return types.Node{}
}

func FatalLexerError(text string, line int, char string, status int) {
	console.Printf(text, line, char)
	Exit(status)
}

func Exit(status int) {
	os.Exit(status)
}
