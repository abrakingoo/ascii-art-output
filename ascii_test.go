package main

import (
	"os"
	"os/exec"
	"testing"

	ascii "ascii/utilities"
)

func TestExecCommand(t *testing.T) {
	type Args struct {
		str  string
		file string
	}

	fileString := ""

	// Create files
	expected, err := os.Create("expected.txt")
	if err != nil {
		t.Errorf("Error :%v", err)
	}
	got, err := os.Create("got.txt")
	if err != nil {
		t.Errorf("Error :%v", err)
	}
	// close the files
	defer expected.Close()
	defer got.Close()

	// Modify here : Input your test input string
	testStrings := []Args{
		{"hello", ""},
		{"hello", "standard"},
		{"hello", "thinkertoy"},
		{"hello", "shadow"},
	}
	// Define the commands to execute here:
	commands := []*exec.Cmd{
		exec.Command("go", "run", ".", "hello"),
		exec.Command("go", "run", ".", "hello", "standard"),
		exec.Command("go", "run", ".", "hello", "thinkertoy"),
		exec.Command("go", "run", ".", "hello", "shadow"),
	}

	//
	for i, val := range testStrings {
		// file to use
		fileData := []string{}
		if val.file == "" {
			file, err := os.ReadFile("standard.txt")
			if err != nil {
				t.Error(err)
				os.Exit(1)
			}
			fileData = append(fileData, ascii.StringContain(string(file))...)

		} else {
			filename := val.file + ".txt"
			file, err := os.ReadFile(filename)
			if err != nil {
				t.Error(err)
				os.Exit(1)
			}
			fileData = append(fileData, ascii.StringContain(string(file))...)
		}

		inputParts, err := ascii.HandleNewLine(testStrings[i].str)
		if err != nil {
			t.Error(err)
		}
		legalChar, err := ascii.CheckIllegalChar(inputParts) // Check for illegal characters in the string
		if err != nil {
			t.Error(err)
		}
		for _, part := range legalChar {
			if part == "" {
				fileString += "\n"
				continue
			}
			for i := 0; i < 8; i++ {
				for _, char := range part {
					startingIndex := ascii.GetStartingIndex(int(char))
					fileString += fileData[startingIndex+i]
				}
				fileString += "\n"
			}

		}
		fileString += "\n"
	}
	expected.WriteString(fileString)

	// Execute each command and store it in a string
	for _, cmd := range commands {
		// commandline output
		output, err := cmd.Output()
		if err != nil {
			t.Error(err)
		}
		got.WriteString(string(output) + "\n")
		//
	}

	// Read the output (1) of the commands from a file
	exTest, err := os.ReadFile("got.txt")
	if err != nil {
		t.Error(err)
	}
	// Read the output (2) of the commands from a file
	gotTest, err := os.ReadFile("expected.txt")
	if err != nil {
		t.Error(err)
	}

	// Match the strings of the file
	if string(exTest) != string(gotTest) {
		t.Errorf("Output does not match. Expected \n%v\n with length %v got \n%v\n with length %v", string(exTest), len(string(exTest)), (string(gotTest)), len(string(gotTest)))
	}
}
