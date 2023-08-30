package mansort

import (
	"strings"
	"unicode"
)

func trimNumber(str string) string {
	return strings.TrimFunc(str, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
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

func set(slice []string) []string {
	m := make(map[string]struct{})

	var result []string

	for _, s := range slice {
		if _, ok := m[s]; !ok {
			m[s] = struct{}{}
			result = append(result, s)
		}
	}

	return result
}
