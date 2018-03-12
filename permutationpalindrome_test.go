package main

import (
	"fmt"
	"testing"
)

func TestPermutePalindrome(t *testing.T) {
	if !anypermutationpalindrome("civic") {
		panic("expected palindrome")
	}
	if !anypermutationpalindrome("ivicc") {
		panic("expected palindrome also")
	}
	if anypermutationpalindrome("civil") {
		panic("no palindrome")
	}
	if anypermutationpalindrome("livci") {
		panic("also no")
	}
	if !anypermutationpalindrome("bob") {
		panic("bob should work")
	}
	if !anypermutationpalindrome("bb") {
		panic("bb")
	}
	if !anypermutationpalindrome("bbb") {
		panic("bbb")
	}
	if !anypermutationpalindrome("bbbb") {
		panic("bbbb")
	}
}

func anypermutationpalindrome(s string) bool {
	letters := make(map[rune]bool)

	for _, c := range s {
		letters[c] = !letters[c]
	}
	// at the end, there should only be at most one 'true' in the map
	var foundOne bool
	for _, v := range letters {
		if v {
			if foundOne {
				fmt.Printf("%s not palindrome\n", s)
				return false // found too many
			} else {
				foundOne = true
			}
		}
	}
	fmt.Printf("%s is palindrome\n", s)
	return true
}
