package main

import "fmt"

type Node struct { //обычный узел
	data      int
	row       int
	col       int
	nextRight *Node
	nextLeft  *Node
	nextUp    *Node
	nextDown  *Node
}

type List struct { //список с узлами
	head *Node
}

type TablePossibleSolutions struct {
	head *Node
}

func main() {
	board := [][]int{
		[]int{0, 0},
		[]int{2, 0},
	}
	var l = List{}
	var t = List{}
	l = fillHeadsCol(l)
	l = fillHeadsRow(l)
	l = fillResrictTable(l)
	printList(l)
	t = fillHeadsCol(t)
	//printList(t)
	doRestrict(board, l, t)
}

func doRestrict(board [][]int, l, t List) {
	for row := 0; row < 2; row++ {
		for col := 0; col < 2; col++ {
			if board[row][col] != 0 {
				numRowTableRestr := 4*row + 2*col + board[row][col] //нашли какая строка должна остаться
				numRowRemove := findRemove(numRowTableRestr)
				fmt.Println("row = ", row, ", col = ", col, ", val = ", board[row][col], ", numRowTableRestr = ", numRowTableRestr, ", numRowRemove = ", numRowRemove)
				l = doCowerRow(l, numRowRemove) //сделаем накрытие строки numRowRemove
			}
		}
	}
	solveSudoku(l, t)
}

func solveSudoku(l, t List) {
	// составим Map с количеством единиц в столбцах
	mapOnesInCols := findMapOnesInCols(l)
	fmt.Println(mapOnesInCols)
	if colNum, ok := mapOnesInCols[1]; ok {
		numRowWithOne := findNumRowWithOne(colNum, l) //найдем номер строки с одной единицей
		fmt.Println("num row with one = ", numRowWithOne)
		t = addToTableProbableSolution(numRowWithOne, l, t)
	}

	/*for i := 1; i <= len(mapOnesInCols); i++ {
		//fmt.Println(i, mapOnesInCols[i])
		if mapOnesInCols[i] == 1 {
			//l.coverCol(i)
			break
		}
	}*/
}

func addToTableProbableSolution(numRowWithOne int, l, t List) List { //будем добавлять строку в таблицу возможных ответов

	return t
}

func findNumRowWithOne(colNum int, l List) int {
	colMain := l.head
	for colMain.col != colNum {
		colMain = colMain.nextRight
	}
	colMain = colMain.nextDown
	return colMain.row
}

func findMapOnesInCols(l List) map[int]int {
	var mapOnesInCols map[int]int = make(map[int]int) //key - количество единиц в столбце, val - номер столбца
	colMain := l.head.nextRight
	for i := 1; i <= 12; i++ {

		cntr := 0
		colCurr := colMain
		for colCurr.nextDown != nil {
			cntr++
			colCurr = colCurr.nextDown
		}
		if cntr > 0 { //в map помещаем только не нулевые столбцы
			mapOnesInCols[cntr] = i
		}
		if i < 12 {
			colMain = colMain.nextRight
		}
	}
	return mapOnesInCols
}

func doCowerRow(l List, numRowRemove int) List { //накрываем строку
	rowCower := l.head
	for rowCower.row != numRowRemove {
		rowCower = rowCower.nextDown
	}
	for {
		if rowCower.nextDown == nil {
			upCell := rowCower.nextUp
			upCell.nextDown = nil
		} else {
			upCell := rowCower.nextUp
			downCell := rowCower.nextDown
			upCell.nextDown = downCell
			downCell.nextUp = upCell
		}

		if rowCower.nextRight != nil {
			rowCower = rowCower.nextRight
		} else {
			break
		}
	}
	return l
}

func findRemove(numRowTableRestr int) int {
	for i := 1; i <= 4; i++ {
		if numRowTableRestr <= i*2 {
			if numRowTableRestr == i*2 {
				return i*2 - 1
			} else {
				return i * 2
			}
		}
	}
	return 0
}

