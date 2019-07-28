package distances

func hamming(seqA, seqB string) int {
	difference := 0
	for i, val := range seqA {
		if string(val) != string(seqB[i]) {
			difference++
		}
	}
	return difference
}
