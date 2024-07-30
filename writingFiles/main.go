package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

			// Convert price to float
			p, err := strconv.ParseFloat(price, 64)
			if err != nil {
				fmt.Println("The price must be a number")
				promptOption(b)
			}
			b.addItem(name, p)

			fmt.Println("Item added -", name, price, "to the bill", b.name)
			promptOption(b)
		case "t":
			tip, _ := getInput("Enter tip amount ($): ", reader)

			// Convert price to float
			t, err := strconv.ParseFloat(tip, 64)
			if err != nil {
				fmt.Println("The tip must be a number")
				promptOption(b)
			}
			b.updateTip(t)

			fmt.Println("Tip added -", tip, "to the bill", b.name)
			promptOption(b)
		case "s":
			b.save()
			fmt.Println("You saved the file", b.name)
		default:
			fmt.Println("That is not a valid option...")
			promptOption(b)
	}
}

func main() {
	mybill := createBill()
	promptOption(mybill)

}