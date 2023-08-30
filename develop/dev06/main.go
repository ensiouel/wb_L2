package main

import (
	"dev06/cut"
	"log"
	"os"
)

func main() {
	err := cut.Exec(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
