// Arrays
package main

import "fmt"

func main() {
	cities := [3]string{"Vancouver", "Boston", "Dallas"}
	a := 10
	b := 20
	display(cities[:])
	details("Ivvy", cities[:], 2)
	details("Lyla", cities[:], 1)
	details("Myra", cities[:], 2)
	fmt.Println(sum(a, b))
}

func display(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func details(name string, arr []string, i int) {
	fmt.Println("Name:", name, " City:", arr[i])
}

func sum(x int, y int) int {
	return x + y
}
