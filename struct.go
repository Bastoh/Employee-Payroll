// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var name string
// 	var id int
// 	var itemName string
// 	var price float64
// 	var quantity int
// 	var discount float64

// 	// Collect input
// 	fmt.Print("Enter customer name: ")
// 	fmt.Scanln(&name)

// 	fmt.Print("Enter customer ID: ")
// 	fmt.Scanln(&id)

// 	fmt.Print("Enter item name: ")
// 	fmt.Scanln(&itemName)

// 	fmt.Print("Enter price per unit: ")
// 	fmt.Scanln(&price)

// 	fmt.Print("Enter quantity: ")
// 	fmt.Scanln(&quantity)

// 	fmt.Print("Enter discount percentage: ")
// 	fmt.Scanln(&discount)

// 	// Create structs
// 	customer := Customer{
// 		Name: name,
// 		ID:   id,
// 	}

// 	item := Item{
// 		Name:     itemName,
// 		Price:    price,
// 		Quantity: quantity,
// 	}

// 	// Calculate subtotal
// 	subtotal, err := item.GenerateTotal()
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	invoice := Invoice{
// 		Customer: customer,
// 		Item:     item,
// 		Subtotal: subtotal,
// 		Discount: discount,
// 	}

// 	// Apply discount
// 	if err := invoice.ApplyDiscount(); err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	// Print invoice
// 	invoice.PrintInvoice()
// }
