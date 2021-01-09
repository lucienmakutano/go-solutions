package revstr

import "strings"

// Reverse takes a string and reverse the words in it
func Reverse(s string) (res string) {
	sl := strings.Split(s, " ")

	for _, w := range sl {
		res = w + " " + res
	}

	res = strings.Trim(res, " ")
	return
}
