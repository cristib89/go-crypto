package http_dh

import (
	"testing"
)

func TestHttpDh(t *testing.T) {
	
	alice := NewHttpDhClient("alice", ":8082", "http://localhost:8083")
	bob := NewHttpDhClient("bob",":8083","http://localhost:8082")
	go func() {
		alice.LaunchServer()
	}()
	go func() {
		bob.LaunchServer()
	}()
	alice.SetSharedSecret()
	if(alice.dhClient.GetSharedSecret().Cmp(bob.dhClient.GetSharedSecret()) != 0){
		t.Error("Exchanged keys do not match!")
	}
}
