package main

import "fmt"

func main() {
	// Variable declaration String
	var nameOne string = "Rias"
	var nameTwo = "Gremory"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Akeno"
	nameTwo = "Lucy"
	nameThree = "Erza"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Koneko"

	fmt.Println(nameFour)

	// Variable declaration Integer
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	fmt.Println(numOne, numTwo, numThree)

	// Float
	var scoreOne float32 = 25.98
	var scoreTwo  float64= 98000000054545311.75
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	
}
