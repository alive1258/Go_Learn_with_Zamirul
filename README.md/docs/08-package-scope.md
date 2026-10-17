# [Golang] 016 — Package Scope (প্যাকেজ স্কোপ)

> **লেভেল:** Beginner  
> **বিষয়:** Go Fundamentals  
> **পূর্বশর্ত:** Local Scope and Block Scope  
> **লক্ষ্য:** Go-তে Package Scope, Exported ও Unexported Identifiers এবং একাধিক প্যাকেজের মধ্যে কোড ব্যবহার করা শেখা।

## ১. Package Scope কী?

Go-তে কোনো প্যাকেজের একদম উপরের স্তরে, অর্থাৎ কোনো ফাংশনের বাইরে Variable, Constant, Function, Type বা Interface ঘোষণা করলে সেগুলো Package Scope-এর অন্তর্ভুক্ত হতে পারে।

সহজভাবে বললে, **Package Scope নির্ধারণ করে একটি প্যাকেজের কোন অংশ থেকে কোনো Identifier ব্যবহার করা যাবে।**

Identifier বলতে বোঝায় কোনো Variable, Constant, Function, Type বা অন্য কোনো নাম, যেটি কোডে ব্যবহার করা হয়।

### উদাহরণ

```go
package main

import "fmt"

const maxTickets = 100

var bookedTickets int

func bookTicket(quantity int) {
    bookedTickets += quantity
}

func main() {
    bookTicket(5)

    fmt.Println("Booked tickets:", bookedTickets)
    fmt.Println("Maximum tickets:", maxTickets)
}
```

**Output:**

```text
Booked tickets: 5
Maximum tickets: 100
```

### কোডের ব্যাখ্যা

- `maxTickets` একটি Package-level Constant।
- `bookedTickets` একটি Package-level Variable।
- `bookTicket()` ফাংশনটি `bookedTickets` পরিবর্তন করতে পারে।
- `main()` ফাংশন থেকেও এই দুইটি Identifier ব্যবহার করা যায়।

কারণ এগুলো একই প্যাকেজের মধ্যে ঘোষণা করা হয়েছে এবং সংশ্লিষ্ট Scope Rules অনুযায়ী অ্যাক্সেসযোগ্য।

## ২. Package Scope বনাম Local Scope

| বৈশিষ্ট্য                  | Package Scope                                  | Local Scope                       |
| -------------------------- | ---------------------------------------------- | --------------------------------- |
| কোথায় ঘোষণা করা হয়       | প্যাকেজের শীর্ষ স্তরে                          | ফাংশন বা ব্লকের ভিতরে             |
| কোথা থেকে ব্যবহার করা যায় | একই প্যাকেজের উপযুক্ত Scope থেকে               | সংশ্লিষ্ট ফাংশন বা ব্লকের ভিতরে   |
| উদাহরণ                     | `var total int`                                | `total := 10`                     |
| অন্য প্যাকেজ থেকে ব্যবহার  | Exported হতে হবে                               | সরাসরি ব্যবহার করা যায় না        |
| সাধারণ ব্যবহার             | প্যাকেজের শেয়ার করা Constant, Function ও Type | নির্দিষ্ট কাজের অস্থায়ী Variable |

### উদাহরণ

```go
package main

import "fmt"

var applicationName = "Ticket Booking"

func main() {
    userName := "Zamirul"

    fmt.Println(applicationName)
    fmt.Println(userName)
}
```

এখানে `applicationName`-এর Package Scope রয়েছে। অন্যদিকে `userName` শুধু `main()` ফাংশনের Local Scope-এর মধ্যে ব্যবহারযোগ্য।

## ৩. Exported ও Unexported Identifiers

Go-তে অন্য প্যাকেজ থেকে কোনো Identifier ব্যবহার করা যাবে কি না, তা নির্ধারণের একটি গুরুত্বপূর্ণ নিয়ম হলো নামের প্রথম অক্ষর।

- **Exported Identifier:** নামের প্রথম অক্ষর বড় হাতের হয়।
- **Unexported Identifier:** নামের প্রথম অক্ষর ছোট হাতের হয়।

