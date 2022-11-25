package main

import "fmt"

func main() {
	var student map[string]int = map[string]int{"张三": 12, "李四": 23, "王麻子": 223, "赵四": 10, "田斌": 23}
	var studentLen = len(student)
	var total int = 0
	for _, i := range student {
		total += i
	}
	var avg int = total / studentLen
	fmt.Println(avg)
}
