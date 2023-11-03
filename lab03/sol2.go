package main

import (
	"fmt"
)

// Problem 1: modifies the s2 string in place, and at the end it is empty and could not be used anymore.
// Problem 2: the complexity is O(n1*n2) because the second string is scanned n1 times. It could be O(n1+n2) if we use a map.
// Problem 3: complexity of the concatenation of strings
// Problem 4: is case sensitive and non-letters are not ignored
func main() {
	var counter int
	var s1, s2 string
	fmt.Scanf("%s", &s1)
	fmt.Scanf("%s", &s2)

	if len(s1) != len(s2) { // check only words with the same length
		fmt.Println("Le due parole non sono anagrammi")
		return
	}

	for i := 0; i < len(s1); i++ { // n1 iterations
		c := s1[i]
		for k := 0; k < len(s2); k++ { // n2 iterations
			if s2[k] == c {
				s2 = s2[:k] + s2[k+1:] // removes the letter from s2
				counter++              // counts the number of letters found
				break
			}
		}
	}

	if counter == len(s1) { // if the letters found are the same as the length of the first word, then the two words are anagrams
		fmt.Println("Le due parole sono anagrammi")
	} else {
		fmt.Println("Le due parole non sono anagrammi")
	}
}
