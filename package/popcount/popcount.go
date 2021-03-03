package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//PopCount возвращает количество единичных битов
//в 64-х битном числе
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(n uint64) int {
	counter := 0

	for i := 0; i < 8; i++ {
		counter += int(pc[byte(n>>(i*8))])
	}

	return counter
}

func PopCountSwap(n uint64) int {
	counter := 0

	for i := 0; i < 64; i++ {

		if (n>>i)&1 == 1 {
			counter++
		}
	}

	return counter
}

func PopCountReset(n uint64) int {
	counter := 0

	for ; n != 0; n &= (n - 1) {
		counter++
	}

	return counter
}
