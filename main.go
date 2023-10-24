package main

import (
	"testing"
)

var sBox = [16]byte{
	0x0, 0x4, 0x8, 0xC,
	0x1, 0x5, 0x9, 0xD,
	0x2, 0x6, 0xA, 0xE,
	0x3, 0x7, 0xB, 0xF,
}

var pBox = [16]byte{
	0, 4, 8, 12,
	1, 5, 9, 13,
	2, 6, 10, 14,
	3, 7, 11, 15,
}

func substitute(input byte) byte {
	return sBox[input]
}

func permute(input byte) byte {
	return pBox[input]
}

func TestSBox(t *testing.T) {
	for i := 0; i < 16; i++ {
		expected := sBox[i]
		actual := substitute(byte(i))
		if expected != actual {
			t.Errorf("S-box substitution for input %x: expected %x, got %x", i, expected, actual)
		}
	}
}

func TestPBox(t *testing.T) {
	for i := 0; i < 16; i++ {
		expected := pBox[i]
		actual := permute(byte(i))
		if expected != actual {
			t.Errorf("P-box permutation for input %x: expected %x, got %x", i, expected, actual)
		}
	}
}

func main() {
	t := &testing.T{}
	TestSBox(t)
	TestPBox(t)
}
