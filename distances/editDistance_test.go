package distances

import "testing"

func TestEditDistance(t *testing.T) {
	//Test Cases
	seqA := "CAAG"
	seqB := "CAGG"

	want := 1
	got := EditDistance(seqA, seqB)

	if want != got {
		t.Errorf("got %d, wanted %d, given %q and %q", got, want, seqA, seqB)
	}
}
