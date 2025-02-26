package logic_03

import "github.com/zaadinda/day7-amartha/utils"

func Soal10(n int) (result [][]int) {
	result = utils.CreateSlice(n)
	mid := (n - 1) / 2

	for j := 0; j <= mid; j++ {
		num := 9
		for i := 0; i <= mid; i++ {
			if i+j >= mid {
				result[j][i] = num
				result[n-1-i][j] = num
				result[i][n-1-j] = num
				result[n-1-i][n-1-j] = num
				num -= 2
			}
		}
	}
	return result
}
