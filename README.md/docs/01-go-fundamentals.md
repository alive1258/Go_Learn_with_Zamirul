# Why Golang? — Go কেন শিখব?

**Go Programming in Bangla: Beginner to Production-Level Engineering**

## 1. Golang কী?

Go, যা Golang নামেও পরিচিত, Google-এ তৈরি একটি open-source, statically typed, compiled programming language। Go ২০০৯ সালে public announcement-এর মাধ্যমে পরিচিত হয়। এর মূল ডিজাইনারদের মধ্যে রয়েছেন Robert Griesemer, Rob Pike এবং Ken Thompson।

Go-এর মূল লক্ষ্য হলো সহজে পড়া যায় এমন কোড, দ্রুত compilation, কার্যকর concurrency এবং maintainable software তৈরি করা।

### Go দিয়ে কী তৈরি করা যায়?

- REST APIs এবং backend services
- Microservices architecture
- Cloud-native applications
- Distributed systems
- DevOps ও infrastructure tools
- Command-line applications (CLI)
- High-concurrency network services

**উদাহরণ:** Docker, Kubernetes এবং Prometheus-এর মতো গুরুত্বপূর্ণ infrastructure projects-এ Go ব্যবহৃত হয়েছে।

## 2. Why Go? — Go কেন শিখব?

ধরো, তুমি একটি e-commerce platform তৈরি করছ, যেখানে একই সময়ে অনেক ব্যবহারকারী পণ্য দেখছে, অর্ডার করছে, পেমেন্ট করছে এবং inventory আপডেট করছে।

এ ধরনের application তৈরি করার সময় কয়েকটি বিষয় গুরুত্বপূর্ণ:

- **Performance:** Application কত দ্রুত কাজ করে?
- **Scalability:** ব্যবহারকারী ও workload বাড়লে system কীভাবে সামলায়?
- **Concurrency:** একই সময়ে একাধিক কাজ কীভাবে পরিচালিত হয়?
- **Maintainability:** কোড পড়া, পরীক্ষা করা ও পরিবর্তন করা কতটা সহজ?
- **Deployment:** Production server-এ application কীভাবে চালু করা হবে?

Go-এর design এই বিষয়গুলোর অনেকগুলো সহজভাবে সমাধান করতে সাহায্য করে।

### Go-এর প্রধান ৬টি সুবিধা

#### 1. Performance

Go source code compile করে native executable তৈরি করা যায়। অনেক backend ও network workload-এ Go ভালো performance দিতে পারে।

তবে বাস্তব performance নির্ভর করে algorithm, database queries, network latency, memory allocation এবং implementation-এর ওপর।

#### 2. Concurrency

Go-তে goroutine, channels এবং synchronization primitives রয়েছে। এগুলো দিয়ে concurrent tasks, background workers এবং network requests পরিচালনা করা যায়।

#### 3. সহজ Deployment

Go সাধারণত native executable binary তৈরি করে। অনেক ক্ষেত্রে আলাদা Go runtime ইনস্টল না করেও application চালানো যায়।

তবে operating system, CPU architecture, external libraries এবং CGO ব্যবহারের বিষয়গুলো deployment-এর সময় বিবেচনা করতে হবে।

#### 4. Simple Syntax

Go ইচ্ছাকৃতভাবে language complexity সীমিত রেখেছে। এর ফলে team-এর সদস্যদের জন্য code review, debugging এবং maintenance সহজ হতে পারে।

#### 5. Cloud ও Microservices

Go-এর networking support, concurrency model এবং deployment সুবিধার কারণে এটি cloud services, APIs ও distributed systems-এর জন্য জনপ্রিয়।

#### 6. Built-in Developer Tools

Go-এর standard toolchain-এ রয়েছে:

- `gofmt` — কোড format করা
- `go test` — automated tests চালানো
- `go vet` — কিছু সম্ভাব্য code issue শনাক্ত করা
- `go mod` — dependencies ও modules পরিচালনা করা
- `go test -bench=.` — benchmark চালানো

## 3. Go vs Node.js + TypeScript

তুমি যদি আগে থেকেই Node.js ও TypeScript জানো, তাহলে Go-এর সঙ্গে তুলনা করে শেখা সহজ হবে।

