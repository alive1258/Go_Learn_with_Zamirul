# **Why Golang**

Why Golang? — Go কেন শিখব?

## ১. Golang কী?

Go Lang logo Gopher - Go - Sticker | TeePublic

teepublic.com

Go বা Golang হলো Google-এর তৈরি একটি open-source, statically typed, compiled programming language। এটি ২০০৯ সালে public release করা হয়। এর ডিজাইনে simplicity, fast compilation, concurrency এবং maintainability-কে গুরুত্ব দেওয়া হয়েছে।

Go তৈরি করার পেছনে গুরুত্বপূর্ণ ব্যক্তিরা হলেন Robert Griesemer, Rob Pike এবং Ken Thompson।

সহজভাবে বললে, Go এমন একটি ভাষা যা দিয়ে তুমি দ্রুত, নির্ভরযোগ্য এবং অনেক ব্যবহারকারী সামলাতে পারে এমন backend software তৈরি করতে পারবে।

Go দিয়ে তৈরি করা যায়:

- REST API ও Backend Services
- Microservices Architecture
- Cloud-native Applications
- Distributed Systems
- DevOps ও Infrastructure Tools
- CLI (Command-Line Interface) Applications
- High-concurrency Network Services

উদাহরণ হিসেবে Docker, Kubernetes এবং Prometheus-এর মতো জনপ্রিয় infrastructure project-এ Go ব্যবহৃত হয়েছে।

## ২. Why Go? — Go কেন তৈরি করা হয়েছিল?

একজন Senior Software Engineer হিসেবে আমি তোমাকে প্রথমে একটি বাস্তব সমস্যা বুঝিয়ে বলি।

ধরো, একটি কোম্পানির e-commerce platform আছে। প্রতিদিন হাজার হাজার ব্যবহারকারী পণ্য দেখছে, অর্ডার করছে, পেমেন্ট করছে এবং inventory আপডেট করছে।

এই ধরনের application তৈরি করতে আমাদের কয়েকটি বিষয় গুরুত্বপূর্ণভাবে বিবেচনা করতে হয়:

- Performance — অ্যাপ্লিকেশন কত দ্রুত কাজ করছে?
- Scalability — ব্যবহারকারী বাড়লে অ্যাপ্লিকেশন কতটা ভালোভাবে কাজ করছে?
- Concurrency — একই সময়ে অনেকগুলো কাজ কীভাবে সামলানো হচ্ছে?
- Maintainability — কোড ভবিষ্যতে পরিবর্তন করা কতটা সহজ?
- Deployment — Production server-এ application চালু করা কতটা সহজ?

Go-এর ডিজাইন এই সমস্যাগুলোর বেশ কয়েকটির সমাধান সহজ করার চেষ্টা করে।

### Go-এর প্রধান ৬টি সুবিধা

Byte Blitz Yellow Lieanr Circle

১. High Performance

Go সরাসরি machine code-এ compile হয়। তাই অনেক ক্ষেত্রে এটি দ্রুত execution এবং কম runtime overhead দিতে পারে। তবে প্রকৃত performance নির্ভর করে algorithm, database, network ও implementation-এর ওপর।

Concurrent Programming and Low-Level Programming - Scaler Topics

২. Concurrency সহজ

Go-তে goroutine এবং channel রয়েছে। এগুলো ব্যবহার করে একসঙ্গে অনেক কাজ পরিচালনা করা সহজ হয়—যেমন একাধিক API request, background jobs এবং network connections।

단테랩스 — 곽지호 | AI교육 · 자동화 · 데이터사이언스

৩. সহজ Deployment

Go সাধারণত native executable binary তৈরি করে। ফলে অনেক ক্ষেত্রে আলাদা language runtime ইনস্টল না করেই application deploy করা যায়। তবে target operating system ও architecture-এর সঙ্গে binary মিলতে হবে।

Coding Interface on Laptop Screen – Royalty-Free Vector | VectorStock

৪. Simple Syntax

Go ইচ্ছাকৃতভাবে language complexity সীমিত রেখেছে। এর ফলে team-এর সদস্যদের জন্য কোড পড়া, review করা এবং maintain করা তুলনামূলক সহজ হতে পারে।

Giải pháp IoT, MES, Chuyển đổi số | Từ Samitech.vn

৫. Cloud ও Microservices

