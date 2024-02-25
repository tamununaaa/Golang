// Declaring Variables and constants
package main

import "fmt"

//Declaring Variables using var keyword without initial value
var num int
var str string

//Declaring Variables using var keyword with initial value
var x = 3
var s = "Ivy"

//Declaring constant
const PI = 3.14    //untyped const
const days int = 7 //typed const

func main() {

	//Using :=
	age := 10
	name := "Lyla"
	fmt.Println(num, str)
	fmt.Println(x, s)
	fmt.Println(age, name)

	//Decalaring multiple variables
	var a, b, c int = 100, 200, 300
	var d, e = 400, "four hundred"
	f, g := "five hundred", 5
	fmt.Println(a, b, c)
	fmt.Println(d, e)
	fmt.Println(f, g)

	//Declaring in a block
	var (
		empID  int    = 198
		empDep string = "Support"
	)
	fmt.Println(empID, empDep)

	//Declaring local constant
	const G = 9.8
	fmt.Println("Value of pi: ", PI)
	fmt.Println("Value of gravitational constant: ", G)
	fmt.Println("Days in a week: ", days)

	//Formatting output
	fmt.Printf("The value is %v and the type is %T ", G, G)
}
