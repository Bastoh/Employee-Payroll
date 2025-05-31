// package main

// import "fmt"

// type Customer struct {
// 	CusName string
// 	CusID int
// }

// type Item struct {
// 	ItName string
// 	Price float64
// 	Quantity int
// }

// type Receipt struct {
// 	Customer
// 	Items []Item
// 	Discount int
// 	Subtotal float64
// 	Total float64
// }

// type ErrorMessage struct {
// 	message string
// }

// func (em *ErrorMessage) Error() string{
// 	return em.message
// }

// func PriceError() error{
// 	return &ErrorMessage{
// 		message: "Invalid price!!!",
// 	}
// }

// func QuantityError() error{
// 	return &ErrorMessage{
// 		message: "Invalid Quantity!!!",
// 	}
// }

// func DiscountError() error{
// 	return &ErrorMessage{
// 		message: "Discount can't be negative!!!",
// 	}
// }

// func (i *Item) SubtotalCost() (float64, error) {
// 	if i.Price <= 0.0 {
// 		return 0.0,PriceError()
// 	} else if i.Quantity <= 0 {
// 		return 0.0, QuantityError()
// 	}
// 	return i.Price * float64(i.Quantity), nil
// }

// func (r *Receipt) ApplyDiscount() error {
// 	if r.Discount <= 0 {
// 		return DiscountError()
// 	}
// 	r.Total = r.Subtotal * (1 - float64(r.Discount/100))
// 	return nil
// }

// func (r *Receipt) PrintReceipt() {
// 	fmt.Println("\nGROCERY SHOPPING RECIEPT")
// 	fmt.Println("----------------------------------")
// 	fmt.Printf("Customer: %s\n", r.CusName)
// 	fmt.Printf("Customer ID: %d\n", r.CusID)

// 	fmt.Println("Items Purchased:")
// 	for i, item := range r.Items {
// 		itemSubtotal := float64(item.Quantity) * item.Price
// 		fmt.Printf("%d. %-15s Qty: %-3d Price: $%-6.2f Subtotal: $%.2f\n",
// 			i+1, item.ItName, item.Quantity, item.Price, itemSubtotal)
// 	}

// 	fmt.Println("-----------------------------------------------")
// 	fmt.Printf("Subtotal: \n", r.Subtotal)
// 	fmt.Printf("Discount: (%d%): -$.2f \n", r.Discount, 1 - float64(r.Discount/100))
// 	fmt.Printf("TOTAL: \n", r.Total)
// 	fmt.Println("----------------------------------------------------")
// 	fmt.Println("Thank you for shopping with us!")
// }

// func main() {
// 	var name string
// 	var Id int
// 	var itemname string
// 	var price float64
// 	var quantity int
// 	var discount int

// 	fmt.Print("Enter Customer's Name: ")
// 	fmt.Scanln(&name)

// 	fmt.Print("Enter Customer's ID: ")
// 	fmt.Scanln(&Id)

// 	var items []Item
// 	var more string
// 	for {
// 		fmt.Print("Enter item name: ")
// 		fmt.Scanln(&itemname)

// 		fmt.Print("Enter item price: ")
// 		fmt.Scanln(&price)

// 		fmt.Print("Enter item quantity: ")
// 		fmt.Scanln(&quantity)
		
// 		items = append(items, Item{
// 			ItName: itemname,
// 			Price: price,
// 			Quantity: quantity,
// 		})

// 		fmt.Println("Do you want to add more items: ")
// 		fmt.Scanln(&more)

// 		if more != "yes"{
// 			break
// 		}
// 	}

// 	fmt.Print("Enter Discount: ")
// 	fmt.Scanln(&discount)

// 	customer := Customer{
// 		CusName: name,
// 		CusID: Id,
// 	}
	
// 	item := Item{
// 		ItName: itemname,
// 		Price: price,
// 		Quantity: quantity,
// 	}

// 	subtotal, err := item.SubtotalCost()
// 	if err != nil{
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	receipt := Receipt{
// 		Customer: customer,
// 		Items: items,
// 		Discount: discount,
// 		Subtotal: subtotal,
// 	}
	
// 	if err := receipt.ApplyDiscount(); err != nil {
// 		fmt.Println("Error: ", err)
// 		return
// 	}

// 	receipt.PrintReceipt()
// }