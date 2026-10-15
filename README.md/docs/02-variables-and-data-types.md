Go Programming বাংলায়: Beginner থেকে Production-Level Engineering
আজকের লেসনে আমরা Go-এর দুটি মৌলিক বিষয় শিখব: Variables (ভেরিয়েবল) এবং Data Types (ডেটা টাইপ)।
তুমি যেহেতু go-learn-with-zamirul repository তৈরি করছ, এই লেসনটি এমনভাবে সাজাচ্ছি যাতে তুমি এটি documentation হিসেবে docs/008-variables-and-data-types.md ফাইলে রাখতে পারো।
একজন Senior Software Engineer হিসেবে আমি শুধু syntax শেখাব না—কোন পরিস্থিতিতে কোন type ব্যবহার করবে, কী কী ভুল হতে পারে এবং production code-এ কেন সঠিক type নির্বাচন গুরুত্বপূর্ণ, সেটিও ব্যাখ্যা করব।
১. Variable কী?
Variable হলো একটি নামযুক্ত storage location, যার মাধ্যমে আমরা program-এর মধ্যে data ব্যবহার করি।
ধরো, তুমি একটি Ticket Booking System তৈরি করছ। সেখানে user-এর নাম, ticket-এর সংখ্যা, ticket-এর দাম এবং booking সফল হয়েছে কি না—এসব তথ্য সংরক্ষণ করতে হবে।
package main

import "fmt"

func main() {
var name string = "Zamirul"
var tickets int = 3
var price float64 = 500.50
var booked bool = true

    fmt.Println(name)
    fmt.Println(tickets)
    fmt.Println(price)
    fmt.Println(booked)

}

Output:
Zamirul
3
500.5
true

এখানে:

- name একটি variable, যার type string।
- tickets একটি variable, যার type int।
- price একটি variable, যার type float64।
- booked একটি variable, যার type bool।
  Go একটি statically typed language। অর্থাৎ variable-এর type compile time-এ নির্ধারিত হয়। Type inference ব্যবহার করলে সব সময় type সরাসরি লিখতে হয় না, কিন্তু Go compiler type যাচাই করে।

২. Variable Declare করার নিয়ম
Go-তে variable declare করার কয়েকটি সাধারণ পদ্ধতি রয়েছে।
পদ্ধতি ১: var দিয়ে declaration
var name string
var age int
var salary float64

এখানে শুধু variable declare করা হয়েছে; কোনো initial value দেওয়া হয়নি।
Go স্বয়ংক্রিয়ভাবে প্রতিটি variable-কে তার type-এর zero value দেয়।
package main

import "fmt"

func main() {
var name string
var age int
var salary float64
var active bool

    fmt.Printf("Name: %q\n", name)
    fmt.Println("Age:", age)
    fmt.Println("Salary:", salary)
    fmt.Println("Active:", active)

}

Output:
Name: ""
Age: 0
Salary: 0
Active: false

Zero value-গুলো মনে রাখবে:
Data type Zero value
string ""
int 0
float64 0
bool false
Pointer nil
Slice nil
Map nil
Channel nil
Interface nil
Struct-এর প্রতিটি field-ও নিজ নিজ type-এর zero value পায়।
Senior tip: Go-তে অনেক type-এর zero value সরাসরি ব্যবহারযোগ্য। তবে nil map-এ read করা যায়, কিন্তু initialized না করা nil map-এ write করা যায় না। এই পার্থক্য পরে collections শেখার সময় বিস্তারিত দেখব।
পদ্ধতি ২: Declaration এবং initialization একসঙ্গে
var name string = "Zamirul"
var age int = 25

এখানে variable declare করার সময় initial value দেওয়া হয়েছে।
পদ্ধতি ৩: Type inference
var name = "Zamirul"
var age = 25
var price = 499.99
var active = true

এখানে Go compiler initial value থেকে type নির্ধারণ করে।
Expression Inferred type
"Zamirul" string
25 int
499.99 float64
true bool
Type inference মানে Go dynamic typing নয়। একবার variable-এর type নির্ধারিত হলে সেটি অন্য type-এর value সরাসরি গ্রহণ করতে পারবে না।
var age = 25
// age = "twenty-five" // Compile error

