package hw03frequencyanalysis

import (
	"regexp"
	"sort"
)

var regex = regexp.MustCompile(`[А-Яа-яA-Za-z\-,.]+`)

func SortByPopularity(words []string) []string {
	WordsPopularity := make(map[string]int)

	for _, word := range words {
		WordsPopularity[word]++
	}
	names := make([]string, 0, len(WordsPopularity))
	for name := range WordsPopularity {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool {
		if WordsPopularity[names[i]] == WordsPopularity[names[j]] {
			if names[j] == "-" {
				return true
			} else if names[i] == "-" {
				return false
			}
			return names[i] < names[j]
		}
		return WordsPopularity[names[i]] > WordsPopularity[names[j]]
	})
	var SortedWordsPopularity []string

	SortedWordsPopularity = append(SortedWordsPopularity, names...)

	return SortedWordsPopularity
}

func Top10(str string) []string {
	if str == "" {
		return nil
	}
	var Words []string
	Words = append(Words, regex.FindAllString(str, -1)...)

	return SortByPopularity(Words)[:10]
}
