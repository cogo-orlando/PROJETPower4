package main

import "fmt"

func main() {
	grid := make([][]int, 6)
	for i := range grid {
		grid[i] = make([]int, 7)
	}
	for i := range grid {
		fmt.Println(grid[i])
	}
}
