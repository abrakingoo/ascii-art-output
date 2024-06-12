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
		file_name := ""     // Initialize variable for file name.
		banner := ""        // Initialize variable for banner.
		ascii_art := ""     // Initialize variable for ASCII art.
		input := ""         // Initialize variable for user input.
		art_switch := true // Initialize switch for ASCII art.
	
		// Check if only one argument is provided from command line.
		if len(os.Args[1:]) == 1 {
			file_name = "bannerfiles/standard.txt"
			input = os.Args[1]
			art_switch = false
		}
	
		// Check if the number of arguments is not equal to 4 and ASCII art switch is on.
		if len(os.Args) != 4 && art_switch {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	
		// Check if more than two arguments are provided.
		if len(os.Args) > 2 {
			input = os.Args[2]                                     // Set user input.
			file_name = "bannerfiles/" + os.Args[3] + ".txt" // Set file name.
		}
	
		flag.StringVar(&banner, "output", "", "") // Set banner to the value of the 'output' flag.
		flag.Parse()                               // Parse command line flags.
	
		// Check if the first argument does not contain '--output=' and ASCII art switch is on.
		if !strings.Contains(os.Args[1], "--output=") && art_switch {
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nEX: go run . --output=<fileName.txt> something standard")
			return
		}
	
		// Check if the first argument does not have a suffix of '.txt' and ASCII art switch is on.
		if !strings.HasSuffix(os.Args[1], ".txt") && art_switch {
			fmt.Println("Usage: --output=<fileName.txt> only files ending with .txt are allowed")
			return
		}
	
	

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
					if !art_switch {
						fmt.Print(fileData[startingIndex+i])
					} else {
						ascii_art += fileData[startingIndex+i] // printing the character line by line
					}
				}
			}
			if !art_switch {
				fmt.Println()
			} else {
				ascii_art += "\n" // printing a new line after printing each line of the charcter
			}
		}
	}

	os.WriteFile(banner, []byte(ascii_art), 0o644)
}
