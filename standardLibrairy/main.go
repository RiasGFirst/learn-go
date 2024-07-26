package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// Strings package
	greeting := "Hello Anime World!"
	fmt.Println(strings.Contains(greeting, "Anime"))
	fmt.Println(strings.ReplaceAll(greeting, "Anime", "Golang"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "Anime"))
	fmt.Println(strings.Split(greeting, " "))

	// The original string is unchanged
	fmt.Println("Original string: ", greeting)

	// Sort package
	ages := []int{45, 20, 35, 30, 75, 60, 50, 25}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 30)
	fmt.Println(index)

	names := []string{"Rias", "Akeno", "Erza", "Lucy", "Koneko"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "Lucy"))
}