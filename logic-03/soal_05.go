package logic_03

import "github.com/zaadinda/day7-amartha/utils"

func Soal05(n int) (result [][]int) {
	result = utils.CreateSlice(n) // Inisialisasi matriks kosong

	num := 1
	mid := n / 2

	for row := 0; row < n; row++ {
		if row < mid {
			for col := 0; col <= row; col++ {
				result[row][col] = num
				result[row][n-1-col] = num
				num += 2
			}
		} else {
			for col := 0; col < n-row; col++ {
				result[row][col] = num
				result[row][n-1-col] = num
				num += 2
			}
		}
	}

	return result
}
