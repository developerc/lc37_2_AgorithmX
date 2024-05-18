package main

import (
	"exerciss/algorithms"
	"fmt"
	"time"
)

func main() {
	board := [][]int{
		[]int{1, 0},
		[]int{0, 0},
	}

	before := time.Now()
	answ := algorithms.SolveAlgX(board)
	after := time.Now()
	dur := after.Sub(before)
	fmt.Println(answ)
	fmt.Printf("duration algorithmX %d us\n", dur.Nanoseconds())

	board2 := [2][2]int{}
	board2[0][0] = 1
	before = time.Now()
	answ2 := algorithms.AlgDFS(board2)
	after = time.Now()
	dur = after.Sub(before)
	fmt.Println(answ2)
	fmt.Printf("duration backtracking %d us\n", dur.Nanoseconds())

	board3 := [2][2]int{}
	//board3[0][0] = 1
	before = time.Now()
	answ3 := algorithms.SolveAlgX2(board3)
	after = time.Now()
	dur = after.Sub(before)
	fmt.Println(answ3)
	fmt.Printf("duration algorithmX2 %d us\n", dur.Nanoseconds())
}