### উদাহরণ

```go
package pricing

const TaxRate = 0.15

const defaultDiscount = 0.05

func CalculateTotal(price float64) float64 {
    return price + (price * TaxRate)
}

func calculateDiscount(price float64) float64 {
    return price * defaultDiscount
}
```

| Identifier          | ধরন                 | অন্য প্যাকেজ থেকে ব্যবহার |
| ------------------- | ------------------- | ------------------------- |
| `TaxRate`           | Exported Constant   | হ্যাঁ                     |
| `defaultDiscount`   | Unexported Constant | না                        |
| `CalculateTotal`    | Exported Function   | হ্যাঁ                     |
| `calculateDiscount` | Unexported Function | না                        |

**মনে রাখবে:** বড় হাতের অক্ষর দিয়ে শুরু হওয়া Identifier-কে Exported বলা হয়। তবে সেটি ব্যবহার করতে হলে সংশ্লিষ্ট প্যাকেজ Import করতে হবে।

## ৪. অন্য প্যাকেজ থেকে Exported Identifier ব্যবহার

একটি Go Project-এ একাধিক প্যাকেজ থাকতে পারে। অন্য প্যাকেজের Exported Function বা Constant ব্যবহার করতে হলে সেটি Import করতে হয়।

### প্রজেক্ট স্ট্রাকচার

```text
go-learn-with-zamirul/
├── go.mod
├── main.go
└── pricing/
    └── pricing.go
```

ধরা যাক, `go.mod` ফাইলে রয়েছে:

```go
module github.com/alive1258/go-learn-with-zamirul

go 1.24
```

তোমার প্রজেক্টে Go-এর অন্য Version ব্যবহার করা হলে সেটি অনুযায়ী Version লিখবে।

### ফাইল: `pricing/pricing.go`

```go
package pricing

const TaxRate = 0.15

func CalculateTotal(price float64) float64 {
    return price + (price * TaxRate)
}
```

### ফাইল: `main.go`

```go
package main

import (
    "fmt"

    "github.com/alive1258/go-learn-with-zamirul/pricing"
)

func main() {
    price := 1000.0

    total := pricing.CalculateTotal(price)

    fmt.Printf("Original Price: %.2f\n", price)
    fmt.Printf("Tax Rate: %.0f%%\n", pricing.TaxRate*100)
    fmt.Printf("Total Price: %.2f\n", total)
}
```

**Output:**

```text
Original Price: 1000.00
Tax Rate: 15%
Total Price: 1150.00
```

### কীভাবে কাজ করছে?

১. `main.go` ফাইলে `pricing` প্যাকেজ Import করা হয়েছে।

২. `CalculateTotal` বড় হাতের অক্ষর দিয়ে শুরু হওয়ায় এটি Exported Function।

৩. `TaxRate`-ও বড় হাতের অক্ষর দিয়ে শুরু হওয়ায় এটি Exported Constant।

৪. `pricing.CalculateTotal()` লিখে অন্য প্যাকেজের Function ব্যবহার করা হয়েছে।

অন্যদিকে `pricing` প্যাকেজের `calculateDiscount()` ফাংশনটি `main.go` থেকে সরাসরি ব্যবহার করা যাবে না, কারণ এটি Unexported।

## ৫. একই প্যাকেজের একাধিক ফাইলে Package Scope

Go-তে একই Directory-র একাধিক `.go` ফাইল একই প্যাকেজের অন্তর্ভুক্ত হতে পারে।

### প্রজেক্ট স্ট্রাকচার

```text
ticketapp/
├── go.mod
├── main.go
├── config.go
└── booking.go
```

### ফাইল: `config.go`

```go
package main

const maxTickets = 100
```

### ফাইল: `booking.go`

```go
package main

var bookedTickets int

func bookTicket(quantity int) {
    bookedTickets += quantity
}
```

### ফাইল: `main.go`

```go
package main

import "fmt"

func main() {
    bookTicket(10)

    fmt.Println("Maximum tickets:", maxTickets)
    fmt.Println("Booked tickets:", bookedTickets)
}
```

**Output:**

```text
Maximum tickets: 100
Booked tickets: 10
```

