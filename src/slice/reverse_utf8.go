package main

import (
	"fmt"
	"math"
	"os"
)

const MaxUnicodeRune = 4
const IncorrectStartRune = 128
const TestBitsOnIncorrectRune = 192
const MaxStartBitsForStartOfRuneWithMaxSize = 240

func reverseUTF8(sourceString []byte) []byte {
	for i, j := 0, len(sourceString)-1; i < j; i, j = i+1, j-1 {
		sizeIrune := getSizeOfFirstByteUTF8Symbol(sourceString[i])
		sizeJrune := getSizeOfUTF8SymbolFromTheEnd(sourceString[:j+1])

		swapFirstAndEndBytes(sourceString[i:j+1], getLower(sizeIrune, sizeJrune))
		swapAllLastBytesOfRune(sourceString[i:j+1], int(sizeIrune-sizeJrune))
		i, j = i+int(getLower(sizeIrune, sizeJrune)), j-int(getLower(sizeIrune, sizeJrune))
	}
	return sourceString
}

func swapFirstAndEndBytes(sourceString []byte, count uint8) []byte {
	fmt.Println(sourceString)
	for i, j := uint8(0), indexOfLastElement(sourceString); i < count; i, j = i+1, j-1 {
		sourceString[i], sourceString[j] = sourceString[j], sourceString[i]
	}
	reverseByteSlice(sourceString[:count])
	reverseByteSlice(sourceString[len(sourceString)-int(count):])
	fmt.Println(sourceString)
	return sourceString
}

func swapAllLastBytesOfRune(stringSegment []byte, differenceBetweenLeftAndRightSize int) []byte {
	fmt.Println(stringSegment)
	if differenceBetweenLeftAndRightSize < 0 {
		reverseByteSlice(stringSegment)
		swapAllLeftLastBytesOfRune(stringSegment[:len(stringSegment)-differenceBetweenLeftAndRightSize], uint(math.Abs(float64(differenceBetweenLeftAndRightSize))))
		reverseByteSlice(stringSegment)
		fmt.Println(stringSegment)
		return stringSegment
	} else {
		swapAllLeftLastBytesOfRune(stringSegment[differenceBetweenLeftAndRightSize:], uint(differenceBetweenLeftAndRightSize))
		fmt.Println(stringSegment)
		return stringSegment
	}
}

func swapAllLeftLastBytesOfRune(stringSegment []byte, count uint) []byte {
	for i := uint(0); i < count; i++ {
		swapLastByteOfRune(stringSegment)
	}
	return stringSegment
}

func swapLastByteOfRune(stringSegmant []byte) []byte {
	LastByteOfRune := stringSegmant[0]
	leftShiftOfSlice(stringSegmant)
	stringSegmant[indexOfLastElement(stringSegmant)] = LastByteOfRune
	return stringSegmant
}

func leftShiftOfSlice(sourceString []byte) []byte {
	copy(sourceString, sourceString[1:])
	return sourceString
}

func reverseByteSlice(arr []byte) []byte {
	for i, j := 0, indexOfLastElement(arr); i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func getSizeOfUTF8SymbolFromTheEnd(sourceString []byte) uint8 {
	for i := len(sourceString) - 1; i >= 0; i-- {
		if sizeRune := getSizeOfFirstByteUTF8Symbol(sourceString[i]); sizeRune != 0 {
			return sizeRune
		}
	}
	fmt.Printf("getSizeOfOutUTF8SymbolFromTheEnd: incorrect symbol\n")
	os.Exit(1)
	return 0
}

func getSizeOfFirstByteUTF8Symbol(char byte) uint8 {
	for i, firstBitsOfRune := uint8(0), uint8(MaxStartBitsForStartOfRuneWithMaxSize); i < MaxUnicodeRune+1; i, firstBitsOfRune = i+1, firstBitsOfRune<<1 {
		if isByteOfRuneWithThisSize(char, firstBitsOfRune) {
			return getCorrectSize(MaxUnicodeRune - i)
		}
	}
	assets(false)
	return 0
}

func isFirstRune(firstByteOfRune byte) bool {
	return firstByteOfRune&TestBitsOnIncorrectRune != IncorrectStartRune
}

func isByteOfRuneWithThisSize(char byte, firstBitsOfRune byte) bool {
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

func indexOfLastElement(array []byte) int {
	return len(array) - 1
}

func getLower(n1, n2 uint8) uint8 {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}

func main() {
	//str := "\u4e16"
	str := "абc"
	sym := []byte(str)
	fmt.Println(string(reverseUTF8(sym)))
}
