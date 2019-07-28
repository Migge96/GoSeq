package distances

import "testing"

func TestHamming(t *testing.T) {
	//Test Cases
	seqA := "Hallo"
	seqB := "Hello"

	want := 1
	got := hamming(seqA, seqB)

	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
