package distances

import "testing"

func TestMaxMatches(t *testing.T) {
	seqA := "cbaabdcb"
	seqB := "abcba"

	want := 2
	got := MaxMatches(seqA, seqB)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
