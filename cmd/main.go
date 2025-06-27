package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/ASPICE/calculator"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: calc <operation> <number1> <number2>")
		fmt.Println("Operations are: add, subtract, multiply, divide")
		os.Exit(1)
	}

	operation := os.Args[1]
	a, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("Invalid number: %s", os.Args[2])
	}

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Fatalf("Invalid number: %s", os.Args[3])
	}

	switch operation {
	case "add":
		fmt.Printf("%d + %d = %d\n", a, b, calculator.Add(a, b))
	case "subtract":
		fmt.Printf("%d - %d = %d\n", a, b, calculator.Subtract(a, b))
	case "multiply":
		fmt.Printf("%d * %d = %d\n", a, b, calculator.Multiply(a, b))
	case "divide":
		result, err := calculator.Divide(a, b)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		fmt.Printf("%d / %d = %d\n", a, b, result)
	default:
		log.Fatalf("Unknown operation: %s", operation)
	}
}
