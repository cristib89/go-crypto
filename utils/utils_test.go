package utils

import (
	"testing"
	"math/big"
)

func TestEverything(t *testing.T) {
	t.Log("Loading test cases")
	cases := []struct {
		bigInt *big.Int;
		hexStr string
	}{
		{GetBigInt("0"), "0"},
		{GetBigInt("2"), "2"},
		{GetBigInt("15"), "f"},
		{GetBigInt("16"), "10"},
		{GetBigInt("981729305719365491283123479147263784617236487234168072647812893419827349871239489710234789263871328979237489273498723293"), 
		"6153b801e736d8410fa50eab54893a5b1edd9402a82894a2acb3c74dbc17a75c52c103b29379375322e02b5723d7284763dd"},
	}
	for _, c := range cases {
		t.Log("Running test case")
		runHexToIntTest(c.bigInt,c.hexStr,t)
		runIntToHex(c.bigInt,c.hexStr,t)
	}
	t.Log("Done")
}



func runHexToIntTest(b *big.Int, s string, t *testing.T) {
	if(s != IntToHex(b)) {
		t.Error("IntToHex failed")
	}

}

func runIntToHex(b *big.Int, s string, t *testing.T) {
	if(b.Cmp(HexToInt(s))!=0) {
		t.Error("HexToInt failed")
	}
}
