package main

import (
	"fmt"
	"os"
)

type Employee struct {
	Name string
	ID int
	Department string
	Salary
}

type Salary struct {
	BaseSalary float64
	BonusPercentage int
}

type ErrorMessage struct{
	Message string
}

func (em *ErrorMessage) Error() string{
	return em.Message
}

func InvalidSalary() error{
	return &ErrorMessage{
		Message: "The salary can't be negative!",
	}
}

func InvalidBonus() error{
	return &ErrorMessage{
		Message: "The bonus can't be negative!",
	}
}

func (e *Employee) ApplyBonus() error {
	if e.BonusPercentage < 0 {
		return InvalidBonus()
	}
	e.BaseSalary += (e.BaseSalary * float64(e.BonusPercentage)/100)
	return nil
}

func (e *Employee) CalculateNetSalary() (float64, error) {
	if e.BaseSalary < 0 {
		return 0.0, InvalidSalary()
	}
	taxRate := 0.05
	if e.BaseSalary > 30000.0 {
		taxRate = 0.1
	}
	return e.BaseSalary * (1 - taxRate), nil
}

func (e *Employee) PrintPaySlip() {
    netSalary, err := e.CalculateNetSalary()
    if err != nil {
        fmt.Printf("\nERROR processing %s's payslip: %v\n", e.Name, err)
        return
    }
    
    taxAmount := e.BaseSalary * 0.05  // Default 5% tax
    if e.BaseSalary > 30000.0 {
        taxAmount = e.BaseSalary * 0.1  // 10% for higher salaries
    }
    
    fmt.Printf("\n=== PAYSLIP ===\n")
    fmt.Printf("Name: %s\n", e.Name)
    fmt.Printf("Department: %s\n", e.Department)
    fmt.Printf("Base Salary: $%.2f\n", e.BaseSalary)
    fmt.Printf("Tax Deducted: $%.2f\n", taxAmount)
    fmt.Printf("Net Salary: $%.2f\n", netSalary)
    fmt.Printf("================\n")
}

func main() {
    employees := []Employee{}
    
    for {
        fmt.Println("\nEmployee Payroll System")
        fmt.Println("1. Add Employee")
        fmt.Println("2. Apply Bonus")
        fmt.Println("3. Generate single payslip using ID")
        fmt.Println("4. Generate Payslips")
        fmt.Println("5. Exit")
        
        var choice int
        fmt.Print("Select option: ")
        fmt.Scanln(&choice)
        
        switch choice {
        case 1:
            var emp Employee
			var basesalary int

			fmt.Print("\nEnter employee's Name: ")
			fmt.Scanln(&emp.Name)

			fmt.Print("\nEnter employee's ID: ")
			fmt.Scanln(&emp.ID)

			fmt.Print("\nEnter employee's Department: ")
			fmt.Scanln(&emp.Department)

			fmt.Print("\nEnter employee's Base Salary: ")
			fmt.Scanln(&basesalary)
			emp.BaseSalary = float64(basesalary)

			fmt.Print("\nEnter employee's Bonus: ")
			fmt.Scanln(&emp.BonusPercentage)

			employees = append(employees, emp)

			fmt.Printf("Successfully added %s\n", emp.Name)

        case 2:
            if len(employees) == 0 {
				fmt.Println("No employees found!!")
				continue
			}

			var id int
			fmt.Print("\n Enter employee's ID: ")
			fmt.Scanln(&id)
			for i := range employees{
				if employees[i].ID == id{
				err := employees[i].ApplyBonus()
					if err != nil{
						fmt.Printf("Error applying bonus: %v", err)
					}else{
						fmt.Printf("Applied %d%% bonus to %s. New base salary: $%.2f\n",
					employees[i].BonusPercentage,
					employees[i].Name,
					employees[i].BaseSalary,
					)
						}
					break
				}
			}
		
		case 3:
			if len(employees) == 0 {
				fmt.Println("No employees found!!!")
				continue
			}
			var id int
			fmt.Print("\nEnter employee's ID: ")
			fmt.Scanln(&id)

			for i := range employees {
				if employees[i].ID == id {
					employees[i].PrintPaySlip()
				}
			}

        case 4:
			if len(employees) == 0 {
				fmt.Printf("No employees found!!!")
				continue
			}
			for _, emp := range employees {
				emp.PrintPaySlip()
			}

        case 5:
			fmt.Println("\nExitting program, Thank you.")
            os.Exit(0)
        default:
            fmt.Println("Invalid choice!")
        }
    }
}