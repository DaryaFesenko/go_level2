package main

import (
	"fmt"
	"go_level_2/fibonacci"
	"os"
)

func main() {
	file, err := os.Create("file.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString("Hello")
	fmt.Println("Success")
	var a int
	fmt.Scanln(&a)
}

func task2() {
	var n int
	fmt.Scanln(&n)

	f := fibonacci.Fibonacci(n)

	fmt.Println(f)
}