| বিষয়             | Go                                    | Node.js + TypeScript                                           |
| ----------------- | ------------------------------------- | -------------------------------------------------------------- |
| Execution         | সাধারণত native executable             | Node.js runtime-এ JavaScript চলে                               |
| Type system       | Statically typed                      | JavaScript dynamic; TypeScript compile-time type checking দেয় |
| Concurrency       | Goroutines, channels, synchronization | Event loop, `async/await`, worker threads                      |
| Memory management | Garbage collection                    | V8 garbage collection                                          |
| Deployment        | প্রায়ই একটি executable binary        | সাধারণত Node.js runtime ও application dependencies             |
| Development style | Explicit, minimal ও convention-driven | Expressive এবং বৃহৎ JavaScript ecosystem                       |
| Common use cases  | APIs, cloud services, infrastructure  | Web applications, APIs, full-stack applications                |

**মনে রাখবে:** Go সব ক্ষেত্রে Node.js-এর চেয়ে দ্রুত নয়। I/O-heavy application-এ database এবং network latency-ই প্রধান bottleneck হতে পারে। CPU-intensive workload বা concurrent connections-এর ক্ষেত্রে Go সুবিধা দিতে পারে, কিন্তু বাস্তব benchmark ছাড়া কোনো প্রযুক্তিকে বিজয়ী বলা ঠিক নয়।

### Example: HTTP Request

Node.js:

```javascript
async function getUser() {
  const response = await fetch("https://example.com/users");

  if (!response.ok) {
    throw new Error(`HTTP error: ${response.status}`);
  }

  return await response.json();
}
```

Go:

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	response, err := client.Get(
		"https://example.com/users",
	)
	if err != nil {
		fmt.Println("Request failed:", err)
		return
	}
	defer response.Body.Close()

	if response.StatusCode < 200 ||
		response.StatusCode >= 300 {
		fmt.Println("HTTP error:", response.Status)
		return
	}

	var users any

	if err := json.NewDecoder(
		response.Body,
	).Decode(&users); err != nil {
		fmt.Println("JSON parsing failed:", err)
		return
	}

	fmt.Println(users)
}
```

দুটি উদাহরণই HTTP request পাঠিয়ে JSON response পড়ে। Go উদাহরণে timeout, HTTP status এবং JSON decoding error সামলানো হয়েছে।

**Production note:** বাস্তব API client-এ response size সীমিত করা, request cancellation, retries-এর নীতি এবং structured error handling-ও বিবেচনা করতে হবে।

Go-তে `if err != nil` বারবার দেখবে। এটি Go-এর explicit error-handling style-এর একটি গুরুত্বপূর্ণ অংশ।

## 4. Goroutine কী?

Goroutine হলো Go runtime দ্বারা পরিচালিত lightweight concurrent execution unit।

উদাহরণ:

```go
package main

import (
	"fmt"
	"sync"
)

