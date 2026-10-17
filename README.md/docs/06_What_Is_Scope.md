# [Golang] 06 — What Is Scope?

## 1. What Is Scope?

**Scope** হলো প্রোগ্রামের সেই নির্দিষ্ট অংশ, যেখানে কোনো variable, constant, function বা identifier-কে নাম দিয়ে access করা যায়।

সহজভাবে বললে, একটি variable কোথায় ব্যবহার করতে পারবে, আর কোথায় পারবে না—সেটি Scope নির্ধারণ করে।

ধরো, তুমি একটি Ticket Booking API তৈরি করছ। সেখানে `userName`, `ticketPrice` এবং `appName`-এর মতো variable থাকবে। সব variable পুরো application থেকে access করা যাবে না। কোন variable কোথায় declare করা হয়েছে, তার ওপর access নির্ভর করে।

### Example

```go
package main

import "fmt"

func main() {
    name := "Zamirul"

    fmt.Println(name)
}
```

**Output:**

```text
Zamirul
```

এখানে `name` variable-টি `main()` function-এর ভেতরে declare করা হয়েছে। তাই এটি ওই function-এর scope-এর মধ্যে ব্যবহার করা যায়।

---

## 2. Why Is Scope Important?

Scope বোঝা গুরুত্বপূর্ণ কারণ এটি:

- Variable কোথায় access করা যাবে তা নির্ধারণ করে।
- একই নামের variable বিভিন্ন block-এ ব্যবহার করতে সাহায্য করে।
- Unnecessary global variables কমাতে সাহায্য করে।
- Code readability ও maintainability বাড়ায়।
- Accidental modification এবং naming conflicts কমায়।
- বড় application-এ code organization সহজ করে।

Production code-এ সঠিক scope ব্যবহার করলে debugging এবং testing-ও সহজ হতে পারে।

## 3. Types of Scope in Go

Go-তে scope বোঝার জন্য প্রধানত এই বিষয়গুলো শিখবে:

1. Local বা function/block scope
2. Block scope
3. Package scope
4. Universe scope
5. Lexical scope এবং shadowing

Go-তে lexical scoping ব্যবহৃত হয়। অর্থাৎ, কোনো identifier কোথায় declare করা হয়েছে এবং তার চারপাশের enclosing blocks কী—তার ওপর identifier-এর visibility নির্ভর করে।

## 4. Local Scope

Function-এর ভেতরে declare করা variable সাধারণত সেই function-এর local scope-এ থাকে।

### Example

```go
package main

import "fmt"

func main() {
    userName := "Zamirul"
    age := 25

    fmt.Println("Name:", userName)
    fmt.Println("Age:", age)
}
```

**Output:**

```text
Name: Zamirul
Age: 25
```

এখানে `userName` ও `age` শুধু `main()` function-এর মধ্যে access করা যায়।

অন্য function থেকে এগুলো সরাসরি access করা যাবে না।

```go
package main

import "fmt"

func main() {
    userName := "Zamirul"
    fmt.Println(userName)
}

func showUser() {
    // fmt.Println(userName) // Compile error
}
```

`showUser()`-এর মধ্যে `userName` নেই, কারণ variable-টি `main()`-এর local scope-এ declare করা হয়েছে।

## 5. Block Scope

Go-তে `{}` দিয়ে তৈরি block-এর মধ্যে declare করা variable সেই block এবং তার ভেতরের nested block-এ access করা যায়।

### Example

```go
package main

import "fmt"

func main() {
    if true {
        message := "Hello from block"
        fmt.Println(message)
    }

    // fmt.Println(message) // Compile error
}
```

**Output:**

```text
Hello from block
```

`message` variable-টি `if` block-এর মধ্যে declare করা হয়েছে। তাই block-এর বাইরে এটি access করা যায় না।

### Nested Block

```go
package main

import "fmt"

func main() {
    name := "Zamirul"

    if true {
        age := 25

        fmt.Println("Name:", name)
        fmt.Println("Age:", age)
    }

    // fmt.Println(age) // Compile error
}
```

