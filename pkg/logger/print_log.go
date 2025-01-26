package logger

import (
	"fmt"

	"github.com/fatih/color"
)

type Options struct {
	Label string
}

var (
	cyan   = color.New(color.FgCyan).SprintFunc()
	green  = color.New(color.FgGreen).SprintFunc()
	blue   = color.New(color.FgBlue).SprintFunc()
	red    = color.New(color.FgRed).SprintFunc()
	yellow = color.New(color.FgYellow).SprintFunc()
)

type ColorOption int

const (
	ColorCyan ColorOption = iota
	ColorGreen
	ColorBlue
	ColorRed
	ColorYellow
)

type PrintOptions struct {
	Label      string
	TagText    string
	TagColor   ColorOption
	TitleColor ColorOption
	MsgColor   ColorOption
}

func getColorFunc(opt ColorOption) func(a ...interface{}) string {
	switch opt {
	case ColorCyan:
		return cyan
	case ColorGreen:
		return green
	case ColorBlue:
		return blue
	case ColorRed:
		return red
	case ColorYellow:
		return yellow
	default:
		return cyan
	}
}

func Printl(title string, message string, options ...PrintOptions) string {
	// Default options
	opt := PrintOptions{
		TagText:    "server",
		TagColor:   ColorGreen,
		TitleColor: ColorBlue,
		MsgColor:   ColorCyan,
	}

	// Override defaults if options provided
	if len(options) > 0 {
		opt = options[0]
	}

	// Apply colors
	tag := getColorFunc(opt.TagColor)(fmt.Sprintf("[%s]:", opt.TagText))

	// Apply title color based on label or specified color
	titleColor := opt.TitleColor
	if opt.Label != "" {
		switch opt.Label {
		case "success":
			titleColor = ColorGreen
		case "warning":
			titleColor = ColorYellow
		case "error":
			titleColor = ColorRed
		}
	}

	coloredTitle := getColorFunc(titleColor)(title)
	coloredMessage := getColorFunc(opt.MsgColor)(message)

	return fmt.Sprintf("%s %s %s", tag, coloredTitle, coloredMessage)
}