func fillHeadsCol(l List) List { //Создаем заголовочные ноды столбцов
	headList := &Node{data: 1, col: 0, row: 0}
	l.head = headList
	currNode := l.head
	for i := 1; i <= 12; i++ {
		newNodeCol := &Node{data: 1, col: i, row: 0}
		currNode.nextRight = newNodeCol
		newNodeCol.nextLeft = currNode
		currNode = newNodeCol
	}
	currNode.nextRight = l.head //соединяем вкруговую
	l.head.nextLeft = currNode
	return l
}

func fillHeadsRow(l List) List { //создаем заголовочные ноды строк
	currNode := l.head
	for i := 1; i <= 8; i++ {
		newNodeCol := &Node{data: 1, col: 0, row: i}
		currNode.nextDown = newNodeCol
		newNodeCol.nextUp = currNode
		currNode = newNodeCol
	}
	currNode.nextDown = l.head //соединяем вкруговую
	l.head.nextUp = currNode
	return l
}

func printList(l List) {
	currNode := l.head
	for i := 0; i <= 12; i++ {
		fmt.Printf("%d\t", currNode.col)
		currNode = currNode.nextRight
	}
	fmt.Printf("\n") //распечатали заголовки столбцов
	currNode = l.head.nextDown
	//for i := 1; i <= 8; i++ {
	for currNode != l.head {
		if currNode == nil {
			break
		}
		fmt.Printf("%d", currNode.row) //заголовок строки
		amountTabs := 0                //номер предыдущего столбца
		currNodeRow := currNode
		for currNodeRow.nextRight != nil {
			amountTabs = currNodeRow.nextRight.col - currNodeRow.col
			for k := 0; k < amountTabs; k++ {
				fmt.Printf("\t")
			}
			fmt.Printf("1")
			currNodeRow = currNodeRow.nextRight
		}
		fmt.Printf("\n")
		currNode = currNode.nextDown
	}
	fmt.Printf("\n") //распечатали заголовки строк
}

func fillResrictTable(l List) List { //заполняем таблицу ограничений для судоку 2х2
	l = addNode(l, 1, 1)
	l = addNode(l, 5, 1)
	l = addNode(l, 9, 1)
	l = addNode(l, 1, 2)
	l = addNode(l, 7, 2)
	l = addNode(l, 11, 2)
	l = addNode(l, 2, 3)
	l = addNode(l, 5, 3)
	l = addNode(l, 10, 3)
	l = addNode(l, 2, 4)
	l = addNode(l, 7, 4)
	l = addNode(l, 12, 4)
	l = addNode(l, 3, 5)
	l = addNode(l, 6, 5)
	l = addNode(l, 9, 5)
	l = addNode(l, 3, 6)
	l = addNode(l, 8, 6)
	l = addNode(l, 11, 6)
	l = addNode(l, 4, 7)
	l = addNode(l, 6, 7)
	l = addNode(l, 10, 7)
	l = addNode(l, 4, 8)
	l = addNode(l, 8, 8)
	l = addNode(l, 12, 8)
	return l
}

func addNode(l List, col, row int) List {
	currNodeDown := l.head
	currNodeRight := l.head
	newNode := &Node{data: 1, col: col, row: row}
	for i := 0; i < row; i++ { //опускаемся на нужную строку
		currNodeDown = currNodeDown.nextDown
	}
	for currNodeDown.nextRight != nil { //в этой строке передвигаемся в крайнее правое положение
		currNodeDown = currNodeDown.nextRight
	}
	currNodeDown.nextRight = newNode //привязываемся к новой ноде справа
	newNode.nextLeft = currNodeDown

	for i := 0; i < col; i++ { //передвигаемся вправо на нужный столбец
		currNodeRight = currNodeRight.nextRight
	}
	for currNodeRight.nextDown != nil { //передвигаемся в столбце в крайнее нижнее положение
		currNodeRight = currNodeRight.nextDown
	}

	newNode.nextUp = currNodeRight
	currNodeRight.nextDown = newNode
	return l
}
