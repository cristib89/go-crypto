package zkp_dh

import(
	"testing"
)

func TestZkpDH(t *testing.T){

	zkpAlice := New("alice","FFFF0000")
	zkpBob := New("bob","FFFF0000")
	
	alicePubKey := zkpAlice.SendPublicKey()
	bobPubKey := zkpBob.SendPublicKey()

	zkpAlice.SetSharedSecret(bobPubKey)
	zkpBob.SetSharedSecret(alicePubKey)

	if(alicePubKey.Cmp(bobPubKey) == 0) {
		t.Error("Something is fishy, why do we have identical public keys?")
	}

	if(zkpAlice.GetSharedSecret().Cmp(zkpBob.GetSharedSecret()) != 0){
		t.Error("Exchanged keys do not match!")
	}

	// TODO - add ZKP proof
	// https://en.wikipedia.org/wiki/Zero-knowledge_proof

}