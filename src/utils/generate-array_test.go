package utils

import "testing"

func TestMakeBigArray(t *testing.T) {
	got := MakeBigArray(MakeBigArrayParams{NumberOfElements: 100, To: 10})
	if len(got) != 100 {
		t.Errorf("MakeBigArray: array length mismatch")
	}
}
