package grep

import (
	"bufio"
	"fmt"
	"os"
)

type Options struct {
	After      int
	Before     int
	Context    int
	Count      bool
	IgnoreCase bool
	Invert     bool
	Fixed      bool
	LineNum    bool
}

func Run(filename string, pattern string, options Options) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var matcher Matcher
	if options.Fixed {
		matcher = NewStringMatcher(pattern, options.IgnoreCase)
	} else {
		matcher = NewRegexpMatcher(pattern, options.IgnoreCase)
	}

	options.Before = max(options.Before, options.Context)
	options.After = max(options.After, options.Context)

	var lines []string
	var matches []int

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		if (matcher.Match(line) && !options.Invert) || (!matcher.Match(line) && options.Invert) {
			matches = append(matches, i)
		}

		lines = append(lines, line)
	}

	if options.Count {
		fmt.Println(len(matches))
		return nil
	}

	visited := make(map[int]struct{})
	matched := make(map[int]struct{})

	for _, m := range matches {
		matched[m] = struct{}{}
	}

	var lastEnd int
	for _, m := range matches {
		if options.Before > 0 || options.After > 0 {
			start := max(0, m-options.Before)
			end := min(len(lines)-1, m+options.After)

			if start-lastEnd >= 2 && lastEnd != 0 {
				fmt.Println("--")
			}

			for ; start <= end; start++ {
				lastEnd = end

				if _, ok := visited[start]; ok {
					continue
				}

				visited[start] = struct{}{}

				line := lines[start]
				if options.LineNum {
					if _, ok := matched[start]; ok {
						line = fmt.Sprintf("%d:%s", start+1, line)
					} else {
						line = fmt.Sprintf("%d-%s", start+1, line)
					}
				}

				fmt.Println(line)
			}

			continue
		}

		line := lines[m]
		if options.LineNum {
			line = fmt.Sprintf("%d:%s", m+1, line)
		}

		fmt.Println(line)
	}

	return nil
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