পদ্ধতি ৪: Short variable declaration — :=
Go-তে এটি খুবই প্রচলিত।
name := "Zamirul"
age := 25
price := 499.99
active := true

:= ব্যবহার করে variable declare এবং initialize দুটোই করা হয়।
এটি সাধারণত function-এর ভেতরে ব্যবহার করা হয়।
func main() {
message := "Learning Go"
fmt.Println(message)
}

মনে রাখবে:
name := "Zamirul" // নতুন variable
name = "Kabir" // আগের variable-এর value পরিবর্তন

:= নতুন variable declaration-এর জন্য, আর = assignment-এর জন্য।
গুরুত্বপূর্ণ: Package-level variable-এ := ব্যবহার করা যায় না।
package main

// name := "Zamirul" // Compile error

var name = "Zamirul" // সঠিক

৩. Multiple Variables একসঙ্গে declare করা
একাধিক variable একসঙ্গে declare করা যায়।
package main

import "fmt"

func main() {
var name, city string = "Zamirul", "Sirajganj"

    age, score := 25, 95.5

    fmt.Println(name, city)
    fmt.Println(age, score)

}

আরও পরিষ্কারভাবে group declaration করা যায়:
var (
appName = "Ticket Booking API"
environment = "development"
port = 8080
)

এটি configuration-এর মতো সম্পর্কিত variable-গুলো একসঙ্গে রাখতে কাজে আসে।
তবে production application-এ configuration-এর মান শুধু variable declare করলেই নিরাপদভাবে পরিচালিত হয় না; environment variables, validation এবং configuration loading-ও প্রয়োজন হতে পারে।

৪. Go Data Types — বিস্তারিত পরিচিতি
Data type নির্ধারণ করে একটি variable কী ধরনের data রাখবে এবং সেই data-এর ওপর কী ধরনের operation করা যাবে।
Go-এর data type-গুলোকে শেখার সুবিধার জন্য কয়েকটি ভাগে ভাগ করতে পারি।

1. Basic Types
   Integers, floating-point numbers, strings, booleans
   int
   int64
   float64
   string
   bool

2. Aggregate Types
   একাধিক value বা field সংগঠিত করে
   array
   struct

3. Reference-like Types
   অন্য data বা collection-এর সঙ্গে কাজ করতে ব্যবহৃত হয়
   pointer
   slice
   map
   channel
   function

4. Other Important Types
   interface
   rune
   complex64
   complex128

এখন আমরা সবচেয়ে বেশি ব্যবহৃত type-গুলো বিস্তারিত শিখব।
৫. Integer Types — পূর্ণসংখ্যা
Integer হলো এমন সংখ্যা যার মধ্যে fractional বা দশমিক অংশ নেই।
উদাহরণ:
age := 25
quantity := 100
temperature := -5

Go-তে signed এবং unsigned integer type রয়েছে।
Type কাজ
int সাধারণ পূর্ণসংখ্যা
int8 ৮-bit signed integer
int16 ১৬-bit signed integer
int32 ৩২-bit signed integer
int64 ৬৪-bit signed integer
uint unsigned integer
uint8 ৮-bit unsigned integer
uint16 ১৬-bit unsigned integer
uint32 ৩২-bit unsigned integer
uint64 ৬৪-bit unsigned integer
uintptr Pointer-এর bit representation ধারণের উপযোগী integer type
Signed integer positive, negative এবং zero ধারণ করতে পারে। Unsigned integer শুধু শূন্য ও positive মান ধারণ করে।
Integer-এর range
Type Minimum Maximum
int8 -128 127
uint8 0 255
int16 -32,768 32,767
uint16 0 65,535
int32 -2,147,483,648 2,147,483,647
uint32 0 4,294,967,295
int এবং uint-এর size platform অনুযায়ী সাধারণত 32-bit বা 64-bit হয়। int64-এর size সব supported platform-এ 64-bit।
Example
package main

import "fmt"

func main() {
var age int = 25
var stock int32 = 1000
var port uint16 = 8080

    fmt.Println("Age:", age)
    fmt.Println("Stock:", stock)
    fmt.Println("Port:", port)

}

কখন কোন integer ব্যবহার করবে?

