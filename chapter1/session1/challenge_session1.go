package main

import "fmt"

func main() {
	var i int = 21

	fmt.Println(i)
	fmt.Printf("%T\n", i)
	fmt.Printf("%%\n")

	var j bool = true

	fmt.Printf("%t\n\n", j)
	fmt.Printf("%b\n", i)
	fmt.Println("\u042f")
	fmt.Printf("%d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%x\n", 15)
	fmt.Printf("%X\n", 15)
	fmt.Printf("%U\n\n", 'Я')

	var k float64 = 123.456

	fmt.Printf("%.6f\n", k)
	fmt.Printf("%e\n", k)
}
