package unpack

import (
	"strconv"
	"strings"
)

func Unpack(s string) string {
	split := strings.Split(s, "")

	var res []string

	for _, r := range split {
		n, err := strconv.Atoi(r)
		if err != nil {
			res = append(res, r)
		} else {
			if len(res) == 0 {
				break
			}

			i := len(res) - 1

			prevRune := res[i]
			if prevRune == `\` {
				if res[i-1] != `\` {
					res[i] = r
					continue
				}

				i -= 1
				n -= 1
			}

			str := strings.Repeat(prevRune, n)

			res[i] = str
		}
	}

	return strings.Join(res, "")
}
