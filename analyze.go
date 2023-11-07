package main

import (
	"strings"

	"golang.org/x/exp/slices"
)

type Word struct {
	text  string
	count int
}

type WordsArr []Word

// cmpWords compares two Word objects based on their count property.
func cmpWords(a Word, b Word) int {
	if a.count < b.count {
		return 1
	} else if a.count > b.count {
		return -1
	}

	return 0
}

// remove_chars takes in a byte slice and removes any characters that are not
// alphanumeric or whitespace, returning the resulting string.
func removeChars(input []byte) string {
	builder := &strings.Builder{}
	builder.Grow(len(input))

	for _, elem := range input {
		if (33 <= elem && elem <= 47) || (58 <= elem && elem <= 64) || (91 <= elem && elem <= 96) || (123 <= elem && elem <= 126) {
			continue
		} else {
			builder.WriteByte(elem)
		}
	}
	return builder.String()
}

// analyze takes in an input byte array, counts the occurrences of each letter in the
// input, and returns the letter count as well as a sorted list of words and their frequencies.
func analyze(input []byte) ([]int, []Word) {
	letterCnt := make([]int, 26)
	wordIds := make(map[string]int)

	str := strings.ToLower(removeChars(input))
	words := strings.Split(strings.ReplaceAll(str, "\n", " "), " ")

	wordsArr := make([]Word, 0, len(words))

	id := 0

	for _, word := range words {
		if strings.TrimSpace(word) == "" {
			continue
		}

		i, ok := wordIds[word]
		if ok {
			wordsArr[i].count++
		} else {
			wordsArr = append(wordsArr, Word{word, 1})
			wordIds[word] = id
			id++
		}

		for _, letter := range word {
			if letter >= 97 && letter <= 122 {
				letterCnt[letter-97]++
			}
		}
	}

	slices.SortFunc(wordsArr, cmpWords)

	if opts.topKWords != 0 && int(opts.topKWords) <= len(wordsArr) {
		wordsArr = wordsArr[:opts.topKWords]
	}

	return letterCnt, wordsArr
}
