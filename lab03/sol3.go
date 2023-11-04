package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

func quali(riga string) map[string]int { // creates a map of letters and their occurrences
	lettere := make(map[string]int)
	for _, char := range riga {
		if unicode.IsLetter(char) {
			lettere[string(char)]++
		}
	}
	return lettere
}

func anagrammi(s1, s2 string) bool {
	return reflect.DeepEqual(quali(s1), quali(s2))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin) // reads from stdin
	scanner.Split(bufio.ScanLines)        // sets the split function for the scanning operation (default is ScanLines)
	scanner.Scan()
	riga := strings.ToLower(scanner.Text()) // scan a line in lower case
	scanner.Scan()
	riga2 := strings.ToLower(scanner.Text()) // scan another line in lower case
	lettere := quali(riga)
	fmt.Println(anagrammi(riga, riga2)) // Basically instantiates two maps and then compares if they are equals; Problem: computes again `riga`
	lettere2 := []string{}
	for key, _ := range lettere { // creates a slice of the keys of the map, the letters
		lettere2 = append(lettere2, key)
	}
	sort.Strings(lettere2) // then sorts the slice
	for _, lett := range lettere2 {
		fmt.Print(lett, " ")
		fmt.Println(strings.Repeat("*", lettere[lett])) // and then, for each letter, that is already the key of the map, prints the number of occurrences...
	}
}
