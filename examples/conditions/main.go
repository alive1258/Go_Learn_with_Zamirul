package conditions

import "fmt"

func main() {
	age := 20

	if age > 18 {
		fmt.Println("Age is greater than 18")
	} else if age < 18 {
		fmt.Println("Age is less than 18")
	} else {
		fmt.Println("Age is exactly 18")
	}
}