package main

import (
	"dev02/unpack"
	"fmt"
)

func main() {
	str := "a4bc2d5e"

	fmt.Println(unpack.Unpack(str))
}
