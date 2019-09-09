package file

import (
	"io/ioutil"
	"os"

	"github.com/i5/i5/src/errors"
)

func Read(fileName string) []byte {
	info, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		errors.FatalError(errors.F(errors.READER_NOT_FOUND, fileName), 1)
	} else if info.IsDir() {
		errors.FatalError(errors.F(errors.READER_DIR, fileName), 1)
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		errors.FatalError(errors.F(errors.READER_CANNOT_READ, fileName), 1)
	}
	return file
}
