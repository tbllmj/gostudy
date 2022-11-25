package main

import "fmt"

func main() {
	var arr [5]int = [5]int{1, 4, 5, 3, 7}
	var newArr [5]int = [5]int{}
	var arrLen int = len(arr)
	var index int = 0
	for i := arrLen - 1; i > -1; i-- {
		fmt.Println(i)
		newArr[index] = arr[i]
		index++
	}
	fmt.Println(newArr)
}
