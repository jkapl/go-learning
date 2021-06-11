package diffiehellman

import (
	"math/big"
	"crypto/rand"
)

func PrivateKey(p *big.Int) *big.Int {
	var max *big.Int

	max.Exp(big.NewInt(2), p, nil).Sub(max, big.NewInt(1))
	n, _ := rand.Int(rand.Reader, max)
	return n
}

func PublicKey(private *big.Int, p *big.Int, g int64) *big.Int {
	var pk *big.Int
	return pk.Exp(g, private, p)

}

func NewPair(p *big.Int, g int64) (private *big.Int, public *big.Int) {
	return big.NewInt(2), big.NewInt(2)
}

func SecretKey(private1 *big.Int, public2 *big.Int, p *big.Int) *big.Int {
	return big.NewInt(2)
}

