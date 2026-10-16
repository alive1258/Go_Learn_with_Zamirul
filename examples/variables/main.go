package variables

import "fmt"

func main() {
	// Integer
	age := 20

	// String
	name := "Zamirul"

	// Float
	height := 5.8

	// Boolean
	isDeveloper := true

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Developer:", isDeveloper)

	// Type checking
	fmt.Printf("Type of name: %T\n", name)
	fmt.Printf("Type of age: %T\n", age)
	fmt.Printf("Type of height: %T\n", height)
	fmt.Printf("Type of isDeveloper: %T\n", isDeveloper)
}