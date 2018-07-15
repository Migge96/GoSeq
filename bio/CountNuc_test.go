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

var testsEdit = []testPairEditDist{
	{ "DBADAD", "BCACD",
	 4,
	},
}

func TestEditDist(t *testing.T) {
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