// Package main provides a set of tests for ASCII utility functions.

package main

import (
	"reflect"
	"testing"

	ascii "ascii/utilities" // Importing ASCII utilities package.
)

// TestCheckIllegalChar tests the CheckIllegalChar function.
func TestCheckIllegalChar(t *testing.T) {
	// Test cases for CheckIllegalChar function.
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
	// Running tests for CheckIllegalChar.
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

// TestHandleNewLine tests the HandleNewLine function.
func TestHandleNewLine(t *testing.T) {
	// Test cases for HandleNewLine function.
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
	// Running tests for HandleNewLine.
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

// TestStringContain tests the StringContain function.
func TestStringContain(t *testing.T) {
	// Test cases for StringContain function.
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
	// Running tests for StringContain.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ascii.StringContain(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("StringContain() = %v, want %v", got, tt.want)
			}
		})
	}
}
