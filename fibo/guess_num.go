package main

import "fmt"

func main() {
	guessNum(1, 100)
}
func guessNum(from uint, to uint) {
	middle := (from + to) / 2
	fmt.Println("你猜的数字是：", middle, "?")
	var guessFlag uint8
	fmt.Print("高了，请输入2，低了，请输入1，正确请输入0：")
	fmt.Scan(&guessFlag)
	switch guessFlag {
	case 2:
		middle = middle - 1
		guessNum(from, middle)
		if from == to {
			fmt.Println("你耍赖！！")
		}
	case 1:
		middle = middle + 1
		guessNum(middle, to)
		if from == to {
			fmt.Println("你耍赖！！")
		}
	case 0:
		fmt.Println("猜对了")
		return
	default:
	}

}
