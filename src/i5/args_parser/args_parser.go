// Copyright 2020 Salif Mehmed
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package args_parser

import (
	"fmt"
	"strings"

	"github.com/i5/i5/src/constants"
)

type ArgumentsParser struct {
	allArguments          []string
	realArguments         []string
	expectedBoolOptions   map[string]*bool
	expectedStringOptions map[string]*string
	helpBuffer            strings.Builder
}

func (s *ArgumentsParser) Init(args []string, usage string) {
	s.allArguments = args[1:]
	s.expectedBoolOptions = make(map[string]*bool, 0)
	s.expectedStringOptions = make(map[string]*string, 0)
	s.realArguments = make([]string, 0)
	s.helpBuffer = strings.Builder{}
	s.helpBuffer.WriteString(fmt.Sprintf("Usage:\n\n    %v\n\nOptions:\n\n", usage))
}

func (s *ArgumentsParser) IsTrue(str string) bool {
	return *s.expectedBoolOptions["--"+str]
}

func (s *ArgumentsParser) Get(str string) string {
	return *s.expectedStringOptions["--"+str]
}

func (s *ArgumentsParser) GetRealArguments() []string {
	return s.realArguments
}

func (s *ArgumentsParser) Bool(arg string, description string) {
	expectedBoolArgument := new(bool)
	s.expectedBoolOptions["--"+arg] = expectedBoolArgument
	s.helpBuffer.WriteString(fmt.Sprintf("    --%-26s%v\n", arg, description))
}

func (s *ArgumentsParser) String(arg string, description string, value string) {
	expectedStringArgument := new(string)
	s.expectedStringOptions["--"+arg] = expectedStringArgument
	s.helpBuffer.WriteString(fmt.Sprintf("    --%-26s%v\n", fmt.Sprintf("%v='%v'", arg, value), description))
}

func (s *ArgumentsParser) Parse() error {
	var options = true
	for _, arg := range s.allArguments {
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
			s.realArguments = append(s.realArguments, arg)
			options = false
		}
	}
	return nil
}

func (s *ArgumentsParser) GetHelp() string {
	return s.helpBuffer.String()
}
