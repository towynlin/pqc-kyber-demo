package main

import (
	"fmt"

	"github.com/cloudflare/circl/kem/kyber/kyber768"
)

func main() {
	// Bob's Kyber768 key pair
	pk, sk, err := kyber768.GenerateKeyPair(nil)
	if err != nil {
		panic("nope")
	}

	// Alice encapsulates a shared secret with Bob's public key
	ct := make([]byte, kyber768.CiphertextSize)
	ssa := make([]byte, kyber768.SharedKeySize)
	pk.EncapsulateTo(ct, ssa, nil)
	fmt.Printf("Alice's shared secret: %x\n", ssa)
	fmt.Printf("\nCiphertext: %x\n\n", ct)

	// Bob decapsulates the shared secret from the ciphertext
	ssb := make([]byte, kyber768.SharedKeySize)
	sk.DecapsulateTo(ssb, ct)
	fmt.Printf("Bob's decapsulated shared secret: %x\n", ssb)
}
