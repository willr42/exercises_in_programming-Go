package util

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func GetInput(r *bufio.Reader, s string) (string, error) {
	fmt.Print(s)

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

func StrToFloat(s string) (float64, error) {
	x, err := strconv.ParseFloat(s, 32)

	if err != nil {
		return 0, err
	}

	return x, nil
}

func GetIntWithPrompt(r *bufio.Reader, prompt string) (int, error) {
	inpStr, err := GetInput(r, prompt)

	if err != nil {
		return 0, err
	}
	inputInt, err := StrToInt(inpStr)

	if err != nil {
		return 0, err
	}

	return inputInt, nil
}

func GetFloatWithPrompt(r *bufio.Reader, prompt string) (float64, error) {
	inpStr, err := GetInput(r, prompt)

	if err != nil {
		return 0, err
	}
	inputInt, err := StrToFloat(inpStr)

	if err != nil {
		return 0, err
	}

	return inputInt, nil
}
