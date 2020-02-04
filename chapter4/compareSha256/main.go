package main

import "crypto/sha256"

import "fmt"

func compareUint8(u1 uint8, u2 uint8) int {
	var result int
	for i := 0; i < 8; i++ {
		b1 := u1 & 1
		b2 := u2 & 1
		u1 >>= 1
		u2 >>= 1
		if b1 == b2 {
			result++
		}
	}
	return result
}

func compareSha256(a1 [32]uint8, a2 [32]uint8) int {
	var result int
	for i := range a1 {
		result += compareUint8(a1[i], a2[i])
	}

	return result
}

func main() {
	c1 := sha256.Sum256([]byte{'x'})
	c2 := sha256.Sum256([]byte{'y'})
	fmt.Printf("c1 = %x\nc2 = %x\ncompare(c1,c2) = %d\n", c1, c2, compareSha256(c1, c2))
}
