package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// Problem 1:
//
//	There is the intention to ignore the case of the letters, but the rune function doesn't save the value.
//
// Problem 2:
//
//	To print the histogram, the program iterates over the 26 letters of the alphabet, and for each letter it iterates over the map, which is quite inefficient.
//
// Problem 3:
//
//	It is actually useless to split the line and then iterate over each word, since the map is populated letter by letter.
func main() {
	scanner := bufio.NewScanner(os.Stdin) // reads from standard input
	scanner.Scan()                        // reads a line
	s := scanner.Text()                   // converts it to a string
	s1 := strings.Split(s, " ")           // splits it by spaces, creating a slice of strings with length = n

	diz := make(map[byte]int) // creates a map: why not using rune as key type?

	for i := 0; i < len(s1); i++ { // n iterations
		parola := s1[i]                    // gets the i-th word
		for k := 0; k < len(parola); k++ { // m iteration, where m is the length of the word
			unicode.ToLower(rune(parola[k]))         // ignore case, but doesn't save the value...
			if parola[k] >= 97 && parola[k] <= 122 { // 2 comparisons, checks if the char is a lower case letter. Using "magic numbers", bad practice
				diz[parola[k]]++
			}
		}
	}

	for i := 'a'; i <= 'z'; i++ { // 26 iterations
		for k, v := range diz {
			if rune(k) == i {
				fmt.Println(string(k), " ", strings.Repeat("*", v))
			}
		}
	}
}
