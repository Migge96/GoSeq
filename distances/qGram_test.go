package distances

import (
	"reflect"
	"testing"
)

func TestQGramDistance(t *testing.T) {

	t.Run("small run on dna with 4 size", func(t *testing.T) {
		seqA := "acaa"
		seqB := "acac"
		alphabet := "dna"

		want := 2
		got := QGramDistance(seqA, seqB, alphabet)

		if got != want {
			t.Errorf("wanted %v got %v , given %v and %v", want, got, seqA, seqB)
		}

	})

	t.Run("small run on proteinsequence with 4 size", func(t *testing.T) {
		seqA := "KQEE"
		seqB := "KQEQ"
		alphabet := "prot"

		want := 2
		got := QGramDistance(seqA, seqB, alphabet)

		if got != want {
			t.Errorf("wanted %v got %v , given %v and %v", want, got, seqA, seqB)
		}

	})

	t.Run("small run on dna with bigger sizes", func(t *testing.T) {
		seqA := "gagccatcattcgctgcttggagtaaggtgcgaatcaggaagctacccggcacaaggcaccgatcgccgggcagcacctgtgacttacaggcggcaccgattatgggcatataaagcggt"
		seqB := "gtgtgtcgtgagaacgcgaaaaggccgcggcgataccgaatacgaggataggtacgatggtcaattaggcacgaagtatggttccggtattgaccctacagcaaaattttgtcaaatcgt"
		alphabet := "dna"

		want := 92
		got := QGramDistance(seqA, seqB, alphabet)

		if got != want {
			t.Errorf("wanted %v got %v , given %v and %v", want, got, seqA, seqB)
		}

	})

}
func TestQGramProfile(t *testing.T) {

	seqA := "ACAA"
	q := 2
	alphabet := "dna"

	want := []int{1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	got := qGramProfile(seqA, q, alphabet)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("wanted %v got %v , given %v", want, got, seqA)
	}
}

func TestRanking(t *testing.T) {
	seqA := "ATACG"
	q := 5
	alphabet := "dna"

	want := 588
	got := ranking(seqA, q, alphabet)

	if got != want {
		t.Errorf("wanted %d, got %d", want, got)
	}
}

func TestPow(t *testing.T) {
	t.Run("Compute 2 ^ 3", func(t *testing.T) {
		want := 8
		got := pow(2, 3)

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})

	t.Run("Compute 5 ^ 0", func(t *testing.T) {
		want := 1
		got := pow(5, 0)

		if got != want {
			t.Errorf("wanted %d, got %d", want, got)
		}
	})
}

func BenchmarkQGramDistance(b *testing.B) {
	seqA := "gagccatcattcgctgcttggagtaaggtgcgaatcaggaagctacccggcacaaggcaccgatcgccgggcagcacctgtgacttacaggcggcaccgattatgggcatataaagcggt"
	seqB := "gtgtgtcgtgagaacgcgaaaaggccgcggcgataccgaatacgaggataggtacgatggtcaattaggcacgaagtatggttccggtattgaccctacagcaaaattttgtcaaatcgt"
	alphabet := "dna"
	b.ReportAllocs()
	// String is often encoded as 8 Byte or 16 Byte. With 2 String as Input we take 32 Bytes Per Operation
	b.SetBytes(32)

	for i := 0; i < b.N; i++ {
		QGramDistance(seqA, seqB, alphabet)
	}
}
