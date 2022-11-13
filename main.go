package main

import (
	"net/http"

	"github.com/iamkshitij/pluralsight-go/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}
