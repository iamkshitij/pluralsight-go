package main

import "fmt"

func main() {
	// d1 := complex(3, 4)
	// d2 := complex(7, 6)

	// var d3 = d1 + d2

	// fmt.Println(d3)

	var firstName *string = new(string)
	*firstName = "Arthur"
	fmt.Println(firstName, *firstName)
}
