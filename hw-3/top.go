package hw_3

import (
	"regexp"
	"sort"
	"strings"
)

func Top(text string) []string  {
	var result []string

	if len(text) == 0 {
		return result
	}

	top := make(map[string]int)

	space := regexp.MustCompile(`\s+`)
	words := strings.Split(space.ReplaceAllString(text, " "), " ")

	for _, word := range words {

		if word == "" {
			continue
		}

		if _, exist := top[word]; exist == false {
			top[word] = 0
		}

		top[word]++
	}

	for word, _ := range top {
		result = append(result, word)
	}

	sort.Slice(result, func(i int, j int) bool {
		return top[result[i]] > top[result[j]]
	})

	return result[:10]
}
