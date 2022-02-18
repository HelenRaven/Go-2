package prime

import "fmt"

// Prime returns array of prime numbers between 0 and given (num)
func Prime(num int) ([]int, error) {
	var result []int

	if num <= 0 {
		return result, fmt.Errorf("number must be positive: %d", num)

	} else if num == 1 {
		return result, fmt.Errorf("there is no prime numbers between 0 and %d", num)

	} else if num == 2 {
		result = append(result, 2)
		return result, nil

	} else {
		result = append(result, 2)

		for i := 3; i <= num; i += 2 {
			isPrime := true
			for j := 3; j < i; j += 2 {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				result = append(result, i)
			}
		}
	}

	return result, nil
}
