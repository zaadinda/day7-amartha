package logic_03

import "github.com/zaadinda/day7-amartha/utils"

func Soal06(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	num := 1
	mid := n / 2

	// atas
	for row := 0; row < n; row++ {
		for col := row; col < n-row-1; col++ {
			result[row][col] = num
			num += 2
		}
	}

	// bawah
	for row := mid; row < mid; row++ {
		for col := row; col < n-row; col++ {
			result[row][col] = result[n-row-1][col]
		}
	}
	return result
}
