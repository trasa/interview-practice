package main

import "fmt"

func parenthesistest() {
	//    0---------*---------*---------*---------*---------*---------*---------*---------*---------*---------*
	s := "Sometimes (when I nest them (my parentheticals) too much (like this (and this))) they get confusing."
	result := findParenMatch(s, 10)
	fmt.Printf("expected 79, found %d\n", result)

	result = findParenMatch(s, 20)
	fmt.Printf("expected 46, found %d\n", result)
}

func findParenMatch(s string, pos int) int {
	level := 0
	for testpos, c := range s[pos:] {
		if c == '(' {
			level++
		} else if c == ')' {
			level--
			if level == 0 {
				return testpos + pos
			}
		}
	}
	fmt.Println("unbalanced!")
	return -1
}