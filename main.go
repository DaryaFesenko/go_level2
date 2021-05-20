package main

import (
	"fmt"

	"go_level_2/task1"
	"go_level_2/task2"
	"go_level_2/task3"
)

func main() {

	fmt.Print("Задание 1: ")
	task1.Run()

	fmt.Println("Задание 2: ")
	err := task2.New("my error")
	fmt.Println(err)

	fmt.Println("Задание 3: ")
	task3.Run()
}