- সাধারণ গণনা ও loop: int
- বড় integer বা নির্দিষ্ট ৬৪-bit value: int64
- নির্দিষ্ট binary format বা সীমিত range: int8, uint8 ইত্যাদি
- Negative value কখনো হবে না—এমন ক্ষেত্রেও uint বেছে নেওয়ার আগে সতর্ক হও। Unsigned arithmetic ও underflow ভুলের কারণ হতে পারে।
  Senior tip: শুধু memory বাঁচানোর জন্য সব জায়গায় int8 বা uint8 ব্যবহার করবে না। আগে data-এর অর্থ, range এবং API/Database-এর requirement বোঝো।

৬. Floating-Point Types — দশমিক সংখ্যা
দশমিক বা fractional সংখ্যা রাখার জন্য Go-তে দুটি প্রধান floating-point type রয়েছে:

- float32
- float64
  package main

import "fmt"

func main() {
var height float32 = 1.75
var productPrice float64 = 499.99

    fmt.Println("Height:", height)
    fmt.Println("Price:", productPrice)

}

Type না লিখলে দশমিক numeric literal থেকে infer করা variable সাধারণত float64 হয়।
price := 499.99 // float64

গুরুত্বপূর্ণ বিষয়: Floating-point precision
এই code দেখো:
package main

import "fmt"

func main() {
result := 0.1 + 0.2
fmt.Println(result)
}

তুমি হয়তো 0.3 আশা করছ, কিন্তু floating-point representation-এর কারণে ফলাফলটি 0.3-এর খুব কাছাকাছি একটি মান হতে পারে।
কারণ কম্পিউটার অনেক দশমিক fraction-কে binary floating-point-এ একদম নির্ভুলভাবে সংরক্ষণ করতে পারে না।
Production-এ কী করবে?
Scientific calculations, measurements এবং graphics-এ floating-point সাধারণত উপযুক্ত।
কিন্তু financial application-এ টাকা হিসাবের জন্য সরাসরি float64 ব্যবহার করলে precision সমস্যা হতে পারে।
উদাহরণস্বরূপ, ticket booking system-এ ticket price যদি ৫০০ টাকা হয়, তাহলে ছোট unit-এ integer হিসেবে সংরক্ষণ করা যায়:
package main

import "fmt"

func main() {
var priceInPaisa int64 = 50000

    fmt.Println("Price in paisa:", priceInPaisa)
    fmt.Println("Price in taka:", priceInPaisa/100)

}

এখানে ১০০ পয়সা = ১ টাকা ধরে হিসাব করা হয়েছে।
তবে fractional paisa, currency conversion, tax এবং rounding-এর মতো জটিল ক্ষেত্রে উপযুক্ত decimal arithmetic ও সুস্পষ্ট rounding rules প্রয়োজন হতে পারে।

৭. String Type — Text Data
String হলো text সংরক্ষণের জন্য ব্যবহৃত type।
package main

import "fmt"

func main() {
var name string = "Zamirul"
city := "Sirajganj"

    fmt.Println(name)
    fmt.Println(city)

}

Go-তে string UTF-8 text ধারণ করতে পারে, তবে একটি গুরুত্বপূর্ণ বিষয় আছে: string-এর indexing byte অনুযায়ী হয়, character অনুযায়ী নয়।
String Concatenation
দুটি string + দিয়ে যুক্ত করা যায়।
package main

import "fmt"

func main() {
firstName := "Zamirul"
lastName := "Kabir"

    fullName := firstName + " " + lastName

    fmt.Println(fullName)

}

Output:
Zamirul Kabir

String-এর Length
package main

import "fmt"

func main() {
name := "Go"

    fmt.Println(len(name))

}

Output:
2

এখানে len() string-এর byte-এর সংখ্যা দেয়।
বাংলা text-এর ক্ষেত্রে এটি গুরুত্বপূর্ণ:
package main

import "fmt"

func main() {
text := "বাংলা"

    fmt.Println(len(text))

}

এখানে output হবে UTF-8 encoded byte-এর সংখ্যা, বাংলা অক্ষরের সংখ্যা নয়।
Rune দিয়ে Unicode character গণনা
package main

import "fmt"

func main() {
text := "বাংলা"

    runes := []rune(text)

    fmt.Println("Bytes:", len(text))
    fmt.Println("Runes:", len(runes))

}

