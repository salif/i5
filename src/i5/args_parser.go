package i5

import (
	"os"
	"strings"

	"github.com/i5/i5/src/errors"
)

type ArgsParser struct {
	arguments      []string
	bools          map[string]*bool
	strings        map[string]*string
	defaultR       *[]string
	defaultStrings []string
}

func InitArgsParser() ArgsParser {
	return ArgsParser{
		os.Args[1:],
		make(map[string]*bool),
		make(map[string]*string),
		new([]string),
		[]string{},
	}
}

func (s *ArgsParser) Empty() bool {
	return len(s.arguments) == 0
}

func (s *ArgsParser) Bool(arg string) (ret *bool) {
	ret = new(bool)
	s.bools["--"+arg] = ret
	return
}

func (s *ArgsParser) String(arg string) (ret *string) {
	ret = new(string)
	s.strings["--"+arg] = ret
	return ret
}

func (s *ArgsParser) Default() (ret *[]string) {
	ret = new([]string)
	s.defaultR = ret
	return
}

func (s *ArgsParser) Parse() {
	var options = true
	for _, arg := range s.arguments {
		if options && strings.HasPrefix(arg, "--") {
			if strings.Contains(arg, "=") {
				var index int = strings.Index(arg, "=")
				var ar, contains = s.strings[arg[:index]]
				if contains {
					*ar = arg[index+1:]
				} else {
					errors.FatalError(errors.F(errors.ARGS_UNKNOWN, arg[:index]), 1)
				}
			} else {
				var ar, contains = s.bools[arg]
				if contains {
					*ar = true
				} else {
					errors.FatalError(errors.F(errors.ARGS_UNKNOWN, arg), 1)
				}
			}
		} else {
			s.defaultStrings = append(s.defaultStrings, arg)
			options = false
		}
	}
	*s.defaultR = s.defaultStrings
}
