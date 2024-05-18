package algorithms

import (
	"fmt"
	"os"
)

var rowNet [12]Node
var colNet [8]Node

func SolveAlgX2(board [2][2]int) [2][2]int {
	fillNet(&board)
	return board
}

func fillNet(board *[2][2]int) {
	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			addNodesCellsRestr(row, col, board)
		}
	}
}

func addNodesCellsRestr(row int, col int, board *[2][2]int) {
	if board[row][col] == 0 { //если val равен 0 то добавляем Node для 1 и 2
		upNode := rowNet[col]
		for val := 1; val <= 2; val++ {
			rNet := 4*row + 2*col + val - 1
			cNet := 2*row + 1*col
			node := &Node{row: rNet, col: cNet}
			colNet[row].nextRight = node
			node.nextLeft = &colNet[row]
			upNode.nextDown = node
			node.nextUp = &upNode
			upNode = *node
		}
	} else { //если val не равен 0 то добавляем Node для 1 или 2
		//
	}
}

func printNetToFile(fileName string) {
	fo, err := os.Create(fileName) // open output file
	if err != nil {
		panic(err)
	}
	defer fo.Close() // close fo on exit and check for its returned error

	for i := 0; i <= 12; i++ {
		_, err = fo.WriteString(fmt.Sprintf("%d\t", i)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
}
