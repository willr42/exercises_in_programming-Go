package util

import (
	"bufio"
	"errors"
	"strings"
)

func GetInput(r *bufio.Reader) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", errors.New("couldn't read input")
	}
	finalLine := strings.Trim(line, "\n")
	return finalLine, nil
}
