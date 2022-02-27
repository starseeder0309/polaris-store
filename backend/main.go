package main

import (
	"log"

	"github.com/kimhyeonu/polaris-store/backend/routes"
)

func main() {
	log.Println("Polaris Store Backend Server")
	routes.RunApi("127.0.0.1:9595")
}
