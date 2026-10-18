# [Golang] 018 — Variable Shadowing

> **Level:** Beginner  
> **Topic:** Go Fundamentals  
> **Prerequisites:** Local Scope, Block Scope, Package Scope  
> **Goal:** Understand Variable Shadowing, how `:=` and `=` behave in nested scopes, and how to avoid shadowing-related bugs in production code.

## 1. What Is Variable Shadowing?

Go-তে **Variable Shadowing** হলো এমন একটি পরিস্থিতি, যেখানে Inner Scope-এর কোনো Variable একই নামের Outer Scope-এর Variable-কে আড়াল করে রাখে।

ফলে Inner Scope-এর মধ্যে একই নাম ব্যবহার করলে সাধারণত Inner Variable-এর Value পাওয়া যায়। Inner Scope শেষ হলে Outer Variable আবার আগের মতো ব্যবহার করা যায়।

### Example 1: Basic Variable Shadowing

```go
package main

import "fmt"

func main() {
    price := 1000

    if true {
        price := 500
        fmt.Println("Inside block:", price)
    }

    fmt.Println("Outside block:", price)
}
```

**Output:**

```text
Inside block: 500
Outside block: 1000
```

### ব্যাখ্যা

1. `price := 1000` বাইরের Scope-এ একটি Variable তৈরি করেছে।
2. `price := 500` `if` Block-এর ভিতরে নতুন Variable তৈরি করেছে।
3. Inner Variable বাইরের `price`-কে Shadow করেছে।
4. `if` Block শেষ হলে বাইরের `price`-এর Value এখনও `1000` থাকে।

**গুরুত্বপূর্ণ:** Shadowing করলে Outer Variable-এর Value স্বয়ংক্রিয়ভাবে পরিবর্তিত হয় না।

## 2. `:=` বনাম `=`

Variable Shadowing বোঝার জন্য এই দুইটি Operator-এর পার্থক্য জানা জরুরি।

| Operator | কাজ                                            |
| -------- | ---------------------------------------------- |
| `:=`     | নতুন Variable ঘোষণা করে এবং Value নির্ধারণ করে |
| `=`      | আগে থেকে থাকা Variable-এর Value পরিবর্তন করে   |

### Example 2: `:=` ব্যবহার করে Shadowing

```go
package main

import "fmt"

func main() {
    balance := 5000

    if true {
        balance := 2000
        fmt.Println("Inner balance:", balance)
    }

    fmt.Println("Outer balance:", balance)
}
```

**Output:**

```text
Inner balance: 2000
Outer balance: 5000
```

এখানে `balance := 2000` একটি নতুন Inner Variable তৈরি করেছে।

### Example 3: `=` ব্যবহার করে Outer Variable পরিবর্তন

```go
package main

import "fmt"

func main() {
    balance := 5000

    if true {
        balance = 2000
        fmt.Println("Inner balance:", balance)
    }

    fmt.Println("Outer balance:", balance)
}
```

**Output:**

```text
Inner balance: 2000
Outer balance: 2000
```

এখানে `balance = 2000` নতুন Variable তৈরি করেনি। এটি বাইরের `balance`-এর Value পরিবর্তন করেছে।

**মনে রাখবে:**

- `:=` — নতুন Variable ঘোষণা।
- `=` — বিদ্যমান Variable-এ Assignment।
- কোন Variable ব্যবহৃত হবে, তা নির্ভর করে Scope এবং Declaration-এর ওপর।

## 3. Function Parameter Shadowing

Function-এর Parameter-ও Function Body-এর Scope-এর অংশ। একই নামের Variable Inner Block-এ ঘোষণা করলে Parameter-টিও Shadow হতে পারে।

```go
package main

import "fmt"

func calculatePrice(price float64) {
    fmt.Println("Original price:", price)

    if true {
        price := 500.0
        fmt.Println("Discounted price:", price)
    }

    fmt.Println("Price after block:", price)
}

func main() {
    calculatePrice(1000)
}
```

**Output:**

