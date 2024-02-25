// Loops
package main

import "fmt"

func main() {

	//for loop
	for i := 1; i < 11; i++ {
		fmt.Printf("%d x 8 = %d \n", i, 8*i)
	}

	//continue statement
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Print(i)
	}
	fmt.Println()

	//break statement
	for i := 0; i < 5; i++ {
		if i == 2 {
			break
		}
		fmt.Print(i)
	}
	fmt.Println()

	//range keyword
	fruits := [4]string{"kiwi", "strawberry", "grapes", "persimmon"}
	for idx, val := range fruits {
		fmt.Printf("%v  %v\n", idx, val)
	}
}
