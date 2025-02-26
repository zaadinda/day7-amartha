package logic_03

import (
	"github.com/zaadinda/day7-amartha/utils"
)

func Soal02(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	num := 1

	for row := 0; row < n; row++ {
		if row%2 == 0 {
			for col := row; col < n; col++ {
				result[row][col] = num
				num += 2
			}
		} else {
			for col := n - 1; col >= row; col-- {
				result[row][col] = num
				num += 2
			}
		}
	}
	return result
}
