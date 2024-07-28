package main

import (
	"fmt"
	"strings"
)

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}
	
	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"
}

func main() {
	fn1, ln1 := getInitials("Rias Gremory")
	fmt.Println(fn1, ln1)

	fn2, ln2 := getInitials("Akeno Himejima")
	fmt.Println(fn2, ln2)

	fn3, ln3 := getInitials("Koneko")
	fmt.Println(fn3, ln3)
}