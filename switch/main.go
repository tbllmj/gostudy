package main

import (
	"fmt"
)

func main() {
	var money interface{} = 20.65
	switch newMoney := money.(type) {
	case int:
		tmpMoney := newMoney + 10.0
		fmt.Println("int money 是", tmpMoney)

	case int64:
		tmpMoney := newMoney + 10.0
		fmt.Println("int64 money 是", tmpMoney)
	case float32:
		tmpMoney := newMoney + 10.111
		fmt.Println("float32 money 是", tmpMoney)
	case float64:
		tmpMoney := newMoney + 10.111
		fmt.Println("float64 money 是", tmpMoney)

	default:
		fmt.Println("未知类型")
	}
}
