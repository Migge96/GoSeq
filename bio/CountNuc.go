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

func hammingDistance(x, y string) int {
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
				diagonal := float64(editMatrix[row-1][col-1] + hammingDistance(string(x[col]), string(y[row])))
				horizontal := float64(editMatrix[row-1][col] + 1)
				vertical := float64(editMatrix[row][col-1] + 1)
				editMatrix[row][col] = int(math.Min(diagonal, math.Min(horizontal, vertical)))
			}
		}
	}
	return editMatrix[len(y)-1][len(x)-1]
}

func editSequence(x, y string) string {
	x = "#" + x
	y = "#" + y
	editMatrix := make([][]int, len(y))
	opMatrix := make([][]int8, len(y))
	var rowIndex int
	var colIndex int
	result := ""
	var CS, D, I, CS_D, CS_I, D_I, CS_D_I int8 = 1,2,3,4,5,6,7
	for k := 0; k < len(y); k++ {
		editMatrix[k] = make([]int, len(x))
		opMatrix[k] = make([]int8, len(x))
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
				hamming := hammingDistance(string(x[col]), string(y[row]))
				diagonal := float64(editMatrix[row-1][col-1] + hamming)
				horizontal := float64(editMatrix[row-1][col] + 1)
				vertical := float64(editMatrix[row][col-1] + 1)
				editMatrix[row][col] = int(math.Min(diagonal, math.Min(horizontal, vertical)))
				if diagonal <= math.Min(horizontal, vertical) {
					opMatrix[row][col] = CS
					if diagonal == horizontal {
						opMatrix[row][col] = CS_D
					}
					if diagonal == vertical {
						opMatrix[row][col] = CS_I
					}
				}
				if horizontal <= vertical {
					opMatrix[row][col] = D
					if horizontal == vertical {
						opMatrix[row][col] = D_I
					}
				}
				if vertical < horizontal {
					opMatrix[row][col] = I
				}
				if (diagonal == horizontal) && (diagonal == vertical) {
					opMatrix[row][col] = CS_D_I
				}
			
			}
			colIndex = col
		}
		rowIndex = row
	}
	
	for rowIndex != 0 && colIndex != 0 {
		switch opMatrix[rowIndex][colIndex] {
		case CS:
			result += "C/S"
			rowIndex --
			colIndex --
		case I:
			result += "I"
			colIndex --
		case D:
			result += "D"
			rowIndex --
		case CS_D:
			result += "C/S"
			rowIndex --
			colIndex --
		case CS_D_I:
			result += "C/S"
			rowIndex --
			colIndex --
		case CS_I:
			result += "C/S"
			rowIndex --
			colIndex --
		case D_I:
			result += "D"
			rowIndex --
		}
		
	}
	return result
}
