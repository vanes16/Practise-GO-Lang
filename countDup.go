package countDup

import "strings"

func duplicate_count(s string) int {
	//your code goes here

	count := make(map[rune]int)
	s = strings.ToLower(s)
	for _, ch := range s {
		count[ch]++
	}

	result := 0

	for _, x := range count {
		if x > 1 {
			result++
		}
	}

	return result
}
