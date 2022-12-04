package main

/*
getBitByNum: Функция из 32-битной руны забирает биты согласно номеру: 0-31
*/
func getBitByNum(value int, num int) int {
	value >>= num
	return value % 2
}

func setBitByNum(value int, value2 int, num int) int {
	value2 <<= num
	return value | value2
}

func clearBitByNum(value int, num int) int {
	bit := 1
	bit <<= num
	value &= 0 ^ bit
	return value
}

func encrypt(data []int, text []int) []int {
	n := 0
	for i := range data {
		if len(text) <= i/32 {
			break
		}
		data[i] = clearBitByNum(data[i], n) | getBitByNum(text[i/32], n)
		n++
		if n == 32 {
			n = 0
			if i+32 > len(data) {
				break
			}
		}
	}
	return data
}

func decrypt(data []int) []int {
	text := make([]int, len(data)/32)
	n := 0
	for i := range data {
		if i >= len(text)*32 {
			break
		}
		text[i/32] = setBitByNum(text[i/32], data[i]%2, n)
		n++
		if n == 32 {
			n = 0
			if i+32 > len(data) {
				break
			}
		}
	}
	for i, v := range text {
		if v == 0 {
			return text[0:i]
		}
	}
	return text
}
