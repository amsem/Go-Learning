package leetcode

func getRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)
	row[0] = 1
	var i, j, m, val int
	for i = 1; i < len(row); i++ {
		m = (i + 1) / 2
		for j = 0; j < m; j++ {
			val = 1
			if j > 0 && j < i {
				val = row[j-1] + row[j]
			}
			row[i-j] = val
		}
		if i-2*m == 0 {
			row[m] = 2 * row[m-1]
		}
		for j = 0; j < m; j++ {
			row[j] = row[i-j]
		}
	}
	return row
}

