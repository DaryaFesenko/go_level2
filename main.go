package main

import (
	//"go_level_2/stream"
	//"go_level_2/task_mutex"
	mapbench "go_level_2/map_bench"
)

func main() {
	//stream.Start(40)

	//task_mutex.Start()

	mapbench.RWMutex(50, 50)
}
