package ascii

import (
	"errors"
	"strings"
)

func Checkflag(arr []string) ([2]string, error) {
	if len(arr) != 3 {
		return [2]string{}, errors.New("usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard")
	}
	if !strings.HasPrefix(arr[0], "--output=") {
		return [2]string{}, errors.New("flag usage: expected --output=<filename.txt>")
	}

	filename := strings.TrimPrefix(arr[0], "--output=")
	// check filename
	if !strings.HasSuffix(filename, ".txt") {
		return [2]string{}, errors.New("only text files are allowed i.e filename")
	}

	if filename == "standard.txt" || filename == "thinkertoy.txt" || filename == "shadow.txt" {
		return [2]string{}, errors.New("error: standard.txt and thinkertoy.txt and shadow.txt are disallowed. use other filename instead")
	}
	if arr[2] != "standard" && arr[2] != "thinkertoy" && arr[2] != "shadow" {
		return [2]string{}, errors.New("expected either standard | thinkertoy | shadow")
	}
	return [2]string{filename, arr[2]}, nil
}
