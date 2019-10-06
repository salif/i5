package file

import (
	"io/ioutil"
	"os"

	"github.com/i5/i5/src/io/console"
)

func Write(fileName string, content string, perm os.FileMode) (result string) {
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		err := ioutil.WriteFile(fileName, []byte(content), perm)
		if err != nil {
			return console.Format(console.FILE_WRITE_CANNOT_WRITE, fileName)
		} else {
			return
		}
	} else {
		return console.Format(console.FILE_WRITE_EXISTS, fileName)
	}
}
