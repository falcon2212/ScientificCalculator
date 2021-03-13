package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqrt(t *testing.T) {
	var C Calculator
	assert := assert.New(t)
	var truePositiveTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{4.0, 2.0},
	}
	var trueNegativeTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{4.0, 1.0},
	}
	var falsePositiveTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{4.0, 1.0},
	}
	var falseNegativeTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{4.0, 2.0},
	}
	// Positive Testing
	for _, test := range truePositiveTests {
		assert.Equal(C.Sqrt(test.input), test.expectedOutput)
	}
	for _, test := range trueNegativeTests {
		assert.NotEqual(C.Sqrt(test.input), test.expectedOutput)
	}
	// Negative Testing
	for _, test := range falsePositiveTests {
		assert.NotEqual(C.Sqrt(test.input), test.expectedOutput)
	}
	for _, test := range falseNegativeTests {
		assert.Equal(C.Sqrt(test.input), test.expectedOutput)
	}
}
func TestFactorial(t *testing.T) {
	var C Calculator
	assert := assert.New(t)
	var truePositiveTests = []struct {
		input          int64
		expectedOutput int64
	}{
		{4, 24},
	}
	var trueNegativeTests = []struct {
		input          int64
		expectedOutput int64
	}{
		{4, 120},
	}
	var falsePositiveTests = []struct {
		input          int64
		expectedOutput int64
	}{
		{4, 120},
	}
	var falseNegativeTests = []struct {
		input          int64
		expectedOutput int64
	}{
		{4, 24},
	}
	// Positive Testing
	for _, test := range truePositiveTests {
		assert.Equal(C.Factorial(test.input), test.expectedOutput)
	}
	for _, test := range trueNegativeTests {
		assert.NotEqual(C.Factorial(test.input), test.expectedOutput)
	}
	// Negative Testing
	for _, test := range falsePositiveTests {
		assert.NotEqual(C.Factorial(test.input), test.expectedOutput)
	}
	for _, test := range falseNegativeTests {
		assert.Equal(C.Factorial(test.input), test.expectedOutput)
	}
}
func TestNaturalLog(t *testing.T) {
	var C Calculator
	assert := assert.New(t)
	var truePositiveTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{10.0, 2.302585092994046},
	}
	var trueNegativeTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{10.0, 1.0},
	}
	var falsePositiveTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{10.0, 1.0},
	}
	var falseNegativeTests = []struct {
		input          float64
		expectedOutput float64
	}{
		{10.0, 2.302585092994046},
	}
	// Positive Testing
	for _, test := range truePositiveTests {
		assert.Equal(C.NaturalLog(test.input), test.expectedOutput)
	}
	for _, test := range trueNegativeTests {
		assert.NotEqual(C.NaturalLog(test.input), test.expectedOutput)
	}
	// Negative Testing
	for _, test := range falsePositiveTests {
		assert.NotEqual(C.NaturalLog(test.input), test.expectedOutput)
	}
	for _, test := range falseNegativeTests {
		assert.Equal(C.NaturalLog(test.input), test.expectedOutput)
	}
}
func TestPower(t *testing.T) {
	var C Calculator
	assert := assert.New(t)
	var truePositiveTests = []struct {
		input struct {
			base float64
			exp  float64
		}
		expectedOutput float64
	}{
		{
			struct {
				base float64
				exp  float64
			}{2.0, 10.0},
			1024.0},
	}
	var trueNegativeTests = []struct {
		input struct {
			base float64
			exp  float64
		}
		expectedOutput float64
	}{
		{
			struct {
				base float64
				exp  float64
			}{2.0, 10.0},
			256.0},
	}
	var falsePositiveTests = []struct {
		input struct {
			base float64
			exp  float64
		}
		expectedOutput float64
	}{
		{
			struct {
				base float64
				exp  float64
			}{2.0, 10.0},
			256.0},
	}
	var falseNegativeTests = []struct {
		input struct {
			base float64
			exp  float64
		}
		expectedOutput float64
	}{
		{
			struct {
				base float64
				exp  float64
			}{2.0, 10.0},
			1024.0},
	}
	// Positive Testing
	for _, test := range truePositiveTests {
		assert.Equal(C.Power(test.input.base, test.input.exp), test.expectedOutput)
	}
	for _, test := range trueNegativeTests {
		assert.NotEqual(C.Power(test.input.base, test.input.exp), test.expectedOutput)
	}
	// Negative Testing
	for _, test := range falsePositiveTests {
		assert.NotEqual(C.Power(test.input.base, test.input.exp), test.expectedOutput)
	}
	for _, test := range falseNegativeTests {
		assert.Equal(C.Power(test.input.base, test.input.exp), test.expectedOutput)
	}
}
