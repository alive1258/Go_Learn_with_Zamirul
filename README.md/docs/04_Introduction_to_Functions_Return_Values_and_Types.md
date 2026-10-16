# 04 — Introduction to Functions, Return Values and Types in Go

## Introduction

In this lesson, we will learn about **Functions, Parameters, Return Values, Multiple Return Values, and Function Types** in Go.

Functions are essential for backend development because they help organize business logic, API operations, database interactions, authentication, and validation.

## 1. What Is a Function?

A function is a reusable block of code that performs a specific task. Instead of writing the same code repeatedly, we can define a function and call it whenever needed.

### Example

```go
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello, Go Developer!")
}

func main() {
	sayHello()
	sayHello()
}
```

**Output:**

```text
Hello, Go Developer!
Hello, Go Developer!
```

### Explanation

- `func` — keyword used to declare a function.
- `sayHello` — function name.
- `()` — parentheses containing parameters, if any.
- `{}` — function body.
- `sayHello()` — function call.

## 2. Functions with Parameters

Parameters allow us to pass data into a function.

### Example

```go
package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	greet("Zamirul")
	greet("Rahim")
	greet("Karim")
}
```

**Output:**

```text
Hello, Zamirul
Hello, Rahim
Hello, Karim
```

Here, `name string` declares a parameter named `name` with the type `string`.

### Multiple Parameters

A function can accept multiple parameters of different types.

```go
package main

import "fmt"

func introduce(name string, age int) {
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
}

func main() {
	introduce("Zamirul", 25)
}
```

When multiple consecutive parameters share the same type, Go allows a shorter declaration:

```go
func add(a, b int) {
	fmt.Println(a + b)
}
```

## 3. Functions with Return Values

A function can return a result instead of only printing it. The `return` keyword sends the result back to the caller.

### Example

```go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	result := add(10, 20)

	fmt.Println("Sum:", result)
}
```

**Output:**

```text
Sum: 30
```

### Function Declaration

```go
func add(a, b int) int
```

- `add` — function name.
- `a, b int` — parameters.
- The final `int` — return type.
- `return a + b` — returns the calculated result.

**Important:** `fmt.Println()` displays a value, while `return` allows the caller to use that value in further calculations or operations.

## 4. Different Return Types

A function can return values of different types. The returned value must match the declared return type.

```go
package main

import "fmt"

func getName() string {
	return "Zamirul"
}

func getAge() int {
	return 25
}

func getPrice() float64 {
	return 199.99
}

func isActive() bool {
	return true
}

func main() {
	fmt.Println(getName())
	fmt.Println(getAge())
	fmt.Println(getPrice())
	fmt.Println(isActive())
}
```

### Common Return Types

| Return Type | Example                        |
| ----------- | ------------------------------ |
| `string`    | `return "Go"`                  |
| `int`       | `return 100`                   |
| `float64`   | `return 99.50`                 |
| `bool`      | `return true`                  |
| `error`     | `return nil` or an error value |

## 5. Multiple Return Values

Go supports returning multiple values from a single function. This feature is especially useful for returning a result and an error together.

### Example: Division with Error Handling

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	result, err := divide(10, 2)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Result:", result)

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
```

**Output:**

```text
Result: 5
Error: cannot divide by zero
```

### Explanation

- `(float64, error)` — declares two return types.
- `nil` — indicates that no error occurred.
- `err != nil` — checks whether an error was returned.
- `_` — ignores a returned value when it is not needed.

**Production rule:** Always handle errors appropriately, especially in payment processing, booking systems, and database operations.

## 6. Named Return Values

Go allows us to name return values in the function declaration.

### Example

```go
package main

import "fmt"

func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b

	return
}

func main() {
	total, multiplication := calculate(5, 3)

	fmt.Println("Sum:", total)
	fmt.Println("Product:", multiplication)
}
```

**Output:**

```text
Sum: 8
Product: 15
```

Here, `sum` and `product` are named return variables. A bare `return` returns their current values.

Named returns can be useful in some situations, but avoid using them when they make a function harder to understand.

## 7. Functions with Conditions

Functions often contain conditions and validation logic in real-world applications.

### Example: Age Validation

```go
package main

import "fmt"

func canRegister(age int) bool {
	if age >= 18 {
		return true
	}

	return false
}

func main() {
	if canRegister(20) {
		fmt.Println("Registration allowed")
	} else {
		fmt.Println("Registration not allowed")
	}
}
```

The same function can be written more concisely:

```go
func canRegister(age int) bool {
	return age >= 18
}
```

Both implementations produce the same result.

> **Note:** This is a programming example, not a complete registration policy. Real applications must follow their applicable age requirements and business rules.

## 8. Variadic Functions

A variadic function accepts a variable number of arguments. In Go, the `...` syntax is used to declare variadic parameters.

### Example

```go
package main

import "fmt"

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())
}
```

**Output:**

```text
3
15
0
```

The parameter `numbers ...int` behaves like a slice of integers inside the function. We use `range` to iterate over the values and calculate their sum.

## 9. Practical Example — Booking System

Let's create a function that validates booking inputs and calculates the total ticket price.

```go
package main

import (
	"errors"
	"fmt"
)

func calculateBookingPrice(ticketCount int, price float64) (float64, error) {
	if ticketCount <= 0 {
		return 0, errors.New("ticket count must be greater than zero")
	}

	if price < 0 {
		return 0, errors.New("price cannot be negative")
	}

	return float64(ticketCount) * price, nil
}

func main() {
	total, err := calculateBookingPrice(3, 500)

	if err != nil {
		fmt.Println("Booking failed:", err)
		return
	}

	fmt.Println("Total price:", total)
}
```

**Output:**

```text
Total price: 1500
```

This function validates the input and calculates the total price.

In a production booking system, additional controls would be necessary for ticket availability, database transactions, payment verification, and concurrency.

For financial calculations, choose a suitable money representation and rounding policy rather than relying on unrestricted floating-point arithmetic.

## 10. Common Mistakes

| Incorrect Code or Practice              | Explanation                                                                       |
| --------------------------------------- | --------------------------------------------------------------------------------- |
| `func add(a, b int) { return a + b }`   | A function returning a value must declare its return type.                        |
| `func getAge() int { return "20" }`     | The function declares `int` but returns a `string`.                               |
| `result := divide(10, 2)`               | `divide` returns two values, so both must be handled or one explicitly discarded. |
| Ignoring errors                         | Failures may go unnoticed.                                                        |
| Putting all logic in one large function | Testing, debugging, and maintenance become harder.                                |

## 11. Practice Tasks

Complete the following exercises to strengthen your understanding of Go functions.

- [ ] Write a `square(number int) int` function that returns the square of a number.
- [ ] Write a `max(a, b int) int` function that returns the larger number.
- [ ] Write an `isEven(number int) bool` function that returns `true` for even numbers.
- [ ] Write a function that returns the sum and difference of two integers.

### Challenge

Create a function with this signature:

```go
func calculate(a, b int) (sum int, difference int)
```

The function should return both the sum and the difference of the two integers.

### How to Run Your Code

Save your code in a Go file, such as `main.go`, and run:

```bash
go run .
```

To run your project's tests:

```bash
go test ./...
```

## Summary

In this lesson, we learned:

- Function declarations and function calls.
- Parameters and arguments.
- Return values and return types.
- Multiple return values.
- Error handling.
- Named return values.
- Functions with conditions.
- Variadic functions.
- Practical booking-price calculation.
- Common mistakes and exercises.

**Next Topic:** Function Parameters, Return Values, Multiple Return Values, and Error Handling in Go.
