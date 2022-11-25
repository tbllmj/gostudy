package main

import "fmt"

func main01() {
	var number []int = []int{1, 43, 5, 6, 7, 343, 254}
	var numberLen int = len(number)
	//var sortNumber []int
	for i := 0; i < numberLen; i++ {
		for j := 0; j < numberLen-i-1; j++ {
			if number[j] < number[j+1] {
				var temp int = number[j]
				number[j] = number[j+1]
				number[j+1] = temp
			}
		}
	}
	fmt.Println(number)
}
