package challenges

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	wordToPattern := make(map[string]string)
	patternVisited := make(map[string]bool)
	words := strings.Split(str, " ")
	patternSplit := strings.Split(pattern, "")
	fmt.Printf("Words are %v\n and Pattern is %v\n", words, pattern)
	if len(patternSplit) != len(words) {
		return false
	}
	for i, ele := range words {
		if _, ok := wordToPattern[ele]; ok {
			if len(patternSplit) <= i || wordToPattern[ele] != patternSplit[i] {
				return false
			}
		} else {
			if len(patternSplit) > i {

				if !patternVisited[patternSplit[i]] {
					patternVisited[patternSplit[i]] = true
					wordToPattern[ele] = patternSplit[i]
					continue
				}
				return false
			}
			return false

		}
	}

	return true
}
