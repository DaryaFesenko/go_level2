package main

import (
	"fmt"
	"go_level_2/transform"
)

func main() {
	RunTask1()
}

func RunTask1() {
	user := &struct {
		Name   string
		DOB    string
		Weight float64
		Age    int
		Height int
	}{}

	data := map[string]interface{}{
		"Name":   "nil",
		"DOB":    "22.11.1998",
		"Weight": "68.5",
		"Age":    22,
		"Height": 160,
	}

	ok := transform.MapToStruct(user, data)

	if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(user)
	}
}
