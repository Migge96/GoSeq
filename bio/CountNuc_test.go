package bio

import "testing"

// test for GC_Content

type testpair struct {
	values string
	content float64
}

var tests = []testpair{
	{"GGCCAATT", 0.5},
	{"GAAT", 0.25},
	{"AATTAAT", 0},
}


func TestGcContent(t *testing.T) {
	for _, pair := range tests {
		v := gcContent(pair.values)
		if v != pair.content {
			t.Error(
				"For", pair.values,
				"expected", pair.content,
				"got", v,
			)
		}

	}
}

type testPairEditDist struct {
	seqX string
	seqY string
	dist int
}

type testPairEditSeq struct {
	seqX string
	seqY string
	editSeq string
}

var testsEdit = []testPairEditDist{
	{ "DBADAD", "BCACD",
	 4,},
	 {"DBADAD", "", 6 },
	 {"", "BCACD", 5},
	 {"", "", 0},
	 {"HALLO", "HALLO", 0},
	 {"AGCAGTTCCTTGACCATGGAGCTGCCTTGGACAGCACTCTTCCTCAGTACCTTCCTCCTGGGTCTTTCAT",
	 "ACATACCTCCTAACAGTTCCTAGAAAATGGAGCTGTCTTGGCATGTAGTCTTTATTGCCCTGCTAAGTTT", 36},
}


var testsEditSeq = []testPairEditSeq{
	{ "DBADAD", "BCACD",
		"ICCI",},
	{"DBADAD", "", "IC" },
	{"", "BCACD", "IIII"},
	{"", "", ""},
	{"HALLO", "HALLO", ""},
}

func TestEditDistance(t *testing.T) {
	for _,pair := range testsEdit {
		v := editDistance(pair.seqX, pair.seqY)
		if v != pair.dist {
			t.Error(
				"For x: ", pair.seqX,
				"for y: ", pair.seqY,
				"expected", pair.dist,
				"got", v,
			)
		}
	}
}

func TestEditSequence(t *testing.T) {
	for _,pair := range testsEditSeq {
		v := editSequence(pair.seqX, pair.seqY)
		if v != pair.editSeq {
			t.Error(
				"For x: ", pair.seqX,
				"for y: ", pair.seqY,
				"expected", pair.editSeq,
				"got", v,
			)
		}
	}
}

func benchmarkDist(x, y string , b *testing.B) {
	for n := 0; n < b.N; n++ {
		editDistance(x, y)
	}
}

func BenchmarkDist1(b *testing.B)  {benchmarkDist(testsEdit[0].seqX, testsEdit[0].seqY, b)}
func BenchmarkDist3(b *testing.B)  {benchmarkDist(testsEdit[3].seqX, testsEdit[3].seqY, b)}
func BenchmarkDist5(b *testing.B)  {benchmarkDist(testsEdit[5].seqX, testsEdit[5].seqY, b)}