package prime_test

import (
	"fmt"
	"go2/prime"
)

func Example() {
	prime7, err := prime.Prime(7)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(prime7)
	}
}
