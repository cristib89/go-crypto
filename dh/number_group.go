package dh

import (
	"math/big"
)

type NumberGroup struct {
	g *big.Int
	p *big.Int
}

func NewNumberGroup(generator *big.Int, prime *big.Int) *NumberGroup {
	return &NumberGroup {generator, prime}
}
