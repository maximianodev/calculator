package calculator

var logMessage = "[LOG]"

// package version
var Version = "1.0"

func internalSum(number int) int {
	return number - 1
}

// sum of two numbers
func Sum(number1, number2 int) int {
	return number1 + number2
}
