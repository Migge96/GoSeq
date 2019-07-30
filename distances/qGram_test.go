package distances

import (
	"reflect"
	"testing"
)

func TestQGramProfile(t *testing.T) {

	seqA := "acaa"

	want := []int{1, 1, 1, 0}
	got := qGramProfile(seqA)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v got %v , given %v", want, got, seqA)
	}
}

func TestRanking(t *testing.T) {
	seqA := "ATACG"
	q := 5
	alphabetSize := 4

	want := 588
	got := ranking(seqA, q, alphabetSize)

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}
