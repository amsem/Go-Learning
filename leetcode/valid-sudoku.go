func isValidSudoku(board [][]byte) bool {
    	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[i][j : j+1]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	for i := 0; i < 9; i++ {
		tmp := [10]int{}
		for j := 0; j < 9; j++ {
			cellVal := board[j][i]
			if string(cellVal) != "." {
				index, _ := strconv.Atoi(string(cellVal))
				if index > 9 || index < 1 {
					return false
				}
				if tmp[index] == 1 {
					return false
				}
				tmp[index] = 1
			}
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			tmp := [10]int{}
			for ii := i * 3; ii < i*3+3; ii++ {
				for jj := j * 3; jj < j*3+3; jj++ {
					cellVal := board[ii][jj]
					if string(cellVal) != "." {
						index, _ := strconv.Atoi(string(cellVal))
						if tmp[index] == 1 {
							return false
						}
						tmp[index] = 1
					}
				}
			}
		}
	}
	return true
}
