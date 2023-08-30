package main

import (
	"dev09/wget"
	"log"
	"os"
)

func main() {
	err := wget.Exec(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
