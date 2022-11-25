package main

import (
	"fmt"
	"strings"
)

func main() {
	var strSlice []string = []string{"中国", "english", "国家", "a"}
	for i, s := range strSlice {
		strSlice[i] = strings.ToUpper(s)
		//if len([]byte(s)) == len([]rune(s)) { //说明不含有中文
		//	strSlice[i] = strings.ToUpper(s)
		//}
	}
	fmt.Println(strSlice)
}
