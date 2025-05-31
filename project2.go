// package main 

// import "fmt"

// type Customer struct {
// 	Name string
// 	ID int
// }

// type Book struct {
// 	Title string
// 	Author string
// 	// rentel per day - RPD
// 	RPD int
// 	// No of days - NOD
// 	NOD int
// }

// type ErrorMessage struct {
// 	message string
// }

// func (em ErrorMessage) Error() string {
// 	return em.message
// }

// // day or price - DOP
// func DOPError() error{
// 	return &ErrorMessage{
// 		message: "Invalid day or price!!!",
// 	}
// }

// func (b Book) GetSubFee() (float64, error) {
// 	if b.NOD < 0 || b.RPD < 0 {
// 		return 0.0, DOPError()
// 	}
// 	return float64(b.NOD) * float64(b.RPD), nil
// }

// func PenaltyError() error{
// 	return &ErrorMessage{
// 		message: "Invalid percentage for penalty fee.",
// 	}
// }

// type Receipt struct {
// 	Customer
// 	Book
// 	SubtotalFee float64
// 	PenaltyFee float64
// 	TotalFee float64
// }

// func (r Receipt) ApplyPenaltyFee() error {
// 	if r.PenaltyFee < 0 || r.PenaltyFee > 50 {
// 		return PenaltyError()
// 	}
// 	r.TotalFee = r.SubtotalFee * (1 + r.PenaltyFee/100)
// 	return nil
// }

// func (r Receipt) PrintReceipt() {
// 	fmt.Println("\n--- Reciept ---")
// 	fmt.Printf("Customer: %s (ID: %d)\n", r.Name, r.ID)
// 	fmt.Printf("Book: %s\n", r.Author)
// 	fmt.Printf("Price per day: $%.2f\n", float64(r.RPD))
// 	fmt.Printf("Days rented: %d\n", r.NOD)
// 	fmt.Printf("Subtotal: $%.2f\n", r.SubtotalFee)
// 	fmt.Printf("Penalty: %.2f%%\n", r.PenaltyFee)
// 	fmt.Printf("Total: $%.2f\n", r.TotalFee)
// 	fmt.Println("----------------")
// 	fmt.Println("Enjoy your reading!")
// }
// func main() {
// 	var name string
// 	var id int 
// 	var title string 
// 	var author string
// 	var rpd int 
// 	var nod int 
// 	var penaltyfee float64

// 	fmt.Print("Enter customer name:")
// 	fmt.Scanln(&name)

// 	fmt.Print("Enter library ID:")
// 	fmt.Scanln(&id)

// 	fmt.Print("Enter book title:")
// 	fmt.Scanln(&title)

// 	fmt.Print("Enter book author:")
// 	fmt.Scanln(&author)

// 	fmt.Print("Enter rental price per day:")
// 	fmt.Scanln(&rpd)

// 	fmt.Print("Enter number of rental days:")
// 	fmt.Scanln(&nod)

// 	fmt.Print("Enter penalty percentage (if late):")
// 	fmt.Scanln(&penaltyfee)

// 	customer := Customer{
// 		Name: name,
// 		ID: id,
// 	}

// 	book := Book{
// 		Title: title,
// 		Author: author,
// 		RPD: rpd,
// 		NOD: nod,
// 	}
	
// 	subtotal, err := book.GetSubFee()
// 	if err != nil {
// 		fmt.Println("Error", err)
// 		return
// 	}

// 	receipt := Receipt{
// 		Customer: customer,
// 		Book: book,
// 		PenaltyFee: penaltyfee ,
// 		SubtotalFee: subtotal,
// 	}

// 	if err := receipt.ApplyPenaltyFee(); err != nil{
// 		fmt.Println("Error", err)
// 		return
// 	}

// 	receipt.PrintReceipt()
// }