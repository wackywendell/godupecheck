package dupechecker

import (
	"math/rand"
)

func CreateStrings(n int, sz int) []string {
	const letters = "abcdefghijklmnopqrstuvwxyz"

	strings := make([]string, 0, n)
	for i := 0; i < n; i++ {
		tag := make([]byte, sz)
		for i := range tag {
			if i == sz/2 {
				tag[i] = ':'
				continue
			}
			tag[i] = letters[rand.Intn(len(letters))]
		}
		strings = append(strings, string(tag))
	}

	return strings
}

func DedupeArraywise(strs []string) []string {
	newList := make([]string, 0, len(strs))

	for _, s := range strs {
		var exists = false
		for _, t := range newList {
			if s == t {
				exists = true
				break
			}
		}

		if !exists {
			newList = append(newList, s)
		}

	}

	return newList
}

func DedupeMapwise(strs []string) []string {
	var exists map[string]bool = make(map[string]bool, len(strs))
	newList := make([]string, 0, len(strs))

	for _, s := range strs {
		found := exists[s]
		if !found {
			newList = append(newList, s)
			exists[s] = true
		}
	}
	return newList
}
