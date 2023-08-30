package anagram

import (
	"sort"
	"strings"
)

func Set(words []string) map[string][]string {
	m := make(map[string][]string)
	exists := make(map[string]struct{})

	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
	}

	for i, word := range words {
		sortedWord := sortWord(word)

		for j := i; j < len(words); j++ {
			if _, ok := exists[words[j]]; ok {
				continue
			}

			if sortedWord == sortWord(words[j]) {
				exists[words[j]] = struct{}{}
				m[word] = append(m[word], words[j])
			}
		}
	}

	for k, v := range m {
		if len(v) < 2 {
			delete(m, k)
			continue
		}

		sort.Strings(v)
	}

	return m
}

func sortWord(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}
