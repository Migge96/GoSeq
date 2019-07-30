package distances

import "math"

func qGramProfile(sequence string) []int {
	return []int{0, 0, 0}
}

func ranking(qgram string, q, alphabetSize int) int {

	rank := 0

	for i, char := range qgram {
		temp := string(char)

		switch temp {
		case "A":
			rank += 0 * int(math.Pow(float64(alphabetSize), float64(i)))
		case "C":
			rank += 1 * int(math.Pow(float64(alphabetSize), float64(i)))
		case "G":
			rank += 2 * int(math.Pow(float64(alphabetSize), float64(i)))
		case "T":
			rank += 3 * int(math.Pow(float64(alphabetSize), float64(i)))
		}

	}

	return rank
}