```text
Original price: 1000
Discounted price: 500
Price after block: 1000
```

এখানে Inner `price` নতুন Variable। এটি Function Parameter-এর Value পরিবর্তন করে না।

## 4. Variable Shadowing with `if` Initializer

Go-তে `if` Statement-এর Initializer-এ Variable ঘোষণা করা যায়। এই Variable `if` Statement এবং সংশ্লিষ্ট `else` Branch-এর Scope-এ উপলব্ধ থাকে।

```go
package main

import "fmt"

func main() {
    status := "pending"

    if status := "approved"; status == "approved" {
        fmt.Println("Inside if:", status)
    } else {
        fmt.Println("Inside else:", status)
    }

    fmt.Println("Outside if:", status)
}
```

**Output:**

```text
Inside if: approved
Outside if: pending
```

### কেন এমন হলো?

- বাইরের `status`-এর Value হলো `"pending"`।
- `if status := "approved"` নতুন একটি Variable তৈরি করেছে।
- `if` ও `else` Branch-এ Inner `status` ব্যবহৃত হয়েছে।
- `if` Statement শেষ হলে বাইরের `status` আবার দৃশ্যমান হয়েছে।

## 5. Shadowing with `for` Loops

Loop-এর ভিতরে একই নামের Variable ঘোষণা করলেও Shadowing হতে পারে।

```go
package main

import "fmt"

func main() {
    count := 10

    for count := 1; count <= 3; count++ {
        fmt.Println("Loop count:", count)
    }

    fmt.Println("Outer count:", count)
}
```

**Output:**

```text
Loop count: 1
Loop count: 2
Loop count: 3
Outer count: 10
```

এখানে `for` Statement-এর `count := 1` একটি নতুন Loop Variable তৈরি করেছে। তাই Loop শেষ হলে বাইরের `count`-এর Value `10` থাকে।

## 6. Real-World Example: Ticket Booking

ধরা যাক, একটি Ticket Booking Application-এ মূল Ticket Price এবং Discounted Price হিসাব করতে হবে।

```go
package main

import "fmt"

func main() {
    ticketPrice := 1000.0
    discountRate := 0.10

    if ticketPrice > 500 {
        ticketPrice := ticketPrice * (1 - discountRate)
        fmt.Println("Discounted ticket price:", ticketPrice)
    }

    fmt.Println("Original ticket price:", ticketPrice)
}
```

**Output:**

```text
Discounted ticket price: 900
Original ticket price: 1000
```

### সমস্যা কোথায়?

এই লাইনটি:

```go
ticketPrice := ticketPrice * (1 - discountRate)
```

Inner Scope-এ নতুন `ticketPrice` তৈরি করেছে। ফলে বাইরের Variable পরিবর্তিত হয়নি।

### আরও পরিষ্কার সমাধান

```go
package main

import "fmt"

func main() {
    ticketPrice := 1000.0
    discountRate := 0.10

    if ticketPrice > 500 {
        discountedPrice := ticketPrice * (1 - discountRate)
        fmt.Println("Discounted ticket price:", discountedPrice)
    }

    fmt.Println("Original ticket price:", ticketPrice)
}
```

এখানে `discountedPrice` নামটি উদ্দেশ্য পরিষ্কার করে এবং অনিচ্ছাকৃত Shadowing এড়ায়।

## 7. Common Mistake: Error Variable Shadowing

Go-তে Error Handling-এর সময় Shadowing একটি গুরুত্বপূর্ণ সমস্যা তৈরি করতে পারে।

### ভুল উদাহরণ

```go
package main

import (
    "errors"
    "fmt"
)

func processOrder() error {
    err := errors.New("database connection failed")

    if true {
        err := errors.New("payment failed")
        fmt.Println("Inner error:", err)
    }

    return err
}

func main() {
    fmt.Println(processOrder())
}
```

**Output:**

```text
Inner error: payment failed
database connection failed
```

### কেন এটি সমস্যা?

