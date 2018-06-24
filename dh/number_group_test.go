package dh

import(
	"math/big"
	"github.com/cristib89/go-crypto/utils"
	"testing"
)

func TestNumberGroup(t *testing.T) {

	generator, ok := new(big.Int).SetString("5",10)
	prime, ok := new(big.Int).SetString("23",10)

	if(!ok) {
		panic("Oopsie")
	}

	group := NewNumberGroup(utils.GetBigInt("5"),utils.GetBigInt("23"))

	if(group.g.Cmp(generator) != 0) {
		t.Error("Variables don't match!")
	}

	if(group.p.Cmp(prime) != 0) {
		t.Error("Variables don't match!")
	}

}