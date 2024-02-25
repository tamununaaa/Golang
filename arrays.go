// Arrays
package main

import "fmt"

func main() {
	//When length of array is defined
	var fruits = [4]string{"strawberry", "persimmon", "papaya", "peach"}
	marks := [3]int{98, 79, 83}

	//When length of array is not defined
	var vegetables = [...]string{"tomato", "brinjal", "peas"}
	days := [...]string{"Mon", "Tues", "Wed", "Thurs", "Fri"}

	fmt.Println(fruits)
	fmt.Println(marks)
	fmt.Println(vegetables)
	fmt.Println(days)

}
