package main

import (
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"golang.org/x/crypto/sha3"
)

func GenKeyPair() ([]byte, []byte) {
	privateKey, err := btcec.NewPrivateKey()

	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PubKey()

	return privateKey.Serialize(), publicKey.SerializeUncompressed()
}

func GenAddress(publicKey []byte) string {
	if len(publicKey) != 65 {
		fmt.Println(len(publicKey))
		panic("Invalid public key length")
	}
	publicKey = publicKey[1:]

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey)
	hashed := hash.Sum(nil)

	return fmt.Sprintf("%x", hashed[12:])
}

func GenWallet() (string, string) {
	pri, pub := GenKeyPair()

	return fmt.Sprintf("%x", pri), GenAddress(pub)
}
