package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

type _JSONLog struct {
	Logger     string
	Level      string
	Msg        string
	Error      string
	Stacktrace string
}

func main() {

	// Init colors
	const ERROR = "error"
	const WARNING = "warning"
	const INFO = "info"
	const TITLE = "title"
	colors := map[string]*color.Color{}
	colors[ERROR] = color.New(color.FgRed, color.Bold)
	colors[WARNING] = color.New(color.FgYellow, color.Bold)
	colors[INFO] = color.New(color.FgGreen, color.Bold)
	colors[TITLE] = color.New(color.FgWhite).Add(color.Underline)

	colorize := func(str string, level string) string {
		if level != "" {
			str = colors[level].Sprint(str)
		}
		return str
	}

	// Start Program
	reader := bufio.NewReader(os.Stdin)

	var output string
	var prevLevel string
	var stacktrace string
	var inputText string
	var _ error
	var inputLen int
	var parseError error
	var addLineBreak bool

	var descriptionLabel = colorize("Description:", TITLE)
	var StrackTraceLabel = colorize("Stacktrace:", TITLE)

	var parsedLog *_JSONLog
	for {
		inputText, _ = reader.ReadString('\n')
		inputLen = len(inputText) - 1
		output = ""

		if inputLen > 0 {
			parseError = json.Unmarshal([]byte(inputText), &parsedLog)
			if parseError != nil || parsedLog == nil || parsedLog.Level == "" {
				addLineBreak = prevLevel != "" && prevLevel != ERROR
				prevLevel = ""
				//output = "/!\\ Failed to parse Input /!\\"
				output = "!> " + strings.Replace(inputText, "\n", "", -1)
			} else {
				addLineBreak = prevLevel != parsedLog.Level && prevLevel != ERROR
				prevLevel = parsedLog.Level
				output += "[" + colorize(strings.ToUpper(parsedLog.Level), parsedLog.Level) + "] " + parsedLog.Msg
				//fmt.Println("->[DEBUG] " + inputText)
				if parsedLog.Error != "" || parsedLog.Stacktrace != "" {
					stacktrace = strings.Replace(parsedLog.Stacktrace, "\t", "  ", -1)
					stacktrace = strings.Replace("\n"+stacktrace, "\n", "\n    ", -1)
					output += " \n  " + descriptionLabel + " " + parsedLog.Error + "\n  " + StrackTraceLabel + " " + stacktrace + "\n"
				}
				//fmt.Printf("->[DEBUG]%+v\n", parsedLog)
			}
			if addLineBreak {
				output = "\n" + output
			}
			fmt.Println(output)
			parsedLog = nil
		}

	}
}
