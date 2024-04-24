package popcount

import (
	"math"
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

func BenchmarkPopcountShiftNonZero64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(math.MaxUint64)
	}
}

func BenchmarkPopcountShiftNonZero32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(math.MaxUint32)
	}
}

func BenchmarkPopcountShiftNonZero16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(math.MaxUint16)
	}
}

func BenchmarkPopcountShiftNonZero8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(math.MaxUint8)
	}
}

func BenchmarkPopcountShiftNonZero4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(15)
	}
}

func BenchmarkPopcountShiftNonZero2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(3)
	}
}

func BenchmarkPopcountShiftNonZero1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(1)
	}
}

func BenchmarkPopcountShiftNonZeroZero(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShiftNonZero(0)
	}
}
