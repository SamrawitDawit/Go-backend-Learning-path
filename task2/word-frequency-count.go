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
func main() {
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
}
