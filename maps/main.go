package main

import "fmt"

func main() {

	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
		"donut": 1.99,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	// loop through map
	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	// int as key
	phonebook := map[int]string{
		666: "Rias",
		667: "Akeno",
		668: "Koneko",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[667])

	// loop through map
	for k, v := range phonebook {
		fmt.Println(k, "-", v)
	}
	
}