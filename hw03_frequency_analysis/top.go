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

var regexpWord = regexp.MustCompile(`\p{L}+`)

func Top10(input string) []string {
	return getTopByCount(10, input)
}

func Top5(input string) []string {
	return getTopByCount(5, input)
}

func Top10Eng(input string) []string {
	return getTopByCount(10, input)
}

func clearPunct(word string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsPunct(r) {
			return -1
		}
		return r
	}, word)
}

func getTopByCount(count int, text string) []string {
	stringSlice := strings.Fields(text)
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

	if len(pairs) < count {
		count = len(pairs)
	}

	// Create result
	result := make([]string, count)
	for i := range count {
		result[i] = pairs[i].word
	}
	return result
}
