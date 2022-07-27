// https://leetcode.com/problems/calculate-digit-sum-of-a-string/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addNumbers(s []string) string {
	total := 0

	for _, str := range s {
		number, _ := strconv.Atoi(str)
		total += number
	}

	return strconv.Itoa(total)
}

func separate(s string, k int) string {

	if len(s) <= k {
		return s
	}

	sarray := strings.Split(s, "")
	var vector []string

	for i := 0; i < len(sarray); i = i + k {
		slice := sarray[i:]
		if cap(slice) < k {
			vector = append(vector, addNumbers(slice))
			break
		} else {
			vector = append(vector, addNumbers(sarray[i:i+k]))
		}
	}

	return separate(strings.Join(vector[:], ""), k)
}

func main() {
	fmt.Print(separate("11111222223", 3))
}
