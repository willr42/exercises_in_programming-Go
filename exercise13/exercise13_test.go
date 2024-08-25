package exercise13_test

import (
	"exercises-in-programming/exercise13"
	"testing"
)

func TestCompoundInterest(t *testing.T) {
	t.Parallel()

	principal := 150000
	rawRate := 4.3
	rate := rawRate / 100
	years := 6
	compoundOccurrence := 4

	got := exercise13.CompoundInterest(principal, rate, years, compoundOccurrence)
	want := 193884

	if got != want {
		t.Errorf("CompoundInterest(%d, %.1f, %d, %d) = %v; want %v", principal, rawRate, years, compoundOccurrence, got, want)
	}
}
