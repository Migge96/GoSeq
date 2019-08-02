package distances

import (
	"math"
	"strings"
)

func qGramDistance(seqA, seqB string, alphabetSize int) int {
	q := int(math.Log(float64(len(seqA))) / math.Log(float64(alphabetSize)))

	if q == 1 {
		q = 2
	}
	profile := make([]int, int(math.Pow(float64(alphabetSize), float64(q))))
	result := 0
	for i := 0; i < len(seqA)-q+1; i++ {
		temp := seqA[i : i+q]
		profile[ranking(temp, q, alphabetSize)]++
	}
	for i := 0; i < len(seqB)-q+1; i++ {
		temp := seqB[i : i+q]
		profile[ranking(temp, q, alphabetSize)]--
	}

	for _, val := range profile {
		result += abs(val)
	}
	return result
}

func qGramProfile(sequence string, q, alphabetSize int) []int {
	profile := make([]int, int(math.Pow(float64(alphabetSize), float64(q)))-1)
	for i := 0; i < len(sequence)-1; i++ {
		temp := sequence[i : i+q]
		profile[ranking(temp, q, alphabetSize)]++

	}

	return profile
}

func ranking(qgram string, q, alphabetSize int) int {

	rank := 0

	for i, char := range qgram {
		temp := strings.ToLower(string(char))

		switch temp {
		case "a":
			rank += 0 * int(math.Pow(float64(alphabetSize), float64(i)))
		case "c":
			rank += 1 * int(math.Pow(float64(alphabetSize), float64(i)))
		case "g":
			rank += 2 * int(math.Pow(float64(alphabetSize), float64(i)))
		case "t":
			rank += 3 * int(math.Pow(float64(alphabetSize), float64(i)))
		}

	}

	return rank
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
