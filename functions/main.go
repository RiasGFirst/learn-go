package main

import (
	"fmt"
	"math"
)

func sayGreetings(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayGoodbye(n string) {
	fmt.Printf("Goodbye %v\n", n)
}

func cycleNames(names []string, f func(string)) {
	for _, name := range names {
		f(name)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {
	names := []string{"Rias", "Akeno", "Koneko", "Asia", "Xenovia"}
	cycleNames(names, sayGreetings)
	cycleNames(names, sayGoodbye)

	a1 := circleArea(10.5)
	a2 := circleArea(15)
	fmt.Printf("Area 1: %0.3f\n", a1)
	fmt.Printf("Area 2: %0.3f\n", a2)
}