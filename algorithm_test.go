package main

import (
	"testing"
)

// just an example of fuzzing, there's nothing interesting here

func IntsByBytes(data []byte) []int {
	data2 := make([]int, len(data)/4)
	for i := range data2 {
		data2[i] = int(data[i*4]) | (int(data[i*4+1]) << 8) |
			(int(data[i*4+2]) << 16) | (int(data[i*4+3]) << 24)
	}
	return data2
}

func FuzzDecrypt(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {
		if got := decrypt(IntsByBytes(data)); got == nil {
			t.Errorf("decrypt() = %v", got)
		}
	})
}

func FuzzEncrypt(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte, text []byte) {
		if got := encrypt(IntsByBytes(data), IntsByBytes(text)); got == nil {
			t.Errorf("encrypt() = %v", got)
		}
	})
}
