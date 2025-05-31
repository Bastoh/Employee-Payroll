// package main

// import "fmt"

// type ErrorMessage struct {
// 	message string
// }

// func (em ErrorMessage) Error() string {
// 	return em.message
// }

// func PriceError() error{
// 		return &ErrorMessage{
// 			message: "Invalid price or quantity",
// 		}
// 	}

// type Customer struct {
// 	Name string 
// 	ID int
// }

// type Item struct{
// 	Name string
// 	Price float64
// 	Quantity int
// }

// func (i *Item) GenerateTotal() (float64, error){
// 	if i.Price <= 0 || i.Quantity <= 0 {
// 		return 0.0, PriceError()
// 	}
// 	return float64(i.Quantity) * i.Price, nil
// }

// type Invoice struct {
// 	Customer
// 	Item
// 	Subtotal float64
// 	Discount float64
// 	Total float64
// }

// func DiscountError() error{
// 	return &ErrorMessage{
// 		message: "Invalid discount percentage",
// 	}
// }

// func (in *Invoice) ApplyDiscount() (error) {
// 	if in.Discount < 0 || in.Discount > 100 {
// 		return DiscountError()
// 	}
// 	in.Total = in.Subtotal * (1 - in.Discount/100)
// 	return nil
// }

// func (in *Invoice) PrintInvoice() {
// 	fmt.Println("\n--- Invoice ---")
// 	fmt.Printf("Customer: %s (ID: %d)\n", in.Customer.Name, in.Customer.ID)
// 	fmt.Printf("Item: %s\n", in.Item.Name)
// 	fmt.Printf("Price per unit: $%.2f\n", in.Item.Price)
// 	fmt.Printf("Quantity: %d\n", in.Item.Quantity)
// 	fmt.Printf("Subtotal: $%.2f\n", in.Subtotal)
// 	fmt.Printf("Discount: %.2f%%\n", in.Discount)
// 	fmt.Printf("Total: $%.2f\n", in.Total)
// 	fmt.Println("----------------")
// 	fmt.Println("Thank you for your purchase!")
// }