func sayHello(name string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Hello,", name)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go sayHello("Zamirul", &wg)

	wg.Wait()

	fmt.Println("Main function finished")
}
```

### কী ঘটছে?

1. `sync.WaitGroup` concurrent task-এর completion-এর জন্য অপেক্ষা করতে সাহায্য করে।
2. `wg.Add(1)` একটি task গণনায় যোগ করে।
3. `go sayHello(...)` নতুন goroutine-এ function চালায়।
4. `defer wg.Done()` function শেষ হলে task-এর count কমায়।
5. `wg.Wait()` task শেষ না হওয়া পর্যন্ত অপেক্ষা করে।

এখানে `time.Sleep()`-এর পরিবর্তে synchronization ব্যবহার করা হয়েছে, তাই এটি আগের introductory example-এর তুলনায় বেশি নির্ভরযোগ্য।

**Senior Engineer tip:** Goroutine তৈরি করা সহজ, কিন্তু production system-এ goroutine leak, race conditions, cancellation, resource limits এবং error propagation সামলানো জরুরি। Concurrency মানেই সব সময় parallel execution নয়।

## 5. Go কোথায় ব্যবহার করবে?

### Backend API

`net/http`, Gin, Echo বা Fiber দিয়ে REST APIs তৈরি করা যায়। Authentication, validation, PostgreSQL এবং Redis যুক্ত করা সম্ভব।

### Cloud ও DevOps

Infrastructure automation, monitoring agents, CLI tools এবং container-related software তৈরি করা যায়।

### Distributed Systems

Concurrent workers, message processing services এবং network-intensive applications তৈরি করা যায়।

### Business Applications

POS, inventory management, order processing, ticket booking এবং SaaS platforms-এর backend তৈরি করা যায়।

তুমি যদি Node.js/NestJS থেকে Go-তে আসো, তাহলে শুরুতে একটি **REST API + PostgreSQL + Authentication + Testing** project তৈরি করাই ভালো।

## 6. Go-এর সীমাবদ্ধতা ও Trade-offs

একজন ভালো engineer শুধু সুবিধা নয়, trade-offs-ও বোঝে।

- **Error handling:** অনেক function-এ `if err != nil` লিখতে হয়।
- **Abstraction:** Go-তে interfaces ও generics আছে, কিন্তু অন্য language-এর সব abstraction pattern হুবহু প্রয়োগ করা উচিত নয়।
- **Frontend:** React বা Next.js-এর মতো frontend application তৈরির জন্য Go সাধারণত প্রধান পছন্দ নয়।
- **Garbage collection:** Memory management সহজ করে, তবে latency-sensitive workload-এ GC behavior বিবেচনা করতে হয়।
- **Concurrency:** Syntax সহজ হলেও synchronization, race conditions ও lifecycle management শেখার জন্য অনুশীলন দরকার।
- **Ecosystem differences:** JavaScript ecosystem-এর কোনো library বা framework-এর Go-তে সরাসরি সমতুল্য নাও থাকতে পারে।

Go-কে অন্য সব language-এর replacement হিসেবে না দেখে engineering toolkit-এর একটি শক্তিশালী সংযোজন হিসেবে বিবেচনা করো।

## 7. Go Backend Engineer Learning Roadmap

প্রতিদিন ১–২ ঘণ্টা নিয়মিত অনুশীলন করতে পারো। সময় তোমার বর্তমান অভিজ্ঞতা ও project complexity-এর ওপর নির্ভর করবে।

### Phase 1 — Fundamentals

- Variables, constants ও data types
- Conditions ও loops
- Functions ও multiple return values
- Arrays, slices ও maps
- Pointers

### Phase 2 — Core Go

- Structs ও methods
- Interfaces ও generics
- Packages ও modules
- Errors ও custom error types
- JSON ও file handling
- Standard library

### Phase 3 — Concurrency

- Goroutines ও channels
- `select`
- `sync.WaitGroup` ও mutex
- `context.Context`
- Race detection
- Cancellation ও graceful shutdown

### Phase 4 — Backend Development

- HTTP server ও REST API
- Routing ও middleware
- Request validation
- Authentication ও authorization
- Structured logging
- Configuration management

### Phase 5 — Database & Testing

- PostgreSQL ও SQL
- Database transactions
- Migrations
- Unit ও integration tests
- Mocking where appropriate
- Benchmarks ও profiling

### Phase 6 — Production Engineering

- Ticket Booking API বা Inventory Management API
- Docker ও deployment
- Environment-based configuration
- Graceful shutdown
- Monitoring ও observability
- CI/CD
- Security, reliability ও performance testing

## 8. আজকের প্রথম অনুশীলন

Go install করা থাকলে `main.go` নামে একটি file তৈরি করো।

```go
package main

import "fmt"

func main() {
	name := "Zamirul"
	language := "Go"

	fmt.Println("Hello,", name)
	fmt.Println("I am learning", language)
}
```

চালাও:

```bash
go run main.go
```

Expected output:

```text
Hello, Zamirul
I am learning Go
```

### Code breakdown

- `package main` — executable program-এর package।
- `import "fmt"` — formatted input/output-এর standard library package।
- `func main()` — executable program-এর entry point।
- `:=` — নতুন local variable declare ও initialize করে।
- `fmt.Println()` — output print করে।

## 9. Learning Goals

এই repository-তে আমি ধাপে ধাপে শিখব:

- Go fundamentals ও idiomatic coding
- Clean, readable এবং maintainable code
- REST APIs ও backend architecture
- PostgreSQL, authentication ও authorization
- Concurrency এবং error handling
- Automated testing ও benchmarking
- Docker, CI/CD ও production deployment
- Real-world backend projects

## Conclusion

Go শেখার উদ্দেশ্য শুধু নতুন syntax শেখা নয়। লক্ষ্য হলো এমন software তৈরি করতে শেখা, যা নির্ভরযোগ্য, maintainable, testable এবং production workload সামলাতে সক্ষম।

এই repository-তে প্রতিটি lesson-এর সঙ্গে practical code examples, exercises এবং বাস্তব backend engineering concepts যোগ করা হবে।

**Let's learn Go step by step!** 💚
