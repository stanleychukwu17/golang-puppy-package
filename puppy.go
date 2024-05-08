package puppy

import (
	dogs "github.com/stanleychukwu17/golang-dogs-package"
)

// Bark returns the sound a puppy makes when barking.
// No parameters.
// Returns a string.
func Bark() string {
	return "ruff ruff"
}

func Barks() string {
	return "ruff ruff ruff"
}

func BigBark() string {
	return dogs.WhenGrownUp("woo woo woo")
}
