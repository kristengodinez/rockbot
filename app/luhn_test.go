package main

import "testing"

func TestLuhn(t *testing.T) {
	// Valid numbers
	// 3379 5135 6110 8795
	// 2769 1483 0405 9987
	t.Run("validating valid numbers for Luhn algorithm", func(t *testing.T) {
		got := Luhn("3379 5135 6110 8795")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}

	})
	// Invalid numbers
	// 3379 5135 6110 8794
	// 2769 1483 0405 9986
}
