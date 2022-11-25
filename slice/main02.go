package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main02() {
	var poker []string = []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	var poker2 []string
	var pokerLen int = len(poker)
	rand.Seed(time.Now().UnixNano())
oneFor:
	for true {
		s := poker[rand.Intn(pokerLen)]
		for _, value := range poker2 {
			if value == s {
				continue oneFor
			}
		}
		poker2 = append(poker2, s)
		if len(poker2) == pokerLen {
			break
		}

	}
	fmt.Println(poker2)
}
