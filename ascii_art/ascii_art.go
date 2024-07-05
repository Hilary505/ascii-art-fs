package ascii_art

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetFile reads from the file specified by filename and returns its contents
func GetFile(filename string) ([]string, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("An error", err)
		os.Exit(1)
	}

	if len(file) == 0 {
		fmt.Println("Error: The banner file is empty")
		os.Exit(1)
	}

	myfile := string(file)
	var contents []string

	// Different line splitting logic based on the file type
	if filename == "thinkertoy.txt" {
		contents = strings.Split(myfile, "\r\n")
	} else {
		contents = strings.Split(myfile, "\n")
	}

	return contents, nil
}

// ColorPicker maps a color name or code to its corresponding ANSI color code
func ColorPicker(color string) (colorCode string) {
	if color == "" {
		return ""
	}

	colorChat := map[string]string{
		"reset":          "\u001b[39m",
		"black":          "\u001b[30m",
		"red":            "\u001b[31m",
		"green":          "\u001b[32m",
		"yellow":         "\u001b[33m",
		"blue":           "\u001b[34m",
		"magenta":        "\u001b[35m",
		"cyan":           "\u001b[36m",
		"white":          "\u001b[37m",
		"bright_black":   "\u001b[90m",
		"bright_red":     "\u001b[91m",
		"bright_green":   "\u001b[92m",
		"bright_yellow":  "\u001b[93m",
		"bright_blue":    "\u001b[94m",
		"bright_magenta": "\u001b[95m",
		"bright_cyan":    "\u001b[96m",
		"bright_white":   "\u001b[97m",
		"orange":         "\033[38;5;208m",
		"violet":         "\033[38;5;129m",
		"indigo":         "\033[38;5;54m",
		"maroon":         "\033[38;5;52m",
		"purple":         "\033[38;5;165m",
		"pink":           "\033[38;5;206m",
		"brown":          "\033[38;5;130m",
		"wheat":          "\033[38;5;229m",
		"tomato":         "\033[38;5;209m",
		"smoke":          "\033[38;5;245m",
		"gray":           "\033[38;5;240m",
		"gold":           "\033[38;5;178m",
		"avocado":        "\033[38;5;58m",
		"oceanblue":      "\033[38;5;27m",
		"navyblue":       "\033[38;5;18m",
		"amber":          "\033[38;5;214m",
	}

	// Retrieve the color code from the map
	colorCode, ok := colorChat[color]

	if ok {
		return colorCode
	}

	// If the color is not found, check if it is an ANSI code
	ansi, err := strconv.Atoi(color)
	if err != nil || ansi < 0 || ansi > 255 {
		fmt.Println("Error: color Not Found")
		os.Exit(1)
	}

	// Return the ANSI color code
	return fmt.Sprintf("\033[38;5;%dm", ansi)
}

// getCi finds the starting index and length of the substring in the text
func getCi(substring, text string) (ci, n int) {
	arrSubSrting := strings.Split(substring, "\\n")
	word := ""

	for _, word = range arrSubSrting {
		ci = strings.Index(text, word)
		if ci != -1 {
			return ci, len(word)
		}
	}

	return ci, len(word)
}

// ProcessInput accepts the contents of the ASCII art file and the input string,
// and processes the input to display the corresponding ASCII art
func ProcessInput(contents []string, input, color, subString string) (strArt string) {
	count := 0
	// Replace newline and tab characters with their respective escape sequences
	strInput := strings.ReplaceAll(input, "\n", "\\n")
	strInput = strings.ReplaceAll(strInput, "\\t", "    ")
	newInput := strings.Split(strInput, "\\n")

	start := -1
	n := 0
	ci := 0

	for _, arg := range newInput {
		if arg == "" {
			count++
			if count < len(newInput) {
				strArt += "\n"
			}
			continue
		}

		for i := 1; i <= 8; i++ {
			// Find the starting index and length of the substring
			if subString != "" && subString != input {
				start, n = getCi(subString, arg)
			}

			if subString == input {
				start = 0
			}

			for j, ch := range arg {
				if ch > 126 {
					fmt.Println("The text contains an unprintable character", ch)
					os.Exit(0)
				}

				index := int(ch-32)*9 + i

				if index >= 0 && index < len(contents) {
					// Apply color if the current index matches the starting index of the substring
					if start == j && color != "" {
						strArt += color
					}

					strArt += (contents[index])

					// Reset color after the substring has been processed
					if start != -1 && start+n-1 == j && j < len(arg)-1 && subString != "" && subString != input {
						strArt += ColorPicker("reset")
						ci, _ = getCi(subString, arg[j+1:])
						start = ci + j + 1
					}

				}
			}
			if color != "" {
				strArt += ColorPicker("reset")
			}
			strArt += "\n"
		}
	}

	return strArt
}
