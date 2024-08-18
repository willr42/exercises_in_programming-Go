package util

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

func GetInput(r *bufio.Reader, s string) (string, error) {
	line, err := r.ReadString('\n')
	if err != nil {
		return "", errors.New("couldn't read input")
	}
	finalLine := strings.Trim(line, "\n")
	return finalLine, nil
}
