package algorithms

import (
	"fmt"
	"os"
)

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
				cNet := 2*row + 1*col // + 8 //столбец
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
				for val := 1; val <= 2; val++ {
					//ограничения в строках
					rNet := 4*row + 2*col + val - 1
					if val == 1 {
						cNet = 4 + row
					} else {
						cNet = 6 + row
					}
					node := &Node{row: rNet, col: cNet, data: 1}
					colRowNet[rNet].nextRight.nextRight = node
					node.nextLeft = colRowNet[rNet].nextRight
					if colRowNet[cNet].nextDown == nil {
						colRowNet[cNet].nextDown = node
						node.nextUp = &colRowNet[cNet]
					} else {
						colRowNet[cNet].nextDown.nextDown = node
						node.nextUp = colRowNet[cNet].nextDown
					}
				}
				for val := 1; val <= 2; val++ {
					//ограничения в столбцах
					rNet := 4*row + 2*col + val - 1
					if val == 1 {
						cNet = 8 + row
					} else {
						cNet = 10 + row
					}
					node := &Node{row: rNet, col: cNet, data: 1}
					colRowNet[rNet].nextRight.nextRight.nextRight = node
					node.nextLeft = colRowNet[rNet].nextRight.nextRight
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
				cNet := 2*row + 1*col // + 8
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
				//ограничения в строках
				//rNet := 4*row + 2*col + val - 1
				if board[row][col] == 1 {
					cNet = 4 + row
				} else {
					cNet = 6 + row
				}
				node = &Node{row: rNet, col: cNet, data: 1}
				colRowNet[rNet].nextRight.nextRight = node
				node.nextLeft = colRowNet[rNet].nextRight
				if colRowNet[cNet].nextDown == nil {
					colRowNet[cNet].nextDown = node
					node.nextUp = &colRowNet[cNet]
				} else {
					colRowNet[cNet].nextDown.nextDown = node
					node.nextUp = colRowNet[cNet].nextDown
				}
				//ограничения в столбцах
				//rNet := 4*row + 2*col + val - 1
				if board[row][col] == 1 {
					cNet = 8 + row
				} else {
					cNet = 10 + row
				}
				node = &Node{row: rNet, col: cNet, data: 1}
				colRowNet[rNet].nextRight.nextRight.nextRight = node
				node.nextLeft = colRowNet[rNet].nextRight.nextRight
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
	fo, err := os.Create(fileName) // open output file
	if err != nil {
		panic(err)
	}
	defer fo.Close() // close fo on exit and check for its returned error
	for i := 0; i < 12; i++ {
		_, err = fo.WriteString(fmt.Sprintf("%d\t", i)) // пишем заголовок столбцов
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
	for i := 0; i < 8; i++ {
		fo.WriteString(fmt.Sprintf("\n"))
		if colRowNet[i].nextRight == nil {
			continue
		}
		numTabs := colRowNet[i].nextRight.col
		for k := 0; k < numTabs; k++ {
			fo.WriteString(fmt.Sprintf("\t"))
		}
		fo.WriteString(fmt.Sprintf("1"))
		//
		//fo.WriteString(fmt.Sprintf("\n"))
		if colRowNet[i].nextRight.nextRight == nil {
			continue
		}
		numTabs = colRowNet[i].nextRight.nextRight.col - colRowNet[i].nextRight.col
		for k := 0; k < numTabs; k++ {
			fo.WriteString(fmt.Sprintf("\t"))
		}
		fo.WriteString(fmt.Sprintf("1"))
		//
		if colRowNet[i].nextRight.nextRight.nextRight == nil {
			continue
		}
		numTabs = colRowNet[i].nextRight.nextRight.nextRight.col - colRowNet[i].nextRight.nextRight.col
		for k := 0; k < numTabs; k++ {
			fo.WriteString(fmt.Sprintf("\t"))
		}
		fo.WriteString(fmt.Sprintf("1"))
	}
}
