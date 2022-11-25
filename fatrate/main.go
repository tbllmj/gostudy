package main

import (
	"fmt"
	"github.com/tbllmj/gostudy/fatrate/bmi"
	genderF "github.com/tbllmj/gostudy/fatrate/bmi/gender"
)

func main() {
	const GENDER_MALE uint8 = 1   //男
	const GENDER_FEMALE uint8 = 0 //女
	for true {
		var name string
		fmt.Print("姓名：")
		_, err2 := fmt.Scan(&name)
		if err2 != nil {
			panic(err2)
			return
		}
		var weight float64
		fmt.Print("体重（KG）：")
		_, err := fmt.Scan(&weight)
		if err != nil {
			panic(err)
			return
		}
		var height float64
		fmt.Print("身高（cm）：")
		_, err = fmt.Scan(&height)
		if err != nil {
			panic(err)
			return
		}
		fmt.Print("年龄：")
		var age uint
		_, err = fmt.Scan(&age)
		if err != nil {
			panic(err)
			return
		}
		var genderName string
		fmt.Print("性别（男/女）")
		_, err = fmt.Scan(&genderName)
		if err != nil {
			panic(err)
			return
		}
		var gender uint8 = 0
		if genderName == "男" {
			gender = GENDER_FEMALE
		}

		var bmiVar float64 = bmi.Bmi(weight, height)
		var bmiRare float64 = bmi.BmiRate(bmiVar, age, gender)
		if GENDER_FEMALE == gender {
			genderF.BmiFemaleTell(bmiRare, age)
		} else {
			genderF.BmiMaleTell(bmiRare, age)
		}
		var continueText string
		fmt.Print("是否继续输入下一个（Y/N）")
		_, err = fmt.Scan(&continueText)
		if continueText == "N" {
			return
		}
	}

}
