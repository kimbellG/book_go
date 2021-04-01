package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

type countsOfTypesUnicode struct {
	leters, puncts, numbers map[rune]int
	invalid                 int
}

func NewCountsOfTypesUnicode() *countsOfTypesUnicode {
	var newStruct countsOfTypesUnicode
	newStruct.leters = make(map[rune]int)
	newStruct.puncts = make(map[rune]int)
	newStruct.numbers = make(map[rune]int)

	return &newStruct
}

func main() {
	result := NewCountsOfTypesUnicode()

	in := bufio.NewReader(os.Stdin)
	for {
		var r rune
		var err error

		if r, err = ReadRune(in); err == io.EOF {
			break
		}

		if err := deftypeAndIncrementCount(result, r); err != nil {
			continue
		}
	}
	outputResult(result)
}

func ReadRune(in *bufio.Reader) (rune, error) {
	r, _, err := in.ReadRune()
	if err == io.EOF {
		return r, io.EOF
	}
	if err != nil {
		log.Fatalln("charcount:", err)
	}

	return r, nil
}

func deftypeAndIncrementCount(result *countsOfTypesUnicode, r rune) error {
	if r == unicode.ReplacementChar {
		result.invalid++
		return errors.New("deftype: Invalid rune!")
	}

	if unicode.IsLetter(r) {
		result.leters[r]++
	}

	if unicode.IsPunct(r) {
		result.puncts[r]++
	}

	if unicode.IsNumber(r) {
		result.numbers[r]++
	}

	return nil
}

func outputResult(res *countsOfTypesUnicode) {
	outputLetter(res.leters)
	outputNumber(res.numbers)
	outputPunts(res.puncts)
	outputInvalidCount(res.invalid)
}

func outputNumber(numbers map[rune]int) {
	fmt.Println("NUMBERS.")
	for c, n := range numbers {
		fmt.Printf("%q\t =\t%d\n", c, n)
	}
}

func outputLetter(letter map[rune]int) {
	fmt.Println("LETTERS.")
	for c, n := range letter {
		fmt.Printf("%q\t =\t%d\n", c, n)
	}
}

func outputPunts(puncts map[rune]int) {
	fmt.Println("PUNCTS.")
	for c, n := range puncts {
		fmt.Printf("%q\t =\t%d\n", c, n)
	}
}

func outputInvalidCount(count int) {
	fmt.Println("INVALID:", count)
}
