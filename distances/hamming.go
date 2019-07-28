package distances

import "errors"

func hamming(seqA, seqB *string) (int, error) {
	if len(*seqA) != len(*seqB) {
		return 0, errors.New("Error: Unequal Length of strings")
	}

	difference := 0
	for i, val := range *seqA {
		if string(val) != string((*seqB)[i]) {
			difference++
		}
	}
	return difference, nil
}
