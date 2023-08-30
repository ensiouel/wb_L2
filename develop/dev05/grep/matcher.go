package grep

import (
	"regexp"
	"strings"
)

type Matcher interface {
	Match(line string) bool
}

type RegexpMatcher struct {
	re *regexp.Regexp
}

func NewRegexpMatcher(pattern string, ignoreCase bool) *RegexpMatcher {
	if ignoreCase {
		pattern = "(?i)" + pattern
	}

	return &RegexpMatcher{
		re: regexp.MustCompile(pattern),
	}
}

func (matcher *RegexpMatcher) Match(line string) bool {
	return matcher.re.MatchString(line)
}

type StringMatcher struct {
	s         string
	equalFunc func(s, t string) bool
}

func NewStringMatcher(pattern string, ignoreCase bool) *StringMatcher {
	equalFunc := strings.Contains

	if ignoreCase {
		equalFunc = strings.EqualFold
	}

	return &StringMatcher{
		s:         pattern,
		equalFunc: equalFunc,
	}
}

func (matcher *StringMatcher) Match(line string) bool {
	return matcher.equalFunc(line, matcher.s)
}
