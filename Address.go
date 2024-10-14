package main

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcutil/base58"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
	"strconv"
)

func GenKeyPair() ([]byte, []byte) {
	privateKey, err := btcec.NewPrivateKey()

	if err != nil {
		panic(err)
	}

	publicKey := privateKey.PubKey()

	return privateKey.Serialize(), publicKey.SerializeUncompressed()
}

func GenETHAddress(publicKey []byte) string {
	if len(publicKey) != 64 {
		panic("Invalid public key length: " + string(rune(len(publicKey))))
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey)
	hashed := hash.Sum(nil)

	return hex.EncodeToString(hashed[12:])
}

func GenTronAddress(publicKey []byte) string {
	if len(publicKey) != 64 {
		panic("Invalid public key length: " + strconv.Itoa(len(publicKey)))
	}

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey)
	hashed := hash.Sum(nil)

	midstAddress := append([]byte{0x41}, hashed[12:]...)
	midst := sha256.Sum256(midstAddress)
	captcha := sha256.Sum256(midst[:])
	byteAddress := append(midstAddress, captcha[:4]...)

	return base58.Encode(byteAddress)
}

func GenKeyETHWallet() (string, string) {
	pri, pub := GenKeyPair()

	return hex.EncodeToString(pri), GenETHAddress(pub[1:])
}

func GenKeyTronWallet() (string, string) {
	pri, pub := GenKeyPair()

	return hex.EncodeToString(pri), GenTronAddress(pub[1:])
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

func KeyPairFromPhrase(phrase string, path []uint32) ([]byte, []byte) {
	seed := bip39.NewSeed(phrase, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		panic(err)
	}
	key, err := masterKey.NewChildKey(path[0])
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(path[1])
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(path[2])
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(path[3])
	if err != nil {
		panic(err)
	}
	key, err = key.NewChildKey(path[4])
	if err != nil {
		panic(err)
	}
	privateKey, publicKey := btcec.PrivKeyFromBytes(key.Key)
	return privateKey.Serialize(), publicKey.SerializeUncompressed()
}

func GenPhraseETHWallet(len int) (string, string) {
	_, phrase, err := GenSecretPhrase(len)
	if err != nil {
		panic(err)
	}

	_, pub := KeyPairFromPhrase(phrase, []uint32{44 + 0x80000000, 60 + 0x80000000, 0 + 0x80000000, 0, 0})
	add := GenETHAddress(pub[1:])

	return phrase, add
}

func GenPhraseTronWallet(len int) (string, string) {
	_, phrase, err := GenSecretPhrase(len)
	if err != nil {
		panic(err)
	}

	_, pub := KeyPairFromPhrase(phrase, []uint32{44 + 0x80000000, 195 + 0x80000000, 0 + 0x80000000, 0, 0})
	add := GenTronAddress(pub[1:])

	return phrase, add
}
