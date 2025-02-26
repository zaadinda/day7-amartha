package logic_03

import "github.com/zaadinda/day7-amartha/utils"

func Soal01(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	num := 1

	for row := 1; row <= n; row++ {
		if row%2 != 0 {
			for col := 0; col < row; col++ {
				result[row-1][col] = num
				num += 2
			}
		} else {
			for col := row - 1; col >= 0; col-- {
				result[row-1][col] = num
				num += 2
			}
		}
	}
	return result
}
