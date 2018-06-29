package dh

import (
	"math/big"
	"math/rand"
	"time"
)

var r = rand.New(rand.NewSource(time.Now().Unix()))
var two = big.NewInt(2)

type DiffieHellmanClient struct {
	group *NumberGroup
	privateKey *big.Int
	sharedKey *big.Int
}

func NewDHClient(group *NumberGroup) *DiffieHellmanClient {
	client := DiffieHellmanClient{group, setPrivateKey(group.p), nil}
	return &client
}

func (dhClient *DiffieHellmanClient) GetSharedSecret() *big.Int {
	return dhClient.sharedKey
}

func (dhClient *DiffieHellmanClient) SetSharedSecret(receivedPublicKey *big.Int) {
	dhClient.sharedKey = secretKey(dhClient.privateKey, receivedPublicKey, dhClient.group.p)
}

func (dhClient *DiffieHellmanClient) SendPublicKey() *big.Int {
	return setPublicKey(dhClient.privateKey, dhClient.group)
}

func (dhClient *DiffieHellmanClient) RefreshPrivateKey() {
	dhClient.privateKey = setPrivateKey(dhClient.group.p)
}

func setPrivateKey(p *big.Int) *big.Int {
	private := new(big.Int).Sub(p, two)
	return private.Add(two, private.Rand(r, private))
}

func setPublicKey(private *big.Int, group *NumberGroup) *big.Int {
	return new(big.Int).Exp(group.g, private, group.p)
}

func secretKey(private1, public2, p *big.Int) *big.Int {
	return new(big.Int).Exp(public2, private1, p)
}
