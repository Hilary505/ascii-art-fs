package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	a "ascii/ascii_art"
)

// Error messages to be displayed when the usage is incorrect
const ErrorText = `Usage: go run . [OPTION] [STRING]

EX: go run . --color=<color> <substring to be colored> "something"`

const fsError = `Usage: go run . [STRING] [BANNER]

EX: go run . something standard`

func main() {
	// Check if there are any arguments passed
	if len(os.Args) <= 1 {
		fmt.Println(fsError)
		os.Exit(0)
	}

	// Default arguments
	input := ""
	bannerFile := "standard"
	subString := input

	// Define the color flag
	flgColor := flag.String("color", "", "Color")
	flag.Parse()
	// Validate color flag
	if !strings.Contains(os.Args[1], "--color=") && flag.NFlag() == 1 {
		fmt.Println(ErrorText)
		os.Exit(1)
	}
	color := a.ColorPicker(*flgColor)

	// Get non-flag arguments
	args := flag.Args()    // Non-flag arguments
	nArgs := len(args)     // Count of non-flag arguments
	nflags := flag.NFlag() // Count of flags

	// Handle arguments based on the number of flags and non-flag arguments
	if nflags == 0 {
		switch nArgs {
		case 1:
			input = args[0]
			subString = input
		case 2:
			input = args[0]
			bannerFile = args[1]
			subString = input
		default:
			fmt.Println(fsError)
			os.Exit(0)
		}
	} else {
		switch nArgs {
		case 1:
			input = args[0]
			subString = input
		case 2:
			subString = args[0]
			input = args[1]
		case 3:
			subString = args[0]
			input = args[1]
			bannerFile = args[2]
		default:
			fmt.Println(ErrorText)
			os.Exit(0)
		}
	}

	if strings.Contains(bannerFile, ".") {
		fmt.Println(fsError)
		os.Exit(0)
	}

	bannerFile += ".txt"

	contents, err := a.GetFile(bannerFile)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Verify the length of the contents
	if len(contents) != 856 {
		fmt.Println("Error")
		return
	}

	// Process and print the input with the specified color and substring
	fmt.Print(a.ProcessInput(contents, input, color, subString))
}
