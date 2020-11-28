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
	words := strings.Split(space.ReplaceAllString(strings.ToLower(text), " "), " ")

	if len(words) == 0 {
		return result
	}

	for _, word := range words {
		word = strings.Trim(word, "!.;\":-,\n\t")

		if word == "" {
			continue
		}

		top[word]++
	}

	if len(top) == 0 {
		return result
	}

	for word, _ := range top {
		result = append(result, word)
	}

	sort.Slice(result, func(i int, j int) bool {
		return top[result[i]] > top[result[j]]
	})

	limit := 10

	if len(result) < limit {
		limit = len(result)
	}

	return result[:limit]
}
