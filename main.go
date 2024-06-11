// Ascii-art is a program which consists in receiving a string as an argument
// and outputting the string in a graphic representation using ASCII.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	ascii "ascii/utilities"
)

func main() {
	file_name := ""
	banner := ""
	ascii_art := ""

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}

	flag.StringVar(&banner, "output", "", "")
	flag.Parse()

	if !strings.Contains(os.Args[1], "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
		return
	}
	if !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Usage: --output=<fileName.txt> only files ending with .txt are allowed")
		return
	}

	input := os.Args[2] // user input

	file_name = "bannerfiles/" + os.Args[3] + ".txt"

	// fmt.Println(len(input))

	if input == "" {
		return
	}
	if input == "\\n" {
		fmt.Println()
		return
	}
	input = ascii.HandleBackspace(input)
	input = strings.ReplaceAll(string(input), "\\t", "   ") // handling the tab sequence

	// Read the ascii art text file
	file, err := os.ReadFile(strings.ToLower(file_name))
	if err != nil {
		fmt.Println(err)
		return
	}

	//
	fileData := ascii.StringContain(string(file))

	// Handling empty files
	if len(fileData) != 856 {
		fmt.Println("Error: >> Banner files  is empty with length of: ", len(file))
		return
	}

	input = strings.ReplaceAll(input, "\\n", "\n")
	inputParts, err := ascii.HandleNewLine(input)
	ascii.ErrHandler(err)

	for _, part := range inputParts {
		if part == "" {
			ascii_art += "\n" // Print a newline if the part is empty (i.e., consecutive newline characters)
			continue
		}
		for i := 0; i < 8; i++ { // this loop is responsible for the height of each character
			for _, char := range part { // iterates through each character of the current word
				startingIndex := ascii.GetStartingIndex(int(char)) // obtaining the starting position of the char
				if startingIndex >= 0 {
					ascii_art += fileData[startingIndex+i] // printing the character line by line
				}
			}
			ascii_art += "\n" // printing a new line after printing each line of the charcter
		}
	}

	os.WriteFile(banner, []byte(ascii_art), 0o644)
}
