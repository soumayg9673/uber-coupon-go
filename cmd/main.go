package main

import (
	"log"
	"net/http"

	"github.com/soumayg9673/uber-coupon-go/packages/server"
)

func main() {
	log.Println("initiating goanubhavbharat backend server")
	if err := server.Run(); err != nil && err != http.ErrServerClosed {
		log.Println(err)
	}
}
