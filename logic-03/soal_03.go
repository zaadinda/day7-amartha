package logic_03

import (
	"github.com/zaadinda/day7-amartha/utils"
)

func Soal03(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	num := 2

	for row := 0; row < n; row++ {
		if row%2 == 0 { // row genap, value nambah
			for col := 0; col <= n-row-1; col++ {
				result[row][col] = num
				num += 3
			}
		} else {
			// row ganjil, value kurang
			num -= 3
			for col := 0; col <= n-row-1; col++ {
				result[row][col] = num
				num -= 3
			}
			num += 3
		}
	}
	return result
}

// polanya udah bener, angkanya masih salah LOL
