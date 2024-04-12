package main

import (
	"cmd-ip/app"
	"fmt"
	"log"
	"os"
)

func main() {
	cmd := app.Gen()

	if error := cmd.Run(os.Args); error != nil {
		fmt.Println("Error on init")
		log.Fatal(error)
	}
}