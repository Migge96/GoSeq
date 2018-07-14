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

func editDistance(x, y string) int  {
	var editMatrix [len(x) + 1][len(y) + 1]int
	for j := 0; j< len(y) + 1; j ++  {
		for i := 0; i < len(x) + 1; i ++ {
			if j == 0 {
				editMatrix[j][i] = i
			}
			if i == 0 {
				editMatrix[j][i] = j
			}
			if i != 0 && j != 0 {
				diagonal := editMatrix[i - 1][j - 1] + hamiltonDistance(string(x[i]), string(y[j]))
				horizontal := editMatrix[i - 1][j] + 1
				vertical := editMatrix[i][j - 1] + 1
				 editMatrix[i][j] = int(math.Min(float64(diagonal), math.Min(float64(horizontal), float64(vertical))))

			}
		}
	}
	result := editMatrix[len(x) + 1][len(y) + 1]
	return result
}

/*
func main(){
    countNuc("AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGCTTCTGAACTG")
    fmt.Println("A:",a, "T:",t, "G:",g, "C:",c)

}

