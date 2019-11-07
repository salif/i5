// SPDX-License-Identifier: GPL-3.0-or-later
package colors

import (
	"fmt"
	"html"

	"github.com/i5/i5/src/constants"
)

var colorize func(string, string) string = defaultColor

func SetColorFormat(format string) error {
	switch format {
	case "no-color":
		colorize = noColor
	case "html":
		colorize = htmlColor
	case "default":
		colorize = defaultColor
	default:
		return fmt.Errorf(constants.ARGS_UNKNOWN_CLR, format)
	}
	return nil
}

func Color(text string, color string) string {
	return colorize(text, color)
}

func Red(text string) string {
	return colorize(text, "red")
}

func Green(text string) string {
	return colorize(text, "green")
}

func Yellow(text string) string {
	return colorize(text, "yellow")
}

func Blue(text string) string {
	return colorize(text, "blue")
}

func Magenta(text string) string {
	return colorize(text, "magenta")
}

func Cyan(text string) string {
	return colorize(text, "cyan")
}

func defaultColor(text string, color string) string {
	if color == "red" {
		return "\x1b[91m" + text + "\x1b[0m"
	} else if color == "green" {
		return "\x1b[92m" + text + "\x1b[0m"
	} else if color == "yellow" {
		return "\x1b[93m" + text + "\x1b[0m"
	} else if color == "blue" {
		return "\x1b[94m" + text + "\x1b[0m"
	} else if color == "magenta" {
		return "\x1b[95m" + text + "\x1b[0m"
	} else if color == "cyan" {
		return "\x1b[96m" + text + "\x1b[0m"
	} else {
		return text
	}
}

func htmlColor(text string, color string) string {
	return "<span style='color:" + color + ";'>" + html.EscapeString(text) + "</span>"
}

func noColor(text string, color string) string {
	return text
}
