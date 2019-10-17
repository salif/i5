// SPDX-License-Identifier: GPL-3.0-or-later
package i5

import (
	"bytes"
	"os"
	"strings"

	"github.com/i5/i5/src/constants"
	"github.com/i5/i5/src/io/console"
)

type ArgumentsParser struct {
	arguments      []string
	bools          map[string]*bool
	strings        map[string]*string
	defaultR       *[]string
	defaultStrings []string
	help           bytes.Buffer
	helpMessage    string
}

func InitArgumentsParser(usage string, options string) ArgumentsParser {
	ap := ArgumentsParser{
		os.Args[1:],
		make(map[string]*bool),
		make(map[string]*string),
		new([]string),
		[]string{},
		bytes.Buffer{},
		"",
	}
	ap.help.WriteString(console.Format("Usage:\n\n    %v\n\n%v:\n\n", usage, options))
	return ap
}

func (s *ArgumentsParser) Empty() bool {
	return len(s.arguments) == 0
}

func (s *ArgumentsParser) Bool(arg string, description string) (ret *bool) {
	ret = new(bool)
	s.bools["--"+arg] = ret
	s.help.WriteString(console.Format("    --%-26s%v\n", arg, description))
	return
}

func (s *ArgumentsParser) String(arg string, description string, value string) (ret *string) {
	ret = new(string)
	s.strings["--"+arg] = ret
	s.help.WriteString(console.Format("    --%-26s%v\n", console.Format("%v='%v'", arg, value), description))
	return ret
}

func (s *ArgumentsParser) Default() (ret *[]string) {
	ret = new([]string)
	s.defaultR = ret
	return
}

func (s *ArgumentsParser) Parse() {
	var options = true
	for _, arg := range s.arguments {
		if options && strings.HasPrefix(arg, "--") {
			if strings.Contains(arg, "=") {
				var index int = strings.Index(arg, "=")
				var ar, contains = s.strings[arg[:index]]
				if contains {
					*ar = arg[index+1:]
				} else {
					console.ThrowError(127, constants.ARGS_UNKNOWN, arg[:index])
				}
			} else {
				var ar, contains = s.bools[arg]
				if contains {
					*ar = true
				} else {
					console.ThrowError(127, constants.ARGS_UNKNOWN, arg)
				}
			}
		} else {
			s.defaultStrings = append(s.defaultStrings, arg)
			options = false
		}
	}
	*s.defaultR = s.defaultStrings
	s.helpMessage = s.help.String()
}

func (s *ArgumentsParser) PrintHelp() {
	console.Println(s.helpMessage)
}
