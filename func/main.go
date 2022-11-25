package main

import "fmt"

func init() {
	fmt.Println(2322323)
}
func main() {
	var testF func(i int)
	if true {
		testF = test
	} else {
		testF = test2
	}
	testF(2)
	str("hello", "world")
}
func test(i int) {
	fmt.Println("test", i)
}
func test2(i int) {
	fmt.Println("test2", i)
}
func str(s ...string) {
	fmt.Println("str：", s)
}
func init() {
	fmt.Println("测试")
}
