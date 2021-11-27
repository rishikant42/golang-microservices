package app

import (
	"fmt"
	"net/http"

	"github.com/rishikant42/golang-microservices/mvc/controllers"
)

func StartApp() {
	fmt.Println("Starting app")
	http.HandleFunc("/users", controllers.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
