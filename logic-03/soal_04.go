package logic_03

import (
	"github.com/zaadinda/day7-amartha/utils"
)

func Soal04(n int) (result [][]int) {
	result = utils.CreateSlice(n)

	num := 1

	for row := 0; row < n; row++ {
		for col := 0; col <= row; col++ {
			if row%2 == 0 {
				result[row][n-1-col] = num // genap- turun - dr kanan ke kiri
			} else {
				result[row][n-1-row+col] = num // ganjil - naik
			}
			num += 2
		}
	}
	return result
}
