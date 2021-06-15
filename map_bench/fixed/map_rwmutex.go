package fixed

import (
	"fmt"
	"sync"
)

type MapRWMutex struct {
	Data  map[int]int
	mutex sync.RWMutex
}

func NewRW(size int) *MapRWMutex {
	data := make(map[int]int, size)
	return &MapRWMutex{Data: data}
}

func (m *MapRWMutex) WriteRWMutex(key int, value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.Data[key] = value
}

func (m *MapRWMutex) ReadRWMutex(key int) error {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	if _, ok := m.Data[key]; !ok {
		return fmt.Errorf("key %d not found", key)
	}

	return nil
}
