package main

import (
	"fmt"
	"math/rand"
)

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


in the hard implementation we will just find the shortest path by adding the
numbers along the way towards the lower number


*/

//look left, look down
var minimum int = 0
var nummap = make(map[*int]int)

func shortestpath(maze [][]int, x int, y int, total int, n int) int {
	if x == n-1 && y == n-1 {
		return maze[x][y]
	}

	if x > n-1 || y > n-1 {
		return 99
	}

	leftTotal := shortestpath(maze, x, y+1, total, n) + maze[x][y]
	downTotal := shortestpath(maze, x+1, y, total, n) + maze[x][y]

	if leftTotal < downTotal {
		return leftTotal
	}
	return downTotal

}
func genMaze(size int) [][]int {
	m := make([][]int, size)

	for i := 0; i < size; i++ {
		m[i] = make([]int, size)
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m[i][j] = rand.Intn(10)
		}
	}

	return m
}

func main() {

	n := 17
	maze := genMaze(n)

	minimum = maze[0][0]
	total := 0

	fmt.Println(shortestpath(maze, 0, 0, total, n))
}
