package main

import (
	"fmt"
	"math/rand"
	"testing"
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
	// fmt.Println(undefinedVariable)

	// Print "Hello, World!" with a random continent
	fmt.Printf("Hello, World from %s!\n", getRandomContinent())
}

// Common error: Missing return type in function declaration
func add(x, y int) int {
	return x + y
}

// getRandomContinent returns a random continent name
func getRandomContinent() string {
	continents := []string{"Africa", "Antarctica", "Asia", "Europe", "North America", "Oceania", "South America"}
	rand.Seed(time.Now().UnixNano())
	return continents[rand.Intn(len(continents))]
}

// Tests

func TestAdd(t *testing.T) {
	result := add(3, 5)
	expected := 8
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGetRandomContinent(t *testing.T) {
	continents := []string{"Africa", "Antarctica", "Asia", "Europe", "North America", "Oceania", "South America"}
	result := getRandomContinent()

	found := false
	for _, continent := range continents {
		if continent == result {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("Unexpected continent: %s", result)
	}
}
