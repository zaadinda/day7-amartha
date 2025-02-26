package logic_03

import "github.com/zaadinda/day7-amartha/utils"

func Soal09(n int) (result [][]int) {
	result = utils.CreateSlice(n)
	mid := (n - 1) / 2

	for col := 0; col <= mid; col++ {
		num := 1
		for row := 0; row <= mid; row++ {
			if row+col >= mid {
				result[col][row] = num
				result[n-1-row][col] = num
				result[row][n-1-col] = num
				result[n-1-row][n-1-col] = num
				num += 2
			}
		}
	}
	return result
}
