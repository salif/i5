package console

import "html"

var clr func(string, string) string = Colorize

func Color(text string, color string) string {
	return clr(text, color)
}

func SetColorizer(fn func(string, string) string) {
	clr = fn
}

func Colorize(text string, color string) string {
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

func HTML(text string, color string) string {
	return "<span style='color:" + color + ";'>" + html.EscapeString(text) + "</span>"
}

func NoColor(text string, color string) string {
	return text
}
