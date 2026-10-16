# [Golang] 015 — Local Scope and Block Scope

## 🎯 Learning Objectives

এই লেসন শেষে তুমি শিখবে:

- Local scope কী এবং কেন প্রয়োজন।
- Block scope কীভাবে কাজ করে।
- Function-এর ভিতরে variable declare ও access করা।
- `if`, `for` এবং `switch` block-এর scope বোঝা।
- Variable shadowing কী এবং কীভাবে এড়িয়ে চলতে হয়।
- Clean, readable এবং maintainable Go code লেখা।

---

## 1. What Is Local Scope?

**Local scope** হলো এমন একটি scope যেখানে কোনো variable নির্দিষ্ট function বা block-এর মধ্যে ব্যবহার করা যায়।

একটি function-এর ভিতরে declare করা local variable সাধারণত সেই function-এর বাইরের code থেকে access করা যায় না।

### Example

```go
package main

import "fmt"

func greet() {
    message := "Hello, Go!"
    fmt.Println(message)
}

func main() {
    greet()
}
```

**Output:**

```text
Hello, Go!
```

এখানে `message` হলো local variable। এটি `greet()` function-এর ভিতরে declare করা হয়েছে।

তাই `greet()` function-এর বাইরে সরাসরি `message` access করা যাবে না।

---

## 2. Local Variable Inside a Function

```go
package main

import "fmt"

func calculateTotal(price, quantity int) int {
    total := price * quantity
    return total
}

func main() {
    total := calculateTotal(500, 3)

    fmt.Println("Total:", total)
}
```

**Output:**

```text
Total: 1500
```

এখানে দুটি আলাদা local variable রয়েছে:

- `total` inside `calculateTotal()`
- `total` inside `main()`

একই নাম হলেও এগুলো আলাদা variable, কারণ এগুলোর scope আলাদা।

**Industry tip:** Function-এর ভিতরে প্রয়োজনীয় variable declare করো। এতে code বোঝা সহজ হয় এবং অপ্রয়োজনীয় shared state এড়ানো যায়।

---

## 3. What Is Block Scope?

Go-তে `{}` দিয়ে তৈরি block-এর মধ্যে declare করা variable সাধারণত সেই block-এর মধ্যেই accessible থাকে।

যেমন:

- Function body
- `if` statement
- `for` loop
- `switch` statement
- Standalone `{}` block

### Example: `if` Block

```go
package main

import "fmt"

func main() {
    age := 20

    if age >= 18 {
        status := "Adult"
        fmt.Println(status)
    }

    // status is not accessible here.
    fmt.Println("Age:", age)
}
```

**Output:**

```text
Adult
Age: 20
```

`status` variable-টি `if` block-এর মধ্যে declare করা হয়েছে। তাই block-এর বাইরে এটি access করা যাবে না।

---

## 4. Variable Scope in a `for` Loop

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 3; i++ {
        message := fmt.Sprintf("Ticket %d", i)
        fmt.Println(message)
    }
}
```

**Output:**

```text
Ticket 1
Ticket 2
Ticket 3
```

এখানে:

- `i` হলো `for` statement-এর scope-এ থাকা variable।
- `message` হলো loop body-এর local variable।
- Loop-এর বাইরে `i` বা `message` ব্যবহার করা যাবে না।

**Best practice:** Loop-এর ভিতরে তৈরি temporary variable যতটা সম্ভব loop-এর ভিতরেই রাখো।

---

## 5. Scope in an `if` Initializer

Go-তে `if` statement-এর শুরুতেই variable declare করা যায়।

```go
package main

import "fmt"

func main() {
    if score := 85; score >= 80 {
        fmt.Println("Grade: A")
    } else {
        fmt.Println("Grade: B")
    }
}
```

**Output:**

```text
Grade: A
```

এখানে `score` variable-টি `if` statement-এর scope-এ রয়েছে। এটি `if` এবং `else` block-এর মধ্যে ব্যবহার করা যায়, কিন্তু পুরো `if-else` statement-এর বাইরে নয়।

এই pattern validation এবং conditional logic-এ বেশ কার্যকর।

---

## 6. Variable Shadowing

**Variable shadowing** ঘটে যখন inner scope-এ একই নামের আরেকটি variable declare করা হয়। এতে বাইরের variable-টি inner scope-এর মধ্যে আড়ালে চলে যায়।

### Example

```go
package main

import "fmt"

func main() {
    message := "Outer message"

    if true {
        message := "Inner message"
        fmt.Println(message)
    }

    fmt.Println(message)
}
```

**Output:**

```text
Inner message
Outer message
```

এখানে `:=` দিয়ে `if` block-এর ভিতরে নতুন `message` তৈরি হয়েছে। এটি বাইরের `message`-কে পরিবর্তন করেনি।

### How to Avoid Accidental Shadowing

যদি বাইরের variable-টি পরিবর্তন করতে চাও, তাহলে একই scope-এ declare করা variable-এ `=` ব্যবহার করো:

```go
package main

