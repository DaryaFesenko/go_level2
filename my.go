package main

import "fmt"

func nmain() {

}

func h() {
	for i := 0; i < 20; i++ {
		go fmt.Println(i)

		for i := 0; i < 20; i++ {
			go fmt.Println(i)

		}

		for i := 0; i < 20; i++ {
			go fmt.Println(i)
		}
	}

	for i := 0; i < 20; i++ {
		go fmt.Println(i)
	}

	go fmt.Println(7)
}

func f() {
	go fmt.Println(7)
}
