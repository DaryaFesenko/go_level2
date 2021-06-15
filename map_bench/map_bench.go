package mapbench

import (
	"math/rand"
	"sync"
)

func Mutex(read int, write int) {
	data := make(map[int]int, 100)

	data[0] = 100

	for i := 1; i < 100; i++ {
		r := rand.Intn(100)

		if r-write < 0 {
			WriteMutex(data, i, i)
		} else {
			ReadMutex(data, 0)
		}
	}
}

func RWMutex(read int, write int) {
	data := make(map[int]int, 100)

	data[0] = 100

	for i := 1; i < 100; i++ {
		r := rand.Intn(100)

		if r-write < 0 {
			WriteRWMutex(data, i, i)
		} else {
			ReadRWMutex(data, 0)
		}
	}
}

func WriteMutex(data map[int]int, key int, value int) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	data[key] = value
}

func ReadMutex(data map[int]int, key int) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	_ = data[key]
}

func WriteRWMutex(data map[int]int, key int, value int) {
	var mutex sync.RWMutex
	mutex.Lock()
	defer mutex.Unlock()

	data[key] = value
}

func ReadRWMutex(data map[int]int, key int) {
	var mutex sync.RWMutex
	mutex.RLock()
	defer mutex.RUnlock()

	_ = data[key]
}
