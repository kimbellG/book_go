package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)

	for input.Scan() {
		word := input.Text()
		seen[word]++
	}

	fmt.Println("WORDS of input file: ")
	for word, count := range seen {
		fmt.Printf("%s\t=\t%d\n", word, count)
	}
}
