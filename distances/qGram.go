package distances

import (
	"math"
	"unicode"
)

const dna = "acgt"
const rna = "acgu"
const prot = "galmfwkqespvicyhrndt"

func selectAlphabet(short string, alphabet *string) {
	switch short {
	case "dna":
		*alphabet = dna
	case "rna":
		*alphabet = rna
	case "prot":
		*alphabet = prot
	}
}

//QGramDistance computes the qGram Distance between seqA, and seqB
func QGramDistance(seqA, seqB, alphabetShort string) int {
	var alphabet string
	selectAlphabet(alphabetShort, &alphabet)
	alphabetSize := len(alphabet)

	q := int(math.Log2(float64(len(seqA))) / math.Log2(float64(alphabetSize)))

	if q <= 1 {
		q = 2
	}
	profile := make([]int, pow(alphabetSize, q))
	result := 0
	for i := 0; i < len(seqA)-q+1; i++ {
		temp := seqA[i : i+q]
		profile[ranking(temp, q, alphabetShort)]++
	}

	for i := 0; i < len(seqB)-q+1; i++ {
		temp := seqB[i : i+q]
		profile[ranking(temp, q, alphabetShort)]--
	}

	for _, val := range profile {
		result += abs(val)
	}
	return result
}

func qGramProfile(sequence string, q int, alphabetShort string) []int {
	var alphabet string
	selectAlphabet(alphabetShort, &alphabet)
	alphabetSize := len(alphabet)

	profile := make([]int, pow(alphabetSize, q))
	for i := 0; i < len(sequence)-1; i++ {
		temp := sequence[i : i+q]
		profile[ranking(temp, q, alphabetShort)]++

	}

	return profile
}

func ranking(qgram string, q int, alphabetShort string) int {
	var alphabet string
	selectAlphabet(alphabetShort, &alphabet)

	var alphabetMap = make(map[rune]int)
	for i, char := range alphabet {
		alphabetMap[char] = i
	}
	alphabetSize := len(alphabet)

	rank := 0

	for i, char := range qgram {
		temp := unicode.ToLower(char)

		rank += alphabetMap[temp] * pow(alphabetSize, i)

	}

	return rank
}

func pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
