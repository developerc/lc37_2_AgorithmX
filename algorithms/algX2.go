package algorithms

var headerRow [12]Node //массив заголовочных узлов строка
var headerCol [8]Node  //массив заголовочных узлов столбец

func SolveAlgX_2(board [2][2]int) [2][2]int {
	for i := 0; i < 12; i++ {
		headerRow[i].col = i
	}
	for i := 0; i < 8; i++ {
		headerCol[i].row = i
	}
	fillNet(&board)
	return board
}

func fillNet(board *[2][2]int) { //заполняем исходную сеть
	var rNet, cNet int //ряд и столбец сети
	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			if board[row][col] == 0 {
				rNet = 4*row + 2*col + board[row][col] - 1
				cNet = 2*row + 1*col
				addNodeToNet(rNet, cNet)
			} else {
				//todo
			}
		}
	}
}

func addNodeToNet(rNet, cNet int) {
	//todo
}
