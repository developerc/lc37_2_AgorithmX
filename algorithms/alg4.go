package algorithms

import (
	"fmt"
	"os"
)

type NodeNet struct {
	colNet [8]Node
	rowNet [12]Node
}

var n int

func SolveAlgX4(board [2][2]int) [2][2]int {
	colRowNet := fillNet4(&board)
	//printNetToFile4("net"+strconv.Itoa(n)+".txt", colRowNet)
	algX(colRowNet, &board)
	return board
}

func algX(colRowNet NodeNet, board *[2][2]int) [2][2]int {
	if isTableEmpty(colRowNet) {
		return *board
	}
	n++
	for i := 0; i < 12; i++ { //будем искать столбцы с одной нодой
		if colRowNet.rowNet[i].nextDown != nil && colRowNet.rowNet[i].nextDown.nextDown == nil {
			colRowNet, *board = coverRows4(colRowNet, i, board)
			//printNetToFile4("net"+strconv.Itoa(n)+".txt", colRowNet)
			algX(colRowNet, board)
			break
		}
	}
	return *board
}

func coverRows4(colRowNet NodeNet, i int, board *[2][2]int) (NodeNet, [2][2]int) {
	rowForCover := colRowNet.rowNet[i].nextDown.row //главная строка для покрытия
	*board = buildBoard(rowForCover, board)
	//fmt.Println("строка возможного ответа: rowForCover = ", rowForCover)
	if colRowNet.colNet[rowForCover].nextRight.nextDown != nil { //смотрим вправо на одну ноду и вниз
		colRowNet = coverSecondRow(colRowNet, colRowNet.colNet[rowForCover].nextRight.nextDown.row)
	}
	if colRowNet.colNet[rowForCover].nextRight.nextUp.data != 0 { //смотрим вправо на одну ноду и вверх
		colRowNet = coverSecondRow(colRowNet, colRowNet.colNet[rowForCover].nextRight.nextUp.row)
	}
	//--
	if colRowNet.colNet[rowForCover].nextRight.nextRight.nextDown != nil { //смотрим вправо на 2 ноду и вниз
		colRowNet = coverSecondRow(colRowNet, colRowNet.colNet[rowForCover].nextRight.nextRight.nextDown.row)
	}
	if colRowNet.colNet[rowForCover].nextRight.nextRight.nextUp.data != 0 { //смотрим вправо на 2 ноду и вверх
		colRowNet = coverSecondRow(colRowNet, colRowNet.colNet[rowForCover].nextRight.nextRight.nextUp.row)
	}
	//--
	if colRowNet.colNet[rowForCover].nextRight.nextRight.nextRight.nextDown != nil { //смотрим вправо на 3 ноду и вниз
		colRowNet = coverSecondRow(colRowNet, colRowNet.colNet[rowForCover].nextRight.nextRight.nextRight.nextDown.row)
	}
	//--
	if colRowNet.colNet[rowForCover].nextRight.nextRight.nextRight.nextUp.data != 0 { //смотрим вправо на 3 ноду и вверх
		colRowNet = coverSecondRow(colRowNet, colRowNet.colNet[rowForCover].nextRight.nextRight.nextRight.nextUp.row)
	}

	colRowNet = coverSecondRow(colRowNet, rowForCover)
	return colRowNet, *board
}

