package task3

import (
	"fmt"
	"os"
)

func Run() {

	file, err := os.Create("file.txt")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	file.WriteString("Hello")
}