### কেন কাজ করছে?

তিনটি ফাইলেই `package main` লেখা হয়েছে। তাই একই প্যাকেজের Package-level Declaration-গুলো পরস্পরের কাছে উপলব্ধ।

এখানে `maxTickets`, `bookedTickets` এবং `bookTicket()` ব্যবহার করতে আলাদা করে Import করার প্রয়োজন নেই।

**গুরুত্বপূর্ণ নিয়ম:**

- প্রতিটি Go ফাইলে Package Declaration থাকতে হবে।
- একই প্যাকেজের ফাইলগুলো সাধারণত একই Directory-তে থাকে।
- প্রতিটি ফাইলকে নিজে যে Package ব্যবহার করবে, সেটির প্রয়োজনীয় Import লিখতে হবে।
- আলাদা প্যাকেজের মধ্যে Identifier ব্যবহার করতে Exported/Unexported নিয়ম মানতে হবে।

## ৬. Ticket Booking Application-এ Package Scope

এবার একটি সাধারণ Ticket Booking উদাহরণ দেখা যাক।

```go
package main

import (
    "errors"
    "fmt"
)

const maxTickets = 100

var bookedTickets int

func bookTicket(quantity int) error {
    if quantity <= 0 {
        return errors.New("quantity must be positive")
    }

    if bookedTickets+quantity > maxTickets {
        return errors.New("not enough tickets available")
    }

    bookedTickets += quantity
    return nil
}

func main() {
    if err := bookTicket(5); err != nil {
        fmt.Println("Booking failed:", err)
        return
    }

    fmt.Println("Booked tickets:", bookedTickets)
    fmt.Println("Available tickets:", maxTickets-bookedTickets)
}
```

**Output:**

```text
Booked tickets: 5
Available tickets: 95
```

### Scope Breakdown

| Identifier      | Scope                | কাজ                                     |
| --------------- | -------------------- | --------------------------------------- |
| `maxTickets`    | Package Scope        | সর্বোচ্চ টিকিটের সংখ্যা নির্ধারণ করে    |
| `bookedTickets` | Package Scope        | বুক হওয়া টিকিটের সংখ্যা রাখে           |
| `quantity`      | Function Parameter   | কতটি টিকিট বুক হবে তা গ্রহণ করে         |
| `err`           | `if` Statement Scope | Booking-এর Error পরীক্ষা করে            |
| `bookTicket`    | Package Scope        | টিকিট বুকিংয়ের Validation ও Update করে |

> **Production Note:** এটি শেখার উদাহরণ, সম্পূর্ণ Production Booking System নয়। একই সময়ে একাধিক Request এলে Ticket Overselling ঠেকাতে উপযুক্ত Synchronization অথবা Database Transaction প্রয়োজন।

## ৭. সাধারণ ভুল

### ভুল ১: অন্য প্যাকেজ থেকে Unexported Identifier ব্যবহার করা

```go
pricing.calculateDiscount(1000)
```

এটি অন্য প্যাকেজ থেকে কাজ করবে না, কারণ `calculateDiscount` ছোট হাতের অক্ষর দিয়ে শুরু হয়েছে।

**সমাধান:** অন্য প্যাকেজের সত্যিই প্রয়োজন হলে Function-টি Exported করো।

### ভুল ২: Package Scope-এ `:=` ব্যবহার করা

```go
package main

total := 100 // ভুল
```

`:=` বা Short Variable Declaration শুধু Function-এর ভিতরে ব্যবহার করা যায়।

**সঠিক পদ্ধতি:**

```go
package main

var total = 100
```

### ভুল ৩: Package-level Variable-কে Global Variable মনে করা

Package-level Variable পুরো Application-এর সব প্যাকেজ থেকে সরাসরি ব্যবহার করা যায় না। এটি মূলত সংশ্লিষ্ট প্যাকেজের Scope-এর মধ্যে থাকে।

### ভুল ৪: `main` প্যাকেজ Import করার চেষ্টা করা

`main` প্যাকেজ সাধারণত Executable Application তৈরি করে। অন্য প্যাকেজে পুনরায় ব্যবহারযোগ্য Logic রাখতে চাইলে সেটিকে আলাদা প্যাকেজে রাখা ভালো।

