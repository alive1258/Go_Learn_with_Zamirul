package localscopeandblock

import "fmt"

const maxTickets = 10

// CalculateTotal calculates the total booking price.
func CalculateTotal(price int64, quantity int) (int64, error) {
	if price <= 0 {
		return 0, fmt.Errorf("price must be positive")
	}

	if quantity < 1 || quantity > maxTickets {
		return 0, fmt.Errorf(
			"quantity must be between 1 and %d",
			maxTickets,
		)
	}

	total := price * int64(quantity)
	return total, nil
}

// Run demonstrates local scope, block scope, and shadowing.
func Run() {
	fmt.Println("=== Local Scope and Block Scope ===")

	// Function-level local variables.
	message := "Learning Go scope"
	fmt.Println(message)

	// Block scope.
	if age := 20; age >= 18 {
		status := "Adult"
		fmt.Println("Status:", status)
		fmt.Println("Age:", age)
	}

	// Loop scope.
	for i := 1; i <= 3; i++ {
		ticketMessage := fmt.Sprintf("Ticket number: %d", i)
		fmt.Println(ticketMessage)
	}

	// Variable shadowing.
	count := 10

	if true {
		count := 20
		fmt.Println("Inner count:", count)
	}

	fmt.Println("Outer count:", count)

	// Assignment instead of shadowing.
	if true {
		count = 30
	}

	fmt.Println("Updated outer count:", count)

	// Practical booking example.
	price := int64(50000)
	quantity := 2

	total, err := CalculateTotal(price, quantity)
	if err != nil {
		fmt.Println("Booking failed:", err)
		return
	}

	fmt.Printf("Booking total: %d paisa (BDT)\n", total)
}
