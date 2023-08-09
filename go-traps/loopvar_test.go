package main

import "testing"

// START OMIT
func TestAllEvenBuggy(t *testing.T) {
	testCases := []int{1, 2, 4, 6}
	for _, tc := range testCases {
		t.Run("sub", func(t *testing.T) {
			t.Parallel()
			if tc%2 != 0 {
				t.Fatal("odd tc", tc)
			}
		})
	}
}

// STOP OMIT
