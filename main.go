package main

import (
	"fmt"
	"go2/prime"
)

func main() {
	var num int
	fmt.Print("Enter number (positive and greater then 1): ")
	fmt.Scanln(&num)

	result, err := prime.Prime(num)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
