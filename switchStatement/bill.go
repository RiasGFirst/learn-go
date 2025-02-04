package main

import (
	"fmt"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Make new bill
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format the bill
func (b *bill) format() string {
	fs := "Bill breakdown: \n"
	var total float64 = 0

	// List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v\n", k+":", v)
		total += v
	}

	// Add tip
	fs += fmt.Sprintf("%-25v ...$%0.2f\n", "Tip:", b.tip)

	// Total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "Total:", total+b.tip)

	return fs
}

// Update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// Add an item to the bill
func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}