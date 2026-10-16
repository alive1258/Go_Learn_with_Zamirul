# 05 — Why Functions Are Needed in Go

## Introduction

Functions are an essential part of Go programming. They help developers write reusable, readable, maintainable, and testable code.

The best way to understand why functions are needed is to solve the same problem in two ways: without functions and with functions.

## 1. Code Reusability

**Code reusability** means writing logic once and using it multiple times instead of repeating the same code.

### Without Function

```go
package main

import "fmt"

func main() {
	price1 := 100
	quantity1 := 2
	fmt.Println(price1 * quantity1)

	price2 := 200
	quantity2 := 3
	fmt.Println(price2 * quantity2)

	price3 := 150
	quantity3 := 4
	fmt.Println(price3 * quantity3)
}
```

In this example, the same calculation is repeated for three customers.

### With Function

```go
package main

import "fmt"

func calculateTotal(price, quantity int) int {
	return price * quantity
}

func main() {
	fmt.Println(calculateTotal(100, 2))
	fmt.Println(calculateTotal(200, 3))
	fmt.Println(calculateTotal(150, 4))
}
```

**Output:**

```text
200
600
600
```

Now, the calculation logic is written only once and reused whenever needed.

**Benefits:**

- Reduces code duplication.
- Makes code easier to maintain.
- Improves consistency.
- Makes future changes easier.

## 2. Maintainability

**Maintainability** means making code easier to modify, update, and improve.

Suppose we need to add a 10% service charge to every order.

```go
package main

import "fmt"

func calculateTotal(price, quantity int) int {
	subtotal := price * quantity
	serviceCharge := subtotal / 10

	return subtotal + serviceCharge
}

func main() {
	fmt.Println(calculateTotal(100, 2))
	fmt.Println(calculateTotal(200, 3))
}
```

**Output:**

```text
220
660
```

The calculation is updated in one place, so every function call uses the new logic.

> **Note:** This example uses integer arithmetic for simplicity. In real payment systems, use an appropriate money representation, such as integer minor units or a decimal type, and define how rounding should work.

## 3. Readability

**Readability** means making code easier for developers to understand.

Instead of placing all logic inside `main()`, we can separate responsibilities into functions.

```go
package main

import "fmt"

func validateAge(age int) bool {
	return age >= 18
}

func registerUser(name string, age int) {
	if !validateAge(age) {
		fmt.Println("Registration not allowed")
		return
	}

	fmt.Println("User registered:", name)
}

func main() {
	registerUser("Zamirul", 20)
	registerUser("Rahim", 16)
}
```

**Output:**

```text
User registered: Zamirul
Registration not allowed
```

The function names describe their responsibilities, making the program easier to understand.

## 4. Testing

Functions make it easier to test individual pieces of logic independently.

### `main.go`

```go
package main

func add(a, b int) int {
	return a + b
}
```

### `main_test.go`

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(10, 20)
	want := 30

	if got != want {
		t.Errorf("add(10, 20) = %d; want %d", got, want)
	}
}
```

Run the test:

```bash
go test
```

If the actual result matches the expected result, the test passes.

**Benefits:**

- Helps identify bugs.
- Verifies expected behavior.
- Makes refactoring safer.
- Improves code reliability.

## 5. Separation of Concerns

**Separation of concerns** means dividing a program into functions with clear and distinct responsibilities.

For example, a booking system may contain the following functions:

| Function            | Responsibility                            |
| ------------------- | ----------------------------------------- |
| `validateBooking()` | Validates booking input                   |
| `calculatePrice()`  | Calculates the booking price              |
| `saveBooking()`     | Saves booking information to the database |
| `processPayment()`  | Coordinates payment processing            |

Each function focuses on a particular responsibility, making the system easier to maintain and test.

In production systems, database transactions, payment verification, error handling, and security must also be designed carefully.

## 6. Main Benefits of Functions

| Benefit         | Explanation                                                           |
| --------------- | --------------------------------------------------------------------- |
| Reusability     | Reuse the same logic multiple times                                   |
| Maintainability | Make code easier to modify                                            |
| Readability     | Make code easier to understand                                        |
| Testability     | Test logic independently                                              |
| Modularity      | Divide large problems into smaller parts                              |
| Debugging       | Identify the source of problems more easily                           |
| Abstraction     | Use functionality without needing to know every implementation detail |

## 7. Practice Task

Write a Go program containing the following three functions:

1. `calculateTotal(price, quantity int) int` — returns the total price.
2. `isEligible(age int) bool` — returns whether the given age is at least 18.
3. `printWelcome(name string)` — prints a welcome message.

Call all three functions from `main()`.

### Expected Concepts

- Function declaration
- Parameters and arguments
- Return values
- Boolean values
- Function calls
- Code reusability

## Senior Developer Tip

Before creating a function, ask yourself:

**"Does this function have one clear responsibility?"**

Functions with clear responsibilities are generally easier to read, test, debug, and maintain. However, avoid creating unnecessary functions that make simple logic harder to follow.

---

**Next Topic:** Function Parameters, Return Values, Multiple Return Values, and Error Handling in Go.
