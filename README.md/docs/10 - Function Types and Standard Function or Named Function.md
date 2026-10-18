# [Golang] 10 — Function Types, Standard Functions, and Named Functions

> **লেভেল:** Beginner  
> **বিষয়:** Go Fundamentals  
> **পূর্বশর্ত:** Variables, Scope, Variable Shadowing, Functions  
> **লক্ষ্য:** Go-তে Function কী, Function Type কী, Named Function কী এবং Standard Library-এর Function কীভাবে ব্যবহার করতে হয় তা শেখা।

## ১. Function কী?

Go-তে Function হলো একটি নির্দিষ্ট কাজ সম্পন্ন করার জন্য লেখা পুনর্ব্যবহারযোগ্য কোডের একটি অংশ।

একটি Function ব্যবহার করে আমরা একই কোড বারবার না লিখে প্রয়োজন অনুযায়ী সেটিকে Call করতে পারি।

### উদাহরণ

```go
package main

import "fmt"

func greet() {
    fmt.Println("Hello, Go Developer!")
}

func main() {
    greet()
    greet()
}
```

**Output:**

```text
Hello, Go Developer!
Hello, Go Developer!
```

এখানে:

- `func` হলো Function Declaration-এর Keyword।
- `greet` হলো Function-এর নাম।
- `()`-এর মধ্যে Function-এর Parameter লেখা হয়।
- `{}`-এর মধ্যে Function-এর Body থাকে।
- `greet()` লিখে Function-টি Call করা হয়।

## ২. Named Function কী?

যে Function-এর নিজস্ব নাম থাকে এবং সেই নাম ব্যবহার করে Function-টিকে Call করা যায়, তাকে সাধারণভাবে Named Function বলা হয়।

Go-তে সাধারণ Function Declaration এভাবেই লেখা হয়।

### Syntax

```go
func functionName(parameters) returnType {
    // Function body
}
```

### উদাহরণ

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(10, 20)

    fmt.Println("Result:", result)
}
```

**Output:**

```text
Result: 30
```

এখানে `add` একটি Named Function।

`a` এবং `b` হলো Parameter। Function-টি দুইটি Integer যোগ করে `int` Return করে।

### একাধিক Parameter-এর সংক্ষিপ্ত Syntax

```go
func add(a, b int) int {
    return a + b
}
```

একই Type-এর একাধিক Parameter থাকলে প্রতিটির Type আলাদাভাবে লেখার প্রয়োজন নেই।

## ৩. Standard Function বলতে কী বোঝায়?

Go-তে “Standard Function” কথাটি একাধিক অর্থে ব্যবহৃত হতে পারে। তাই দুইটি ধারণা আলাদা করে বোঝা গুরুত্বপূর্ণ।

### A. সাধারণ Function Declaration

কোনো নির্দিষ্ট কাজের জন্য নিজে Function লিখতে পারো।

```go
func calculateSquare(number int) int {
    return number * number
}
```

### B. Go Standard Library-এর Function

Go-এর সঙ্গে সরবরাহ করা Standard Library-তে অনেক আগে থেকে তৈরি Function রয়েছে। এগুলো ব্যবহার করে সাধারণ কাজ সহজে করা যায়।

উদাহরণ:

```go
package main

import (
    "fmt"
    "math"
    "strings"
)

func main() {
    fmt.Println(strings.ToUpper("golang"))
    fmt.Println(math.Sqrt(64))
}
```

**Output:**

```text
GOLANG
8
```

এখানে:

- `fmt.Println()` — Output দেখায়।
- `strings.ToUpper()` — String-এর অক্ষরগুলো Uppercase করে।
- `math.Sqrt()` — একটি সংখ্যার Square Root হিসাব করে।

**মনে রাখবে:** Named Function এবং Standard Library Function একে অপরের বিপরীত শ্রেণি নয়। Standard Library-তেও নামযুক্ত Function থাকে।

## ৪. Function Type কী?

Go-তে Function-ও একটি Type-এর সঙ্গে সম্পর্কিত হতে পারে। কোনো Function-এর Parameter এবং Return Type একই হলে, তাদের Function Type-ও একই হতে পারে।

যেমন:

```go
func(int, int) int
```

এর অর্থ হলো এমন একটি Function, যা দুইটি `int` Parameter গ্রহণ করে এবং একটি `int` Return করে।

### উদাহরণ

```go
package main

