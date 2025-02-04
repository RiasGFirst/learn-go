package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	
	return strings.TrimSpace(input), err
}

// Create a new bill
func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOption(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
		case "a":
			name, _ := getInput("Item name: ", reader)
			price, _ := getInput("Item price: ", reader)
			fmt.Println(name, price)
		case "t":
			tip, _ := getInput("Enter tip amount ($): ", reader)
			fmt.Println(tip)
		case "s":
			fmt.Println("Saving the bill")
		default:
			fmt.Println("That is not a valid option...")
			promptOption(b)
	}
}

func main() {
	mybill := createBill()
	promptOption(mybill)

}