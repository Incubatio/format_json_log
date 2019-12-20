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
	//Logger     string
	Level      string `json:",omitempty"`
	Msg        string `json:",omitempty"`
	Error      string `json:",omitempty"`
	Stacktrace string `json:",omitempty"`
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	// Init colors
	const ERROR = "error"
	const WARNING = "warning"
	const INFO = "info"
	const TITLE = "title"
	const JSON = "json"
	colors := map[string]*color.Color{}
	colors[ERROR] = color.New(color.FgRed, color.Bold)
	colors[WARNING] = color.New(color.FgYellow, color.Bold)
	colors[INFO] = color.New(color.FgGreen, color.Bold)
	colors[JSON] = color.New(color.FgMagenta, color.Bold)
	colors[TITLE] = color.New(color.FgWhite).Add(color.Underline)
	levels := []string{ERROR, WARNING, INFO}

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
	var result map[string]interface{}
	var data []byte
	for {
		inputText, _ = reader.ReadString('\n')
		inputLen = len(inputText) - 1
		output = ""

		if inputLen > 0 {
			parseError = json.Unmarshal([]byte(inputText), &result)
			if parseError != nil || result == nil {
				addLineBreak = prevLevel != "" && prevLevel != ERROR
				prevLevel = ""
				//output = "/!\\ Failed to parse Input /!\\"
				output = "!> " + strings.Replace(inputText, "\n", "", -1)
			} else {
				parseError = json.Unmarshal([]byte(inputText), &parsedLog)
				if parsedLog.Level == "" || !stringInSlice(parsedLog.Level, levels) {
					//output = "!> " + strings.Replace(inputText, "\n", "", -1)
					//parseError = json.UnmarshalIndent([]byte(inputText), &parsedLog)
					data, parseError = json.MarshalIndent(result, "", "  ")
					if parseError != nil {
						output = "!> " + strings.Replace(inputText, "\n", "", -1)
					} else {
						prevLevel = "json"
						output = "[" + colorize(prevLevel, prevLevel) + "]\n" + string(data)
					}
				} else {
					addLineBreak = prevLevel != parsedLog.Level && prevLevel != ERROR
					prevLevel = parsedLog.Level
					output += "[" + colorize(strings.ToUpper(parsedLog.Level), parsedLog.Level) + "] " + parsedLog.Msg
					//fmt.Println("->[DEBUG] " + inputText)
					if parsedLog.Error != "" {
						output += " \n  " + descriptionLabel + " " + parsedLog.Error
					}
					if parsedLog.Stacktrace != "" {
						stacktrace = strings.Replace(parsedLog.Stacktrace, "\t", "  ", -1)
						stacktrace = strings.Replace("\n"+stacktrace, "\n", "\n    ", -1)
						output += "\n  " + StrackTraceLabel + " " + stacktrace + "\n"
					}
					//fmt.Printf("->[DEBUG]%+v\n", parsedLog)
				}
			}
			if addLineBreak {
				output = "\n" + output
			}
			fmt.Println(output)
			parsedLog = nil
			result = nil
		}

	}
}