import "fmt"

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func main() {
    var operation func(int, int) int

    operation = add
    fmt.Println("Addition:", operation(10, 5))

    operation = subtract
    fmt.Println("Subtraction:", operation(10, 5))
}
```

**Output:**

```text
Addition: 15
Subtraction: 5
```

### ব্যাখ্যা

```go
var operation func(int, int) int
```

এখানে `operation` হলো একটি Variable, যার Type হলো Function Type।

এটি এমন Function গ্রহণ করতে পারে, যার দুইটি `int` Parameter এবং একটি `int` Return Value রয়েছে।

তাই `add` এবং `subtract`—দুটিই `operation`-এ Assign করা যায়।

**গুরুত্বপূর্ণ:** Function-কে Variable-এ Assign করার সময় `add` লিখবে, `add()` নয়। `add()` লিখলে Function-টি Call করা হবে।

## ৫. Function-কে Argument হিসেবে পাঠানো

Go-তে Function-কে অন্য Function-এর Parameter হিসেবেও পাঠানো যায়। একে অনেক সময় Higher-Order Function-এর ধারণার সঙ্গে যুক্ত করা হয়।

### উদাহরণ

```go
package main

import "fmt"

func calculate(a, b int, operation func(int, int) int) int {
    return operation(a, b)
}

func add(a, b int) int {
    return a + b
}

func multiply(a, b int) int {
    return a * b
}

func main() {
    fmt.Println("Addition:", calculate(10, 5, add))
    fmt.Println("Multiplication:", calculate(10, 5, multiply))
}
```

**Output:**

```text
Addition: 15
Multiplication: 50
```

এখানে `calculate()` Function-টি `operation` Parameter-এর মাধ্যমে ভিন্ন Function গ্রহণ করছে।

ফলে একই `calculate()` Function ব্যবহার করে যোগ ও গুণ—দুই ধরনের কাজ করা যাচ্ছে।

## ৬. Function Type-এর জন্য `type` ব্যবহার

Function Type বারবার লিখতে হলে একটি Named Type তৈরি করা যায়।

### উদাহরণ

```go
package main

import "fmt"

type Operation func(int, int) int

func add(a, b int) int {
    return a + b
}

func subtract(a, b int) int {
    return a - b
}

func calculate(a, b int, operation Operation) int {
    return operation(a, b)
}

func main() {
    addition := calculate(20, 10, add)
    subtraction := calculate(20, 10, subtract)

    fmt.Println("Addition:", addition)
    fmt.Println("Subtraction:", subtraction)
}
```

**Output:**

```text
Addition: 30
Subtraction: 10
```

এখানে:

```go
type Operation func(int, int) int
```

`Operation` নামে একটি Named Function Type তৈরি করেছে।

এটি ব্যবহার করলে Function Parameter-এর Type আরও পরিষ্কারভাবে প্রকাশ করা যায়।

## ৭. Named Function বনাম Function Type

| বৈশিষ্ট্য         | Named Function              | Function Type                                 |
| ----------------- | --------------------------- | --------------------------------------------- |
| কী?               | নামসহ Function Declaration  | Function-এর Parameter ও Return Type-এর কাঠামো |
| উদাহরণ            | `func add(a, b int) int`    | `func(int, int) int`                          |
| কাজ               | নির্দিষ্ট কাজ সম্পন্ন করা   | কোন ধরনের Function গ্রহণযোগ্য তা নির্ধারণ করা |
| ব্যবহার           | `add(2, 3)`                 | `var op func(int, int) int`                   |
| Variable-এ Assign | Function-কে Assign করা যায় | উপযুক্ত Function-এর Type নির্ধারণ করে         |

## ৮. বাস্তব উদাহরণ: Ticket Booking Fee

ধরা যাক, একটি Ticket Booking Application-এ ভিন্ন ধরনের Fee হিসাব করতে হবে।

```go
package main

import "fmt"

type FeeCalculator func(float64) float64

func standardFee(price float64) float64 {
    return price * 0.05
}

func premiumFee(price float64) float64 {
    return price * 0.10
}

