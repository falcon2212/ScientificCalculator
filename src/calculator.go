package main

import (
	"math"
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
