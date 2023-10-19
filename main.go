package main

import (
	app "app-command-line/pkg"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Running aplication")

	aplication := app.Generated()
	if error := aplication.Run(os.Args); error != nil {
		log.Fatal(error)
	}

}