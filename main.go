// main.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Your main code here...
	fmt.Printf("Hello, World from %s!\n", getRandomContinent())
}

// add is a function that adds two integers.
func add(x, y int) int {
	return x + y
}

// getRandomContinent returns a random continent name
func getRandomContinent() string {
	continents := []string{"Africa", "Antarctica", "Asia", "Europe", "North America", "Oceania", "South America"}
	rand.Seed(time.Now().UnixNano())
	return continents[rand.Intn(len(continents))]
}
