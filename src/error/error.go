package error

import (
	"fmt"
	"github.com/i5-lang/i5/src/types"
	"os"
)

func FatalError(describtion interface{}, status int) types.Node {
	fmt.Println(describtion)
	Exit(status)
	return types.Node{}
}

func FatalLexerError(describtion string, line int, char string, status int) {
	fmt.Printf(describtion, line, char)
	Exit(status)
}

func Exit(status int) {
	os.Exit(status)
}

func Print(str ...interface{}) {
	for i := 0; i < len(str); i++ {
		fmt.Print(str[i])
	}
	fmt.Println("")
}
