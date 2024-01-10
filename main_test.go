// main_test.go
package main

import "testing"

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
