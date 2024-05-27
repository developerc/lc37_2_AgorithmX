package algorithms

func SolveAlgX3(board [2][2]int) [2][2]int {
	//var colRowNet [20]Node
	//var accumulator [12]int       //накапливаем значения строк
	colRowNet := fillNet3(&board) //var colRowNet [20]Node
	printNetToFile3("net.txt", colRowNet)

	return board
}

func fillNet3(board *[2][2]int) [20]Node {
	var colRowNet [20]Node
	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			if board[row][col] == 0 { //если val равен 0 то добавляем Node для 1 и 2
				cNet := 2*row + 1*col + 8 //столбец
				for val := 1; val <= 2; val++ {
					//ограничения в ячейках
					rNet := 4*row + 2*col + val - 1 //ряд
					node := &Node{row: rNet, col: cNet, data: 1}
					colRowNet[rNet].nextRight = node
					node.nextLeft = &colRowNet[rNet]
					if colRowNet[cNet].nextDown == nil {
						colRowNet[cNet].nextDown = node
						node.nextUp = &colRowNet[cNet]
					} else {
						colRowNet[cNet].nextDown.nextDown = node
						node.nextUp = colRowNet[cNet].nextDown
					}

				}
			} else { //если val не равен 0 то добавляем Node для 1 или 2
				//ограничения в ячейках
				cNet := 2*row + 1*col + 8
				rNet := 4*row + 2*col + board[row][col] - 1
				node := &Node{row: rNet, col: cNet, data: 1}
				colRowNet[rNet].nextRight = node
				node.nextLeft = &colRowNet[rNet]
				if colRowNet[cNet].nextDown == nil {
					colRowNet[cNet].nextDown = node
					node.nextUp = &colRowNet[cNet]
				} else {
					colRowNet[cNet].nextDown.nextDown = node
					node.nextUp = colRowNet[cNet].nextDown
				}
			}
		}
	}
	return colRowNet
}

func printNetToFile3(fileName string, colRowNet [20]Node) {

}
