package bio

import "testing"

type testpair struct {
	values string
	content float64
}

var tests = []testpair{
	{"GGCCAATT", 0.5},
	{"GAAT", 0.25},
	{"AATTAAT", 0},
}

func TestCountNuc(t *testing.T) {
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
