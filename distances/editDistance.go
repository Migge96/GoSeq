package distances

// EditDistance computes the edit distance betwwen two strings a and b.
func EditDistance(seqA, seqB string) int {
	lenA := (len(seqA) + 1)
	lenB := (len(seqB) + 1)
	// Initiliaze the Edit Matrix
	editMatrix := make([][]int, lenA)
	for i := 0; i < lenA; i++ {
		editMatrix[i] = make([]int, lenB)
	}

	for i := 0; i < lenA; i++ {
		editMatrix[i][0] = i

	}

	for j := 0; j < lenB; j++ {
		editMatrix[0][j] = j

	}

	for i := 1; i < lenA; i++ {
		for j := 1; j < lenB; j++ {
			vertical := editMatrix[i-1][j] + 1
			horizontal := editMatrix[i][j-1] + 1
			diagonal := editMatrix[i-1][j-1]

			if seqA[i-1] != seqB[j-1] {
				diagonal++
			}

			editMatrix[i][j] = minOf3(vertical, horizontal, diagonal)

		}
	}

	return editMatrix[lenA-1][lenB-1]
}

func minOf3(a, b, c int) int {
	min := a

	if b < min {
		min = b
	}

	if c < min {
		min = c
	}

	return min
}
