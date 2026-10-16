package main

import "fmt"

func calculate(a, b int) (sum int, product int) {
	sum = a + b
	product = a * b

	return
}

func main() {
	total, multiplication := calculate(5, 3)

	fmt.Println("Sum:", total)
	fmt.Println("Product:", multiplication)
}
