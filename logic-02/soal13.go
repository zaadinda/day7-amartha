package logic_02

import "github.com/zaadinda/day7-amartha/utils"

func Soal13(n int) (result [][]int) {

	//create slice
	result = utils.CreateSlice(n)

	// isi slice
	for row := 0; row < n; row++ {
		num := 1
		for col := 0; col < n; col++ {
			if row <= col && row+col <= n-1 {
				result[row][col] = num
				result[n-1-row][col] = num
			}
			num += 2
		}
	}
	return result
}
