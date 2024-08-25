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
	return strings.TrimSpace(line), nil
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

func GetCentsWithPrompt(r *bufio.Reader, prompt string) (int, error) {
	inpStr, err := GetInput(r, prompt)

	if err != nil {
		return 0, err
	}

	cents, err := StrToCents(inpStr)

	if err != nil {
		return 0, err
	}

	return cents, nil
}

var (
	ErrInvalidFormat = fmt.Errorf("must be in whole dollars or dollar.cents fmt")
	ErrInvalidNumber = fmt.Errorf("couldn't convert dollars to int")
)

func StrToCents(s string) (int, error) {
	parts := strings.Split(s, ".")
	if len(parts) > 2 {
		return 0, ErrInvalidFormat
	}

	dollars := 0
	cents := 0
	var err error

	if len(parts) > 0 {
		dollars, err = strconv.Atoi(parts[0])
		if err != nil {
			return 0, ErrInvalidNumber
		}
	}

	// Cents exists
	if len(parts) == 2 {
		if len(parts[1]) > 2 {
			parts[1] = parts[1][:2]
		}
		cents, err = strconv.Atoi(parts[1])
		if err != nil {

			return 0, ErrInvalidNumber
		}

	}
	return dollars*100 + cents, nil
}

func FormatCents(cents int) string {
	dollars := cents / 100
	remainingCents := cents % 100
	return fmt.Sprintf("%d.%02d", dollars, remainingCents)
}
