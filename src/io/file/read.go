package file

import (
	"io/ioutil"
	"os"

	"github.com/i5/i5/src/io/console"
)

func Read(fileName string) []byte {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		console.ThrowError(1, console.FILE_READ_NOT_FOUND, fileName)
	} else if info.IsDir() {
		console.ThrowError(1, console.FILE_READ_DIR, fileName)
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		console.ThrowError(1, console.FILE_READ_CANNOT_READ, fileName)
	}
	return file
}
