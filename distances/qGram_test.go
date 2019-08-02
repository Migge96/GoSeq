package distances

import (
	"reflect"
	"testing"
)

func TestQGramDistance(t *testing.T) {

	t.Run("small run on dna with 4 size", func(t *testing.T) {
		seqA := "acaa"
		seqB := "acac"
		alphabetSize := 4

		want := 2
		got := qGramDistance(seqA, seqB, alphabetSize)

		if got != want {
			t.Errorf("wanted %v got %v , given %v and %v", want, got, seqA, seqB)
		}

	})

	t.Run("small run on dna with bigger sizes", func(t *testing.T) {
		seqA := "gagccatcattcgctgcttggagtaaggtgcgaatcaggaagctacccggcacaaggcaccgatcgccgggcagcacctgtgacttacaggcggcaccgattatgggcatataaagcggt"
		seqB := "gtgtgtcgtgagaacgcgaaaaggccgcggcgataccgaatacgaggataggtacgatggtcaattaggcacgaagtatggttccggtattgaccctacagcaaaattttgtcaaatcgt"
		alphabetSize := 4

		want := 92
		got := qGramDistance(seqA, seqB, alphabetSize)

		if got != want {
			t.Errorf("wanted %v got %v , given %v and %v", want, got, seqA, seqB)
		}

	})

}
func TestQGramProfile(t *testing.T) {

	seqA := "ACAA"
	q := 2
	alphabetSize := 4

	want := []int{1, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	got := qGramProfile(seqA, q, alphabetSize)

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

func BenchmarkQGramDistance(b *testing.B) {
	seqA := "gagccatcattcgctgcttggagtaaggtgcgaatcaggaagctacccggcacaaggcaccgatcgccgggcagcacctgtgacttacaggcggcaccgattatgggcatataaagcggt"
	seqB := "gtgtgtcgtgagaacgcgaaaaggccgcggcgataccgaatacgaggataggtacgatggtcaattaggcacgaagtatggttccggtattgaccctacagcaaaattttgtcaaatcgt"

	b.ReportAllocs()
	// String is often encoded as 8 Byte or 16 Byte. With 2 String as Input we take 32 Bytes Per Operation
	b.SetBytes(32)

	for i := 0; i < b.N; i++ {
		qGramDistance(seqA, seqB, 4)
	}
}
