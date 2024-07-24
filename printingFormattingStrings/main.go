package main

import "fmt"

func main() {
	age := 20
	name := "Rias"

	// Print
	fmt.Print("Hello, ")
	fmt.Print("World! \n")
	fmt.Print("New line \n")

	// Println
	fmt.Println("Hello, World!")
	fmt.Println("goodbye, World!")
	fmt.Println("My age is", age, "and my name is", name)

	// Printf (formatted strings) %_ = format specifier
	fmt.Printf("My age is %v and my name is %v \n", age, name)
	fmt.Printf("My age is %q and my name is %q \n", age, name)
	fmt.Printf("Age is of type %T and name is of type %T \n", age, name)
	fmt.Printf("You scored %0.1f points! \n", 225.55)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)
	fmt.Println("The saved string is:", str)
}
