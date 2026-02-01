package main

import "fmt"

func main() {
	names := []string{"John", "Jane", "Doe"}
	fmt.Println(names)
	fmt.Println(names[0])
	names[0] = "Alice"
	fmt.Println(names)
}
