package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"log"
	"math/big"
)

type pubkey struct {
	x, y *big.Int
}

// Coords returns the x and y public coordinates for key
func (p *pubkey) Coords() (x, y *big.Int) {
	return p.x, p.y
}

// Bytes returns the concatenation of the x and y public coordinates for the key
func (p *pubkey) Bytes() []byte {
	return append(p.x.Bytes(), p.y.Bytes()...)
}

// NewPubkey returns a new pubkey for the given x and y public coordinates
func NewPubKey(x, y *big.Int) *pubkey {
	return &pubkey{x, y}
}

type keypair struct {
	priv []byte
	pub  *pubkey
}

// ComputeShared computes the shared secret for it and another public key
func (kp *keypair) ComputeShared(curve elliptic.Curve, pub *pubkey) ([]byte, error) {
	x, y := pub.Coords()
	secret, _ := curve.ScalarMult(x, y, kp.priv)
	return secret.Bytes(), nil
}

// NewKeypair returns a new keypair
func NewKeypair(priv []byte, pub *pubkey) *keypair {
	return &keypair{priv, pub}
}

func main() {
	curve := elliptic.P256()

	// step 0a) generate first ephemeral ECDH keypair from pseudo-random generator
	priv, x, y, err := elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	pub1 := NewPubKey(x, y)
	kp1 := NewKeypair(priv, pub1)

	// step 0b) generate another ephermeral ECDH keypair from pseudo-random generator
	priv, x, y, err = elliptic.GenerateKey(curve, rand.Reader)
	if err != nil {
		log.Fatal(err)
	}
	pub2 := NewPubKey(x, y)
	kp2 := NewKeypair(priv, pub2)

	// step 1) exchange public keys (implied via function scope)
	// step 2) perform key agreement (implied via function scope)

	// step 3) extract shared secret (they should match)
	ss1, err := kp1.ComputeShared(curve, pub2)
	ss2, err := kp2.ComputeShared(curve, pub1)

	fmt.Printf("First secret: %x\n", ss1)
	fmt.Printf("Secret secret: %x\n", ss2)

	// step 4) derive final key by hashing secret with the two pubkeys (as advised in RFC7748)
	pubBytes := append(pub1.Bytes(), pub2.Bytes()...)
	k1 := sha256.Sum256(append(ss1, pubBytes...))
	k2 := sha256.Sum256(append(ss2, pubBytes...))

	fmt.Printf("First key: %x\n", k1)
	fmt.Printf("Second key: %x\n", k2)
}
