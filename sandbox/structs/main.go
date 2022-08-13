package main

import "fmt"

type testStruct struct {
	Val int
	Str string
}

func main() {

	y := testStruct{5, "hello"}
	o := testStruct{Val: 7}

	fmt.Println(y.Str)
	fmt.Println(o)

}
