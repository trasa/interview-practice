package main

import (
	"fmt"
	"strings"
	"unicode"
)

func worldcloudtest() {
	c := buildWordCloud("After beating the eggs, Dana read the next step:")
	fmt.Printf("first:\n")
	for k, v := range c {
		fmt.Printf("%s: %d\n", k, v)
	}

	fmt.Printf("second\n")
	c = buildWordCloud("Add milk and eggs, then add flour and sugar.")
	for k, v := range c {
		fmt.Printf("%s : %d\n", k, v)
	}

}

func splitWords(s string) (result []string) {
	start := 0
	wlen := 0
	for pos, ch := range s {
		if unicode.IsLetter(ch) {
			if wlen == 0 {
				start = pos
			}
			wlen++
		} else {
			// found a word
			w := strings.ToLower(s[start : start+wlen])
			fmt.Println(w)
			result = append(result, w)
			wlen = 0
		}
	}
	return
}

func buildWordCloud(s string) (result map[string]int) {
	result = make(map[string]int)
	for _, w := range splitWords(s) {
		result[w] = result[w] + 1
	}
	return
}
