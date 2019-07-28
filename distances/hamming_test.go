package distances

import "testing"

func TestHamming(t *testing.T) {

	t.Run("Two Strings of equal length", func(t *testing.T) {
		//Test Cases
		seqA := "CAG"
		seqB := "CAA"

		want := 1
		got, _ := hamming(&seqA, &seqB)

		if got != want {
			t.Errorf("Got %d, wanted %d", got, want)
		}
	})

	t.Run("unequal lengths", func(t *testing.T) {
		seqA := "CAG"
		seqB := "CA"

		want := "Error: Unequal Length of strings"
		_, got := hamming(&seqA, &seqB)

		if got == nil {
			t.Error("want an error for unequal length")
		}

		if got.Error() != want {
			t.Errorf("got %q, wanted %q", got, want)
		}

	})
}
