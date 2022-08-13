package main

import "fmt"

func main() {

	x := 5
	y := &x

	fmt.Println(y)
	fmt.Println(*y)
	fmt.Println(&y)

}
