package main

import "fmt"

/*
This is my attempt at solving a shortest path algorithm.
It's inspired by a google interview question that I bombed and this
is practice for it. I'm attempting to solve it in a dynamic programming
manner


using the randommaize.py program we can easily generate a 2d array for this
exercise

The problem as for us to print out the lowest number we encounter on our way
from the top left side of the array to the bottom right side of the array

i.e.
if we start at 9 at the top left maze[0][0], then on our way to the bottom
right of the array maze[3][3], then the lowest number should be 0


*/

func main() {

	maze := [4][4]int{
		{9, 5, 7, 7},
		{5, 3, 1, 4},
		{8, 2, 6, 8},
		{4, 9, 2, 9},
	}

	fmt.Print(maze)

}
