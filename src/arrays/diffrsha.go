package main

import (
	"crypto/sha256"
	"fmt"

	"../package/popcount"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	var countOfDiference int
	for i, _ := range c1 {
		countOfDiference += popcount.ForByte(c1[i] ^ c2[i])
	}

	fmt.Printf("%x\n%x\n%T\n", c1, c2, c1)
	fmt.Printf("number of diference bit: %d\n", countOfDiference)
}
