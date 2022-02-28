package main

import "fmt"

func WorkerPool(i int) int {
	//pool := make(chan struct{}, 1)
	res := make(chan int, 1)
	res <- 0

	for j := 0; j <= i; j++ {
		//pool <- struct{}{}

		go func() {

			r := <-res
			r++
			defer func() {
				res <- r
			}()
		}()
	}
	r := <-res
	return r
}

func main() {
	res := WorkerPool(1000)
	fmt.Println("Result:", res)
}
