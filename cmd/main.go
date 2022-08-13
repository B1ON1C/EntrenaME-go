package main

import (
	"log"
	"time"
)

func main() {
	log.Println("Starting EntrenaME Server")
	timeElapsedStart := time.Now()

	timeElapsedEnd := time.Since(timeElapsedStart)
	log.Println("EntrenaME Served started in " + timeElapsedEnd.String())
}
