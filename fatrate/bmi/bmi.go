package bmi

func Bmi(weight float64, height float64) float64 {
	return weight / (height * height)
}

func BmiRate(bmi float64, age uint, gender uint8) float64 {
	return 1.2*bmi + 0.23*float64(age) - 5.4 - 10.8*float64(gender)/100
}
