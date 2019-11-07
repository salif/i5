// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"bufio"
	"fmt"
	"os"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/file"
)

func InitModule() error {
	if file.Exists(constants.I5_MOD_FILE_NAME) {
		return fmt.Errorf(constants.FILE_FILE_EXISTS, constants.I5_MOD_FILE_NAME)
	}
	reader := bufio.NewReader(os.Stdin)
	fline := input(reader, "main package directory (./) ", "./")
	sline := input(reader, "modules directory (./i5_modules/) ", "./i5_modules/")
	return file.Write(constants.I5_MOD_FILE_NAME, fmt.Sprintf("%v\n%v\n", fline, sline), 0644)
}

func input(b *bufio.Reader, str string, def string) string {
	fmt.Print(str)
	answer, _ := b.ReadString('\n')
	if answer == "\n" {
		return def
	} else {
		return answer[:len(answer)-1]
	}
}
