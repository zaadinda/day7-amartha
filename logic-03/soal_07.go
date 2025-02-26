package logic_03

import "github.com/zaadinda/day7-amartha/utils"

func Soal07(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	mid := (n - 1) / 2
	num := 1

	// isi slice
	for col := 0; col < n; col++ {
		for row := 0; row < mid; row++ {
			// 3,7
			if row+col >= mid && col-row <= mid {
				if row%2 == 1 {
					result[row][col] = num
					result[n-1-row][col] = num
				} else {
					result[row][col] = num
					result[n-1-row][n-1-col] = num
				}
				num += 2
			}
		}
	}
	return result
}
