package algorithms

import (
	"fmt"
	"os"
	//"sync"
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
			//go func() {
			if board[row][col] == 0 { //если val равен 0 то добавляем Node для 1 и 2
				cNet := 2*row + 1*col
				for val := 1; val <= 2; val++ {
					rNet := 4*row + 2*col + val - 1
					node := &Node{row: rNet, col: cNet}
					colNet[rNet].nextRight = node
					node.nextLeft = &colNet[rNet]
					if rowNet[cNet].nextDown == nil {
						rowNet[cNet].nextDown = node
						node.nextUp = &rowNet[cNet]
					} else {
						rowNet[cNet].nextDown.nextDown = node
						node.nextUp = rowNet[cNet].nextDown
					}
				}
			} else { //если val не равен 0 то добавляем Node для 1 или 2
				//
			}

		}
	}

	printCol()
	//fmt.Println("-------")
	//printRow()
	printNetToFile("net.txt")
}

func printCol() {
	for i := 0; i < 8; i++ {
		fmt.Println(colNet[i].nextRight)
	}
}
func printRow() {
	for i := 0; i < 12; i++ {
		fmt.Println(rowNet[i].nextDown.nextDown)
	}
}

func printNetToFile(fileName string) {
	fo, err := os.Create(fileName) // open output file
	if err != nil {
		panic(err)
	}
	defer fo.Close() // close fo on exit and check for its returned error

	for i := 0; i < 12; i++ {
		_, err = fo.WriteString(fmt.Sprintf("%d\t", i)) // writing...
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}
	for i := 0; i < 8; i++ {
		fo.WriteString(fmt.Sprintf("\n"))
		numTabs := colNet[i].nextRight.col
		//fmt.Println("numTabs", numTabs)
		for k := 0; k < numTabs; k++ {
			fo.WriteString(fmt.Sprintf("\t"))
		}
		fo.WriteString(fmt.Sprintf("1"))
		//---
		if colNet[i].nextRight.nextRight != nil {
			numTabs = colNet[i].nextRight.nextRight.col - numTabs
			for k := 0; k < numTabs; k++ {
				fo.WriteString(fmt.Sprintf("\t"))
			}
			fo.WriteString(fmt.Sprintf("1"))
			//---
			if colNet[i].nextRight.nextRight.nextRight != nil {
				numTabs = colNet[i].nextRight.nextRight.nextRight.col - numTabs
				for k := 0; k < numTabs; k++ {
					fo.WriteString(fmt.Sprintf("\t"))
				}
				fo.WriteString(fmt.Sprintf("1"))
			}
		}

	}
}
