package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

// dec binary   population
//
//	 0 00000000 0
//	 1 00000001 1
//	 2 00000010 1
//	 3 00000011 2
//	 4 00000100 1
//	 5 00000101 2
//	 6 00000110 2
//	 7 00000111 3
//	 8 00001000 1
//	50 00110010 3
//
// 100 01100100 3
// 200 11001000 3
// 255 11111111 8
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("i = %3d:%08b, i/2 = %3v, pc[i/2] = %v, i&1 = %v, pop = %v\n", i, i, i/2, pc[i/2], i&1, pc[i])
	}
}

// PopCount returns the population count (number of set bits) of x.
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

func PopCountLoop(x uint64) int {
	var pop byte = 0
	for i := 0; i < 8; i++ {
		pop += pc[byte(x>>(i*8))]
	}
	return int(pop)
}

func PopCountShift64(x uint64) int {
	var pop int = 0
	for i := 0; i < 64; i++ {
		// shift right then bitwise-and the rightmost bit with 1 to count set bits
		pop += int(x >> i & 1)
	}
	return int(pop)
}

func PopCountShiftNonZero(x uint64) int {
	if x == 0 {
		return 0
	}

	var pop int = 0
	last := x

	for {
		// shift right then bitwise-and the rightmost bit with 1 to count set bits
		next := last & (last - 1)
		if next == last {
			return pop
		}
		last = next
		pop += 1
	}
}