func coverSecondRow(colRowNet NodeNet, row int) NodeNet {
	//fmt.Println("coverSecondRow row = ", row)
	//-- 1
	if colRowNet.colNet[row].nextRight.nextUp.data != 0 { //если выше не шапка
		if colRowNet.colNet[row].nextRight.nextDown != nil {
			colRowNet.colNet[row].nextRight.nextDown.nextUp = colRowNet.colNet[row].nextRight.nextUp
			colRowNet.colNet[row].nextRight.nextUp.nextDown = colRowNet.colNet[row].nextRight.nextDown
		} else {
			colRowNet.colNet[row].nextRight.nextUp.nextDown = nil
		}
	} else { //если выше шапка
		col := colRowNet.colNet[row].nextRight.nextUp.col
		if colRowNet.colNet[row].nextRight.nextDown != nil {
			colRowNet.rowNet[col].nextDown.nextDown.nextUp = &colRowNet.rowNet[col]
			colRowNet.rowNet[col].nextDown = colRowNet.colNet[row].nextRight.nextDown
		} else {
			colRowNet.rowNet[col].nextDown = nil
		}
	}
	//-- 2
	if colRowNet.colNet[row].nextRight.nextRight.nextUp.data != 0 { //если выше не шапка
		if colRowNet.colNet[row].nextRight.nextRight.nextDown != nil {
			colRowNet.colNet[row].nextRight.nextRight.nextDown.nextUp = colRowNet.colNet[row].nextRight.nextRight.nextUp
			colRowNet.colNet[row].nextRight.nextRight.nextUp.nextDown = colRowNet.colNet[row].nextRight.nextRight.nextDown
		} else {
			colRowNet.colNet[row].nextRight.nextRight.nextUp.nextDown = nil
		}
	} else { //если выше шапка
		col := colRowNet.colNet[row].nextRight.nextRight.nextUp.col
		if colRowNet.colNet[row].nextRight.nextRight.nextDown != nil {
			colRowNet.rowNet[col].nextDown.nextDown.nextUp = &colRowNet.rowNet[col]
			colRowNet.rowNet[col].nextDown = colRowNet.colNet[row].nextRight.nextRight.nextDown
		} else {
			colRowNet.rowNet[col].nextDown = nil
		}
	}
	//-- 3
	if colRowNet.colNet[row].nextRight.nextRight.nextRight.nextUp.data != 0 { //если выше не шапка
		if colRowNet.colNet[row].nextRight.nextRight.nextRight.nextDown != nil {
			colRowNet.colNet[row].nextRight.nextRight.nextRight.nextDown.nextUp = colRowNet.colNet[row].nextRight.nextRight.nextRight.nextUp
			colRowNet.colNet[row].nextRight.nextUp.nextDown = colRowNet.colNet[row].nextRight.nextDown
		} else {
			colRowNet.colNet[row].nextRight.nextRight.nextRight.nextUp.nextDown = nil
		}
	} else { //если выше шапка
		col := colRowNet.colNet[row].nextRight.nextRight.nextRight.nextUp.col
		if colRowNet.colNet[row].nextRight.nextRight.nextRight.nextDown != nil {
			colRowNet.rowNet[col].nextDown.nextDown.nextUp = &colRowNet.rowNet[col]
			colRowNet.rowNet[col].nextDown = colRowNet.colNet[row].nextRight.nextRight.nextRight.nextDown
		} else {
			colRowNet.rowNet[col].nextDown = nil
		}
	}

	colRowNet.colNet[row].nextRight = nil
	return colRowNet
}

