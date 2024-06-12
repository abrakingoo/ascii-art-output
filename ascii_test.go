package main

import (
	"os"
	"os/exec"
	"reflect"
	"testing"

	ascii "ascii/utilities"
)

func TestCheckIllegalChar(t *testing.T) {
	tests := []struct {
		name    string
		input   []string
		want    []string
		wantErr bool
	}{
		{
			name:    "Valid ASCII",
			input:   []string{"Hello", "World"},
			want:    []string{"Hello", "World"},
			wantErr: false,
		},
		{
			name:    "Contains illegal character",
			input:   []string{"Hello", "W\x1fold", "こんにちは、アーロンさん"},
			want:    []string{},
			wantErr: true,
		},
		{
			name:    "Empty string",
			input:   []string{""},
			want:    []string{""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ascii.CheckIllegalChar(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckIllegalChar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CheckIllegalChar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleNewLine(t *testing.T) {

	tests := []struct {
		name    string
		input   string
		want    []string
		wantErr bool
	}{
		{
			name:    "No newlines",
			input:   "Hello World",
			want:    []string{"Hello World"},
			wantErr: false,
		},
		{
			name:    "Single newline",
			input:   "Hello\\nWorld",
			want:    []string{"Hello", "World"},
			wantErr: false,
		},
		{
			name:    "Multiple newlines",
			input:   "Hello\\nWorld\\nfrom\\nGo",
			want:    []string{"Hello", "World", "from", "Go"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ascii.HandleNewLine(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleNewLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleNewLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringContain(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "Contains 'o'",
			s:    "Hello\r\nWorld",
			want: []string{"Hello", "World"},
		},
		{
			name: "Does not contain 'o'",
			s:    "Hi\nThere",
			want: []string{"Hi", "There"},
		},
		{
			name: "Empty string",
			s:    "",
			want: []string{""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ascii.StringContain(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringContain() = %v, want %v", got, tt.want)
			}
		})
	}
}
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
	testStrings := [4]Args{
		{"hello", ""},
		{"hello", "standard"},
		{"hello", "thinkertoy"},
		{"hello", "shadow"},
	}
	// Define the commands to execute here:
	commands := [4]*exec.Cmd{
		exec.Command("go", "run", ".", "--output=test01.txt", "hello", "standard"),
		exec.Command("go", "run", ".", "--output=test02.txt", "hello", "thinkertoy"),
		exec.Command("go", "run", ".", "--output=test03.txt", "hello", "shadow"),
		exec.Command("go", "run", ".", "--output=test04.txt", "hello", "standard"),
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
