package main

import "fmt"
import "math"

/*

When this program is called it should find the maximum amount of money that
could be made from cutting a rod of size x

./main 7
result = 18

./main 5
result = 13

*/

var maxValue int = 0
var inNumber int = 10

// n being the length of the rod
// a being the array holding the values
func rodC(n int, a []int) int {

	if n <= 1 {
		return a[0] + a[n-1]
	}

	currValue := a[n-1] + a[inNumber-n-1]
	newValue := rodC(n-1, a)
	potential := math.Max(float64(currValue), float64(newValue))

	if potential > float64(maxValue) {
		maxValue = int(potential)
	}

	return maxValue
}

func main() {
	// the value the pipes as x length
	values := []int{1, 5, 8, 9, 10, 17, 17, 20, 24, 30}
	maxValue = values[inNumber-1]

	fmt.Print(rodC(inNumber-1, values))
}
