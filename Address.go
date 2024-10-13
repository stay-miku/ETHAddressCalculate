package main

import (
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
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
	if len(publicKey) != 64 {
		panic("Invalid public key length: " + string(rune(len(publicKey))))
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey)
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed[12:])
}

func GenKeyWallet() (string, string) {
	pri, pub := GenKeyPair()

	return hex.EncodeToString(pri), GenAddress(pub[1:])
}

func GenSecretPhrase(len int) ([]byte, string, error) {
	entropy, err := bip39.NewEntropy(len)
	if err != nil {
		return nil, "", err
	}

	phrase, err := bip39.NewMnemonic(entropy)

	if err != nil {
		return nil, "", err
	}

	return entropy, phrase, nil
}

func KeyPairFromPhrase(phrase string) ([]byte, []byte) {
	seed := bip39.NewSeed(phrase, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}
	key, err := masterKey.NewChildKey(44 + 0x80000000)
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(60 + 0x80000000)
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(0 + 0x80000000)
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(0)
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(0)
	if err != nil {
		panic(err)
	}
	privateKey, publicKey := btcec.PrivKeyFromBytes(key.Key)
	return privateKey.Serialize(), publicKey.SerializeUncompressed()
}

func GenPhraseWallet(len int) (string, string) {
	_, phrase, err := GenSecretPhrase(len)
	if err != nil {
		panic(err)
	}

	_, pub := KeyPairFromPhrase(phrase)
	add := GenAddress(pub[1:])

	return phrase, add
}
