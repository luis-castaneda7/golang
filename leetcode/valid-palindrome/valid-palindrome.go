// https://leetcode.com/problems/valid-palindrome

package main

import (
	"fmt"
	"regexp"
	"strings"
)

func fixString(sentence string) string {
	reg, _ := regexp.Compile("[^a-zA-Z0-9]+")

	processedString := reg.ReplaceAllString(sentence, "")

	return strings.ToLower(processedString)
}

func isPalindrome(s string) bool {

	sentence := fixString(s)
	stringLength := len(sentence)

	for i, j := 0, stringLength-1; i < len(sentence); i, j = i+1, j-1 {
		if i == j {
			return true
		}
		if sentence[i] != sentence[j] {
			return false
		}
	}

	return true

}

func main() {
	sentence := "."

	fmt.Print(isPalindrome(sentence))
}
