package actions_test

import (
	"log"

	"clement.actions/version"
)

const v = "1.0.0"

func main() {
	currentVersion, err := version.New(v)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(currentVersion)
}
