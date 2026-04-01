package utils

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// makes first letter uppercase, makes the rest lowercase
func CapitalizeName(name string) string {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		return ""
	}

	runes := []rune(name)
	first := []rune(strings.ToUpper(string(runes[0])))[0]

	for i := 1; i < len(runes); i++ {
		runes[i] = []rune(strings.ToLower(string(runes[i])))[0]
	}

	runes[0] = first
	return string(runes)
}

// makes sure the name of the character is written with letters
func IsAlpha(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		// autorise lettres, espace, '-' et '\'' (apostrophe)
		if unicode.IsLetter(r) || r == ' ' || r == '-' || r == '\'' {
			continue
		}
		return false
	}
	return true
}

// pauses execution for given seconds
func DelaySeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}

// Prints grid
func PrintGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Println(row)
	}
}

// verifies if the grid is full
func GridFull(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				return false
			}
		}
	}
	fmt.Println("Grid full! Game Over")
	return true
}
