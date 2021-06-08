package stream

import (
	"fmt"
	"sync"
)

func Start(n int) {
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func(j int) {
			fmt.Println("stream ", j)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
}
