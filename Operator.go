// Unary operator
package main

import "fmt"

var x int = 40

func main() {
	x--
	fmt.Println(x)
}
// Arithmetic operator
package main

import "fmt"

func main() {
	a := 10
	b := 20
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}
// Logical operator
package main

import "fmt"

func main() {
	a := true
	b := false
	fmt.Println(a || b)
	fmt.Println(a && b)
	fmt.Println(!a)
}
