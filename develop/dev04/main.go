package main

import (
	"dev04/anagram"
	"fmt"
)

func main() {
	words := []string{"пятак", "Пятка", "тяпка", "лиСток", "слитОк", "столик", "гаир", "юруис", "рисую"}

	set := anagram.Set(words)
	for k, value := range set {
		fmt.Println(k, value)
	}
}
