package main

import "fmt"

func main() {
	var transactions []float64
	for {
		transaction := scanTransaction()
		if transaction == 0.0 {
			break
		}
		transactions = append(transactions, transaction)
	}
	fmt.Printf("Result: %.2f", transactions)
}

func scanTransaction() float64 {
	var transaction float64
	fmt.Print("Enter transaction (N to exit): ")
	fmt.Scan(&transaction)
	return transaction
}