Nested block-এর মধ্যে বাইরের block-এর `name` access করা যায়। কিন্তু বাইরের block থেকে ভেতরের `age` access করা যায় না।

**মনে রাখবে:** Inner block বাইরের scope-এর identifier ব্যবহার করতে পারে, যদি সেগুলো সেখানে visible হয়। কিন্তু বাইরের block inner scope-এর local variable access করতে পারে না।

## 6. Package Scope

একটি package-এর top level-এ declare করা variable, constant, type এবং function সেই package-এর অন্যান্য source file থেকেও ব্যবহার করা যায়।

### Example

`config.go`

```go
package main

var appName = "Ticket Booking API"
var appVersion = "1.0.0"
```

`main.go`

```go
package main

import "fmt"

func main() {
    fmt.Println(appName)
    fmt.Println(appVersion)
}
```

দুটি file একই directory-তে এবং একই package-এ থাকলে package-level identifier ব্যবহার করা যায়।

চালানোর command:

```bash
go run .
```

এখানে মনে রাখবে, `go run .` চালানোর জন্য ওই directory-তে একটি executable `main` package থাকতে হবে এবং package-এ একটি `main()` entry point থাকতে হবে।

### Package-level Variables কখন ব্যবহার করবে?

```go
package main

const appName = "Ticket Booking API"
```

Application-এর স্থির নামের মতো value constant হিসেবে রাখা যুক্তিযুক্ত হতে পারে।

কিন্তু প্রতিটি request-এর user data package-level mutable variable-এ রাখা উচিত নয়:

```go
// Avoid this design for request-specific data.
var currentUser string
```

একাধিক request একই সময়ে এই variable পরিবর্তন করলে race condition এবং data leakage-এর ঝুঁকি তৈরি হতে পারে।

Request-specific data function parameter, request context বা উপযুক্ত data structure-এর মাধ্যমে পরিচালনা করো।

## 7. Exported vs Unexported Identifiers

Go-তে identifier-এর নামের প্রথম অক্ষর uppercase না lowercase—তার ওপর অন্য package থেকে access করা যাবে কি না নির্ভর করে।

```go
package config

const AppName = "Ticket Booking API"

const appVersion = "1.0.0"
```

- `AppName` exported; অন্য package থেকে `config.AppName` হিসেবে access করা যায়।
- `appVersion` unexported; অন্য package থেকে সরাসরি access করা যায় না।

এটি package-level visibility-এর নিয়ম। Function-এর ভেতরে uppercase দিয়ে local variable-এর নাম রাখলেই সেটি অন্য package থেকে access করা যাবে না।

## 8. Variable Shadowing

**Shadowing** হলো যখন একটি inner scope-এ একই নামের নতুন variable declare করা হয়, ফলে বাইরের scope-এর variable সাময়িকভাবে আড়াল হয়ে যায়।

### Example

```go
package main

import "fmt"

func main() {
    name := "Zamirul"

    if true {
        name := "Kabir"
        fmt.Println("Inside:", name)
    }

    fmt.Println("Outside:", name)
}
```

**Output:**

```text
Inside: Kabir
Outside: Zamirul
```

এখানে দুটি আলাদা `name` variable আছে:

- বাইরের `name`-এর value `"Zamirul"`।
- `if` block-এর ভেতরের `name`-এর value `"Kabir"`।

Inner variable বাইরের variable-টিকে পরিবর্তন করেনি।

### Shadowing-এর সম্ভাব্য সমস্যা

```go
package main

import "fmt"

func main() {
    total := 100

    if true {
        total := 200
        fmt.Println(total)
    }

    fmt.Println(total)
}
```

**Output:**

```text
200
100
```

কখনো কখনো এই আচরণ bug তৈরি করতে পারে, বিশেষ করে `err`, `total` বা database result-এর মতো গুরুত্বপূর্ণ variable-এর ক্ষেত্রে।

## 9. `:=` এবং `=`-এর পার্থক্য

Go-তে scope ও shadowing শেখার সময় এই পার্থক্যটি খুব গুরুত্বপূর্ণ।

