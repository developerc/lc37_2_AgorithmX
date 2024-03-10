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
	fmt.Printf("duration algorithmX %d us\n", dur.Microseconds())

	board2 := [2][2]int{}
	board2[0][0] = 1
	before = time.Now()
	answ2 := algorithms.AlgDFS(board2)
	after = time.Now()
	dur = after.Sub(before)
	fmt.Println(answ2)
	fmt.Printf("duration backtracking %d us\n", dur.Microseconds())
}
