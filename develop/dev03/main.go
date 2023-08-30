package main

import (
	"bufio"
	"dev03/mansort"
	"flag"
	"fmt"
	"log"
	"os"
)

type Options struct {
	Column          int
	IsNumericColumn bool
	IsReverse       bool
	IsUnique        bool
	Output          string
}

func main() {
	options := Options{}

	flagSet := flag.NewFlagSet("", flag.ContinueOnError)
	flagSet.IntVar(&options.Column, "k", 0, "specifying the column to sort")
	flagSet.BoolVar(&options.IsNumericColumn, "n", false, "sort by numeric value")
	flagSet.BoolVar(&options.IsReverse, "r", false, "sort in reverse order")
	flagSet.BoolVar(&options.IsUnique, "u", false, "do not output duplicate lines")
	flagSet.StringVar(&options.Output, "o", "", "output file")

	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("invalid params")
	}

	filename := args[0]
	if filename == "" {
		log.Fatal("filename is required")
	}

	flagSet.Parse(args[1:])

	lines, err := parseFile(filename)
	if err != nil {
		log.Fatalf("failed to parse file: %s", err)
	}

	lines = mansort.Sort(lines, options.Column, options.IsNumericColumn, options.IsReverse, options.IsUnique)

	output := os.Stdout
	if options.Output != "" {
		output, err = os.Create(options.Output)
		if err != nil {
			log.Fatalf("failed to create output file: %s", err)
		}
	}

	for _, line := range lines {
		fmt.Fprintln(output, line)
	}
}

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}
