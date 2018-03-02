package main

import "fmt"

func teststringrecur() {
	/*
	   "abc" -> "abc", "acb", "bac", "bca" ...
	*/

	//known := map[string]string{
	//	"ab": "ab",
	//	"ba": "ba",
	//}
	//result := permutestringrecur(known, "c")
	//fmt.Printf("%v", result)
	result := permutestring("abc")
	fmt.Printf("%v", result)
}

func permutestring(s string) map[string]string {
	if len(s) <= 1 {
		return map[string]string{s: s}
	}
	front := s[0 : len(s)-1]
	last := s[len(s)-1:]

	allbutlast := permutestring(front)
	result := make(map[string]string)
	for k := range allbutlast {
		for i := 0; i < len(k)+1; i++ {
			s := k[i:] + last + k[:i]
			result[s] = s
		}
	}
	return result
}