ছোট binary, standard networking library এবং concurrency support-এর কারণে Go cloud services, APIs ও distributed systems-এর জন্য জনপ্রিয়।

Playwright E2E testing cheatsheet - DEV Community

৬. Built-in Developer Tools

Go-তে formatting (`gofmt`), testing (`go test`), dependency management এবং benchmarking-এর জন্য built-in tooling রয়েছে। আলাদা tools-এর ওপর নির্ভরতা কমিয়ে development workflow সহজ করা যায়।

## ৩. Go vs JavaScript/Node.js — তোমার জন্য গুরুত্বপূর্ণ তুলনা

তুমি যেহেতু Node.js এবং TypeScript জানো, তাই এই তুলনাটি ভালোভাবে বোঝা জরুরি।

| বিষয়             | Go                                    | Node.js + TypeScript                                           |
| ----------------- | ------------------------------------- | -------------------------------------------------------------- |
| Execution         | Compiled native binary                | Node.js runtime-এ JavaScript চলে                               |
| Type system       | Statically typed                      | JavaScript dynamic; TypeScript compile-time type checking দেয় |
| Concurrency       | Goroutines, channels, sync primitives | Event loop, async/await, worker threads                        |
| Memory management | Garbage collection                    | V8 garbage collection                                          |
| Deployment        | সাধারণত একটি executable binary        | সাধারণত Node.js runtime ও application dependencies লাগে        |
| Development style | Explicit ও minimal                    | Expressive; ecosystem অনেক বিস্তৃত                             |
| Common use cases  | Cloud services, APIs, infrastructure  | Web apps, APIs, full-stack applications                        |

মনে রাখবে: Go সব ক্ষেত্রে Node.js-এর চেয়ে দ্রুত নয়। I/O-heavy API, database query কিংবা network latency-র ক্ষেত্রে দুই প্রযুক্তির performance কাছাকাছি হতে পারে। CPU-intensive কাজ বা অনেক concurrent connection-এর ক্ষেত্রে Go ভালো ফল দিতে পারে, কিন্তু benchmark ছাড়া সিদ্ধান্ত নেওয়া উচিত নয়।

### একটি ছোট উদাহরণ

Node.js-এ:

```
async function getUser() {
  const response = await fetch(
    "https://example.com/users"
  );

  return await response.json();
}
```

Go-তে:

```
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

func main() {
    response, err := http.Get(
        "https://example.com/users",
    )
    if err != nil {
        fmt.Println("Request failed:", err)
        return
    }
    defer response.Body.Close()

    var users any

    if err := json.NewDecoder(
        response.Body,
    ).Decode(&users); err != nil {
        fmt.Println("JSON error:", err)
        return
    }

    fmt.Println(users)
}
```

এগুলো একই ধরনের কাজ করে—HTTP request পাঠিয়ে JSON response পড়ে। উদাহরণগুলোতে network error ও JSON parsing error সামলানো হয়েছে; বাস্তব production code-এ timeout, status code এবং response size-ও বিবেচনা করতে হবে।

Go-তে `if err != nil` বারবার দেখবে। এটি Go-এর explicit error-handling style-এর গুরুত্বপূর্ণ অংশ।

## ৪. Goroutine কী? Go-এর সবচেয়ে আকর্ষণীয় বিষয়গুলোর একটি

ধরো, তোমার API-তে ১০০টি request এসেছে। প্রতিটি request-এর জন্য আলাদা কাজ করতে হবে।

Go-তে goroutine দিয়ে concurrent কাজ পরিচালনা করা যায়।

```
package main

import (
    "fmt"
    "time"
)

func sayHello(name string) {
    fmt.Println("Hello,", name)
}

func main() {
    go sayHello("Zamirul")

    time.Sleep(100 * time.Millisecond)

    fmt.Println("Main function finished")
}
```

এখানে:

- `sayHello()` একটি function।
- `go sayHello("Zamirul")` নতুন goroutine-এ function চালায়।
- `time.Sleep()` শুধু এই introductory example-এ goroutine-কে কাজ করার সময় দেওয়ার জন্য।
- বাস্তব application-এ synchronization-এর জন্য `sync.WaitGroup`, channels বা অন্য উপযুক্ত পদ্ধতি ব্যবহার করবে।

