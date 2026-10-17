package main

import "fmt"

// Package Scope
const companyName = "Tech Solutions"

var employeeCount = 10

func calculateSalary(basicSalary float64) float64 {
	// Function Scope
	bonus := 5000.0

	if basicSalary >= 50000 {
		// Block Scope
		tax := basicSalary * 0.10
		fmt.Println("Tax:", tax)
	}

	// tax এখানে ব্যবহার করা যাবে না।
	// কারণ tax শুধু if block-এর মধ্যে উপলব্ধ।

	return basicSalary + bonus
}

func main() {
	// main Function Scope
	employeeName := "Zamirul"
	salary := 60000.0

	fmt.Println("Company:", companyName)
	fmt.Println("Employees:", employeeCount)
	fmt.Println("Employee:", employeeName)

	totalSalary := calculateSalary(salary)
	fmt.Println("Total Salary:", totalSalary)

	// bonus এখানে ব্যবহার করা যাবে না।
	// কারণ bonus calculateSalary() ফাংশনের Local Variable।
}