package main

import (
	bootstrap "api-http-codely/cmd/api/boostrap"
	"log"
)

func main() {
	err := bootstrap.Run()
	if err != nil {
		log.Fatal(err)
	}
}
