package main

import "fmt"

func main() {
	var age int = 4
	fmt.Println(age)
	var pointer *int = &age
	*pointer = 50
	fmt.Println(age)
	fmt.Println(pointer)

}
