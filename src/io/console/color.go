package console

import "html"

var clr func(string, string) Color = defaultColor

type Color struct {
	Value string
}

func (c Color) String() string {
	return c.Value
}

func (c Color) ValueOf(color string) Color {
	return clr(c.Value, color)
}

func (c Color) Red() Color {
	return clr(c.Value, "red")
}

func (c Color) Green() Color {
	return clr(c.Value, "green")
}

func (c Color) Yellow() Color {
	return clr(c.Value, "yellow")
}

func (c Color) Blue() Color {
	return clr(c.Value, "blue")
}

func (c Color) Magenta() Color {
	return clr(c.Value, "magenta")
}

func (c Color) Cyan() Color {
	return clr(c.Value, "cyan")
}

func defaultColor(text string, color string) Color {
	if color == "red" {
		return Color{"\x1b[91m" + text + "\x1b[0m"}
	} else if color == "green" {
		return Color{"\x1b[92m" + text + "\x1b[0m"}
	} else if color == "yellow" {
		return Color{"\x1b[93m" + text + "\x1b[0m"}
	} else if color == "blue" {
		return Color{"\x1b[94m" + text + "\x1b[0m"}
	} else if color == "magenta" {
		return Color{"\x1b[95m" + text + "\x1b[0m"}
	} else if color == "cyan" {
		return Color{"\x1b[96m" + text + "\x1b[0m"}
	} else {
		return Color{text}
	}
}

func htmlColor(text string, color string) Color {
	return Color{"<span style='color:" + color + ";'>" + html.EscapeString(text) + "</span>"}
}

func noColor(text string, color string) Color {
	return Color{text}
}
