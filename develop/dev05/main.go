package main

import (
	"dev05/grep"
	"flag"
	"log"
	"os"
)

func main() {
	options := grep.Options{}

	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	flagSet.IntVar(&options.After, "A", 0, "print +N lines after a match")
	flagSet.IntVar(&options.Before, "B", 0, "print +N lines to match")
	flagSet.IntVar(&options.Context, "C", 0, "(A+B) print ±N lines around the match")
	flagSet.BoolVar(&options.Count, "c", false, "number of rows")
	flagSet.BoolVar(&options.IgnoreCase, "i", false, "ignore case")
	flagSet.BoolVar(&options.Invert, "v", false, "instead of a match, exclude")
	flagSet.BoolVar(&options.Fixed, "F", false, "exact match with a string, not a pattern")
	flagSet.BoolVar(&options.LineNum, "n", false, "print line number")

	args := os.Args[1:]
	if len(args) < 2 {
		log.Fatal("invalid params")
	}

	pattern := args[0]
	if pattern == "" {
		log.Fatal("pattern is required")
	}

	filename := args[1]
	if filename == "" {
		log.Fatal("filename is required")
	}

	flagSet.Parse(args[2:])

	err := grep.Run(filename, pattern, options)
	if err != nil {
		log.Fatal(err)
	}
}
