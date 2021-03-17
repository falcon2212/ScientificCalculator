package main

import (
	"fmt"
	"math"
	"strconv"
)

type Calculator struct {
	CurrentState float64
}

func (C Calculator) Sqrt(param1 float64) (result float64) {
	result = float64(math.Sqrt(param1))
	C.CurrentState = result
	return result
}
func (C Calculator) Factorial(param1 int64) (result int64) {
	result = 1
	for i := int64(1); i <= param1; i++ {
		result = result * i
	}
	C.CurrentState = float64(result)
	return result
}
func (C Calculator) NaturalLog(param1 float64) (result float64) {
	result = float64(math.Log(param1))
	C.CurrentState = result
	return result
}
func (C Calculator) Power(b float64, exp float64) (result float64) {
	result = float64(math.Pow(b, exp))
	C.CurrentState = result
	return result
}
func main() {
	fmt.Println("Welcome to the Scientific Calculator!")
	var C Calculator
	for true {
		fmt.Println("Select the operation that you would like to perform:")
		fmt.Println("1. Sqrt")
		fmt.Println("2. Factorial")
		fmt.Println("3. Natural log")
		fmt.Println("4. Power")
		fmt.Println("5. Exit")
		var quit bool = false
		var response string
		fmt.Scanln(&response)
		switch response {
		case "1":
			fmt.Print("Enter the number you want compute sqrt for: ")
			var strparam string
			fmt.Scanln(&strparam)
			param, _ := strconv.ParseFloat(strparam, 64)
			var result float64 = C.Sqrt(param)
			fmt.Println("Result:", result)
		case "2":
			fmt.Print("Enter the number you want compute factorial for: ")
			var strparam string
			fmt.Scanln(&strparam)
			param, _ := strconv.ParseInt(strparam, 10, 64)
			var result int64 = C.Factorial(param)
			fmt.Println("Result:", result)
		case "3":
			fmt.Print("Enter the number you want compute natural log for: ")
			var strparam string
			fmt.Scanln(&strparam)
			param, _ := strconv.ParseFloat(strparam, 64)
			var result float64 = C.NaturalLog(param)
			fmt.Println("Result:", result)
		case "4":
			fmt.Print("Enter the base: ")
			var strparam1 string
			fmt.Scanln(&strparam1)
			param1, _ := strconv.ParseFloat(strparam1, 64)
			fmt.Print("Enter the exponent: ")
			var strparam2 string
			fmt.Scanln(&strparam2)
			param2, _ := strconv.ParseFloat(strparam2, 64)
			var result float64 = C.Power(param1, param2)
			fmt.Println("Result:", result)
		case "5":
			fmt.Println("Thank you for using this calculator")
			quit = true
			break
		}
		if quit {
			break
		}
	}

}
