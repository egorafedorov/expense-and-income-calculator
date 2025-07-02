package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	var transactions []float64
	for {
		transaction := scanTransaction()
		if transaction == 0.0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	result := calculateBalance(transactions)
	color.Blue("Result: %.2f", result)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Enter transaction (N to exit): ")
	fmt.Scan(&transaction)
	return transaction
}

func calculateBalance(arr []float64) float64 {
	var balance float64
	for _, value := range arr {
		balance = balance + value
	}
	return balance
}
