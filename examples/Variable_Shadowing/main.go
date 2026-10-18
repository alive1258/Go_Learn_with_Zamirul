package variableshadowing

import "fmt"

func main() {
    price := 1000

    if true {
        price := 500
        fmt.Println("Inside block:", price)
    }

    fmt.Println("Outside block:", price)
}