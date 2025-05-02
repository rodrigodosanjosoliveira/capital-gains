package main

import (
	"fmt"

	"github.com/rodrigodosanjosoliveira/capital-gains/internal/calculator"
	"github.com/rodrigodosanjosoliveira/capital-gains/internal/io"
)

func main() {
	operations, err := io.ReadInputFlexible()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	taxes := calculator.CalculateCapitalGains(operations)

	err = io.WriteOutput(taxes)
	if err != nil {
		fmt.Println("Error generating output:", err)
		return
	}
}
