package main

import "fmt"

func greet() {
	fmt.Println("Hello")

}

func main() {
	greet()
}

package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	fmt.Println(add(5, 3))
}

package main

import "fmt"

func sqrt(x int) int {
	return x * x
}
func main() {
	fmt.Println(sqrt(5))
}
