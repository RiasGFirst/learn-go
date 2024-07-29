package main

import "fmt"

func updateName(x string, y string) string {
	x = y
	return x
}

func updateMenu(x map[string]float64) {
	x["coffee"] = 2.99
}

func main() {

	name := "Akeno"

	fmt.Println("Before updateName:", name)
	name = updateName(name, "Rias")
	fmt.Println("After updateName:", name)

	menu := map[string]float64{
		"pie":       5.99,
		"ice cream": 3.99,
	}

	fmt.Println("Before updateMenu:", menu)
	updateMenu(menu)
	fmt.Println("After updateMenu:", menu)
}
