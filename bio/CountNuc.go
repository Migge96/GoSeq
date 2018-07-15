/*
Various implementations of bioinformat sequence analysis algorithmn.
 */

package bio

import (
	"strings"
	"math"
)

var A, T, G, C int

func countNuc(param string) {
	A = strings.Count(param, "A")
	T = strings.Count(param, "T")
	G = strings.Count(param, "G")
	C = strings.Count(param, "C")

}

func gcContent(param string) float64 {
	countNuc(param)
	content := float64(G+C) / float64(A+C+G+T)
	return content
}

func hamiltonDistance(x, y string) int {
	if x == y {
		return 0
	} else {
		return 1
	}
}

func editDistance(x, y string) int {
	x = "#" + x
	y = "#" + y
	editMatrix := make([][]int, len(y))
	for k := 0; k < len(y); k++ {
		editMatrix[k] = make([]int, len(x))
	}
	for row, h := range editMatrix {
		for col := range h {
			if row == 0 {
				editMatrix[row][col] = col
			}
			if col == 0 {
				editMatrix[row][col] = row
			}
			if col != 0 && row != 0 {
				diagonal := float64(editMatrix[row-1][col-1] + hamiltonDistance(string(x[col]), string(y[row])))
				horizontal := float64(editMatrix[row-1][col] + 1)
				vertical := float64(editMatrix[row][col-1] + 1)
				editMatrix[row][col] = int(math.Min(diagonal, math.Min(horizontal, vertical)))
			}
		}
	}
	return editMatrix[len(y)-1][len(x)-1]
}
