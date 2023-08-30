package unpack

import "testing"

func TestUnpack(t *testing.T) {
	var cases = map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"°5€1":     "°°°°°€",
		"":         "",
		"45":       "",
		"5oijsd":   "",
		`qwe\4\5`:  "qwe45",
		`qwe\45`:   "qwe44444",
		`qwe\\5`:   `qwe\\\\\`,
	}

	for in, want := range cases {
		if got := Unpack(in); got != want {
			t.Errorf("Unpack(%q) = %q, want %q", in, got, want)
		}
	}
}
