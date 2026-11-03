package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
	"unicode"
)

type pair struct {
	word  string
	count int
}

var regexpWord = regexp.MustCompile("[А-Яа-я]")

func Top10(input string) []string {
	stringSlice := strings.Fields(input)
	wordMap := make(map[string]int)

	for _, word := range stringSlice {
		if regexpWord.MatchString(word) {
			clearString := clearPunct(word)
			wordMap[strings.ToLower(clearString)]++
		}
	}

	pairs := make([]pair, 0, len(wordMap))
	for word, count := range wordMap {
		pairs = append(pairs, pair{word, count})
	}

	// Sort
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count != pairs[j].count {
			return pairs[i].count > pairs[j].count
		}
		return pairs[i].word < pairs[j].word
	})

	top10 := 10
	if len(pairs) < top10 {
		top10 = len(pairs)
	}

	// Create result
	result := make([]string, top10)
	for i := range top10 {
		result[i] = pairs[i].word
	}
	return result
}

func clearPunct(word string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, word)
}
