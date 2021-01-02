package anagram

import (
	"strings"
)

// AreAnagram checks if 2 strings are anagram of each other
func AreAnagram(s, t string) bool {

	for _, sChar := range s {

		if !strings.Contains(t, string(sChar)) {
			return false
		}
	}

	return true
}
