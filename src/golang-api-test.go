package main

import (
	"golang-api-test/app"
	"log"
)

func main() {
	router := app.SetRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}