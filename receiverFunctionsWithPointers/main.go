package main

import "fmt"

func main() {

	mybill := newBill("rias's bill")

	mybill.updateTip(10)
	mybill.addItem("Onion soup", 4.50)
	mybill.addItem("Meat pie", 8.95)
	mybill.addItem("Lemonade", 2.50)

	fmt.Println(mybill.format())
}