import "fmt"

func main() {
    message := "Outer message"

    if true {
        message = "Updated message"
        fmt.Println(message)
    }

    fmt.Println(message)
}
```

**Output:**

```text
Updated message
Updated message
```

মনে রাখবে:

- `:=` নতুন variable declare করে।
- `=` আগে থেকে থাকা assignable variable-এ নতুন value assign করে।

`:=` ব্যবহার করলেই সবসময় shadowing হয় না। এটি নতুন variable declare করে; কোন scope-এ তা হচ্ছে, সেটিই গুরুত্বপূর্ণ।

---

## 7. Practical Example: Ticket Booking

এবার local scope ও block scope একটি বাস্তব booking example-এ ব্যবহার করি।

```go
package main

import "fmt"

const maxTickets = 10

func calculateTotal(price int64, quantity int) (int64, error) {
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

func main() {
    price := int64(50000)
    quantity := 2

    total, err := calculateTotal(price, quantity)
    if err != nil {
        fmt.Println("Booking failed:", err)
        return
    }

    fmt.Printf("Total: %d paisa\n", total)
}
```

**Output:**

```text
Total: 100000 paisa
```

### Scope Breakdown

| Variable                 | Scope                                |
| ------------------------ | ------------------------------------ |
| `maxTickets`             | Package block                        |
| `price` parameter        | `calculateTotal` function            |
| `quantity` parameter     | `calculateTotal` function            |
| `total`                  | `calculateTotal` function            |
| `price` inside `main`    | `main` function                      |
| `quantity` inside `main` | `main` function                      |
| `total` inside `main`    | `main` function                      |
| `err` inside `main`      | `main` function-এর declaration scope |

একই নামের `price`, `quantity` ও `total` আলাদা function-এ থাকলে সেগুলো আলাদা local variable।

**Note:** বাস্তব payment বা booking system-এ বড় সংখ্যার multiplication-এর overflow-ও validate করতে হবে।

---

## 8. Common Mistakes

### Mistake 1: Accessing a Variable Outside Its Scope

```go
if true {
    status := "Success"
}

fmt.Println(status) // Compile error
```

**Fix:** Variable-টি block-এর বাইরে প্রয়োজন হলে উপযুক্ত outer scope-এ declare করো।

### Mistake 2: Accidentally Shadowing a Variable

```go
count := 10

if true {
    count := 20
}

fmt.Println(count) // 10
```

Inner `count` একটি নতুন variable। বাইরের `count` পরিবর্তন হয়নি।

### Mistake 3: Declaring Variables Too Early

অপ্রয়োজনীয়ভাবে function-এর শুরুতে অনেক variable declare না করে যেখানে প্রয়োজন, সেখানে declare করো। এতে scope ছোট ও code পরিষ্কার থাকে।

---

## 9. Run the Code

Repository-এর root directory থেকে:

```powershell
go run ./examples/scope
```

তোমার lesson repository-তে এই lesson-এর জন্য আলাদা folder ব্যবহার করতে চাইলে:

```text
docs/
└── 015-local-scope-and-block-scope.md

examples/
└── scope/
    └── main.go
```

একটি বিষয় মনে রাখবে: `examples/scope/main.go`-তে বর্তমানে আগের scope lesson-এর code থাকলে নতুন practical example দিয়ে replace করতে পারো, অথবা আলাদা example directory ব্যবহার করতে পারো।

---

## 10. Homework

নিজে code লিখে নিচের কাজগুলো করো:

1. `calculateDiscount()` নামে একটি function তৈরি করো।
2. Function-এর ভিতরে `discount` নামে local variable declare করো।
3. `if` block ব্যবহার করে discount-এর পরিমাণ নির্ধারণ করো।
4. একই নামের variable inner scope-এ declare করে shadowing-এর output দেখাও।
5. বাইরের variable পরিবর্তন করতে `:=` ও `=`-এর পার্থক্য পরীক্ষা করো।

### Challenge

কেন নিচের code compile হয় না? কারণটি ব্যাখ্যা করো এবং সঠিক code লেখো।

```go
func main() {
    if true {
        result := "Success"
    }

    fmt.Println(result)
}
```

---

## Key Takeaways

- Local variable নির্দিষ্ট function বা block-এর মধ্যে ব্যবহার করা যায়।
- Block scope `{}`-এর সঙ্গে সম্পর্কিত।
- Inner scope-এর variable outer scope-এর variable-কে shadow করতে পারে।
- `:=` declaration করে; `=` assignment করে।
- ছোট ও স্পষ্ট scope maintainable Go code লিখতে সাহায্য করে।

**Next lesson:** `[Golang] 016 — Package Scope and Visibility (Exported vs Unexported Identifiers)`
