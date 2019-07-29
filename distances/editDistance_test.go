package distances

import "testing"

func TestEditDistance(t *testing.T) {
	t.Run("Simple DNA Example with len 5 and 6", func(t *testing.T) {
		//Test Cases
		seqA := "BCACD"
		seqB := "DBADAD"

		want := 4
		got := EditDistance(seqA, seqB)

		if want != got {
			t.Errorf("got %d, wanted %d, given %q and %q", got, want, seqA, seqB)
		}

	})

	t.Run("One Sequence Empty Should be only insertions", func(t *testing.T) {
		//Test Cases
		seqA := "AAAAA"
		seqB := ""

		want := 5
		got := EditDistance(seqA, seqB)

		if want != got {
			t.Errorf("got %d, wanted %d, given %q and %q", got, want, seqA, seqB)
		}
	})

	t.Run("Simple english Sentence with Typo", func(t *testing.T) {
		//Test Cases
		seqA := "Hello, my name is Lukas."
		seqB := "Hallo, my name is Lukas."

		want := 1
		got := EditDistance(seqA, seqB)

		if want != got {
			t.Errorf("got %d, wanted %d, given %q and %q", got, want, seqA, seqB)
		}
	})

}
