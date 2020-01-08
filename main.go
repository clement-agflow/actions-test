package main

import (
	"log"
	"time"

	"clement.actions/version"
)

const v = "1.0.0"

func main() {
	currentVersion, err := version.New(v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Version:", currentVersion)

	t := time.NewTicker(2 * time.Second)
	for range t.C {
		log.Println("tick")
	}
}
