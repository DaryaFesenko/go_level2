package mutex_stream

import (
	"os"
	"runtime/trace"
	"sync"
)

// Похожее задание было в прошлом дз
// Только я не разобралась, как выполнить трассировку
// Может дело в том, что пытаюсь из windows запускать
// Пишет, что файл trace.out не подходит

func Start() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	data := make(map[int]int, 20)
	var mutex sync.Mutex

	for i := 0; i < 10; i++ {
		go func(key int, value int) {
			mutex.Lock()
			data[key] = value
			mutex.Unlock()
		}(i, i)
	}

	for i := 10; i < 20; i++ {
		go func(key int, value int) {
			mutex.Lock()
			data[key] = value

			mutex.Unlock()
		}(i, i)
	}
}
