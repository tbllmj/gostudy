package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"gostudy/fatrate/bmi"
	genderF "gostudy/fatrate/bmi/gender"
)

func main() {
	var name, sex string
	var tall, weight float64
	var age uint8

	cmd := &cobra.Command{
		Use:   "fatrate",
		Short: "计算个人体脂",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("姓名：", name)
			fmt.Println("性别：", sex)
			fmt.Println("身高：", tall)
			fmt.Println("体重：", weight)
			fmt.Println("年龄：", age)
			var bmiVar float64 = bmi.Bmi(weight, tall)
			var sexUint uint8
			if sex == "男" {
				sexUint = 1
			} else {
				sexUint = 0
			}
			var bmiRare float64 = bmi.BmiRate(bmiVar, uint(age), sexUint)
			if 1 == sexUint {
				genderF.BmiFemaleTell(bmiRare, uint(age))
			} else {
				genderF.BmiMaleTell(bmiRare, uint(age))
			}
		},
	}

	cmd.Flags().StringVar(&name, "name", "", "姓名")
	cmd.Flags().StringVar(&sex, "sex", "", "性别")
	cmd.Flags().Float64Var(&tall, "tall", 0.0, "身高")
	cmd.Flags().Float64Var(&weight, "weight", 0.0, "体重")
	cmd.Flags().Uint8Var(&age, "age", 0.0, "年龄")
	err := cmd.Execute()
	if err != nil {
		return
	}
}
