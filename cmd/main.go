package main

import (
	"geolocator-backend/internal/router"
	"log"
)

func main() {
	r := router.New()
	if err := r.Run(); err != nil {
		log.Fatalln(err)
	}
}
