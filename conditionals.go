// Conditional Statements
package main

import "fmt"

func main() {
	a := 18
	b := 30
	c := 26

	// if statement
	if a > b {
		fmt.Println("a is greater than b")
	}

	// if-else statement
	if b > a {
		fmt.Println("b is an adult")
	} else {
		fmt.Println("b is not an adult")
	}

	// nested if-else if statement
	if b > a {
		if c > b {
			fmt.Println("C is the largest")
		} else {
			fmt.Println("B is the largest")
		}

	}

}