### `:=` — Short Variable Declaration

```go
name := "Zamirul"
```

এটি নতুন local variable declare করে এবং initial value দেয়।

### `=` — Assignment

```go
name := "Zamirul"
name = "Kabir"
```

এখানে দ্বিতীয় লাইনে আগের variable-এর value পরিবর্তন হয়েছে।

### Shadowing Example

```go
package main

import "fmt"

func main() {
    age := 25

    if true {
        age := 30
        fmt.Println("Inner age:", age)
    }

    fmt.Println("Outer age:", age)
}
```

Output:

```text
Inner age: 30
Outer age: 25
```

যদি inner block-এ `age = 30` লিখতে, তাহলে নতুন variable তৈরি না হয়ে বাইরের variable-এর value পরিবর্তন হতো।

## 10. Scope in `if` Short Statements

Go-তে `if` statement-এর আগে variable declare করা যায়। সেই variable `if` এবং তার সংশ্লিষ্ট `else` blocks-এর মধ্যে ব্যবহারযোগ্য।

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    input := "123"

    if number, err := strconv.Atoi(input); err == nil {
        fmt.Println("Number:", number)
    } else {
        fmt.Println("Invalid number:", err)
    }

    // fmt.Println(number) // Compile error
}
```

**Output:**

```text
Number: 123
```

`number` এবং `err` শুধু `if` statement-এর scope-এর মধ্যে ব্যবহারযোগ্য।

এটি Go-তে error handling-এর খুব সাধারণ pattern।

## 11. Universe Scope

Go language specification অনুযায়ী কিছু predeclared identifier universe block-এ থাকে। এগুলো সাধারণভাবে Go program-এর যেকোনো scope থেকে ব্যবহার করা যায়।

উদাহরণ:

- `int`, `string`, `bool`
- `true`, `false`
- `nil`
- `len()`, `cap()`, `make()`, `new()`
- `error`, `any`, `comparable`

Example:

```go
package main

import "fmt"

func main() {
    numbers := []int{10, 20, 30}

    fmt.Println(len(numbers))
}
```

এখানে `int` একটি predeclared type এবং `len` একটি predeclared built-in function।

## 12. Scope vs Lifetime

Scope এবং lifetime একই বিষয় নয়।

- **Scope:** Source code-এর কোন অংশ থেকে identifier-কে নাম দিয়ে access করা যায়।
- **Lifetime:** Program execution-এর কোন সময় পর্যন্ত variable বা সংশ্লিষ্ট storage বিদ্যমান থাকে।

উদাহরণস্বরূপ, function থেকে pointer return করা হলে local variable-এর storage প্রয়োজন অনুযায়ী function শেষ হওয়ার পরও বেঁচে থাকতে পারে।

```go
package main

import "fmt"

func createCounter() *int {
    count := 10
    return &count
}

func main() {
    counter := createCounter()
    fmt.Println(*counter)
}
```

**Output:**

```text
10
```

এখানে `count`-এর lexical scope `createCounter()` function-এর মধ্যে সীমাবদ্ধ। কিন্তু return করা pointer-এর মাধ্যমে value-টি পরে access করা যায়। Go compiler escape analysis-এর মাধ্যমে storage কীভাবে পরিচালিত হবে তা নির্ধারণ করে।

## 13. Real-World Example — Ticket Booking

ধরো, একটি booking-এর মোট মূল্য হিসাব করতে হবে।

```go
package main

import "fmt"

const appName = "Ticket Booking API"

func calculateTotal(price int64, quantity int) int64 {
    total := price * int64(quantity)
    return total
}

