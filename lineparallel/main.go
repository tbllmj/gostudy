package main

import "fmt"

func main() {
	var line1 [2]float64 = [2]float64{2, 2}
	var line2 [2]float64 = [2]float64{3, 3}
	var slope1 float64 = line1[0] / line1[1]
	var slope2 float64 = line2[0] / line2[1]
	if slope2 == slope1 {
		fmt.Println("两条直线平行")
	} else {
		fmt.Println("两条直线不平行")
	}
}
