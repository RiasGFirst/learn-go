package main

import "fmt"

func main() {
	x := 0
	for x < 5 {
		fmt.Println("Value of x is:", x)
		x++
	}
	fmt.Println("=====================================")

	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is:", i)
	}
	fmt.Println("=====================================")

	names := []string{"Rias", "Akeno", "Koneko", "Asia", "Xenovia"}
	for i := 0; i < len(names); i++ {
		fmt.Println("Value of names is:", names[i])
	}
	fmt.Println("=====================================")

	for index, value := range names {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}

	// If you don't want to use index, you can use _ instead of index
	for _, value := range names {
		fmt.Printf("Value: %s\n", value)
	}
}
