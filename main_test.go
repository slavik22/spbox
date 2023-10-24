package main

import "testing"

func TestSBlock(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{0x53, 0xED},
		{0x0F, 0x76},
		{0x00, 0x63},
		{0xFF, 0x16},
	}

	for _, test := range tests {
		output := sBlock(test.input)
		if output != test.expected {
			t.Errorf("S-Box failed for input 0x%x. Expected: 0x%x, Got: 0x%x", test.input, test.expected, output)
		}
	}
}

func TestReverseSBlock(t *testing.T) {
	tests := []struct {
		input    byte
		expected byte
	}{
		{0xED, 0x53},
		{0x23, 0x32},
		{0x63, 0x00},
		{0xA8, 0x6F},
	}

	for _, test := range tests {
		output := reverseSBlock(test.input)
		if output != test.expected {
			t.Errorf("Reverse S-Box failed for input 0x%x. Expected: 0x%x, Got: 0x%x", test.input, test.expected, output)
		}
	}
}

func TestPBlock(t *testing.T) {
	input := [4][4]byte{
		{0x32, 0x88, 0x31, 0xE0},
		{0x43, 0x5A, 0x11, 0x3A},
		{0xB0, 0x0F, 0xE6, 0x04},
		{0x23, 0x7C, 0x9A, 0x17},
	}

	expected := [4][4]byte{
		{0x32, 0x88, 0x31, 0xE0},
		{0x43, 0x5A, 0x11, 0x3A},
		{0xB0, 0x0F, 0xE6, 0x04},
		{0x23, 0x7C, 0x9A, 0x17},
	}

	output := pBlock(input)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if output[i][j] != expected[i][j] {
				t.Errorf("P-Block failed. Expected: 0x%x, Got: 0x%x", expected[i][j], output[i][j])
			}
		}
	}
}

func TestReversePBlock(t *testing.T) {
	input := [4][4]byte{
		{0x32, 0x88, 0x31, 0xE0},
		{0x43, 0x5A, 0x11, 0x3A},
		{0xB0, 0x0F, 0xE6, 0x04},
		{0x23, 0x7C, 0x9A, 0x17},
	}

	expected := [4][4]byte{
		{0x32, 0x88, 0x31, 0xE0},
		{0x43, 0x5A, 0x11, 0x3A},
		{0xB0, 0x0F, 0xE6, 0x04},
		{0x23, 0x7C, 0x9A, 0x17},
	}

	output := reversePBlock(input)

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if output[i][j] != expected[i][j] {
				t.Errorf("Reverse P-Block failed. Expected: 0x%x, Got: 0x%x", expected[i][j], output[i][j])
			}
		}
	}
}

func TestMain(m *testing.M) {
	m.Run()
}
