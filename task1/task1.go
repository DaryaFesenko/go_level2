package task1

import "fmt"

func Run() {
	defer func() {
		if p := recover(); p != nil {
			fmt.Println("panic was caught ", p)
		}
	}()

	fmt.Println("Программа завершается без аварий")
}

func forPanic() {
	arr := [3]int{1, 2, 3}

	for i := 0; i < 5; i++ {
		fmt.Println(arr[i])
	}
}
