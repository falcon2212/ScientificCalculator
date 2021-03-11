package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Scientific calculator!")
	reader := bufio.NewReader(os.Stdin)
	var C Calculator
	for true {
		fmt.Println("Select the operation that you would like to perform:")
		fmt.Println("1. Sqrt")
		fmt.Println("2. Factorial")
		fmt.Println("3. Natural log")
		fmt.Println("4. Power")
		fmt.Println("5. Exit")
		var quit bool = false
		response, _ := reader.ReadString('\n')
		switch response {
		case "1\n":
			fmt.Print("Enter the number you want compute sqrt for: ")
			strparam, _ := reader.ReadString('\n')
			strparam = strings.TrimSuffix(strparam, "\n")
			param, _ := strconv.ParseFloat(strparam, 64)
			var result float64 = C.Sqrt(param)
			fmt.Println("Result:", result)
		case "2\n":
			fmt.Print("Enter the number you want compute factorial for: ")
			strparam, _ := reader.ReadString('\n')
			strparam = strings.TrimSuffix(strparam, "\n")
			param, _ := strconv.ParseInt(strparam, 10, 64)
			var result int64 = C.Factorial(param)
			fmt.Println("Result:", result)
		case "3\n":
			fmt.Print("Enter the number you want compute natural log for: ")
			strparam, _ := reader.ReadString('\n')
			strparam = strings.TrimSuffix(strparam, "\n")
			param, _ := strconv.ParseFloat(strparam, 64)
			var result float64 = C.NaturalLog(param)
			fmt.Println("Result:", result)
		case "4\n":
			fmt.Print("Enter the base: ")
			strparam1, _ := reader.ReadString('\n')
			strparam1 = strings.TrimSuffix(strparam1, "\n")
			param1, _ := strconv.ParseFloat(strparam1, 64)
			fmt.Print("Enter the exponent: ")
			strparam2, _ := reader.ReadString('\n')
			strparam2 = strings.TrimSuffix(strparam2, "\n")
			param2, _ := strconv.ParseFloat(strparam2, 64)
			var result float64 = C.Power(param1, param2)
			fmt.Println("Result:", result)
		case "5\n":
			fmt.Println("Thank you for using this calculator")
			quit = true
			break
		}
		if quit {
			break
		}
	}

}
