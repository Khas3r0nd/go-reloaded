package main

import (
	"go-reloaded/utils"
	"testing"
)

func TestBinaryToDecimal(t *testing.T) {
	binToDec := func(binary string, expectedDecimal string) {
		gotDecimal := utils.BinaryToDecimal(binary)
		if gotDecimal != expectedDecimal {
			t.Errorf("binaryToDecimal(%s) = %s; want %s", binary, expectedDecimal, gotDecimal)
		}
	}

	binToDec("0", "0")
	binToDec("10", "2")
	binToDec("1010", "10")
}

func TestHexToDecimal(t *testing.T) {
	hexToDec := func(hex string, expectedDecimal string) {
		gotDecimal := utils.HexToDecimal(hex)
		if gotDecimal != expectedDecimal {
			t.Errorf("hexToDecimal(%s) = %s; want %s", hex, expectedDecimal, gotDecimal)
		}
	}

	hexToDec("0", "0")
	hexToDec("3e8", "1000")
}
