package main

import (
	"dev10/telnet"
	"log"
	"os"
)

func main() {
	err := telnet.Run(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
