package main

import (
	"fmt"
	"sync"
)

func MatrixMultiplication(a, b [][]int) {
	var wg = sync.WaitGroup{}

	matrix := make([][]int, len(a))
	for i := range matrix {
		matrix[i] = make([]int, len(a))
	}

	for i := range a {
		for j := range a {
			for k := range a[i] {
				wg.Add(1)
				go func(i, j, k int) {
					defer wg.Done()
					matrix[i][j] += a[i][k] * b[k][j]
				}(i, j, k)
			}
		}
	}
	wg.Wait()

	for _, m := range matrix {
		fmt.Println(m)
	}

}

func main() {
	a := [][]int{
		{1, 2, 3},
		{1, 2, 3},
	}
	b := [][]int{
		{2, 3},
		{2, 3},
		{2, 3},
	}

	MatrixMultiplication(a, b)

}
