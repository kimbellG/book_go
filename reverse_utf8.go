package main

import (
	"fmt"
	"os"
)

const MaxUnicodeRune = 4
const IncorrectStartRune = 128
const TestBitsOnIncorrectRune = 192
const MaxStartBitsForStartOfRuneWithMaxSize = 240

func reverseUTF8(sourceString []byte) []byte {
	for i, j = 0, len(sourceString)-1; i < j; i, j = i+1, j-1 {
		sizeIrune := getSizeOfFirstByteUTF8Symbol(sourceString[i])
		sizeJrune := getSizeOfUTF8SymbolFromTheEnd(sourceString[:j])

	}

}

func getSizeOfUTF8SymbolFromTheEnd(sourceString []byte) int {
	for i := len(sourceString) - 1; i >= 0; i-- {
		if sizeRune := getSizeOfFirstByteUTF8Symbol(sourceString[i]); sizeRune != 0 {
			return sizeRune
		}
	}
}

func getSizeOfFirstByteUTF8Symbol(char byte) uint8 {
	for i, firstBitsOfRune := uint8(0), uint8(MaxStartBitsForStartOfRuneWithMaxSize); i < MaxUnicodeRune+1; i, firstBitsOfRune = i+1, firstBitsOfRune<<1 {
		if isByteOfRuneWithThisSize(char, firstBitsOfRune) {
			return getCorrectSize(MaxUnicodeRune - i)
		}
	}
	assets(false)
}

func isFirstRune(firstByteOfRune byte) bool {
	return firstByteOfRune&TestBitsOnIncorrectRune != IncorrectStartRune
}

func isByteOfRuneWithThisSize(char byte, firstBitsOfRune byte) bool {
	fmt.Printf("%b %b\n", uint8(char&firstBitsOfRune), firstBitsOfRune)
	return char&firstBitsOfRune == firstBitsOfRune
}

func getCorrectSize(countOfLeftShift uint8) uint8 {
	switch countOfLeftShift {
	case 0:
		return 1
	case 1:
		return 0
	default:
		return countOfLeftShift
	}
}

func leftShiftWithOneFilingIn(number uint8) uint8 {
	return (number >> 1) | (1 << 7)
}

func assets(condition bool) {
	if !condition {
		os.Exit(1)
	}
}

func main() {
	//str := "\u4e16"
	str := "a"
	sym := []byte(str)
	fmt.Println(getSizeOfFirstByteUTF8Symbol(sym[0]))
}