func main() {
    ticketPrice := int64(50000)
    quantity := 2

    total := calculateTotal(ticketPrice, quantity)

    fmt.Println("Application:", appName)
    fmt.Println("Total price:", total, "paisa")
}
```

**Output:**

```text
Application: Ticket Booking API
Total price: 100000 paisa
```

এই উদাহরণে:

- `appName` হলো package-level constant।
- `ticketPrice` ও `quantity` হলো `main()`-এর local variables।
- `price` ও `quantity` হলো `calculateTotal()`-এর parameters।
- `total` হলো `calculateTotal()`-এর local variable।

Production application-এ user input validate করবে এবং `price * quantity` হিসাবের সম্ভাব্য integer overflow বিবেচনা করবে।

## 14. Common Mistakes

### Mistake 1: Local Variable বাইরে ব্যবহার করা

```go
func main() {
    if true {
        message := "Hello"
        println(message)
    }

    // println(message) // Compile error
}
```

**সমাধান:** Variable-টি যে scope-এ দরকার, সেই অনুযায়ী declare করো।

### Mistake 2: ভুল করে Shadowing তৈরি করা

```go
total := 100

if true {
    total := 200
}
```

এখানে বাইরের `total` পরিবর্তন হয় না।

**সমাধান:** নতুন variable দরকার হলে `:=` এবং আগের variable পরিবর্তন করতে চাইলে `=` ব্যবহার করো।

### Mistake 3: সবকিছু Global করা

সব variable package-level-এ রাখলে code-এর dependencies অস্পষ্ট হতে পারে এবং concurrency-related bug তৈরি হতে পারে।

**সমাধান:** Local variables, function parameters এবং explicit dependencies-কে অগ্রাধিকার দাও।

### Mistake 4: Package Scope ও Exported Visibility গুলিয়ে ফেলা

একই package-এর অন্য file থেকে unexported identifier ব্যবহার করা যায়। কিন্তু অন্য package থেকে যায় না।

**সমাধান:** Go-তে package boundary এবং identifier capitalization-এর নিয়ম বুঝে ব্যবহার করো।

## 15. Best Practices — Senior Engineer Tips

- Variable যতটা সম্ভব ছোট ও প্রয়োজনীয় scope-এ রাখো।
- Request-specific mutable data global variable-এ রাখবে না।
- Shadowing এড়িয়ে চলবে, যদি সেটি code-কে অস্পষ্ট করে।
- Function parameters দিয়ে dependencies স্পষ্ট করো।
- Package-level variable শুধু প্রকৃত প্রয়োজন থাকলেই ব্যবহার করো।
- Exported identifiers-এর নাম অর্থবহ রাখো এবং public API সচেতনভাবে ডিজাইন করো।
- `go vet` এবং `go test` দিয়ে code যাচাই করো।
- Concurrent code-এর জন্য `go test -race ./...` ব্যবহার করো।

## 16. Knowledge Check

1. Scope কী?
2. Local variable কি অন্য function থেকে সরাসরি access করা যায়?
3. Block scope বলতে কী বোঝায়?
4. Package-level variable কি একই package-এর অন্য file থেকে ব্যবহার করা যায়?
5. Shadowing কী?
6. `:=` ও `=`-এর মধ্যে পার্থক্য কী?
7. Scope এবং lifetime কেন এক নয়?
8. কেন request-specific data global mutable variable-এ রাখা ঝুঁকিপূর্ণ?

## 17. Homework

তোমার `go-learn-with-zamirul` repository-তে নিচের কাজগুলো করো:

1. `docs/014-what-is-scope.md` ফাইলে এই lesson সংরক্ষণ করো।
2. `examples/scope/main.go` তৈরি করো।
3. Local scope এবং nested block scope-এর উদাহরণ লেখো।
4. Variable shadowing-এর একটি example তৈরি করে output ব্যাখ্যা করো।
5. দুটি source file ব্যবহার করে package-level scope পরীক্ষা করো।
6. `calculateTotal()` function-এ parameters ও local variables চিহ্নিত করো।
7. নিজের ভাষায় Scope এবং Lifetime-এর পার্থক্য লেখো।

### Run the examples

Repository root থেকে:

```bash
go run ./examples/scope
go test ./...
go vet ./...
```

Race detection সমর্থিত হলে:

```bash
go test -race ./...
```

## Next Lesson

**[Golang] 015 — What Are Pointers?**

পরের লেসনে আমরা শিখব memory address, pointer declaration, dereferencing এবং Go-তে pointers কখন ব্যবহার করা উচিত।