func fillNet4(board *[2][2]int) NodeNet {
	nodeNet := NodeNet{}

	for i := 0; i < 12; i++ {
		nodeNet.rowNet[i].col = i
	}

	var cNet, rNet int
	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			valEnd := 1
			cNet = 2*row + 1*col      //столбец
			if board[row][col] == 0 { //если val равен 0 то добавляем Node для 1 и 2
				valEnd = 2
			}
			//ограничения в ячейках
			for val := 1; val <= valEnd; val++ {
				if valEnd == 1 {
					rNet = 4*row + 2*col + board[row][col] - 1
				} else {
					rNet = 4*row + 2*col + val - 1 //ряд
				}
				node := &Node{row: rNet, col: cNet, data: 1}
				nodeNet.colNet[rNet].nextRight = node
				node.nextLeft = &nodeNet.colNet[rNet]
				if nodeNet.rowNet[cNet].nextDown == nil {
					nodeNet.rowNet[cNet].nextDown = node
					node.nextUp = &nodeNet.rowNet[cNet]
				} else {
					nodeNet.rowNet[cNet].nextDown.nextDown = node
					node.nextUp = nodeNet.rowNet[cNet].nextDown
				}
			}
			//ограничения в строках
			for val := 1; val <= valEnd; val++ {
				if valEnd == 1 {
					rNet = 4*row + 2*col + board[row][col] - 1
				} else {
					rNet = 4*row + 2*col + val - 1 //ряд
				}
				if valEnd == 2 {
					if val == 1 {
						cNet = 4 + row
					} else {
						cNet = 6 + row
					}
				} else {
					if board[row][col] == 1 {
						cNet = 4 + row
					} else {
						cNet = 6 + row
					}
				}

				node := &Node{row: rNet, col: cNet, data: 1}
				nodeNet.colNet[rNet].nextRight.nextRight = node
				node.nextLeft = nodeNet.colNet[rNet].nextRight
				if nodeNet.rowNet[cNet].nextDown == nil {
					nodeNet.rowNet[cNet].nextDown = node
					node.nextUp = &nodeNet.rowNet[cNet]
				} else {
					nodeNet.rowNet[cNet].nextDown.nextDown = node
					node.nextUp = nodeNet.rowNet[cNet].nextDown
				}
			}
			//ограничения в столбцах
			for val := 1; val <= valEnd; val++ {
				if valEnd == 1 {
					rNet = 4*row + 2*col + board[row][col] - 1
				} else {
					rNet = 4*row + 2*col + val - 1 //ряд
				}
				if valEnd == 2 {
					if val == 1 {
						cNet = 8 + (board[row][col]/2)*2 + col
					} else {
						cNet = 10 + (board[row][col]/2)*2 + col
					}
				} else {
					if board[row][col] == 1 {
						cNet = 8 + col
					} else {
						cNet = 10 + col
					}
				}

				node := &Node{row: rNet, col: cNet, data: 1}
				nodeNet.colNet[rNet].nextRight.nextRight.nextRight = node
				node.nextLeft = nodeNet.colNet[rNet].nextRight.nextRight
				if nodeNet.rowNet[cNet].nextDown == nil {
					nodeNet.rowNet[cNet].nextDown = node
					node.nextUp = &nodeNet.rowNet[cNet]
				} else {
					nodeNet.rowNet[cNet].nextDown.nextDown = node
					node.nextUp = nodeNet.rowNet[cNet].nextDown
				}
			}
		}
	}

	return nodeNet
}

func isTableEmpty(colRowNet NodeNet) bool {
	for i := 0; i < len(colRowNet.colNet); i++ {
		if colRowNet.colNet[i].nextRight != nil {
			return false
		}
	}
	return true
}

func buildBoard(rowForCover int, board *[2][2]int) [2][2]int {
	switch rowForCover {
	case 0:
		board[0][0] = 1
	case 1:
		board[0][0] = 2
	case 2:
		board[0][1] = 1
	case 3:
		board[0][1] = 2
	case 4:
		board[1][0] = 1
	case 5:
		board[1][0] = 2
	case 6:
		board[1][1] = 1
	case 7:
		board[1][1] = 2
	}
	return *board
}

func printNetToFile4(fileName string, colRowNet NodeNet) {
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
		fo.WriteString("\n")
		if colRowNet.colNet[i].nextRight == nil {
			continue
		}
		numTabs := colRowNet.colNet[i].nextRight.col
		for k := 0; k < numTabs; k++ {
			fo.WriteString("\t")
		}
		fo.WriteString("1")
		//--
		numTabs = colRowNet.colNet[i].nextRight.nextRight.col - colRowNet.colNet[i].nextRight.col
		for k := 0; k < numTabs; k++ {
			fo.WriteString("\t")
		}
		fo.WriteString("1")
		//--
		numTabs = colRowNet.colNet[i].nextRight.nextRight.nextRight.col - colRowNet.colNet[i].nextRight.nextRight.col
		for k := 0; k < numTabs; k++ {
			fo.WriteString("\t")
		}
		fo.WriteString("1")
	}
}
