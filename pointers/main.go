package main

import "fmt"

func updateName(x *string) {
	*x = "Rias"
}


func main() {

	name := "Akeno"

	fmt.Println("Before updateName:", name)

	m := &name
	fmt.Println("Memory Address:", m)
	fmt.Println("Value at Memory Address:", *m)

	updateName(m)
	fmt.Println("After updateName:", name)
}