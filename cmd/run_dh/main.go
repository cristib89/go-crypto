package main

import (
	"os"
	"github.com/cristib89/go-crypto/dh"
)

func main() {
	id := os.Args[1]
	localPort := os.Args[2]
	remoteAddr := os.Args[3]
	alice := dh.NewHttpDhClient(id, localPort, remoteAddr)
	alice.LaunchServer()
	
}