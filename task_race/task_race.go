package task_race

func Start() {
	data := make(map[int]int, 20)

	for i := 0; i < 10; i++ {
		go func(key int, value int) {
			data[key] = value
		}(i, i)
	}
}