Senior Engineer-এর পরামর্শ: Goroutine তৈরি করা সহজ, কিন্তু হাজার হাজার goroutine চালানোর সময় memory, synchronization, cancellation এবং error handling বুঝতে হবে। শুধু `go` keyword ব্যবহার করলেই application স্বয়ংক্রিয়ভাবে scalable হয়ে যায় না।

## ৫. Go কোথায় ব্যবহার করবে?

Building High-Performance Microservices with C# and .NET: A Practical Guide to Scalability and Reliability | by Sidharth CP | Medium

Backend API

`net/http`, Gin, Echo বা Fiber দিয়ে REST API তৈরি করা যায়। Authentication, validation, PostgreSQL এবং Redis যুক্ত করা সম্ভব।

Kubernetes Monitoring | Grafana Cloud | Grafana Labs

Cloud & DevOps

Infrastructure automation, monitoring agents, CLI tools এবং container-related software তৈরি করতে Go উপযোগী।

An overview of distributed systems architectures | by Neeraj Kushwaha | Medium

Distributed Systems

Message processing, concurrent workers এবং network services তৈরি করা যায়।

Inventory Software: The Architecture Choices That Matter

Business Software

POS, inventory, order management, ticket booking এবং SaaS application-এর backend তৈরি করা যায়।

তুমি যদি Node.js/NestJS থেকে Go-তে আসো, তাহলে শুরুতে REST API + PostgreSQL + Authentication + Testing দিয়ে একটি বাস্তব project তৈরি করাই ভালো।

## ৬. Go-এর কিছু সীমাবদ্ধতা

একজন ভালো engineer শুধু সুবিধা নয়, trade-off-ও বোঝে।

- Error handling: অনেক সময় `if err != nil` বারবার লিখতে হয়।
- Generics ও abstraction: আছে, কিন্তু Go-এর design ইচ্ছাকৃতভাবে তুলনামূলক সরল; অন্য language-এর কিছু abstraction pattern সরাসরি প্রয়োগ করা যায় না।
- Frontend ecosystem: React বা Next.js-এর মতো frontend তৈরির জন্য Go সাধারণত পছন্দের প্রধান tool নয়।
- Garbage collection: Memory management সহজ করে, তবে latency-sensitive system-এ GC behavior বিবেচনা করতে হয়।
- Learning curve: Syntax সহজ হলেও concurrency, race conditions, interfaces এবং production architecture শেখার জন্য অনুশীলন প্রয়োজন।

তাই Go-কে সব ভাষার replacement হিসেবে নয়, বরং তোমার engineering toolkit-এর আরেকটি শক্তিশালী ভাষা হিসেবে দেখবে।

## ৭. তোমার Go শেখার Roadmap

## Go Backend Engineer Learning Path

প্রস্তাবিত পরিকল্পনা: প্রতিদিন ১–২ ঘণ্টা নিয়মিত অনুশীলন। সময় তোমার বর্তমান অভিজ্ঞতা ও অনুশীলনের ওপর নির্ভর করবে।

1. Phase 1 — Fundamentals

   Variables, constants, data types, conditions, loops, functions, arrays, slices, maps ও pointers।

2. Phase 2 — Core Go

   Structs, methods, interfaces, packages, modules, errors, JSON এবং file handling।

3. Phase 3 — Concurrency

   Goroutines, channels, `select`, mutex, `WaitGroup`, context ও race detection।

4. Phase 4 — Backend Development

   HTTP server, REST API, routing, middleware, validation, authentication এবং structured logging।

5. Phase 5 — Database & Testing

   PostgreSQL, SQL, transactions, migrations, integration tests এবং benchmarks।

6. Phase 6 — Production Project

   একটি Ticket Booking API বা Inventory Management API তৈরি করে Docker, configuration, monitoring ও deployment শেখা।

## ৮. আজকের প্রথম অনুশীলন

Go ইনস্টল করা থাকলে প্রথমে এই program চালিয়ে দেখো:

```
package main

import "fmt"

func main() {
    name := "Zamirul"
    language := "Go"

    fmt.Println("Hello,", name)
    fmt.Println("I am learning", language)
}
```

চালানোর command:

```
go run main.go
```

Expected output:

```
Hello, Zamirul
I am learning Go
```

এখানে `package main` executable program-এর package নির্দেশ করে, `func main()` হলো program-এর entry point এবং `:=` দিয়ে নতুন local variable ঘোষণা করা হয়েছে।
