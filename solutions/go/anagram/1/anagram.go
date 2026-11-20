package anagram

import (
	"maps"
	"strings"
	"unicode"
)

func Detect(subject string, candidates []string) []string {
	var result []string
	lowerSubject := strings.ToLower(subject)
	subjectFreq := charFrequency(lowerSubject)

	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)

		if lowerCandidate == lowerSubject {
			continue
		}

		if maps.Equal(subjectFreq, charFrequency(lowerCandidate)) {
			result = append(result, candidate)
		}
	}

	return result
}

func charFrequency(s string) map[rune]int {
	freq := make(map[rune]int)
	for _, r := range s {
		freq[unicode.ToLower(r)]++
	}
	return freq
}