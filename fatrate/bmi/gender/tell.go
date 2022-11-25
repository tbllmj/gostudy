package gender

import "fmt"

func BmiMaleTell(bmiRate float64, age uint) {
	if age >= 18 && age <= 39 {
		if bmiRate > 0 && bmiRate <= 1.0 {
			fmt.Println("偏廋")
		} else if bmiRate > 1.0 && bmiRate <= 1.6 {
			fmt.Println("标准")
		} else if bmiRate > 1.6 && bmiRate <= 2.1 {
			fmt.Println("偏重")
		} else if bmiRate > 2.1 && bmiRate <= 2.6 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("严重肥胖")
		}
	} else if age >= 40 && age <= 59 {
		if bmiRate > 0 && bmiRate <= 1.1 {
			fmt.Println("偏廋")
		} else if bmiRate > 1.1 && bmiRate <= 1.7 {
			fmt.Println("标准")
		} else if bmiRate > 1.7 && bmiRate <= 2.2 {
			fmt.Println("偏重")
		} else if bmiRate > 2.2 && bmiRate <= 2.7 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("严重肥胖")
		}
	} else {
		if bmiRate > 0 && bmiRate <= 1.3 {
			fmt.Println("偏廋")
		} else if bmiRate > 1.3 && bmiRate <= 1.9 {
			fmt.Println("标准")
		} else if bmiRate > 1.9 && bmiRate <= 2.4 {
			fmt.Println("偏重")
		} else if bmiRate > 2.4 && bmiRate <= 2.9 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("严重肥胖")
		}
	}
}
func BmiFemaleTell(bmiRate float64, age uint) {
	if age >= 18 && age <= 39 {
		if bmiRate > 0 && bmiRate <= 2.0 {
			fmt.Println("偏廋")
		} else if bmiRate > 2.0 && bmiRate <= 2.7 {
			fmt.Println("标准")
		} else if bmiRate > 2.7 && bmiRate <= 3.4 {
			fmt.Println("偏重")
		} else if bmiRate > 3.4 && bmiRate <= 3.9 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("严重肥胖")
		}
	} else if age >= 40 && age <= 59 {
		if bmiRate > 0 && bmiRate <= 2.1 {
			fmt.Println("偏廋")
		} else if bmiRate > 2.1 && bmiRate <= 2.8 {
			fmt.Println("标准")
		} else if bmiRate > 2.8 && bmiRate <= 3.5 {
			fmt.Println("偏重")
		} else if bmiRate > 3.5 && bmiRate <= 4.0 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("严重肥胖")
		}
	} else {
		if bmiRate > 0 && bmiRate <= 2.2 {
			fmt.Println("偏廋")
		} else if bmiRate > 2.2 && bmiRate <= 2.9 {
			fmt.Println("标准")
		} else if bmiRate > 2.9 && bmiRate <= 3.6 {
			fmt.Println("偏重")
		} else if bmiRate > 3.6 && bmiRate <= 4.1 {
			fmt.Println("肥胖")
		} else {
			fmt.Println("严重肥胖")
		}
	}
}
