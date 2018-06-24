package utils

import (
    "fmt"
	"math/big"
)

func GetBigInt(s string) *big.Int {
	number,ok := new(big.Int).SetString(s,10)
	if(!ok){
		panic("Oopsie")
	}
	return number
}

func HexToInt(h string)  *big.Int {
	p, ok := new(big.Int).SetString(h, 16)
	if !ok {
		panic("invalid hex: " + h)
	}
	return p
}

func IntToHex(i *big.Int) string {
	p := fmt.Sprintf("%x", i)
	return p
}
