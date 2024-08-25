package pkg

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

func PrintLog(title string, message string, options ...Options) string {
	var label string
	var tag string

	// default tag
	tag = green("[server]:")

	if options == nil {
		label = blue(title)
	}

	for _, opt := range options {
		// label condition
		switch opt.Label {
		case "success":
			label = green(title)

		case "warning":
			label = yellow(title)

		case "error":
			label = red(title)

		case "custom":
			label = blue(title)

		default:
			label = blue(title)
		}
	}

	logMessage := cyan(message)
	result := fmt.Sprintf("%s %s %s", tag, label, logMessage)

	return result
}
