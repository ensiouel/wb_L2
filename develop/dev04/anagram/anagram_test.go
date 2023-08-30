package anagram_test

import (
	"dev04/anagram"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet(t *testing.T) {
	testCases := []struct {
		name string
		in   []string
		want map[string][]string
	}{
		{
			name: "normal",
			in:   []string{"пятак", "Пятка", "тяпка", "лиСток", "слитОк", "столик", "гаир", "юруис", "рисую"},
			want: map[string][]string{
				"пятак":  {"пятак", "пятка", "тяпка"},
				"листок": {"листок", "слиток", "столик"},
				"юруис":  {"рисую", "юруис"},
			},
		},
		{
			name: "empty",
			in:   []string{},
			want: map[string][]string{},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := anagram.Set(tc.in)
			assert.Equal(t, tc.want, got)
		})
	}
}