Inner Scope-এর `err :=` নতুন Error Variable তৈরি করেছে। ফলে `return err` বাইরের Error-টিই Return করেছে, Inner Error নয়।

এটি কখনও কখনও Bug তৈরি করতে পারে, বিশেষ করে Nested `if` Block বা Error Handling-এর সময়।

### ভালো অভ্যাস

যদি নতুন Error Variable প্রয়োজন হয়, তাহলে এমন নাম ব্যবহার করো যা তার উদ্দেশ্য স্পষ্ট করে।

```go
package main

import (
    "errors"
    "fmt"
)

func processOrder() error {
    err := errors.New("database connection failed")

    if true {
        paymentErr := errors.New("payment failed")
        fmt.Println("Inner error:", paymentErr)
    }

    return err
}

func main() {
    fmt.Println(processOrder())
}
```

এখানে `paymentErr` নামটি পরিষ্কার করে যে এটি Payment-সম্পর্কিত Error।

## 8. Shadowing শনাক্ত করার উপায়

Go-এর Shadowing সমস্যা ধরতে নিচের অভ্যাসগুলো অনুসরণ করতে পারো:

- Inner Scope-এ অপ্রয়োজনীয়ভাবে একই Variable Name ব্যবহার কোরো না।
- `err`, `result`, `value`-এর মতো সাধারণ নাম ব্যবহারের সময় সতর্ক থাকো।
- `:=` ব্যবহারের আগে যাচাই করো নতুন Variable তৈরি হচ্ছে কি না।
- Variable-এর উদ্দেশ্য বোঝায় এমন নাম ব্যবহার করো।
- Nested Block-এ Return Value এবং Error Handling ভালোভাবে পরীক্ষা করো।

Go-এর `go vet` এবং Static Analysis Tools কিছু ধরনের সমস্যায় সাহায্য করতে পারে, তবে সব ধরনের Variable Shadowing স্বয়ংক্রিয়ভাবে শনাক্ত করবে—এমন নিশ্চয়তা নেই।

## 9. Practice Exercises

### Exercise 1: Output Predict করো

```go
package main

import "fmt"

func main() {
    value := 100

    if true {
        value := 50
        value += 10
        fmt.Println(value)
    }

    fmt.Println(value)
}
```

প্রশ্ন: Output কী হবে? কেন?

### Exercise 2: Bug Fix করো

```go
package main

import "fmt"

func main() {
    total := 100

    if true {
        total := total + 50
        fmt.Println("Updated total:", total)
    }

    fmt.Println("Final total:", total)
}
```

লক্ষ্য: বাইরের `total`-এর Value `150` করা, যাতে Final Output-ও `150` হয়।

### Exercise 3: Error Handling

একটি Function তৈরি করো যেখানে:

1. বাইরের Scope-এ `err` ঘোষণা করা হবে।
2. একটি Nested Block-এ নতুন Error তৈরি হবে।
3. Shadowing-এর কারণে কোন Error Return হচ্ছে তা পরীক্ষা করবে।
4. Variable Name বা Assignment পরিবর্তন করে সমস্যাটি সমাধান করবে।

## 10. Key Takeaways

- Variable Shadowing ঘটে যখন Inner Scope-এর একই নামের Variable Outer Variable-কে আড়াল করে।
- `:=` নতুন Variable তৈরি করতে পারে।
- `=` বিদ্যমান Variable-এর Value পরিবর্তন করে।
- Shadowing করলে Outer Variable-এর Value সাধারণত অপরিবর্তিত থাকে।
- Error Handling-এর সময় Shadowing বিশেষভাবে সতর্কতার সঙ্গে ব্যবহার করতে হয়।
- Production Code-এ স্পষ্ট Variable Name ব্যবহার করলে Bug কমানো সহজ হয়।

## Next Lesson

**[Golang] 019 — Constants in Go**

পরবর্তী লেসনে শিখবে Go-তে Constants, `const` Keyword, Typed ও Untyped Constants এবং বাস্তব প্রজেক্টে Constant ব্যবহারের নিয়ম।
