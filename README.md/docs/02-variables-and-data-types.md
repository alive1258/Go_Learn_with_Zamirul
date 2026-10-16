# Go Programming in Bangla 🇧🇩

### Beginner to Production-Level Engineering

Welcome to my Go (Golang) learning journey! This repository contains structured lessons, practical examples, exercises, and production-oriented concepts as I progress from basic Go syntax to backend engineering.

## 📚 Learning Roadmap

| Lesson | Topic                                   | Documentation                                       |
| ------ | --------------------------------------- | --------------------------------------------------- |
| 008    | Variables and Data Types                | [Read Lesson](docs/008-variables-and-data-types.md) |
| 009    | If...else and switch                    | [Read Lesson](docs/009-if-else-and-switch.md)       |
| 010    | Loops, Arrays, Slices and Maps          | Coming Soon                                         |
| 011    | Why Functions Are Needed                | Coming Soon                                         |
| 012    | Functions, Parameters and Return Values | Coming Soon                                         |

## 🎯 Lesson 02 — Variables and Data Types

In this lesson, I learned:

- Variable declaration using `var` and `:=`
- Type inference and zero values
- Integer types: `int`, `int32`, `int64`, and unsigned integers
- Floating-point types: `float32` and `float64`
- Strings, Unicode, bytes, and runes
- Boolean values and constants
- Explicit type conversion
- Variable scope and exported identifiers
- Formatting output using the `fmt` package
- Type safety and common programming mistakes

### Example

```go
package main

import "fmt"

const AppName = "Go Ticket Booking"

type Ticket struct {
    ID       int
    UserName string
    Price    int64
    Quantity int
    Paid     bool
}

func main() {
    ticket := Ticket{
        ID:       101,
        UserName: "Zamirul",
        Price:    50000,
        Quantity: 2,
        Paid:     true,
    }

    total := ticket.Price * int64(ticket.Quantity)

    fmt.Println("Application:", AppName)
    fmt.Println("Ticket ID:", ticket.ID)
    fmt.Println("User:", ticket.UserName)
    fmt.Printf("Total: %d paisa\n", total)
    fmt.Println("Paid:", ticket.Paid)
}
```

Run the example from its directory:

```bash
go run .
```

## 🔀 Lesson 009 — If...else and switch

In this lesson, I learned:

- `if`, `else if`, and `else`
- Comparison operators: `==`, `!=`, `>`, `<`, `>=`, `<=`
- Logical operators: `&&`, `||`, `!`
- Short-circuit evaluation
- Short statements in `if`
- `switch`, multiple `case` values, and `default`
- Expressionless `switch`
- `fallthrough` and type switches
- Input validation and error handling
- Ticket booking business rules

### Example

```go
package main

import (
    "errors"
    "fmt"
)

type BookingRequest struct {
    EventActive   bool
    Quantity      int
    Available     int
    PaymentStatus string
}

func ValidateBooking(req BookingRequest) error {
    if !req.EventActive {
        return errors.New("event is not active")
    }

    if req.Quantity <= 0 {
        return errors.New("quantity must be positive")
    }

    if req.Quantity > req.Available {
        return errors.New("not enough tickets available")
    }

    switch req.PaymentStatus {
    case "paid":
        return nil
    case "pending":
        return errors.New("payment is pending")
    case "failed":
        return errors.New("payment failed")
    default:
        return errors.New("invalid payment status")
    }
}

func main() {
    req := BookingRequest{
        EventActive:   true,
        Quantity:      2,
        Available:     10,
        PaymentStatus: "paid",
    }

    if err := ValidateBooking(req); err != nil {
        fmt.Println("Booking rejected:", err)
        return
    }

    fmt.Println("Booking validation successful")
}
```

## 🛠️ Repository Structure

```text
go-learn-with-zamirul/
├── go.mod
├── README.md
├── docs/
│   ├── 008-variables-and-data-types.md
│   ├── 009-if-else-and-switch.md
│   └── 010-loops-arrays-slices-and-maps.md
└── examples/
    ├── variables/
    │   └── main.go
    ├── conditions/
    │   └── main.go
    └── switch/
        └── main.go
```

_Note: The Lesson 010 documentation and example files will be added as I progress._

## 🧪 Practice and Testing

I am strengthening my understanding through hands-on coding, practical exercises, and testing.

Run all package tests from the repository root:

```bash
go test ./...
```

Run a specific example from its directory:

```bash
cd examples/variables
go run .
```

## 🚀 Learning Goals

- Build a strong foundation in Go
- Write clean, readable, maintainable code
- Understand functions, data structures, interfaces, and error handling
- Develop REST APIs and backend services
- Work with PostgreSQL, databases, and authentication
- Learn concurrency with goroutines and channels
- Write unit and integration tests
- Learn Docker, CI/CD, and production deployment practices
- Build real-world backend projects

## 💡 Learning Philosophy

> Learn the fundamentals, practice with real examples, test your assumptions, and gradually build production-ready software.

This repository is a work in progress. Each lesson includes notes, code examples, exercises, and practical backend engineering concepts.

**Let's learn Go step by step!** 💚
