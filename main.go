package main

import (
	"log"

	birdactyl "github.com/pizzlad/birdactyl-go-sdk"
)

func main() {
	plugin := birdactyl.New("hello-world", "1.0.0")

	plugin.OnStart(func() {
		plugin.API().Log("info", "Hello Birdactyl!")
	})

	if err := plugin.Start("localhost:50050", 50105); err != nil {
		log.Fatal(err)
	}
}