- len(text) byte-এর সংখ্যা দেয়।
- len([]rune(text)) Unicode code point-এর সংখ্যা দেয়।
  মনে রাখবে, Unicode code point এবং মানুষের চোখে দেখা character সব সময় এক নয়। কিছু character একাধিক code point দিয়ে তৈরি হয়।
  Double Quotes এবং Backticks
  name := "Zamirul"
  message := "Hello\nWorld"

query := `SELECT * FROM users
WHERE active = true`

- " " হলো interpreted string literal; \n newline বোঝায়।
- Backtick-এর মধ্যে raw string literal লেখা যায়; newline ও backslash সাধারণ text হিসেবে থাকে।

৮. Rune Type — Unicode Code Point
Go-তে rune হলো int32-এর একটি alias। এটি একটি Unicode code point উপস্থাপন করতে ব্যবহৃত হয়।
package main

import "fmt"

func main() {
var english rune = 'A'
var bangla rune = 'অ'

    fmt.Println(english)
    fmt.Println(bangla)

    fmt.Printf("%c\n", english)
    fmt.Printf("%c\n", bangla)

}

Output:
65
2437
A
অ

এখানে:

- 'A' একটি rune literal, যার Unicode code point 65।
- 'অ' একটি rune literal, যার Unicode code point 2437।
- %c integer-টিকে সংশ্লিষ্ট character হিসেবে print করে।
  একটি গুরুত্বপূর্ণ পার্থক্য:
  letter := 'A' // rune, অর্থাৎ int32 alias
  text := "A" // string

Single quotes এবং double quotes একই ধরনের value তৈরি করে না।

৯. Boolean Type — True অথবা False
Boolean type-এর মাত্র দুটি value:

- true
- false
  package main

import "fmt"

func main() {
isLoggedIn := true
hasTicket := false

    fmt.Println(isLoggedIn)
    fmt.Println(hasTicket)

    if isLoggedIn {
        fmt.Println("User is logged in")
    }

}

Backend application-এ Boolean ব্যবহারের উদাহরণ:
isActive := true
emailVerified := false
paymentCompleted := true

Boolean variable condition, authorization decision এবং status tracking-এ কাজে লাগে।
তবে isAdmin, isAuthorized বা paymentCompleted-এর মতো security-critical state শুধু client-এর পাঠানো Boolean থেকে বিশ্বাস করবে না। Server-side validation এবং database-backed authorization প্রয়োজন।

১০. Constants — const
const দিয়ে compile-time constant declare করা হয়। Constant-এর মান program চলার সময় পরিবর্তন করা যায় না।
package main

import "fmt"

const AppName = "Ticket Booking API"
const MaxTickets = 10

func main() {
fmt.Println(AppName)
fmt.Println(MaxTickets)
}

একসঙ্গে constants
const (
StatusPending = "pending"
StatusConfirmed = "confirmed"
StatusCancelled = "cancelled"
)

এভাবে status-এর মতো নির্দিষ্ট value এক জায়গায় রাখা যায়।
তবে constants-এর নাম meaningful হওয়া উচিত এবং একই business concept-এর জন্য একাধিক জায়গায় আলাদা literal লিখে duplication তৈরি করা উচিত নয়।
const এবং var-এর পার্থক্য
const var
Constant value পরিবর্তনযোগ্য variable
Compile-time constant expression দরকার Runtime-এও value নির্ধারণ হতে পারে
Function call-এর result সরাসরি constant হিসেবে নেওয়া যায় না Function call-এর result রাখা যায়
সাধারণত primitive constant value যেকোনো valid variable type
Go-তে const দিয়ে array, slice, map বা struct-এর মতো composite value declare করা যায় না।

১১. Type Conversion — এক Data Type থেকে অন্য Type
Go-তে একটি গুরুত্বপূর্ণ নিয়ম হলো: ভিন্ন numeric type-এর মধ্যে automatic conversion সাধারণত হয় না। প্রয়োজনে explicit conversion করতে হবে।
ভুল উদাহরণ:
var age int = 25
var total int64

// total = age // Compile error

সঠিক উদাহরণ:
package main

import "fmt"