func calculateTotal(price float64, calculator FeeCalculator) float64 {
    return price + calculator(price)
}

func main() {
    price := 1000.0

    standardTotal := calculateTotal(price, standardFee)
    premiumTotal := calculateTotal(price, premiumFee)

    fmt.Printf("Standard Total: %.2f\n", standardTotal)
    fmt.Printf("Premium Total: %.2f\n", premiumTotal)
}
```

**Output:**

```text
Standard Total: 1050.00
Premium Total: 1100.00
```

### কেন এই পদ্ধতি উপকারী?

- Fee হিসাবের Logic আলাদা রাখা যায়।
- একই Calculation Function পুনর্ব্যবহার করা যায়।
- নতুন Fee Policy যোগ করা সহজ হয়।
- Function-কে Argument হিসেবে পাঠিয়ে আচরণ পরিবর্তন করা যায়।

বাস্তব আর্থিক সিস্টেমে Fee Policy-এর পাশাপাশি Currency, Rounding, Validation এবং Error Handling-এর বিষয়ও বিবেচনা করতে হবে।

## ৯. Anonymous Function-এর সংক্ষিপ্ত পরিচয়

যে Function-এর আলাদা Declaration Name নেই, তাকে Anonymous Function বলা হয়।

```go
package main

import "fmt"

func main() {
    multiply := func(a, b int) int {
        return a * b
    }

    fmt.Println(multiply(4, 5))
}
```

**Output:**

```text
20
```

এখানে Function-টির আলাদা নাম নেই। এটি `multiply` Variable-এ Assign করা হয়েছে।

Named Function এবং Anonymous Function-এর বিস্তারিত পার্থক্য পরবর্তী Advanced Function Lessons-এ শেখা যাবে।

## ১০. সাধারণ ভুল

### ভুল ১: Function-কে Call করে Variable-এ Assign করা

```go
operation := add(10, 20)
```

এটি Function Type Assign করে না। `add()` Call করে তার Return Value Assign করে।

Function-টিকে Assign করতে:

```go
operation := add
```

### ভুল ২: ভিন্ন Signature-এর Function Assign করা

```go
func add(a, b int) int {
    return a + b
}

func greet(name string) {
    fmt.Println(name)
}
```

`add` এবং `greet`-এর Parameter ও Return Type আলাদা। তাই এদের একই Function Type-এর Variable-এ Assign করা যাবে না।

### ভুল ৩: Return Type ভুল বোঝা

```go
func greet() {
    fmt.Println("Hello")
}
```

এই Function কোনো Value Return করে না। তাই এর Return Value কোনো Variable-এ Assign করা যাবে না।

## ১১. Practice Exercises

### Exercise 1 — Named Function

`calculateArea(width, height float64) float64` নামে একটি Function লেখো, যা Rectangle-এর Area Return করবে।

### Exercise 2 — Function Type

`func(int, int) int` Function Type ব্যবহার করে `add` এবং `multiply` Function Assign করে ফলাফল দেখাও।

### Exercise 3 — Custom Function Type

`type DiscountCalculator func(float64) float64` তৈরি করো। এরপর ৫% ও ১০% Discount হিসাব করার Function লিখো।

### Exercise 4 — Higher-Order Function

এমন একটি Function তৈরি করো, যা দুইটি সংখ্যা এবং একটি Operation Function গ্রহণ করে হিসাবের ফলাফল Return করবে।

## ১২. Key Takeaways

- Named Function-এর নিজস্ব নাম থাকে।
- Go Standard Library-তে অনেক প্রস্তুত Function রয়েছে।
- Function Type নির্ধারণ করে Function-এর Parameter ও Return Type কেমন হবে।
- Function-কে Variable-এ Assign করা যায়।
- Function-কে অন্য Function-এর Argument হিসেবেও পাঠানো যায়।
- `type` ব্যবহার করে Named Function Type তৈরি করা যায়।
- Anonymous Function-এর আলাদা Declaration Name থাকে না।

## Next Lesson

**[Golang] 020 — Anonymous Functions and Function Expressions**

পরবর্তী লেসনে Anonymous Function, Function Expression এবং Function-কে Variable-এ সংরক্ষণ করার ব্যবহারিক উদাহরণ শেখা হবে।
