package fibonacci

import "fmt"

var cached map[int]int

func init() {
	cached = map[int]int{
		1: 1,
		2: 1,
	}
}

// Fibonacci finds a numeric tracker by its index
func Fibonacci(n int) int {

	if val, ok := cached[n]; ok {
		defer func(val, n int) {
			fmt.Printf("Из кэша использовано: [%v] = %v\n", n, val)
		}(val, n)
		return val
	}

	result := calculationFibonacci(n)
	cached[n] = result
	return result
}

func calculationFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return calculationFibonacci(n-1) + calculationFibonacci(n-2)
}
