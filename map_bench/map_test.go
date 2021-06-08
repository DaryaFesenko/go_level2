package mapbench

import (
	"testing"
)

func BenchmarkMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mutex(50, 50)
	}
}

func BenchmarkRWMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RWMutex(50, 50)
	}
}

// 10:90 - записи больше, чем чтения
// Mutex -   15550 ns/op	    3589 B/op	     103 allocs/op  1.337s
// RWMutex - 16632 ns/op	    5173 B/op	     103 allocs/op  1.289s

// 50:50
// Mutex -   14115 ns/op	    3506 B/op	     101 allocs/op  1.359s
// RWMutex - 15034 ns/op	    5090 B/op	     101 allocs/op  1.354s

// 90:10 - чтения больше, чем записи
// Mutex -   11446 ns/op	    3504 B/op	     101 allocs/op  1.293s
// RWMutex - 12931 ns/op	    5088 B/op	     101 allocs/op  1.398s
