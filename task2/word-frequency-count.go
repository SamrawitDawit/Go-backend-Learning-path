package main

import (
	"fmt"
	"strings"
	"unicode"
)

func frequencyCounter(String string) map[string]int {
	words := strings.FieldsFunc(String, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}
	return wordCount
}
func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	characters := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
	new_s := strings.Join(characters, "")
	n := len(new_s)
	for i := 0; i < n/2; i++ {
		if new_s[i] != new_s[n-i-1] {
			return false
		}
	}
	return true
}
func main() {

	// testing frequencyCounter

	test := frequencyCounter("Hello, this is a test. Hello, this is a test.")
	fmt.Println(test)
	//test 1
	if test["Hello"] != 2 {
		fmt.Println("Test 1 failed")
		return
	}
	//test 2
	if test["Hello,"] != 0 {
		fmt.Println("Test 2 failed")
		return
	}
	//test 3
	if test["This"] != 0 {
		fmt.Println("Test 3 failed")
		return
	}
	fmt.Println("All tests passed")

	// testing isPalindrome

	//test 1
	test1 := isPalindrome("Hello.ll/eh")
	if !test1 {
		fmt.Println("Test 1 failed")
		return
	}
	//test 2
	test2 := isPalindrome("123-321")
	if !test2 {
		fmt.Println("Test 2 failed")
		return
	}
	//test 3
	test3 := isPalindrome("Hello, this is a test.")
	if test3 {
		fmt.Println("Test 3 failed")
		return
	}
	fmt.Println("All tests passed")
}
