package main

import (
	"fmt"
	"github.com/golang-collections/collections/stack"
	"testing"
)

func TestParenthesis(t *testing.T) {
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

func testbracketvalidator() {
	if !bracketvalidator("{[]()}") {
		panic("should be valid")
	}
	if bracketvalidator("{[(])}") {
		panic("should NOT be valid")
	}
	if bracketvalidator("{[}") {
		panic("should ALSO NOT be valid")
	}
	if bracketvalidator("{") {
		panic("should ALSO NOT be valid")
	}
}

func bracketvalidator(s string) bool {
	expected := stack.New()

	for _, c := range s {
		switch c {
		case '(', '{', '[':
			fmt.Printf("found %#U\n", c)
			expected.Push(c)
		case ')':
			if !checkcloser(expected, '(') {
				return false
			}

		case ']':
			if !checkcloser(expected, '[') {
				return false
			}

		case '}':
			if !checkcloser(expected, '{') {
				return false
			}
		}
	}
	return expected.Len() == 0
}

func checkcloser(expected *stack.Stack, wantedChar rune) bool {
	if expected.Len() == 0 {
		return false
	}
	opener := expected.Pop()
	return opener == wantedChar
}
