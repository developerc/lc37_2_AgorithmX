package algorithms

import (
	"fmt"
	"os"
	//"sync"
)

var rowNet [12]Node
var colNet [8]Node
var accumulator [12]int //накапливаем значения строк

func SolveAlgX2(board [2][2]int) [2][2]int {
	fillNet(&board)
	//сделаем в цикле
	numColOneNode := findOneInCol()
	fmt.Println("numColOneNode = ", numColOneNode)
	rowToAccum(numColOneNode)

	return board
}

func rowToAccum(numColOneNode int) { //помещаем строку в аккумулятор
	//fmt.Println("colNet[numColOneNode].nextRight.col = ", colNet[rowNet[numColOneNode].nextDown.row].nextRight.col)
	accumulator[colNet[rowNet[numColOneNode].nextDown.row].nextRight.col]++
	accumulator[colNet[rowNet[numColOneNode].nextDown.row].nextRight.nextRight.col]++
	accumulator[colNet[rowNet[numColOneNode].nextDown.row].nextRight.nextRight.nextRight.col]++
	fmt.Println(accumulator)
	row := rowNet[numColOneNode].nextDown.row //находим номер строки одинокой ноды
	nextNode := colNet[row]
	for nextNode.nextRight != nil {
		coverRows(nextNode)
		nextNode = *nextNode.nextRight
	}
}

func coverRows(nextNode Node) {

}

func findOneInCol() int { //находим столбец с одной нодой
	numCol := 0
	for ; numCol < 12; numCol++ {
		if rowNet[numCol].nextDown != nil && rowNet[numCol].nextDown.nextDown == nil {
			break
		}
	}
	return numCol
}

func fillNet(board *[2][2]int) { //заполняем сеть

	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			//go func() {
			if board[row][col] == 0 { //если val равен 0 то добавляем Node для 1 и 2
				cNet := 2*row + 1*col
				for val := 1; val <= 2; val++ {
					//ограничения в ячейках
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
				for val := 1; val <= 2; val++ {
					//ограничения в строках
					rNet := 4*row + 2*col + val - 1
					if val == 1 {
						cNet = 4 + row
					} else {
						cNet = 6 + row
					}
					node := &Node{row: rNet, col: cNet}
					colNet[rNet].nextRight.nextRight = node
					node.nextLeft = colNet[rNet].nextRight
					if rowNet[cNet].nextDown == nil {
						rowNet[cNet].nextDown = node
						node.nextUp = &rowNet[cNet]
					} else {
						rowNet[cNet].nextDown.nextDown = node
						node.nextUp = rowNet[cNet].nextDown
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
					node := &Node{row: rNet, col: cNet}
					colNet[rNet].nextRight.nextRight.nextRight = node
					node.nextLeft = colNet[rNet].nextRight.nextRight
					if rowNet[cNet].nextDown == nil {
						rowNet[cNet].nextDown = node
						node.nextUp = &rowNet[cNet]
					} else {
						rowNet[cNet].nextDown.nextDown = node
						node.nextUp = rowNet[cNet].nextDown
					}
				}
			} else { //если val не равен 0 то добавляем Node для 1 или 2
				//ограничения в ячейках
				cNet := 2*row + 1*col
				rNet := 4*row + 2*col + board[row][col] - 1
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
				//ограничения в строках
				//rNet := 4*row + 2*col + val - 1
				if board[row][col] == 1 {
					cNet = 4 + row
				} else {
					cNet = 6 + row
				}
				node = &Node{row: rNet, col: cNet}
				colNet[rNet].nextRight.nextRight = node
				node.nextLeft = colNet[rNet].nextRight
				if rowNet[cNet].nextDown == nil {
					rowNet[cNet].nextDown = node
					node.nextUp = &rowNet[cNet]
				} else {
					rowNet[cNet].nextDown.nextDown = node
					node.nextUp = rowNet[cNet].nextDown
				}
				//ограничения в столбцах
				//rNet := 4*row + 2*col + val - 1
				if board[row][col] == 1 {
					cNet = 8 + row
				} else {
					cNet = 10 + row
				}
				node = &Node{row: rNet, col: cNet}
				colNet[rNet].nextRight.nextRight.nextRight = node
				node.nextLeft = colNet[rNet].nextRight.nextRight
				if rowNet[cNet].nextDown == nil {
					rowNet[cNet].nextDown = node
					node.nextUp = &rowNet[cNet]
				} else {
					rowNet[cNet].nextDown.nextDown = node
					node.nextUp = rowNet[cNet].nextDown
				}
			}

		}
	}

	//printCol()
	//fmt.Println("-------")
	printRow()
	printNetToFile("net.txt")
}

func printCol() {
	for i := 0; i < 8; i++ {
		fmt.Println(colNet[i].nextRight)
	}
}
func printRow() {
	fmt.Println(rowNet[11].nextDown)
	fmt.Println(rowNet[11].nextDown.nextDown)
	/*for i := 0; i < 12; i++ {
		fmt.Println(rowNet[i].nextDown.nextDown)
	}*/
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
		if colNet[i].nextRight == nil {
			continue
		}
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
				numTabs = colNet[i].nextRight.nextRight.nextRight.col - colNet[i].nextRight.nextRight.col
				for k := 0; k < numTabs; k++ {
					fo.WriteString(fmt.Sprintf("\t"))
				}
				fo.WriteString(fmt.Sprintf("1"))
			}
		}

	}
}
