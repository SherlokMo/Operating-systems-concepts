package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

var a int = 0

func test(str string) {

	for i := 0; i < 1000000; i++ {
		mu.Lock()
		a++
		mu.Unlock()
	}
	fmt.Println("Finished " + str)
}

func main() {

	go test("a")
	go test("b")

	fmt.Scanln()

	fmt.Println(a)
}
