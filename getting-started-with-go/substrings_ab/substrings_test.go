package substringsab

import "testing"

func TestNumberOfSubstrings(t *testing.T) {
	input := "ccbaacbabbcbab"
	res := NumberOfSubstrings(&input)

	if res != 15 {
		t.Errorf("Expected 15, got %d", res)
	}
}
