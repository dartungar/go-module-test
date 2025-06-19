// My test package.
package main

import "golang.org/x/exp/constraints"

func main() {

}

type Number interface {
	constraints.Float | constraints.Integer
}

// Adds two integers together and returns the result.
func Add[T Number](a T, b T) T {
	return a + b
}
