package main

import (
	"log"
	"os"

	"github.com/matheusgb/go-ip-finder/app"
)

func main() {
	app := app.Generate()

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
