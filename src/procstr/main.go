package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
	"unicode/utf8"
)

func isNotStart(index int) bool {
	return index > 0
}

func IsPlaceForComma(index int, CommaPeriod int) bool {
	if isNotStart(index) && index%CommaPeriod == 0 {
		return true
	} else {
		return false
	}
}

func aligmentStringOfNumbersByNumber(StringOfNumbers string, AlignmentBoundary int) (bytes.Buffer, int) {
	var StartStringOfNumbersWithComma bytes.Buffer
	var StartAlignmentString int

	for StartAlignmentString = 0; StartAlignmentString < len(StringOfNumbers)%AlignmentBoundary; StartAlignmentString++ {
		StartStringOfNumbersWithComma.WriteByte(StringOfNumbers[StartAlignmentString])
	}

	if isNotStart(StartAlignmentString) {
		StartStringOfNumbersWithComma.WriteByte(',')
	}

	return StartStringOfNumbersWithComma, StartAlignmentString
}

func addCommaInStringOfIntegerNumbers(StringOfNumbers string, CommaPeriod int) string {
	StringOfNumbersWithComma, StartAlignmentString := aligmentStringOfNumbersByNumber(StringOfNumbers, CommaPeriod)

	for i := 0; i+StartAlignmentString < len(StringOfNumbers); i++ {
		if IsPlaceForComma(i, CommaPeriod) {
			StringOfNumbersWithComma.WriteByte(',')
		}
		StringOfNumbersWithComma.WriteByte(StringOfNumbers[i+StartAlignmentString])
	}

	return StringOfNumbersWithComma.String()
}

func addCommaInStringOfNumbers(StringOfNumbers string, CommaPeriod int) string {
	var StringOfNumbersWithComma bytes.Buffer

	StringOfNumbersWithoutSign := AddSignInStringOfNumbers(StringOfNumbers, &StringOfNumbersWithComma)
	IntegerPart, FractionalPart := GetIntegerFractionalParts(StringOfNumbersWithoutSign)

	IntegerPartWithComma := addCommaInStringOfIntegerNumbers(IntegerPart, CommaPeriod)

	StringOfNumbersWithComma.WriteString(IntegerPartWithComma)
	StringOfNumbersWithComma.WriteByte('.')
	StringOfNumbersWithComma.WriteString(FractionalPart)

	return StringOfNumbersWithComma.String()
}

func GetOptionalSignFromStringOfNumbers(StringOfNumbers string) (string, byte) {
	var signOfNumber rune
	if signOfNumber, _ = utf8.DecodeRuneInString(StringOfNumbers); unicode.IsDigit(signOfNumber) {
		return StringOfNumbers, ' '
	} else if signOfNumber == '+' || signOfNumber == '-' {
		return StringOfNumbers[1:], byte(signOfNumber)
	}

	fmt.Fprintf(os.Stderr, "Sign: %v\n", signOfNumber)
	os.Exit(1)
	return "", ' '
}

func AddSignInStringOfNumbers(StringOfNumbers string, SignDest *bytes.Buffer) string {
	StringOfNumbersWithoutSign, sign := GetOptionalSignFromStringOfNumbers(StringOfNumbers)

	if sign != ' ' {
		SignDest.WriteByte(sign)
	}

	return StringOfNumbersWithoutSign
}

func GetIntegerFractionalParts(StringOfNumbers string) (string, string) {
	const PointOfFloat = '.'

	for i := 0; i < len(StringOfNumbers); i++ {
		if StringOfNumbers[i] == PointOfFloat {
			return StringOfNumbers[:i], StringOfNumbers[i+1:]
		}
	}

	return StringOfNumbers, ""
}

func isAnagram(test string) bool {
	if len(test) < 2 {
		return true
	}

	if test[0] != test[len(test)-1] {
		return false
	} else {
		return isAnagram(test[1 : len(test)-1])
	}
}

func main() {
	fmt.Printf("%s\n", addCommaInStringOfNumbers("-1234567890.5678", 3))
	fmt.Printf("%t\n", isAnagram("abcxba"))
}
