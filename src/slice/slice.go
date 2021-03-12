package main

import (
	"fmt"
	"unicode"
)

func testSwapAdjacentSpaceChar() {
	testString := "abc\t\tcde\n"
	fmt.Printf("Source string: %q\n", testString)

	testString = string(swapBlocksOfSpaceCharToOneSpace([]byte(testString)))
	fmt.Printf("Result string: %q\n", testString)
}

func testReverseArray() {
	testArray := [32]int{1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Source Array: %v\n", testArray)

	testArray = *reverseArray(&testArray)
	fmt.Printf("Reverse Array: %v\n", testArray)
}

func testDeleteAdjacentDup() {
	testSlice := []int{1, 2, 2, 2, 2, 3, 3, 3, 3, 4, 5, 6, 7, 8}
	fmt.Printf("Source Array: %v\n", testSlice)

	testSlice = deleteAdjacentDup(testSlice)
	fmt.Printf("Array without adjacent duplication: %v\n", testSlice)
}

func reverseArray(sourceArray *[32]int) *[32]int {
	for i, j := 0, len(sourceArray)-1; i < j; i, j = i+1, j-1 {
		sourceArray[i], sourceArray[j] = sourceArray[j], sourceArray[i]
	}

	return sourceArray
}

func removeElementFromSlice(src []int, index int) []int {
	copy(src[index:], src[index+1:])

	return src[:len(src)-1]
}

func deleteAdjacentDup(sourceIntSlice []int) []int {
	for i := 0; i < len(sourceIntSlice)-1; i++ {
		for sourceIntSlice[i] == sourceIntSlice[i+1] {
			sourceIntSlice = removeElementFromSlice(sourceIntSlice, i+1)
		}
	}

	return sourceIntSlice
}

func swapBlocksOfSpaceCharToOneSpace(sourceString []byte) []byte {
	for i := 0; i < len(sourceString); i++ {
		if unicode.IsSpace(rune(sourceString[i])) {
			sourceString = rmAdjcentSpaceCharFromHere(sourceString, i)
			sourceString[i] = ' '
		}
	}

	return sourceString
}

func rmAdjcentSpaceCharFromHere(sourceString []byte, startBlockOfChar int) []byte {

	countRemovedSpaceChar := offsetAfterAdjcentSpaceCharFromHere(sourceString[startBlockOfChar:])
	return decreaseStringAfterOffset(sourceString, countRemovedSpaceChar)
}

func offsetAfterAdjcentSpaceCharFromHere(startSpaceChar []byte) int {
	for i, char := range startSpaceChar {
		if !unicode.IsSpace(rune(char)) {
			startSpaceChar = rmBlockOfChar(startSpaceChar, i-1)
			return i
		}
	}
	return 0
}

func decreaseStringAfterOffset(sourceString []byte, numberOfOffsetChar int) []byte {
	return sourceString[:len(sourceString)-numberOfOffsetChar]
}

func rmBlockOfChar(startStringBlock []byte, endOfBlock int) []byte {
	copy(startStringBlock, startStringBlock[endOfBlock+1:])
	return startStringBlock[:len(startStringBlock)-endOfBlock]
}

func main() {
	testSwapAdjacentSpaceChar()
}
