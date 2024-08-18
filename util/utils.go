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

func StrToInt(s string) (int, error) {
	x, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return x, nil
}
