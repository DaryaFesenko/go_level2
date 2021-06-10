package task_scheduler

import (
	"fmt"
	"runtime"
)

// Без runtime.Gosched() скорее всего вывод был бы таким: 0 1 2 3 4 5 6 7 8 9 10

func Start() {
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println(i)
			runtime.Gosched()
		}
	}()

	for i := 5; i < 10; i++ {
		runtime.Gosched()
		fmt.Println(i)
	}
}
