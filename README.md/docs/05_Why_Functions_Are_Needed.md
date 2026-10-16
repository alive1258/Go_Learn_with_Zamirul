— Why Functions Are Needed

Go-তে Functions কেন প্রয়োজন, তা বোঝার সবচেয়ে ভালো উপায় হলো একই problem দুইভাবে solve করা—Function ছাড়া এবং Function ব্যবহার করে।

1. Code Reusability — একই code বারবার না লেখা
   ধরো, তোমাকে তিনজন customer-এর total price calculate করতে হবে।
   Without Function

package main

import "fmt"

func main() {
price1 := 100
quantity1 := 2
fmt.Println(price1 \* quantity1)

    price2 := 200
    quantity2 := 3
    fmt.Println(price2 * quantity2)

    price3 := 150
    quantity3 := 4
    fmt.Println(price3 * quantity3)

}

এখানে একই calculation বারবার লিখতে হচ্ছে।
With Function

package main

import "fmt"

func calculateTotal(price, quantity int) int {
return price \* quantity
}

func main() {
fmt.Println(calculateTotal(100, 2))
fmt.Println(calculateTotal(200, 3))
fmt.Println(calculateTotal(150, 4))
}

Output:
200
600
600

এখন calculation-এর logic একবারই লেখা হয়েছে। প্রয়োজনমতো function call করে ব্যবহার করা যাচ্ছে। 2. Maintainability — সহজে code পরিবর্তন করা
ধরো, প্রতিটি order-এর total price-এর সঙ্গে ১০% service charge যোগ করতে হবে।
Function ব্যবহার করলে এক জায়গায় পরিবর্তন করলেই সব call-এ নতুন calculation ব্যবহার হবে।

package main

import "fmt"

func calculateTotal(price, quantity int) int {
subtotal := price \* quantity
serviceCharge := subtotal / 10

    return subtotal + serviceCharge

}

func main() {
fmt.Println(calculateTotal(100, 2))
fmt.Println(calculateTotal(200, 3))
}

Output:
220
660

এটি integer-based উদাহরণ। বাস্তব payment system-এ অর্থ হিসাবের জন্য floating-point rounding এড়িয়ে উপযুক্ত decimal বা integer minor units ব্যবহার করা ভালো। 3. Readability — code সহজে বোঝা
একটি বড় main() function-এর মধ্যে সব logic লিখলে code বোঝা কঠিন হয়।
Functions ব্যবহার করে কাজগুলো আলাদা করা যায়।

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

Output:
User registered: Zamirul
Registration not allowed

Function-এর নাম দেখেই বোঝা যায় কোন function কী কাজ করে। 4. Testing — সহজে bug খুঁজে বের করা
ছোট function আলাদাভাবে test করা সহজ।
main.go:

package main

func add(a, b int) int {
return a + b
}

main_test.go:

package main

import "testing"

func TestAdd(t \*testing.T) {
got := add(10, 20)
want := 30

    if got != want {
    	t.Errorf("add(10, 20) = %d; want %d", got, want)
    }

}

একই directory-তে দুটি file রাখবে। তারপর চালাবে:
go test

ফলাফল ঠিক থাকলে test pass করবে। Function ছোট ও নির্দিষ্ট দায়িত্বের হলে bug কোথায় হয়েছে তা শনাক্ত করা সহজ হয়। 5. Separation of Concerns — আলাদা কাজ আলাদা function-এ রাখা
একটি booking system-এ বিভিন্ন কাজের জন্য আলাদা function থাকতে পারে।
Create Booking
মূল booking operation

validateBooking()
Input validation

calculatePrice()
Price calculation

saveBooking()
Database operation

processPayment()
Payment processing

প্রতিটি function-এর আলাদা দায়িত্ব থাকলে code maintain ও test করা সহজ হয়। তবে বাস্তবে payment processing এবং database transaction-এর মতো কাজের জন্য অতিরিক্ত security ও consistency ব্যবস্থা দরকার। 6. Functions-এর প্রধান সুবিধা
সুবিধা কেন দরকার
Reusability একই logic বারবার ব্যবহার করা
Maintainability পরিবর্তন সহজ করা
Readability code বোঝা সহজ করা
Testability আলাদাভাবে test করা
Modularity বড় problem-কে ছোট অংশে ভাগ করা
Debugging সমস্যার উৎস শনাক্ত করা
Abstraction ভেতরের implementation না জেনেও function ব্যবহার করা
Practice task
একটি Go program লেখো যেখানে তিনটি function থাকবে:

1. calculateTotal(price, quantity int) int
2. isEligible(age int) bool
3. printWelcome(name string)
   তারপর main() থেকে প্রতিটি function call করবে।
   Senior Developer tip: Function তৈরি করার সময় নিজেকে প্রশ্ন করবে—এই function-এর একটি পরিষ্কার দায়িত্ব আছে কি? সাধারণত ছোট, নির্দিষ্ট দায়িত্বের function পড়তে, test করতে এবং পরিবর্তন করতে সহজ হয়। তবে অকারণে অতিরিক্ত ছোট function তৈরি করাও প্রয়োজন নেই।
