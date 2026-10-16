package main

import "fmt"

func main() {
	age := 20
	hasNID := true
	hasJob := false

	// 1. Greater than (>)
	if age > 18 {
		fmt.Println("Age is greater than 18")
	}

	// 2. Less than (<)
	if age < 25 {
		fmt.Println("Age is less than 25")
	}

	// 3. Greater than or equal to (>=)
	if age >= 20 {
		fmt.Println("Age is greater than or equal to 20")
	}

	// 4. Less than or equal to (<=)
	if age <= 20 {
		fmt.Println("Age is less than or equal to 20")
	}

	// 5. Equal to (==)
	if age == 20 {
		fmt.Println("Age is exactly 20")
	}

	// 6. Not equal to (!=)
	if age != 18 {
		fmt.Println("Age is not 18")
	}

	// 7. Logical AND (&&)
	if age >= 18 && hasNID {
		fmt.Println("Adult age and has NID")
	}

	// 8. Logical OR (||)
	if hasNID || hasJob {
		fmt.Println("Has NID or has a job")
	}

	// 9. Logical NOT (!)
	if !hasJob {
		fmt.Println("Does not have a job")
	}

	// 10. AND + OR
	if age >= 18 && (hasNID || hasJob) {
		fmt.Println("Meets the conditions")
	} else {
		fmt.Println("Does not meet the conditions")
	}

	// 11. If-else-if-else
	if age < 18 {
		fmt.Println("Minor")
	} else if age == 18 {
		fmt.Println("Exactly 18")
	} else {
		fmt.Println("Older than 18")
	}
}
