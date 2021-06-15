package task_mutex

import (
	"fmt"
	"sync"
)

func Start() {
	data := make(map[int]int, 20)

	for i := 0; i < 20; i++ {
		go Write(data, i, i)

	}

	for i := 0; i < 20; i++ {
		Read(data, i)
	}
}

func Write(data map[int]int, key int, value int) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	data[key] = value
}

func Read(data map[int]int, key int) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	if val, ok := data[key]; ok {
		fmt.Println(val)
	}
}