func main() {
var age int = 25
var total int64 = int64(age)

    fmt.Println(total)

}

এখানে int64(age) দিয়ে int থেকে int64-এ conversion করা হয়েছে।
Integer থেকে Float
var quantity int = 5
var price float64 = 99.50

total := float64(quantity) \* price

fmt.Println(total)

Output:
497.5

Float থেকে Integer
var price float64 = 99.99

wholePrice := int(price)

fmt.Println(wholePrice)

Output:
99

এখানে fractional অংশ বাদ পড়েছে। এটি rounding নয়।
Senior Engineer warning: Narrower integer type-এ conversion করলে value representable range-এর বাইরে হলে ফলাফল প্রত্যাশার বাইরে হতে পারে। Conversion নিজে থেকেই সব ক্ষেত্রে range validation করে না।

১২. Type Inference বনাম Explicit Typing
দুটিই Go-তে বৈধ।
var userName string = "Zamirul"

age := 25

প্রথমটিতে type স্পষ্টভাবে লেখা হয়েছে। দ্বিতীয়টিতে compiler type অনুমান করেছে।
কখন কোনটি ব্যবহার করবে?

- Type স্পষ্ট হলে এবং declaration সহজ হলে := ব্যবহার করতে পারো।
- Zero value দিয়ে শুরু করতে হলে var ব্যবহার করো।
- নির্দিষ্ট numeric type দরকার হলে type স্পষ্ট করে লেখা ভালো।
- Public API, domain model এবং database boundary-তে type নির্বাচন সচেতনভাবে করো।
  var retryCount int
  var timeoutSeconds int64 = 30

কোডকে অপ্রয়োজনীয়ভাবে verbose করা যেমন উচিত নয়, তেমনি গুরুত্বপূর্ণ type-এর অর্থ অস্পষ্ট রাখাও উচিত নয়।

১৩. fmt Package — Variable Print ও Debugging
Go-তে fmt package output এবং formatting-এর জন্য অনেক কাজে লাগে।
package main

import "fmt"

func main() {
name := "Zamirul"
age := 25
active := true

    fmt.Println(name, age, active)

    fmt.Printf("Name: %s\n", name)
    fmt.Printf("Age: %d\n", age)
    fmt.Printf("Active: %t\n", active)

    fmt.Printf("Name type: %T\n", name)
    fmt.Printf("Age type: %T\n", age)

}

দরকারি formatting verbs
Verb ব্যবহার
%v Default format-এ value
%T Value-এর type
%s String
%d Decimal integer
%f Floating-point number
%.2f দশমিকের পরে দুই ঘর
%t Boolean
%q Quoted string
%c Unicode code point থেকে character
উদাহরণ:
price := 499.999

fmt.Printf("%.2f\n", price)

Output:
500.00

এখানে fmt output-কে দুই decimal place-এ format করেছে। কিন্তু display rounding আর financial calculation-এর নির্ভুলতা এক জিনিস নয়।

১৪. Scope — Variable কোথায় ব্যবহার করা যাবে?
Scope বোঝা গুরুত্বপূর্ণ, বিশেষ করে বড় Go project-এ।
Local Variable
func main() {
message := "Hello Go"
fmt.Println(message)
}

message শুধু main() function-এর ভেতরে accessible।
Package-level Variable
package main

var appName = "Ticket API"

একই package-এর অন্য file থেকেও এই variable ব্যবহার করা যায়, যদি সেখানে একই package থাকে।
Exported এবং Unexported Names
var AppName = "Ticket API"
var appVersion = "1.0.0"

- AppName uppercase দিয়ে শুরু হওয়ায় অন্য package থেকে accessible।
- appVersion lowercase দিয়ে শুরু হওয়ায় অন্য package থেকে সরাসরি accessible নয়।
  Go-তে visibility নিয়ন্ত্রণের একটি গুরুত্বপূর্ণ নিয়ম এটি।
  Production tip: সব variable global করে রাখবে না। Global mutable state testing, concurrency এবং maintenance জটিল করতে পারে।

১৫. Common Mistakes — যে ভুলগুলো এড়াবে
Mistake 1
ভুল capitalization
fmt.println("Hello")

সঠিক:
fmt.Println("Hello")

Go case-sensitive।

