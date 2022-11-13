package main

import (
	"fmt"

	"github.com/iamkshitij/pluralsight-go/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Alex",
		LastName:  "Hales",
	}
	fmt.Println(u)
}
