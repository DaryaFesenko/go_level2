package fixed

import (
	"fmt"
	"sync"
	"testing"
)

type reader func(int) error

type writer func(int, int)

var (
	map_size     = 10
	benchWorkers = 1000
	values       = []int{0, 10, 50, 90, 100}
)

func baseSetBench(b *testing.B, writeProbability int, reader reader, writer writer) {
	writers := benchWorkers * writeProbability / 100
	readers := benchWorkers - writers

	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(benchWorkers)

		for g := 0; g < readers; g++ {
			go func(i int) {
				defer wg.Done()
				reader(i)
			}(g)
		}

		for g := 0; g < writers; g++ {
			go func(i int) {
				defer wg.Done()

				writer(i, i)
			}(g)
		}
		wg.Wait()
	}
}

func BenchmarkMutexSet(b *testing.B) {
	for _, writeProbability := range values {
		benchDescr := fmt.Sprintf("test set write/read with %d%% writers", writeProbability)
		b.Run(benchDescr, func(b *testing.B) {
			test := New(map_size)
			baseSetBench(b, writeProbability, test.ReadMutex, test.WriteMutex)
		})
	}
}

func BenchmarkRWMutexSet(b *testing.B) {
	for _, writeProbability := range values {
		benchDescr := fmt.Sprintf("test set write/read with %d%% writers", writeProbability)
		b.Run(benchDescr, func(b *testing.B) {
			test := NewRW(map_size)
			baseSetBench(b, writeProbability, test.ReadRWMutex, test.WriteRWMutex)
		})
	}
}

/*
BenchmarkMutexSet/test_set_write/read_with_0%_writers-8         	     537	   2258405 ns/op	   45662 B/op	    2748 allocs/op
BenchmarkMutexSet/test_set_write/read_with_10%_writers-8        	     530	   2173392 ns/op	   37486 B/op	    2247 allocs/op
BenchmarkMutexSet/test_set_write/read_with_50%_writers-8        	     704	   1649957 ns/op	     148 B/op	       2 allocs/op
BenchmarkMutexSet/test_set_write/read_with_90%_writers-8        	     697	   1706902 ns/op	     172 B/op	       0 allocs/op
BenchmarkMutexSet/test_set_write/read_with_100%_writers-8       	     722	   1690562 ns/op	     170 B/op	       0 allocs/op
*/

/*
BenchmarkRWMutexSet/test_set_write/read_with_0%_writers-8         	    1654	    752268 ns/op	   45263 B/op	    2744 allocs/op
BenchmarkRWMutexSet/test_set_write/read_with_10%_writers-8        	    1443	    876893 ns/op	   37277 B/op	    2245 allocs/op
BenchmarkRWMutexSet/test_set_write/read_with_50%_writers-8        	     909	   1143862 ns/op	     117 B/op	       1 allocs/op
BenchmarkRWMutexSet/test_set_write/read_with_90%_writers-8        	     722	   1620832 ns/op	     230 B/op	       1 allocs/op
BenchmarkRWMutexSet/test_set_write/read_with_100%_writers-8       	     742	   1738007 ns/op	     155 B/op	       0 allocs/op
*/
