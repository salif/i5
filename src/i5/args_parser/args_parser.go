// SPDX-License-Identifier: GPL-3.0-or-later
package args_parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/i5/i5/src/constants"
)

type ArgumentsParser struct {
	realArguments         []string
	expectedBoolOptions   map[string]*bool
	expectedStringOptions map[string]*string
	notOptions            []string
	helpBuffer            strings.Builder
	finalHelpMessage      string
}

func (s *ArgumentsParser) Init(usage string, options string) {
	s.realArguments = os.Args[1:]
	s.expectedBoolOptions = make(map[string]*bool)
	s.expectedStringOptions = make(map[string]*string)
	s.notOptions = []string{}
	s.helpBuffer = strings.Builder{}
	s.finalHelpMessage = ""

	s.helpBuffer.WriteString(fmt.Sprintf("Usage:\n\n    %v\n\n%v:\n\n", usage, options))
}

func (s *ArgumentsParser) IsEmpty() bool {
	return len(s.realArguments) == 0
}

func (s *ArgumentsParser) IsTrue(str string) bool {
	return *s.expectedBoolOptions["--"+str]
}

func (s *ArgumentsParser) Get(str string) string {
	return *s.expectedStringOptions["--"+str]
}

func (s *ArgumentsParser) GetNotOptions() []string {
	return s.notOptions
}

func (s *ArgumentsParser) Bool(arg string, description string) {
	expectedBoolArgument := new(bool)
	s.expectedBoolOptions["--"+arg] = expectedBoolArgument
	s.helpBuffer.WriteString(fmt.Sprintf("    --%-26s%v\n", arg, description))
	return
}

func (s *ArgumentsParser) String(arg string, description string, value string) (ret *string) {
	expectedStringArgument := new(string)
	s.expectedStringOptions["--"+arg] = expectedStringArgument
	s.helpBuffer.WriteString(fmt.Sprintf("    --%-26s%v\n", fmt.Sprintf("%v='%v'", arg, value), description))
	return ret
}

func (s *ArgumentsParser) Parse() error {
	var options = true
	for _, arg := range s.realArguments {
		if options && strings.HasPrefix(arg, "--") {
			if strings.Contains(arg, "=") {
				var index int = strings.Index(arg, "=")
				if ar, contains := s.expectedStringOptions[arg[:index]]; contains {
					*ar = arg[index+1:]
				} else {
					return fmt.Errorf(constants.ARGS_UNKNOWN, arg[:index])
				}
			} else {
				if ar, contains := s.expectedBoolOptions[arg]; contains {
					*ar = true
				} else {
					return fmt.Errorf(constants.ARGS_UNKNOWN, arg)
				}
			}
		} else {
			s.notOptions = append(s.notOptions, arg)
			options = false
		}
	}
	s.finalHelpMessage = s.helpBuffer.String()
	return nil
}

func (s *ArgumentsParser) GetHelp() string {
	return s.finalHelpMessage
}