## ৮. Industry Best Practices

- প্রয়োজন ছাড়া Package-level Mutable Variable ব্যবহার করবে না।
- কোনো Variable শুধু একটি Function-এ প্রয়োজন হলে সেটিকে Local রাখবে।
- স্থির মানের জন্য Constant ব্যবহার করবে।
- অন্য প্যাকেজের প্রয়োজন নেই এমন Function বা Type Export করবে না।
- Package-এর নাম ছোট, পরিষ্কার ও অর্থবহ রাখবে।
- প্রতিটি ফাইলের জন্য আলাদা Package তৈরি না করে কাজের দায়িত্ব অনুযায়ী Package সাজাবে।
- `utils` বা `common`-এর মতো অস্পষ্ট Package Name অপ্রয়োজনীয়ভাবে ব্যবহার করবে না।
- **Scope এবং Lifetime এক বিষয় নয়।** Scope নির্ধারণ করে Identifier কোথায় ব্যবহার করা যাবে; Lifetime নির্ধারণ করে কোনো Value বা তার Storage কতক্ষণ প্রাসঙ্গিক থাকবে।

## ৯. অনুশীলনী

### Exercise 1 — Scope শনাক্ত করো

নিচের কোডে প্রতিটি Identifier-এর Scope নির্ধারণ করো।

```go
package main

const appName = "TicketApp"

var totalBookings int

func createBooking(quantity int) {
    bookingCount := quantity
    totalBookings += bookingCount
}
```

শনাক্ত করো:

- `appName`
- `totalBookings`
- `quantity`
- `bookingCount`

### Exercise 2 — Exported নাকি Unexported?

```go
func CalculatePrice() {}
func calculateTax() {}

type Booking struct{}
type bookingStatus struct{}

const MaxCapacity = 100
const defaultCapacity = 50
```

প্রতিটি Identifier-কে Exported অথবা Unexported হিসেবে শ্রেণিবদ্ধ করো।

### Exercise 3 — নিজের Package তৈরি করো

একটি `pricing` প্যাকেজ তৈরি করো, যেখানে থাকবে:

- `TaxRate` নামে Exported Constant।
- `CalculateTotal` নামে Exported Function।
- `calculateTax` নামে Unexported Helper Function।

এরপর `main.go` থেকে প্যাকেজটি Import করে একটি পণ্যের Tax-সহ মোট মূল্য হিসাব করো।

## ১০. কোড রান ও যাচাই

প্রজেক্টের Root Directory থেকে চালাও:

```bash
go run .
```

প্রজেক্টের Package Structure এবং Module Configuration পরীক্ষা করতে:

```bash
go env GOMOD
go list ./...
```

মনে রাখবে, Multi-package Example-এ Import Path অবশ্যই `go.mod`-এর Module Path-এর সঙ্গে মিলতে হবে।

## ১১. আজকের মূল শিক্ষা

- Function-এর বাইরে Declaration করলে Package Scope তৈরি হতে পারে।
- একই প্যাকেজের একাধিক ফাইল Package-level Declaration ব্যবহার করতে পারে।
- বড় হাতের অক্ষর দিয়ে শুরু হওয়া Identifier Exported।
- ছোট হাতের অক্ষর দিয়ে শুরু হওয়া Identifier Unexported।
- অন্য প্যাকেজ থেকে Exported Identifier ব্যবহার করতে Import করতে হয়।
- Package-level Declaration-এর জন্য `:=` ব্যবহার করা যায় না।
- ভালো Package Design-এর লক্ষ্য হলো প্রয়োজনীয় Public API প্রকাশ করা এবং অপ্রয়োজনীয় Implementation Details আড়ালে রাখা।

## পরবর্তী লেসন

**[Golang] 017 — Package Organization and Imports**

পরবর্তী লেসনে শিখবে কীভাবে Go Project-এ Package সাজাতে হয়, অর্থবহ Package Name নির্বাচন করতে হয় এবং Maintainable Application-এর জন্য Import পরিচালনা করতে হয়।
