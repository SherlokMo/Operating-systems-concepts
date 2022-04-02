package main

import "fmt"

func main() {
	fmt.Printf("main routine virtual adress %p \n", main)
	fmt.Printf("Heap location is %p \n", &[]int{1, 2, 3, 4})
	x := 5
	fmt.Printf("Stack location of x is %p \n", &x)
}
