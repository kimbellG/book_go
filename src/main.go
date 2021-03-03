package main

import (
	"fmt"

	"../package/popcount"
)

func main() {
	fmt.Printf("Loop: %d\nDefault: %d\nSwap: %d\nReset: %d\n", popcount.PopCountLoop(228), popcount.PopCount(228),
		popcount.PopCountSwap(228), popcount.PopCountReset(228))
}
