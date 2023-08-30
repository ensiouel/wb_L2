package main

import (
	"flag"
	"fmt"
	"github.com/beevik/ntp"
	"log"
)

var addr = flag.String("addr", "0.beevik-ntp.pool.ntp.org", "ntp server address")

func main() {
	flag.Parse()

	time, err := ntp.Time(*addr)
	if err != nil {
		log.Fatalf("failed to get time: %s", err)
	}

	fmt.Println("time:", time)
}
