package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	num1 := flag.Float64("num1", 3, "The first number")
	num2 := flag.Float64("num2", 6, "The Second number")
	op := flag.String("op", "+", "The Operator('+','-','*','/')")

	flag.Parse()

	var result float64

	switch *op {
	case "+":
			result = *num1 + *num2
	case "-":
			result = *num1 - *num2
	case "*":
			result = *num1 * *num2
	case "/":
			if *num2 == 0 {
				fmt.Fprintln(os.Stderr, "Error: Divisible by zero is not allowed.")
				os.Exit(1)
			}
			result = *num1 / *num2
	default:
			fmt.Fprintf(os.Stderr, "Error: Invalid operator %s. use +, -, *, /", *op)
			os.Exit(1)
	}
	fmt.Printf("Result: %g\n", result)
}