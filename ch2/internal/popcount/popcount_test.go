package popcount

import (
	"math"
	"strconv"
	"testing"
)

var testCases = []struct {
	name     string
	value    uint64
	expected int
}{
	{name: "0", value: 0, expected: 0},
	{name: "1", value: 1, expected: 1},
	{name: "2", value: 2, expected: 1},
	{name: "3", value: 3, expected: 2},
	{name: "4", value: 4, expected: 1},
	{name: "8", value: 8, expected: 1},
	{name: "100", value: 100, expected: 3},
	{name: "255", value: 255, expected: 8},
	{name: "MaxUint32", value: math.MaxUint32, expected: 32},
	{name: "MaxUint64", value: math.MaxUint64, expected: 64},
}

func TestPopcount(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PopCount(tc.value)
			if got != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, got)
			}
		})
	}
}

func TestPopcountLoop(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PopCountLoop(tc.value)
			if got != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, got)
			}
		})
	}
}

func TestPopcountShift64(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PopCountShift64(tc.value)
			if got != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, got)
			}
		})
	}
}

func TestPopcountShiftNonZero(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := PopCountShiftNonZero(tc.value)
			if got != tc.expected {
				t.Errorf("expected %v but got %v", tc.expected, got)
			}
		})
	}
}

func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(math.MaxUint64)
	}
}

func BenchmarkPopcountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(math.MaxUint64)
	}
}

func BenchmarkPopcountShift64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift64(math.MaxUint64)
	}
}

func BenchmarkPopcountShiftNonZero(b *testing.B) {
	// because the performance of this method is based on how many bits are set,
	// we run the test for a somewhat representative sample, tho biased toward
	// smaller numbers
	for _, val := range []uint64{
		math.MaxUint64, math.MaxUint32,
		math.MaxUint16, math.MaxUint8, 15, 3, 1, 0} {
		b.Run(strconv.FormatUint(val, 10), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				PopCountShiftNonZero(val)
			}
		})
	}
}
