package util_test

import (
	"bufio"
	"bytes"
	"errors"
	"exercises-in-programming/util"
	"math"
	"testing"
)

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestGetInput(t *testing.T) {
	t.Parallel()

	input := "test input\n"
	want := "test input"

	r := bufio.NewReader(bytes.NewBufferString(input))

	got, err := util.GetInput(r, "Prompt here")

	if err != nil {
		t.Fatalf("GetInput() returned an error: %v", err)
	}

	if got != want {
		t.Errorf("GetInput() = %v; want %v", got, want)
	}
}

func TestStrToInt(t *testing.T) {
	t.Parallel()

	input := "1234"
	want := 1234

	got, err := util.StrToInt(input)

	if err != nil {
		t.Fatalf("StrToInt(%s) returned an error: %v", input, err)
	}

	if got != want {
		t.Errorf("StrToInt(%s) = %v; want %v", input, got, want)
	}
}

func TestStrToFloat(t *testing.T) {
	t.Parallel()

	input := "12.34"
	want := 12.34

	got, err := util.StrToFloat(input)

	if err != nil {
		t.Fatalf("StrToFloat(%s) returned an error: %v", input, err)
	}

	if !closeEnough(got, want, 0.001) {
		t.Errorf("StrToFloat(%s) = %v; want %v", input, got, want)
	}
}

func TestStrToCents(t *testing.T) {
	t.Parallel()

	type testCase struct {
		input string
		want  int
	}

	testCases := []testCase{
		{input: "12.34", want: 1234},
		{input: "12", want: 1200},
	}

	for _, test := range testCases {
		got, err := util.StrToCents(test.input)

		if err != nil {
			t.Fatalf("StrToCents(%s) returned an error: %v", test.input, err)
		}

		if got != test.want {
			t.Errorf("StrToCents(%s) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestStrToCentsInvalid(t *testing.T) {
	t.Parallel()
	type testCase struct {
		input string
		want  error
	}

	testCases := []testCase{
		{input: "A sentence of some length", want: util.ErrInvalidNumber},
		{input: "12.3.4", want: util.ErrInvalidFormat},
	}

	for _, test := range testCases {
		d, err := util.StrToCents(test.input)

		if err == nil {
			t.Fatalf("StrToCents(%d) should error", d)
		}

		if !errors.Is(err, test.want) {
			t.Errorf("StrToCents expected error %v, got %v", test.want, err)
		}
	}
}
