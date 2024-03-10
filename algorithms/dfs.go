package algorithms

func AlgDFS(board [2][2]int) [2][2]int {
	backtrack(&board)
	return board
}

func backtrack(board *[2][2]int) bool {
	if !hasEmptyCell(board) {
		return true
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if board[i][j] == 0 {
				for candidate := 2; candidate >= 1; candidate-- {
					board[i][j] = candidate
					if isBoardValid(board) {
						if backtrack(board) {
							return true
						}
						board[i][j] = 0
					} else {
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func hasEmptyCell(board *[2][2]int) bool {
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if board[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func isBoardValid(board *[2][2]int) bool {

	//check duplicates by row
	for row := 0; row < 2; row++ {
		counter := [3]int{}
		for col := 0; col < 2; col++ {
			counter[board[row][col]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}

	//check duplicates by column
	for row := 0; row < 2; row++ {
		counter := [3]int{}
		for col := 0; col < 2; col++ {
			counter[board[col][row]]++
		}
		if hasDuplicates(counter) {
			return false
		}
	}
	return true
}

func hasDuplicates(counter [3]int) bool {
	for i, count := range counter {
		if i == 0 {
			continue
		}
		if count > 1 {
			return true
		}
	}
	return false
}
