package main

import "fmt"

func main01() {
	fmt.Println(fibonacci(50))
}
func fibonacci(i uint) uint {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
