package fixed

import (
	"fmt"
	"sync"
)

type MapMutex struct {
	Data  map[int]int
	mutex sync.Mutex
}

func New(size int) *MapMutex {
	data := make(map[int]int, size)
	return &MapMutex{Data: data}
}

func (m *MapMutex) WriteMutex(key int, value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Data[key] = value
}

func (m *MapMutex) ReadMutex(key int) error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, ok := m.Data[key]; !ok {
		return fmt.Errorf("key %d not found", key)
	}

	return nil
}