Mistake 2
ভুল quote
name := 'Zamirul'

সঠিক:
name := "Zamirul"

Single-quoted rune literal একটি Unicode code point-এর জন্য; string-এর জন্য double quotes ব্যবহার করবে।

Mistake 3
Type mismatch
var age int = 25
var count int64 = age

সঠিক:
var count int64 = int64(age)

Mistake 4
Unused variable
func main() {
age := 25
}

Go compiler-এ function-এর ভেতরে declared কিন্তু ব্যবহৃত না হওয়া local variable compile error তৈরি করে।

Mistake 5
ভুল numeric assumption
var price float64 = 99.99
var count int = 2

হিসাবের আগে type conversion ও precision-এর প্রয়োজন বোঝো। Money-এর জন্য floating-point ব্যবহার করার সিদ্ধান্ত সচেতনভাবে নিতে হবে।

১৬. Practice Project — Ticket Booking Data Model
এবার শেখা বিষয়গুলো একসঙ্গে ব্যবহার করো।
package main

import "fmt"

const AppName = "Go Ticket Booking"

type Ticket struct {
ID int
UserName string
Price int64
Quantity int
Paid bool
}

func main() {
ticket := Ticket{
ID: 101,
UserName: "Zamirul",
Price: 50000,
Quantity: 2,
Paid: true,
}

    total := ticket.Price * int64(ticket.Quantity)

    fmt.Println("Application:", AppName)
    fmt.Println("Ticket ID:", ticket.ID)
    fmt.Println("User:", ticket.UserName)
    fmt.Printf("Total: %d paisa\n", total)
    fmt.Println("Paid:", ticket.Paid)

}

এখানে Price-কে পয়সায় integer হিসেবে রাখা হয়েছে এবং Quantity-কে int64-এ convert করে মোট দাম হিসাব করা হয়েছে।
এই উদাহরণে একটি বিষয় মনে রাখবে: int64 arithmetic-এ overflow হতে পারে। Production code-এ input validation এবং overflow protection-এর প্রয়োজনীয়তা বিবেচনা করতে হবে।
১৭. Interactive Quiz — নিজের শেখা যাচাই করো
Go Knowledge Check
0 / 5 answered

১. `age := 25` লিখলে সাধারণত কোন type তৈরি হয়?

int

int64

string

float64

২. Go-তে string-এর `len()` কী গণনা করে?

Character-এর সংখ্যা সবসময়

UTF-8 byte-এর সংখ্যা

শুধু word-এর সংখ্যা

String-এর memory address

৩. `var active bool`-এর zero value কী?

nil

0

false

true

৪. `var x int = 10; var y int64 = x` কেন compile error দেয়?

int64 ব্যবহার করা যায় না

Variable x immutable

Numeric type-এর explicit conversion প্রয়োজন

fmt import করা হয়নি

৫. টাকা হিসাবের ক্ষেত্রে float64 নিয়ে সতর্ক থাকতে হয় কেন?

এটি শুধু negative number রাখে

Floating-point representation-এ precision সমস্যা হতে পারে

এটি string type

এতে কোনো arithmetic করা যায় না

উত্তর যাচাই করো
১৮. আজকের Homework
তোমার go-learn-with-zamirul repository-তে নিচের কাজগুলো করো:

1. docs/008-variables-and-data-types.md ফাইলে আজকের notes সংরক্ষণ করো।
2. examples/variables/main.go তৈরি করে Ticket Booking Data Model-এর code চালাও।
3. int, float64, string, bool এবং rune নিয়ে আলাদা ছোট example লেখো।
4. বাংলা string-এর len() এবং len([]rune(...)) তুলনা করে দেখো।
5. নিজের ভাষায় লিখো: Go-তে type safety কেন গুরুত্বপূর্ণ?
   কোড চালানোর command:
   go run .

একটি directory-তে একাধিক .go file থাকলে go run . পুরো package চালায়। তবে ওই directory-তে একটি মাত্র main entry point রাখবে।
পরের লেসন: Arrays, Slices, Maps এবং Loops—যেখানে আমরা শিখব কীভাবে Go-তে data collection নিয়ে কাজ করতে হয় এবং এগুলো দিয়ে বাস্তব backend logic তৈরি করতে হয়।
