package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Common error: Unused variable
	unusedVariable := "This variable is not used"

	// Common error: Unused import
	// import "fmt"

	// Common error: Printf format string without arguments
	// fmt.Printf("Hello, World!")

	// Common error: Missing return statement
	// result := add(3, 5)

	// Common error: Variable shadowing
	// num := 10
	// if num > 5 {
	//     num := 5
	//     fmt.Println(num)
	// }

	// Common error: Infinite loop
	// for {
	//     fmt.Println("Hello, World!")
	// }

	// Uncomment the following line to introduce a runtime error
	 fmt.Println(undefinedVariable)

	// Print "Hello, World!" with a random continent
	fmt.Printf("Hello, World from %s!\n", getRandomContinent())
}

// Common error: Missing return type in function declaration
 func add(x, y) {
     return x + y
 }

// getRandomContinent returns a random continent name
func getRandomContinent() string {
	continents := []string{"Africa", "Antarctica", "Asia", "Europe", "North America", "Oceania", "South America"}
	rand.Seed(time.Now().UnixNano())
	return continents[rand.Intn(len(continents))]
}
